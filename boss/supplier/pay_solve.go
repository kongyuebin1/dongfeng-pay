/***************************************************
 ** @Desc : 处理订单状态，用户加款等核心业务
 ** @Time : 2019/10/31 11:44
 ** @Author : yuebin
 ** @File : pay_solve
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/31 11:44
 ** @Software: GoLand
****************************************************/
package controller

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"dongfeng-pay/service/common"
	"dongfeng-pay/service/message_queue"
	"dongfeng-pay/service/models"
	"dongfeng-pay/service/utils"
	url2 "net/url"
	"strconv"
)

type PaySolveController struct {
}

//处理支付成功的加款等各项操作
func (c *PaySolveController) SolvePaySuccess(bankOrderId string, factAmount float64, trxNo string) bool {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error("SolvePaySuccess fail， call rollback")
		}
	}(o)

	var orderInfo models.OrderInfo
	if err := o.Raw("select * from order_info where bank_order_id = ? for update", bankOrderId).QueryRow(&orderInfo); err != nil || orderInfo.BankOrderId == "" {
		o.Rollback()
		logs.Error("不存在该订单，或者select for update出错")
		return false
	}

	if orderInfo.Status != "wait" {
		o.Rollback()
		logs.Error("该订单已经处理，订单号=", bankOrderId)
		return false
	}

	if factAmount <= common.ZERO {
		factAmount = orderInfo.OrderAmount
	}

	var orderProfitInfo models.OrderProfitInfo
	if err := o.Raw("select * from order_profit_info where bank_order_id = ? for update", bankOrderId).QueryRow(&orderProfitInfo); err != nil || orderProfitInfo.BankOrderId == "" {
		logs.Error("select order_profit_info for update fail: ", err)
		o.Rollback()
		return false
	}

	if orderProfitInfo.BankOrderId == "" {
		logs.Error("solve pay success, get orderProfit fail, bankOrderId = ", bankOrderId)
		o.Rollback()
		return false
	}

	comp := c.CompareOrderAndFactAmount(factAmount, orderInfo)
	//如果实际支付金额比订单金额大或者小，那么重新按照实际金额金额利润计算
	if comp != 0 {
		orderProfitInfo.FactAmount = factAmount
		orderProfitInfo.SupplierProfit = orderInfo.FactAmount * orderProfitInfo.SupplierRate
		orderProfitInfo.PlatformProfit = orderInfo.FactAmount * orderProfitInfo.PlatformRate
		orderProfitInfo.AgentProfit = orderInfo.FactAmount * orderProfitInfo.AgentRate
		orderProfitInfo.AllProfit = orderProfitInfo.SupplierProfit + orderProfitInfo.PlatformProfit + orderProfitInfo.AgentProfit
		orderProfitInfo.UserInAmount = orderProfitInfo.FactAmount - orderProfitInfo.AllProfit
		orderProfitInfo.UpdateTime = utils.GetBasicDateTime()

		orderInfo.FactAmount = factAmount
		//如果实际支付金额跟订单金额有出入，那么需要重新更新利润记录
		if _, err := o.Update(orderProfitInfo); err != nil {
			logs.Info("solve pay success fail：", err)
		}
	}

	orderInfo.Status = common.SUCCESS
	orderInfo.BankTransId = trxNo
	orderInfo.UpdateTime = utils.GetBasicDateTime()
	if _, err := o.Update(&orderInfo); err != nil || orderInfo.BankOrderId == "" {
		logs.Error(fmt.Sprintf("solve pay success, update order info fail: %s, bankOrderId = %s", err, bankOrderId))
		o.Rollback()
		return false
	}

	//插入一条待结算记录
	settAmount := orderProfitInfo.FactAmount - orderProfitInfo.SupplierProfit - orderProfitInfo.PlatformProfit - orderProfitInfo.AgentProfit
	if settAmount <= 0.00 {
		logs.Error(fmt.Sprintf("订单id=%s，计算利润存在异常", bankOrderId))
		o.Rollback()
		return false
	}
	orderSettleInfo := models.OrderSettleInfo{PayTypeCode: orderInfo.PayTypeCode, PayProductCode: orderInfo.PayProductCode, RoadUid: orderInfo.RoadUid,
		PayProductName: orderInfo.PayProductName, PayTypeName: orderInfo.PayTypeName, MerchantUid: orderInfo.MerchantUid, MerchantOrderId: orderInfo.MerchantOrderId,
		MerchantName: orderInfo.MerchantName, BankOrderId: bankOrderId, SettleAmount: settAmount, IsAllowSettle: common.YES,
		IsCompleteSettle: common.NO, UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

	if _, err := o.Insert(&orderSettleInfo); err != nil {
		logs.Error(fmt.Sprintf("solve pay success，insert order settle info fail: %s, bankOrderId = %s", err, bankOrderId))
		o.Rollback()
		return false
	}

	//做账户的加款操作，最重要的一部
	var accountInfo models.AccountInfo
	if err := o.Raw("select * from account_info where account_uid = ? for update", orderInfo.MerchantUid).QueryRow(&accountInfo); err != nil || accountInfo.AccountUid == "" {
		logs.Error(fmt.Sprintf("solve pay success, raw account info fail: %s, bankOrderId = %s", err, bankOrderId))
		o.Rollback()
		return false
	}
	if _, err := o.QueryTable(models.ACCOUNT_INFO).Filter("account_uid", orderInfo.MerchantUid).
		Update(orm.Params{"balance": accountInfo.Balance + settAmount, "wait_amount": accountInfo.WaitAmount + settAmount}); err != nil {
		logs.Error(fmt.Sprintf("solve pay success, update account info fail: %s, bankOrderId = %s", err, bankOrderId))
		o.Rollback()
		return false
	}

	//添加一条动账记录
	accountHistory := models.AccountHistoryInfo{AccountUid: orderInfo.MerchantUid, AccountName: orderInfo.MerchantName,
		Type: common.PLUS_AMOUNT, Amount: settAmount, Balance: accountInfo.Balance + settAmount,
		UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}
	if _, err := o.Insert(&accountHistory); err != nil {
		logs.Error(fmt.Sprintf("solve pay success，insert account history fail：%s, bankOrderId = %s", err, bankOrderId))
		o.Rollback()
		return false
	}

	//更新通道信息
	roadInfo := models.GetRoadInfoByRoadUid(orderInfo.RoadUid)
	roadInfo.UpdateTime = utils.GetBasicDateTime()
	roadInfo.RequestSuccess += 1
	roadInfo.TotalIncome += orderInfo.FactAmount
	roadInfo.TodayIncome += orderInfo.FactAmount
	roadInfo.TodayProfit += orderProfitInfo.PlatformProfit + orderProfitInfo.AgentProfit
	roadInfo.TotalProfit += orderProfitInfo.PlatformProfit + orderProfitInfo.AgentProfit
	roadInfo.UpdateTime = utils.GetBasicDateTime()
	if _, err := o.Update(&roadInfo); err != nil {
		logs.Error(fmt.Sprintf("solve pay success, update road info fail: %s, bankOrderId = %s", err, bankOrderId))
		o.Rollback()
		return false
	}

	//更新订单利润表
	orderProfitInfo.Status = common.SUCCESS
	orderProfitInfo.UpdateTime = utils.GetBasicDateTime()
	if _, err := o.Update(&orderProfitInfo); err != nil {
		logs.Error(fmt.Sprintf("solve pay success, update order profit info fail:  %s, bankOrderId = %s", err, bankOrderId))
		o.Rollback()
		return false
	}

	if err := o.Commit(); err != nil {
		logs.Error(fmt.Sprintf("订单bankOrderId = %s，加款失败！", bankOrderId))
		logs.Error("失败原因：", err)
		return false
	} else {
		logs.Info("账户加款成功，并记录了账户历史！")
		//如果处理成功，发送到消息队列，进行商户的回调操作
		go c.CreateOrderNotifyInfo(orderInfo, common.SUCCESS)
	}
	return true
}

