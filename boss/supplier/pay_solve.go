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
	"boss/common"
	"boss/message_queue"
	"boss/models"
	"boss/utils"
	"context"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	url2 "net/url"
	"strconv"
)

type PaySolveController struct {
}

//处理支付成功的加款等各项操作
func (c *PaySolveController) SolvePaySuccess(bankOrderId string, factAmount float64, trxNo string) bool {

	o := orm.NewOrm()

	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		var orderInfo models.OrderInfo
		if err := txOrm.Raw("select * from order_info where bank_order_id = ? for update", bankOrderId).QueryRow(&orderInfo); err != nil || orderInfo.BankOrderId == "" {
			logs.Error("不存在该订单，或者select for update出错")
			return err
		}

		if orderInfo.Status != "wait" {
			logs.Error("该订单已经处理，订单号=", bankOrderId)
			return errors.New(fmt.Sprintf("该订单已经处理，订单号= %s", bankOrderId))
		}

		if factAmount <= common.ZERO {
			factAmount = orderInfo.OrderAmount
		}

		var orderProfitInfo models.OrderProfitInfo
		if err := txOrm.Raw("select * from order_profit_info where bank_order_id = ? for update", bankOrderId).QueryRow(&orderProfitInfo); err != nil || orderProfitInfo.BankOrderId == "" {
			logs.Error("select order_profit_info for update fail: ", err)
			return err
		}

		if orderProfitInfo.BankOrderId == "" {
			logs.Error("solve pay success, get orderProfit fail, bankOrderId = ", bankOrderId)
			return errors.New(fmt.Sprintf("solve pay success, get orderProfit fail, bankOrderId = %s", bankOrderId))
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
			if _, err := txOrm.Update(orderProfitInfo); err != nil {
				logs.Info("solve pay success fail：", err)
				return err
			}
		}

		orderInfo.Status = common.SUCCESS
		orderInfo.BankTransId = trxNo
		orderInfo.UpdateTime = utils.GetBasicDateTime()
		if _, err := txOrm.Update(&orderInfo); err != nil || orderInfo.BankOrderId == "" {
			logs.Error(fmt.Sprintf("solve pay success, update order info fail: %s, bankOrderId = %s", err, bankOrderId))
			return err
		}

		//插入一条待结算记录
		settAmount := orderProfitInfo.FactAmount - orderProfitInfo.SupplierProfit - orderProfitInfo.PlatformProfit - orderProfitInfo.AgentProfit
		if settAmount <= 0.00 {
			logs.Error(fmt.Sprintf("订单id=%s，计算利润存在异常", bankOrderId))
			return errors.New(fmt.Sprintf("订单id=%s，计算利润存在异常", bankOrderId))
		}
		orderSettleInfo := models.OrderSettleInfo{PayTypeCode: orderInfo.PayTypeCode, PayProductCode: orderInfo.PayProductCode, RoadUid: orderInfo.RoadUid,
			PayProductName: orderInfo.PayProductName, PayTypeName: orderInfo.PayTypeName, MerchantUid: orderInfo.MerchantUid, MerchantOrderId: orderInfo.MerchantOrderId,
			MerchantName: orderInfo.MerchantName, BankOrderId: bankOrderId, SettleAmount: settAmount, IsAllowSettle: common.YES,
			IsCompleteSettle: common.NO, UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

		if _, err := txOrm.Insert(&orderSettleInfo); err != nil {
			logs.Error(fmt.Sprintf("solve pay success，insert order settle info fail: %s, bankOrderId = %s", err, bankOrderId))
			return err
		}

		//做账户的加款操作，最重要的一部
		var accountInfo models.AccountInfo
		if err := txOrm.Raw("select * from account_info where account_uid = ? for update", orderInfo.MerchantUid).QueryRow(&accountInfo); err != nil || accountInfo.AccountUid == "" {
			logs.Error(fmt.Sprintf("solve pay success, raw account info fail: %s, bankOrderId = %s", err, bankOrderId))
			return err
		}
		if _, err := txOrm.QueryTable(models.ACCOUNT_INFO).Filter("account_uid", orderInfo.MerchantUid).
			Update((orm.Params{"balance": accountInfo.Balance + settAmount, "wait_amount": accountInfo.WaitAmount + settAmount})); err != nil {
			logs.Error(fmt.Sprintf("solve pay success, update account info fail: %s, bankOrderId = %s", err, bankOrderId))
			return err
		}

		//添加一条动账记录
		accountHistory := models.AccountHistoryInfo{AccountUid: orderInfo.MerchantUid, AccountName: orderInfo.MerchantName,
			Type: common.PLUS_AMOUNT, Amount: settAmount, Balance: accountInfo.Balance + settAmount,
			UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}
		if _, err := txOrm.Insert(&accountHistory); err != nil {
			logs.Error(fmt.Sprintf("solve pay success，insert account history fail：%s, bankOrderId = %s", err, bankOrderId))
			return err
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
		if _, err := txOrm.Update(&roadInfo); err != nil {
			logs.Error(fmt.Sprintf("solve pay success, update road info fail: %s, bankOrderId = %s", err, bankOrderId))
			return err
		}

		//更新订单利润表
		orderProfitInfo.Status = common.SUCCESS
		orderProfitInfo.UpdateTime = utils.GetBasicDateTime()
		if _, err := txOrm.Update(&orderProfitInfo); err != nil {
			logs.Error(fmt.Sprintf("solve pay success, update order profit info fail:  %s, bankOrderId = %s", err, bankOrderId))
			return err
		}

		// 给下游发送回调通知
		go c.CreateOrderNotifyInfo(orderInfo, common.SUCCESS)

		return nil
	})

	if err != nil {
		logs.Error("SolvePaySuccess失败：", err)
		return false
	}

	logs.Info("SolvePaySuccess处理成功")
	return true
}

