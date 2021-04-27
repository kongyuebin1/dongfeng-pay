/***************************************************
 ** @Desc : This file for 状态常量
 ** @Time : 19.11.30 11:12
 ** @Author : Joker
 ** @File : status
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.11.30 11:12
 ** @Software: GoLand
****************************************************/
package enum

// 成功与否
const (
	FailedFlag    = -9
	FailedString  = "操作失败! "
	FailedToAdmin = "系统内部错误，请联系管理员!"
	SuccessFlag   = 9
	SuccessString = "操作成功! "
)

// 用户状态
const (
	ACTIVE   = "active"
	FREEZE   = "FREEZE"
	UNACTIVE = "unactive"
)

var userStatus = map[string]string{
	ACTIVE:   "激活",
	FREEZE:   "冻结",
	UNACTIVE: "冻结",
}

// 用户状态
func GetUserStatus() map[string]string {
	return userStatus
}

// 充值订单状态
const (
	SUCCESS         = "success"
	FAILED          = "failed"
	WAITING_PAYMENT = "wait"
)

var orderStatus = map[string]string{
	SUCCESS:         "交易成功",
	FAILED:          "交易失败",
	WAITING_PAYMENT: "等待支付",
}

// 充值订单状态
func GetOrderStatus() map[string]string {
	return orderStatus
}

// 投诉订单状态
const (
	YES = "yes"
	NO  = "no"
)

var orderComStatus = map[string]string{
	YES: "冻结",
	NO:  "未冻结",
}

// 投诉订单状态
func GetComOrderStatus() map[string]string {
	return orderComStatus
}

// 结算订单状态
const (
	WAIT_CONFIRM  = "payfor_confirm"
	REMITTING     = "payfor_solving"
	REMIT_FAIL    = "failed"
	BANK_DEALING  = "payfor_banking"
	REMIT_SUCCESS = "success"
)

var settlementStatus = map[string]string{
	WAIT_CONFIRM:  "等待审核",
	REMITTING:     "打款中",
	REMIT_FAIL:    "打款失败",
	BANK_DEALING:  "银行处理中",
	REMIT_SUCCESS: "打款成功",
}

// 结算订单状态
func GetSettlementStatus() map[string]string {
	return settlementStatus
}

// 充值订单状态
const (
	RECHARGE = "recharge"
	REFUND   = "refund"
	FREEZER  = "freeze"
	UNFREEZE = "unfreeze"
)

var rechargeStatus = map[string]string{
	RECHARGE: "充值",
	REFUND:   "退款",
	FREEZER:  "冻结",
	UNFREEZE: "解冻",
}

// 充值订单状态
func GetRechargeStatus() map[string]string {
	return rechargeStatus
}

// 历史记录状态
const (
	PLUS_AMOUNT     = "plus_amount"
	SUB_AMOUNT      = "sub_amount"
	FREEZE_AMOUNT   = "freeze_amount"
	UNFREEZE_AMOUNT = "unfreeze_amount"
)

var historyStatus = map[string]string{
	PLUS_AMOUNT:     "加款",
	SUB_AMOUNT:      "减款",
	FREEZE_AMOUNT:   "冻结",
	UNFREEZE_AMOUNT: "解冻",
}

// 历史记录状态
func GetHistoryStatus() map[string]string {
	return historyStatus
}