//处理支付失败
func (c *PaySolveController) SolvePayFail(orderInfo models.OrderInfo, str string) bool {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error("SolvePaySuccess fail， call rollback")
		}
	}(o)

	var orderTmp models.OrderInfo
	bankOrderId := orderInfo.BankOrderId
	if err := o.Raw("select * from order_info where bank_order_id = ?", bankOrderId).QueryRow(&orderTmp); err != nil || orderTmp.BankOrderId == "" {
		o.Rollback()
		return false
	}

	if orderTmp.Status != "wait" {
		o.Rollback()
		return false
	}
	_, err1 := o.QueryTable(models.ORDER_INFO).Filter("bank_order_id", bankOrderId).Update(orm.Params{"status": str, "bank_trans_id": orderInfo.BankTransId})
	_, err2 := o.QueryTable(models.ORDER_PROFIT_INFO).Filter("bank_order_id", bankOrderId).Update(orm.Params{"status": str, "bank_trans_id": orderInfo.BankTransId})
	if err1 != nil || err2 != nil {
		logs.Error("SolvePayFail fail: ", err1, err2)
		o.Rollback()
		return false
	} else {
		o.Commit()
		go c.CreateOrderNotifyInfo(orderInfo, common.FAIL)
		return true
	}
}

//处理订单冻结
func (c *PaySolveController) SolveOrderFreeze(bankOrderId string) bool {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if err := recover(); err != nil {
			o.Rollback()
			return
		}
	}(o)

	var orderInfo models.OrderInfo
	if err := o.Raw("select * from order_info where bank_order_id = ? for update", bankOrderId).QueryRow(&orderInfo); err != nil || orderInfo.BankOrderId == "" {
		logs.Error("solve order freeze 不存在这样的订单记录，bankOrderId = ", bankOrderId)
		o.Rollback()
		return false
	}

	if orderInfo.Status != common.SUCCESS {
		o.Rollback()
		return false
	}

	orderInfo.Freeze = common.YES
	orderInfo.FreezeTime = utils.GetBasicDateTime()
	orderInfo.UpdateTime = utils.GetBasicDateTime()
	if _, err := o.Update(&orderInfo); err != nil {
		logs.Error("solve order freeze fail: ", err)
		o.Rollback()
		return false
	}

	//账户的冻结金额里面加入相应的金额
	orderProfitInfo := models.GetOrderProfitByBankOrderId(bankOrderId)
	var accountInfo models.AccountInfo
	if err := o.Raw("select * from account_info where account_uid = ? for update", orderInfo.MerchantUid).QueryRow(&accountInfo); err != nil || accountInfo.AccountUid == "" {
		logs.Error(fmt.Sprintf("solve pay fail select acount fail：%s", err))
		o.Rollback()
		return false
	}
	accountInfo.UpdateTime = utils.GetBasicDateTime()
	accountInfo.FreezeAmount = accountInfo.FreezeAmount + orderProfitInfo.UserInAmount
	if _, err := o.Update(&accountInfo); err != nil {
		logs.Error("solve order freeze fail: ", err)
		o.Rollback()
		return false
	}
	//插入一条动账记录
	accountHistoryInfo := models.AccountHistoryInfo{AccountName: accountInfo.AccountName, AccountUid: accountInfo.AccountUid,
		Type: common.FREEZE_AMOUNT, Amount: orderProfitInfo.UserInAmount, Balance: accountInfo.Balance, UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}
	if _, err := o.Insert(&accountHistoryInfo); err != nil {
		logs.Error("solve order freeze fail: ", err)
		o.Rollback()
		return false
	}

	if err := o.Commit(); err != nil {
		logs.Error("SolveOrderFreeze fail")
	} else {
		logs.Info("冻结处理成功")
	}
	return true
}