//处理支付失败
func (c *PaySolveController) SolvePayFail(orderInfo models.OrderInfo, str string) bool {
	o := orm.NewOrm()
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		var orderTmp models.OrderInfo
		bankOrderId := orderInfo.BankOrderId
		if err := txOrm.Raw("select * from order_info where bank_order_id = ?", bankOrderId).QueryRow(&orderTmp); err != nil || orderTmp.BankOrderId == "" {
			return err
		}

		if orderTmp.Status != "wait" {
			return errors.New("订单已经处理，不要重复加款")
		}
		if _, err := txOrm.QueryTable(models.ORDER_INFO).Filter("bank_order_id", bankOrderId).Update(orm.Params{"status": str, "bank_trans_id": orderInfo.BankTransId}); err != nil {
			logs.Error("更改订单状态失败：", err)
			return err
		}
		if _, err := txOrm.QueryTable(models.ORDER_PROFIT_INFO).Filter("bank_order_id", bankOrderId).Update(orm.Params{"status": str, "bank_trans_id": orderInfo.BankTransId}); err != nil {
			logs.Error("更改订单状态失败：", err)
			return err
		}

		go c.CreateOrderNotifyInfo(orderInfo, common.FAIL)

		return nil
	})

	if err != nil {
		logs.Error("SolvePayFail：", err)
		return false
	}

	logs.Info("SolvePayFail成功")
	return true
}

