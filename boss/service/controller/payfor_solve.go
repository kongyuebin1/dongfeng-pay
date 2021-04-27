/***************************************************
 ** @Desc : 代付处理
 ** @Time : 2019/11/28 18:52
 ** @Author : yuebin
 ** @File : payfor_solve
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/28 18:52
 ** @Software: GoLand
****************************************************/
package controller

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"dongfeng/service/common"
	"dongfeng/service/message_queue"
	"dongfeng/service/models"
	"dongfeng/service/utils"
	"strings"
)

func PayForFail(payFor models.PayforInfo) bool {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if err := recover(); err != nil {
			logs.Error("pay for fail,rollback....")
			o.Rollback()
		}
		return
	}(o)

	var tmpForPay models.PayforInfo
	if err := o.Raw("select * from payfor_info where bank_order_id = ? for update", payFor.BankOrderId).QueryRow(&tmpForPay); err != nil || tmpForPay.PayforUid == "" {
		logs.Error("solve pay fail select fail：", err)
		o.Rollback()
		return false
	}

	if tmpForPay.Status == common.PAYFOR_FAIL || tmpForPay.Status == common.PAYFOR_SUCCESS {
		logs.Error(fmt.Sprintf("该代付订单uid=%s，状态已经是最终结果", payFor.PayforUid))
		o.Rollback()
		return false
	}
	//更新payfor记录的状态
	tmpForPay.Status = common.PAYFOR_FAIL
	tmpForPay.UpdateTime = utils.GetBasicDateTime()
	if _, err := o.Update(&tmpForPay); err != nil {
		logs.Error("PayForFail update payfor_info fail: ", err)
		o.Rollback()
		return false
	}

	var account models.AccountInfo
	if err := o.Raw("select * from account_info where account_uid = ? for update", payFor.MerchantUid).QueryRow(&account); err != nil || account.AccountUid == "" {
		logs.Error("payfor select account fail：", err)
		o.Rollback()
		return false
	}
	account.UpdateTime = utils.GetBasicDateTime()
	if account.PayforAmount < (payFor.PayforAmount + payFor.PayforFee) {
		logs.Error(fmt.Sprintf("商户uid=%s，账户中待代付金额小于代付记录的金额", payFor.MerchantUid))
		o.Rollback()
		return false
	}
	//将正在打款中的金额减去
	account.PayforAmount = account.PayforAmount - payFor.PayforAmount - payFor.PayforFee

	if _, err := o.Update(&account); err != nil {
		logs.Error("PayForFail update account fail: ", err)
		o.Rollback()
		return false
	}

	if err := o.Commit(); err != nil {
		logs.Error("代付失败处理出错，fail：", err)
	} else {
		logs.Info("代付处理成功")
	}

	return true
}

func PayForSuccess(payFor models.PayforInfo) bool {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if err := recover(); err != nil {
			logs.Error("pay for success,rollback....")
			o.Rollback()
		}
		return
	}(o)

	var tmpPayFor models.PayforInfo
	if err := o.Raw("select * from payfor_info where bank_order_id = ? for update", payFor.BankOrderId).QueryRow(&tmpPayFor); err != nil || tmpPayFor.PayforUid == "" {
		logs.Error("payfor success select payfor fail：", err)
		o.Rollback()
		return false
	}
	if tmpPayFor.Status == common.PAYFOR_FAIL || tmpPayFor.Status == common.PAYFOR_SUCCESS {
		logs.Error(fmt.Sprintf("该代付订单uid=%s，已经是最终结果，不需要处理", payFor.PayforUid))
		o.Rollback()
		return false
	}
	tmpPayFor.UpdateTime = utils.GetBasicDateTime()
	tmpPayFor.Status = common.PAYFOR_SUCCESS
	_, err := o.Update(&tmpPayFor)
	if err != nil {
		logs.Error("PayForSuccess update payfor fail: ", err)
		o.Rollback()
		return false
	}

	var account models.AccountInfo
	if err := o.Raw("select * from account_info where account_uid = ? for update", payFor.MerchantUid).QueryRow(&account); err != nil || account.AccountUid == "" {
		logs.Error("payfor success select account fail：", err)
		o.Rollback()
		return false
	}

	account.UpdateTime = utils.GetBasicDateTime()
	if account.PayforAmount < (payFor.PayforAmount + payFor.PayforFee) {
		logs.Error(fmt.Sprintf("商户uid=%s，账户中待代付金额小于代付记录的金额", payFor.MerchantUid))
		o.Rollback()
		return false
	}

	//代付打款中的金额减去
	account.PayforAmount = account.PayforAmount - payFor.PayforAmount - payFor.PayforFee
	//减去余额，减去可用金额
	account.Balance = account.Balance - payFor.PayforAmount - payFor.PayforFee
	//已结算金额减去
	account.SettleAmount = account.SettleAmount - payFor.PayforAmount - payFor.PayforFee

	if _, err := o.Update(&account); err != nil {
		logs.Error("PayForSuccess udpate account fail：", err)
		return false
	}

	//添加一条动账记录
	accountHistory := models.AccountHistoryInfo{AccountUid: payFor.MerchantUid, AccountName: payFor.MerchantName,
		Type: common.SUB_AMOUNT, Amount: payFor.PayforAmount + payFor.PayforFee, Balance: account.Balance,
		UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}
	if _, err := o.Insert(&accountHistory); err != nil {
		logs.Error("PayForSuccess insert account history fail: ", err)
		o.Rollback()
		return false
	}

	if err := o.Commit(); err != nil {
		logs.Error("代付成功处理失败，fail：", err)
	} else {
		logs.Info("代付处理成功")
	}

	return true
}