//订单解冻
func (c *PaySolveController) SolveOrderUnfreeze(bankOrderId string) bool {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if err := recover(); err != nil {
			o.Rollback()
			return
		}
	}(o)

	orderInfo := new(models.OrderInfo)
	if err := o.Raw("select * from order_info where bank_order_id = ? for update", bankOrderId).QueryRow(orderInfo); err != nil || orderInfo.BankOrderId == "" {
		logs.Error("solve order unfreeze 不存在这样的订单记录，bankOrderId = ", bankOrderId)
		return false
	}

	orderInfo.Freeze = ""
	orderInfo.Unfreeze = common.YES
	orderInfo.UnfreezeTime = utils.GetBasicDateTime()
	orderInfo.UpdateTime = utils.GetBasicDateTime()
	if _, err := o.Update(orderInfo); err != nil {
		logs.Error("solve order unfreeze fail: ", err)
		o.Rollback()
		return false
	}

	orderProfitInfo := models.GetOrderProfitByBankOrderId(bankOrderId)

	accountInfo := new(models.AccountInfo)
	if err := o.Raw("select * from account_info where account_uid = ? for update", orderInfo.MerchantUid).QueryRow(accountInfo); err != nil || accountInfo.AccountUid == "" {
		logs.Error(fmt.Sprintf("unfreeze select account fail: %s", err))
		o.Rollback()
		return false
	}
	accountInfo.UpdateTime = utils.GetBasicDateTime()
	accountInfo.FreezeAmount = accountInfo.FreezeAmount - orderProfitInfo.UserInAmount

	if _, err := o.Update(accountInfo); err != nil {
		logs.Error("solve order unfreeze fail: ", err)
		o.Rollback()
		return false
	}

	accountHistoryInfo := models.AccountHistoryInfo{AccountUid: accountInfo.AccountUid, AccountName: accountInfo.AccountName, Type: common.UNFREEZE_AMOUNT,
		Amount: orderProfitInfo.UserInAmount, Balance: accountInfo.Balance, UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

	if _, err := o.Insert(&accountHistoryInfo); err != nil {
		logs.Error("solve order unfreeze fail: ", err)
		o.Rollback()
		return false
	}

	if err := o.Commit(); err != nil {
		logs.Error(fmt.Sprintf("unfreeze commit fail: %s", err))
		return false
	} else {
		logs.Info("解冻成功")
	}
	return true
}

