/***************************************************
 ** @Desc : This file for 配置常量
 ** @Time : 2019.04.01 11:45
 ** @Author : Joker
 ** @File : strings
 ** @Last Modified by : Joker
 ** @Last Modified time: 2019-11-29 11:05:48
 ** @Software: GoLand
****************************************************/
package enum

// 对接云片
// 短信配置
const (
	ApiKey     = "fd264ab6c43c02c52s40eab1ba"
	TPL1       = 332236
	SendSmsUrl = "https://sms.yunpian.com/v2/sms/tpl_single_send.json"
)

// session 配置
const (
	SessionPath         = "./sys/temp"   // 保存路径
	SessionExpireTime   = 9600           // 有效时间,秒
	CookieExpireTime    = 1800           // 有效时间,秒
	SmsCookieExpireTime = 60             // 有效时间,秒
	LocalSessionName    = "JOKERSession" // 客户端session名称
)

// 提现限制金额
const (
	WithdrawalMaxAmount = 45000
	WithdrawalMinAmount = 2
	SettlementFee       = 2 // 提现单笔手续费
)