//处理订单冻结
func (c *PaySolveController) SolveOrderFreeze(bankOrderId string) bool {
	o := orm.NewOrm()

	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		var orderInfo models.OrderInfo
		if err := txOrm.Raw("select * from order_info where bank_order_id = ? for update", bankOrderId).QueryRow(&orderInfo); err != nil || orderInfo.BankOrderId == "" {
			logs.Error("solve order freeze 不存在这样的订单记录，bankOrderId = ", bankOrderId)
			return err
		}

		if orderInfo.Status != common.SUCCESS {
			logs.Error("非成功订单不能进行冻结")
			return errors.New("非成功订单不能进行冻结")
		}

		orderInfo.Freeze = common.YES
		orderInfo.FreezeTime = utils.GetBasicDateTime()
		orderInfo.UpdateTime = utils.GetBasicDateTime()
		if _, err := txOrm.Update(&orderInfo); err != nil {
			logs.Error("solve order freeze fail: ", err)
			return err
		}

		//账户的冻结金额里面加入相应的金额
		orderProfitInfo := models.GetOrderProfitByBankOrderId(bankOrderId)
		var accountInfo models.AccountInfo
		if err := txOrm.Raw("select * from account_info where account_uid = ? for update", orderInfo.MerchantUid).QueryRow(&accountInfo); err != nil || accountInfo.AccountUid == "" {
			logs.Error(fmt.Sprintf("solve pay fail select acount fail：%s", err))
			return err
		}
		accountInfo.UpdateTime = utils.GetBasicDateTime()
		accountInfo.FreezeAmount = accountInfo.FreezeAmount + orderProfitInfo.UserInAmount
		if _, err := txOrm.Update(&accountInfo); err != nil {
			logs.Error("solve order freeze fail: ", err)
			return err
		}
		//插入一条动账记录
		accountHistoryInfo := models.AccountHistoryInfo{AccountName: accountInfo.AccountName, AccountUid: accountInfo.AccountUid,
			Type: common.FREEZE_AMOUNT, Amount: orderProfitInfo.UserInAmount, Balance: accountInfo.Balance, UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}
		if _, err := txOrm.Insert(&accountHistoryInfo); err != nil {
			logs.Error("solve order freeze fail: ", err)
			return err
		}

		return nil
	})

	if err != nil {
		logs.Error("SolveOrderFreeze：", err)
		return false
	}

	logs.Info("SolveOrderFreeze")

	return true
}

//订单解冻
func (c *PaySolveController) SolveOrderUnfreeze(bankOrderId string) bool {
	o := orm.NewOrm()

	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		orderInfo := new(models.OrderInfo)
		if err := txOrm.Raw("select * from order_info where bank_order_id = ? for update", bankOrderId).QueryRow(orderInfo); err != nil || orderInfo.BankOrderId == "" {
			logs.Error("solve order unfreeze 不存在这样的订单记录，bankOrderId = ", bankOrderId)
			return err
		}

		orderInfo.Freeze = ""
		orderInfo.Unfreeze = common.YES
		orderInfo.UnfreezeTime = utils.GetBasicDateTime()
		orderInfo.UpdateTime = utils.GetBasicDateTime()
		if _, err := txOrm.Update(orderInfo); err != nil {
			logs.Error("solve order unfreeze fail: ", err)
			return err
		}

		orderProfitInfo := models.GetOrderProfitByBankOrderId(bankOrderId)

		accountInfo := new(models.AccountInfo)
		if err := txOrm.Raw("select * from account_info where account_uid = ? for update", orderInfo.MerchantUid).QueryRow(accountInfo); err != nil || accountInfo.AccountUid == "" {
			logs.Error(fmt.Sprintf("unfreeze select account fail: %s", err))
			return err
		}
		accountInfo.UpdateTime = utils.GetBasicDateTime()
		accountInfo.FreezeAmount = accountInfo.FreezeAmount - orderProfitInfo.UserInAmount

		if _, err := txOrm.Update(accountInfo); err != nil {
			logs.Error("solve order unfreeze fail: ", err)
			return err
		}

		accountHistoryInfo := models.AccountHistoryInfo{AccountUid: accountInfo.AccountUid, AccountName: accountInfo.AccountName, Type: common.UNFREEZE_AMOUNT,
			Amount: orderProfitInfo.UserInAmount, Balance: accountInfo.Balance, UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

		if _, err := txOrm.Insert(&accountHistoryInfo); err != nil {
			return err
		}

		return nil
	}); err != nil {
		logs.Error("SolveOrderUnfreeze失败：", err)
		return false
	}

	return true
}

