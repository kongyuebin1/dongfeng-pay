/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/12/8 22:15
 ** @Author : yuebin
 ** @File : send_notify_merchant
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/12/8 22:15
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"dongfeng-pay/service/common"
	"dongfeng-pay/service/models"
	"strings"
)

type SendNotify struct {
	beego.Controller
}

func (c *SendNotify) SendNotifyToMerchant() {
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = -1
	orderInfo := models.GetOrderByBankOrderId(bankOrderId)
	if orderInfo.Status == common.WAIT {
		keyDataJSON.Msg = "该订单不是成功状态，不能回调"
	} else {
		notifyInfo := models.GetNotifyInfoByBankOrderId(bankOrderId)
		notifyUrl := notifyInfo.Url
		logs.Info(fmt.Sprintf("boss管理后台手动触发订单回调，url=%s", notifyUrl))
		req := httplib.Post(notifyUrl)
		response, err := req.String()
		if err != nil {
			logs.Error("回调发送失败，fail：", err)
			keyDataJSON.Msg = fmt.Sprintf("该订单回调发送失败，订单回调，fail：%s", err)
		} else {
			if !strings.Contains(strings.ToLower(response), "success") {
				keyDataJSON.Msg = fmt.Sprintf("该订单回调发送成功，但是未返回success字段， 商户返回内容=%s", response)
			} else {
				keyDataJSON.Code = 200
				keyDataJSON.Msg = fmt.Sprintf("该订单回调发送成功")
			}
		}
	}
	c.Data["json"] = keyDataJSON
	c.ServeJSON()
}

func (c *SendNotify) SelfSendNotify() {
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))

	notifyInfo := models.GetNotifyInfoByBankOrderId(bankOrderId)

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = 200

	req := httplib.Post(notifyInfo.Url)

	response, err := req.String()
	if err != nil {
		keyDataJSON.Msg = fmt.Sprintf("订单 bankOrderId=%s，已经发送回调出错：%s", bankOrderId, err)
	} else {
		keyDataJSON.Msg = fmt.Sprintf("订单 bankOrderId=%s，已经发送回调，商户返回内容：%s", bankOrderId, response)
	}

	c.Data["json"] = keyDataJSON
	c.ServeJSON()
}
