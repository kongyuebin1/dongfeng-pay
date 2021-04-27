/***************************************************
 ** @Desc : This file for 发送短信验证码
 ** @Time : 2019.04.04 9:37
 ** @Author : Joker
 ** @File : send_messages
 ** @Last Modified by : Joker
 ** @Last Modified time: 2019-11-29 11:05:41
 ** @Software: GoLand
****************************************************/
package utils

import (
	"fmt"
	"merchant/sys/enum"
	"net/http"
	"net/url"
)

// 发送提现通知
func SendSmsForPay(mobile, code string) bool {
	tplValue := url.Values{"#code#": {code}}.Encode()
	dataTplSms := url.Values{
		"apikey":    {enum.ApiKey},
		"mobile":    {mobile},
		"tpl_id":    {fmt.Sprintf("%d", enum.TPL1)},
		"tpl_value": {tplValue}}
	_, err := http.PostForm(enum.SendSmsUrl, dataTplSms)
	if err != nil {
		LogError("sms send fail,err:", err)
		return false
	}
	LogInfo(fmt.Sprintf("sms send to %s is success", mobile))
	return true
}
