/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/25 14:14
 ** @Author : yuebin
 ** @File : consts.go
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/25 14:14
 ** @Software: GoLand
****************************************************/
package common

const (
	ACTIVE          = "active"
	UNACTIVE        = "unactive"
	DELETE          = "delete"
	REFUND          = "refund"
	ORDERROLL       = "order_roll"
	WAIT            = "wait"
	SUCCESS         = "success"
	FAIL            = "fail"
	YES             = "yes"
	NO              = "no"
	ZERO            = 0.0  //0元手续费
	VERIFY_CODE_LEN = 4    //验证码的长度
	PAYFOR_FEE      = 2.00 //代付手续费
	PAYFOR_INTERVAL = 5    //每过5分钟执行一次代付

	PLUS_AMOUNT     = "plus_amount"     //加款操作
	SUB_AMOUNT      = "sub_amount"      //减款操作
	FREEZE_AMOUNT   = "freeze_amount"   //冻结操作
	UNFREEZE_AMOUNT = "unfreeze_amount" //解冻操作

	PAYFOR_COMFRIM = "payfor_confirm" //下发带审核
	PAYFOR_SOLVING = "payfor_solving" //发下处理中
	PAYFOR_HANDING = "payfor_handing" //手动打款中
	PAYFOR_BANKING = "payfor_banking" //银行处理中
	PAYFOR_FAIL    = "payfor_fail"    //代付失败
	PAYFOR_SUCCESS = "payfor_success" //代付成功

	PAYFOR_ROAD   = "payfor_road"   //通道打款
	PAYFOR_HAND   = "payfor_hand"   //手动打款
	PAYFOR_REFUSE = "payfor_refuse" // 拒绝打款

	SELF_API      = "self_api"      //自助api系统下发
	SELF_MERCHANT = "self_merchant" //管理手动处理商户下发
	SELF_HELP     = "self_help"     //管理自己提现

	PUBLIC  = "public"  //对公卡
	PRIVATE = "private" //对私卡
)