func (c *PaySolveController) SolveRefund(bankOrderId string) bool {
	o := orm.NewOrm()
	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		orderInfo := new(models.OrderInfo)
		if err := txOrm.Raw("select * from order_info where bank_order_id = ? for update", bankOrderId).QueryRow(orderInfo); err != nil || orderInfo.BankOrderId == "" {
			logs.Error("solve refund 不存在这样的订单，bankOrderId = " + bankOrderId)
			return err
		}

		orderInfo.UpdateTime = utils.GetBasicDateTime()
		orderInfo.Refund = common.YES
		orderInfo.RefundTime = utils.GetBasicDateTime()

		orderProfitInfo := models.GetOrderProfitByBankOrderId(bankOrderId)
		account := new(models.AccountInfo)
		if err := txOrm.Raw("select * from account_info where account_uid = ? for update", orderInfo.MerchantUid).QueryRow(account); err != nil || account.AccountUid == "" {
			return err
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

		if _, err := txOrm.Update(orderInfo); err != nil {
			logs.Error("solve order refund update order info fail: ", err)
			return err
		}
		if _, err := txOrm.Update(account); err != nil {
			logs.Error("solve order refund update account fail: ", err)
			return err
		}

		accountHistoryInfo := models.AccountHistoryInfo{AccountName: account.AccountName, AccountUid: account.AccountUid,
			Type: common.REFUND, Amount: orderProfitInfo.UserInAmount, Balance: account.Balance,
			UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

		if _, err := txOrm.Insert(&accountHistoryInfo); err != nil {
			logs.Error("solve order refund insert account history fail: ", err)
			return err
		}

		return nil
	}); err != nil {
		logs.Error("SolveRefund 成功：", err)
		return false
	}
	return true
}

func (c *PaySolveController) SolveOrderRoll(bankOrderId string) bool {
	o := orm.NewOrm()
	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		orderInfo := new(models.OrderInfo)

		if err := txOrm.Raw("select * from order_info where bank_order_id = ? for update", bankOrderId).QueryRow(orderInfo); err != nil {
			logs.Error("solve order roll fail： ", err)
			return err
		}

		if orderInfo.Status != common.SUCCESS {
			logs.Error("solve order roll 订单不存在或者订单状态不是success, bankOrderId=", bankOrderId)
			return errors.New("solve order roll failed")
		}
		orderInfo.UpdateTime = utils.GetBasicDateTime()

		orderProfitInfo := models.GetOrderProfitByBankOrderId(bankOrderId)

		account := new(models.AccountInfo)
		if err := txOrm.Raw("select * from account_info where account_uid = ? for update", orderInfo.MerchantUid).QueryRow(account); err != nil || account.AccountUid == "" {
			return err
		}

		account.UpdateTime = utils.GetBasicDateTime()
		if orderInfo.Refund == common.YES {
			account.Balance = account.Balance + orderProfitInfo.UserInAmount
			account.SettleAmount = account.SettleAmount + orderProfitInfo.UserInAmount
			orderInfo.Refund = common.NO
		}

		if _, err := txOrm.Update(orderInfo); err != nil {
			logs.Error("solve order roll fail update order info fail:  ", err)
			return err
		}
		if _, err := txOrm.Update(account); err != nil {
			logs.Error("solve order roll update account fail: ", err)
			return err
		}

		accountHistoryInfo := models.AccountHistoryInfo{AccountUid: account.AccountUid, AccountName: account.AccountName,
			Type: common.PLUS_AMOUNT, Amount: orderProfitInfo.UserInAmount, Balance: account.Balance,
			UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

		if _, err := txOrm.Insert(&accountHistoryInfo); err != nil {
			logs.Error("solve order roll insert account history fail: ", err)
			return err
		}

		return nil

	}); err != nil {
		logs.Error("SolveOrderRoll处理失败：", err)
		return false
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