func (c *PaySolveController) SolveRefund(bankOrderId string) bool {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if err := recover(); err != nil {
			o.Rollback()
			return
		}
	}(o)

	orderInfo := new(models.OrderInfo)
	if err := o.Raw("select * from order_info where bank_order_id = ? for update", bankOrderId).QueryRow(orderInfo); err != nil || orderInfo.BankOrderId == "" {
		logs.Error("solve refund 不存在这样的订单，bankOrderId = " + bankOrderId)
		return false
	}

	orderInfo.UpdateTime = utils.GetBasicDateTime()
	orderInfo.Refund = common.YES
	orderInfo.RefundTime = utils.GetBasicDateTime()

	orderProfitInfo := models.GetOrderProfitByBankOrderId(bankOrderId)
	account := new(models.AccountInfo)
	if err := o.Raw("select * from account_info where account_uid = ? for update", orderInfo.MerchantUid).QueryRow(account); err != nil || account.AccountUid == "" {
		o.Rollback()
		return false
	}

	account.UpdateTime = utils.GetBasicDateTime()
	account.SettleAmount = account.SettleAmount - orderProfitInfo.UserInAmount
	account.Balance = account.Balance - orderProfitInfo.UserInAmount

	if orderInfo.Freeze == common.YES {
		account.FreezeAmount = account.FreezeAmount - orderProfitInfo.UserInAmount
		if account.FreezeAmount < 0 {
			account.FreezeAmount = common.ZERO
		}
		orderInfo.Freeze = ""
	}

	if _, err := o.Update(orderInfo); err != nil {
		logs.Error("solve order refund update order info fail: ", err)
		o.Rollback()
		return false
	}
	if _, err := o.Update(account); err != nil {
		logs.Error("solve order refund update account fail: ", err)
		o.Rollback()
		return false
	}

	accountHistoryInfo := models.AccountHistoryInfo{AccountName: account.AccountName, AccountUid: account.AccountUid,
		Type: common.REFUND, Amount: orderProfitInfo.UserInAmount, Balance: account.Balance,
		UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

	if _, err := o.Insert(&accountHistoryInfo); err != nil {
		logs.Error("solve order refund insert account history fail: ", err)
		o.Rollback()
		return false
	}

	if err := o.Commit(); err != nil {
		logs.Error("退款处理失败，fail： ", err)
	} else {
		logs.Info("退款处理成功")
	}
	return true
}