/*
* 自动审核代付订单
 */
func SolvePayForConfirm() {
	params := make(map[string]string)
	beforeOneDay := utils.GetDateTimeBeforeDays(1)
	nowDate := utils.GetBasicDateTime()
	params["create_time__lte"] = beforeOneDay
	params["create_time__gte"] = nowDate
	params["status"] = common.PAYFOR_COMFRIM
	payForList := models.GetPayForListByParams(params)
	for _, p := range payForList {
		if p.Type == common.SELF_HELP || p.Type == common.SELF_MERCHANT {
			//系统后台提交的，人工审核
			continue
		}
		//判断商户是否开通了自动代付功能
		merchant := models.GetMerchantByUid(p.MerchantUid)
		//判断商户是否开通了自动代付
		if merchant.AutoPayFor == common.NO || merchant.AutoPayFor == "" {
			logs.Notice(fmt.Sprintf("该商户uid=%s， 没有开通自动代付功能", p.MerchantUid))
			continue
		}
		//找自动代付通道
		findPayForRoad(p, merchant)
	}
}

func findPayForRoad(payFor models.PayforInfo, merchant models.MerchantInfo) bool {
	o := orm.NewOrm()

	//检查是否单独填写了每笔代付的手续费
	if merchant.PayforFee > common.ZERO {
		logs.Info(fmt.Sprintf("商户uid=%s，有单独的代付手续费。", merchant.MerchantUid))
		payFor.PayforFee = merchant.PayforFee
		payFor.PayforTotalAmount = payFor.PayforFee + payFor.PayforAmount
	}

	if merchant.SinglePayForRoadUid != "" {
		payFor.RoadUid = merchant.SinglePayForRoadUid
		payFor.RoadName = merchant.SinglePayForRoadName
	} else {
		//到轮询里面寻找代付通道
		if merchant.RollPayForRoadCode == "" {
			logs.Notice(fmt.Sprintf("该商户没有配置代付通道"))
			return false
		}
		roadPoolInfo := models.GetRoadPoolByRoadPoolCode(merchant.RollPayForRoadCode)
		roadUids := strings.Split(roadPoolInfo.RoadUidPool, "||")
		roadInfoList := models.GetRoadInfosByRoadUids(roadUids)
		if len(roadUids) == 0 || len(roadInfoList) == 0 {
			logs.Error(fmt.Sprintf("通道轮询池=%s, 没有配置通道", merchant.RollPayForRoadCode))
			return false
		}
		payFor.RoadUid = roadInfoList[0].RoadUid
		payFor.RoadName = roadInfoList[0].RoadName
	}

	o.Begin()
	defer func(interface{}) {
		if err := recover(); err != nil {
			o.Rollback()
			return
		}
	}(o)

	var tmpPayFor models.PayforInfo
	if err := o.Raw("select * from payfor_info where payfor_uid = ? for update", payFor.PayforUid).QueryRow(&tmpPayFor); err != nil || tmpPayFor.PayforUid == "" {
		logs.Error("find payfor road select payfor fail：", err)
		o.Rollback()
		return false
	}
	if tmpPayFor.Status != common.PAYFOR_COMFRIM {
		logs.Notice(fmt.Sprintf("该代付记录uid=%s，已经被审核", payFor.PayforUid))
		return false
	}
	tmpPayFor.UpdateTime = utils.GetBasicDateTime()
	tmpPayFor.Status = common.PAYFOR_SOLVING
	tmpPayFor.GiveType = common.PAYFOR_ROAD
	if _, err := o.Update(&tmpPayFor); err != nil {
		logs.Error(fmt.Sprintf("该代付记录uid=%s，从审核更新为正在处理出错: %s", payFor.PayforUid, err))
		o.Rollback()
		return false
	}

	if err := o.Commit(); err != nil {
		logs.Error("挑选代付通道失败，fail：", err)
	} else {
		logs.Info("挑选代付通道成功")
	}

	return true
}