func (c *PaySolveController) SolveOrderRoll(bankOrderId string) bool {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if err := recover(); err != nil {
			o.Rollback()
			return
		}
	}(o)

	orderInfo := new(models.OrderInfo)

	if err := o.Raw("select * from order_info where bank_order_id = ? for update", bankOrderId).QueryRow(orderInfo); err != nil {
		logs.Error("solve order roll fail： ", err)
		o.Rollback()
		return false
	}

	if orderInfo.Status != common.SUCCESS {
		logs.Error("solve order roll 订单不存在或者订单状态不是success, bankOrderId=", bankOrderId)
		o.Rollback()
		return false
	}
	orderInfo.UpdateTime = utils.GetBasicDateTime()

	orderProfitInfo := models.GetOrderProfitByBankOrderId(bankOrderId)

	account := new(models.AccountInfo)
	if err := o.Raw("select * from account_info where account_uid = ? for update", orderInfo.MerchantUid).QueryRow(account); err != nil || account.AccountUid == "" {
		logs.Error("solve order roll get account is nil, accountUid = ", orderInfo.MerchantUid)
		o.Rollback()
		return false
	}

	account.UpdateTime = utils.GetBasicDateTime()
	if orderInfo.Refund == common.YES {
		account.Balance = account.Balance + orderProfitInfo.UserInAmount
		account.SettleAmount = account.SettleAmount + orderProfitInfo.UserInAmount
		orderInfo.Refund = common.NO
	}

	if _, err := o.Update(orderInfo); err != nil {
		logs.Error("solve order roll fail update order info fail:  ", err)
		o.Rollback()
		return false
	}
	if _, err := o.Update(account); err != nil {
		logs.Error("solve order roll update account fail: ", err)
		o.Rollback()
		return false
	}

	accountHistoryInfo := models.AccountHistoryInfo{AccountUid: account.AccountUid, AccountName: account.AccountName,
		Type: common.PLUS_AMOUNT, Amount: orderProfitInfo.UserInAmount, Balance: account.Balance,
		UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}
	if _, err := o.Insert(&accountHistoryInfo); err != nil {
		logs.Error("solve order roll insert account history fail: ", err)
		o.Rollback()
		return false
	}

	if err := o.Commit(); err != nil {
		logs.Error("处理订单回滚失败,fail: ", err)
	} else {
		logs.Info("处理订单回滚成功")
	}

	return true
}

//比较订单金额和实际支付金额的大小
func (c *PaySolveController) CompareOrderAndFactAmount(factAmount float64, orderInfo models.OrderInfo) int {
	orderAmount := orderInfo.OrderAmount
	//将金额放大1000倍
	oa := int64(orderAmount * 1000)
	fa := int64(factAmount * 1000)
	if oa > fa {
		//如果实际金额大，返回1
		return 1
	} else if oa == fa {
		return 0
	} else {
		return 2
	}
}

//支付完成后，处理给商户的回调信息
func (c *PaySolveController) CreateOrderNotifyInfo(orderInfo models.OrderInfo, tradeStatus string) {

	notifyInfo := new(models.NotifyInfo)
	notifyInfo.Type = "order"
	notifyInfo.BankOrderId = orderInfo.BankOrderId
	notifyInfo.MerchantOrderId = orderInfo.MerchantOrderId
	notifyInfo.Status = "wait"
	notifyInfo.Times = 0
	notifyInfo.UpdateTime = utils.GetBasicDateTime()
	notifyInfo.CreateTime = utils.GetBasicDateTime()

	merchantInfo := models.GetMerchantByUid(orderInfo.MerchantUid)

	params := make(map[string]string)
	params["orderNo"] = orderInfo.MerchantOrderId
	params["orderPrice"] = strconv.FormatFloat(orderInfo.OrderAmount, 'f', 2, 64)
	params["factPrice"] = strconv.FormatFloat(orderInfo.FactAmount, 'f', 2, 64)
	params["orderTime"] = utils.GetDateTimeNot()

	if orderInfo.BankTransId != "" {
		params["trxNo"] = orderInfo.BankTransId
	} else {
		params["trxNo"] = orderInfo.BankOrderId
	}
	params["statusCode"] = "00"
	params["tradeStatus"] = tradeStatus
	params["payKey"] = merchantInfo.MerchantKey

	params["sign"] = utils.GetMD5Sign(params, utils.SortMap(params), merchantInfo.MerchantSecret)

	url := url2.Values{}
	for k, v := range params {
		url.Add(k, v)
	}

	notifyInfo.Url = orderInfo.NotifyUrl + "?" + url.Encode()

	if models.InsertNotifyInfo(*notifyInfo) {
		logs.Info(fmt.Sprintf("订单bankOrderId=%s，已经将回调地址插入数据库", orderInfo.BankOrderId))
	} else {
		logs.Error(fmt.Sprintf("订单bankOrderId=%s，插入回调数据库失败", orderInfo.BankOrderId))
	}
	//将订单发送到消息队列，给下面的商户进行回调
	message_queue.SendMessage(common.MqOrderNotify, orderInfo.BankOrderId)
}