/*
* 执行逻辑
 */
func SolvePayFor() {
	//取出一天之内的没有做处理并且不是手动打款的代付记录
	params := make(map[string]string)
	beforeOneDay := utils.GetDateTimeBeforeDays(1)
	nowDate := utils.GetBasicDateTime()
	params["create_time__lte"] = nowDate
	params["create_time__gte"] = beforeOneDay
	params["is_send"] = "no"
	params["status"] = common.PAYFOR_SOLVING
	params["give_type"] = common.PAYFOR_ROAD

	payForList := models.GetPayForListByParams(params)
	for _, p := range payForList {
		if p.Type == common.SELF_HELP {
			//如果后台管理人员，通过任意下发，不涉及到商户减款操作，直接发送代付请求
			solveSelf(p)
		} else {
			SendPayFor(p)
		}
	}
}

func solveSelf(payFor models.PayforInfo) {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if err := recover(); err != nil {
			o.Rollback()
		}
	}(o)

	var tmpPayFor models.PayforInfo
	if err := o.Raw("select * from payfor_info where payfor_uid = ? for update", payFor.PayforUid).QueryRow(&tmpPayFor); err != nil || tmpPayFor.PayforUid == "" {
		logs.Error("solve self payfor fail：", err)
		o.Rollback()
		return
	}

	if tmpPayFor.IsSend == common.YES {
		o.Rollback()
		return
	}

	tmpPayFor.UpdateTime = utils.GetBasicDateTime()
	if payFor.RoadUid == "" {
		tmpPayFor.Status = common.PAYFOR_FAIL
	} else {
		tmpPayFor.Status = common.PAYFOR_BANKING
		tmpPayFor.RequestTime = utils.GetBasicDateTime()
		tmpPayFor.IsSend = common.YES
	}
	_, err := o.Update(&tmpPayFor)
	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
		RequestPayFor(payFor)
	}
}

func SendPayFor(payFor models.PayforInfo) bool {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if err := recover(); err != nil {
			o.Rollback()
			return
		}
	}(o)

	var tmpPayFor models.PayforInfo
	if err := o.Raw("select * from payfor_info where payfor_uid = ? for update", payFor.PayforUid).QueryRow(&tmpPayFor); err != nil || tmpPayFor.PayforUid == "" {
		logs.Error("send payfor select payfor fail: ", err)
		o.Rollback()
		return false
	}

	var account models.AccountInfo
	if err := o.Raw("select * from account_info where account_uid = ? for update", payFor.MerchantUid).QueryRow(&account); err != nil || account.AccountUid == "" {
		logs.Error("send payfor select account fail：", err)
		o.Rollback()
		return false
	}

	//支付金额不足，将直接判定为失败，不往下面邹逻辑了
	if account.SettleAmount-account.PayforAmount < tmpPayFor.PayforAmount+tmpPayFor.PayforFee {
		tmpPayFor.Status = common.PAYFOR_FAIL
		tmpPayFor.UpdateTime = utils.GetBasicDateTime()
		o.Update(&tmpPayFor)
		o.Commit()
		return false
	}

	account.UpdateTime = utils.GetBasicDateTime()
	account.PayforAmount = account.PayforAmount + payFor.PayforAmount + payFor.PayforFee

	_, err := o.Update(&account)
	if err != nil {
		logs.Error(fmt.Sprintf("商户uid=%s，在发送代付给上游的处理中，更新账户表出错, err: %s", payFor.MerchantUid, err))
		o.Rollback()
		return false
	}

	tmpPayFor.IsSend = common.YES
	tmpPayFor.Status = common.PAYFOR_BANKING //变为银行处理中
	tmpPayFor.RequestTime = utils.GetBasicDateTime()
	tmpPayFor.UpdateTime = utils.GetBasicDateTime()

	_, err = o.Update(&tmpPayFor)
	if err != nil {
		logs.Error(fmt.Sprintf("商户uid=%s，在发送代付给上游的处理中，更代付列表出错， err：%s", payFor.MerchantUid, err))
		o.Rollback()
		return false
	} else {
		o.Commit()
		RequestPayFor(payFor)
	}
	return true
}

func RequestPayFor(payFor models.PayforInfo) {
	if payFor.RoadUid == "" {
		return
	}
	roadInfo := models.GetRoadInfoByRoadUid(payFor.RoadUid)
	supplierCode := roadInfo.ProductUid
	supplier := GetPaySupplierByCode(supplierCode)
	res := supplier.PayFor(payFor)
	logs.Info(fmt.Sprintf("代付uid=%s，上游处理结果为：%s", payFor.PayforUid, res))
	//将代付订单号发送到消息队列
	message_queue.SendMessage(common.MQ_PAYFOR_QUERY, payFor.BankOrderId)
}
