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
	"boss/service"
	"github.com/beego/beego/v2/server/web"
	"strings"
)

type SendNotify struct {
	web.Controller
}

func (c *SendNotify) SendNotifyToMerchant() {
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))
	se := new(service.SendNotifyMerchantService)
	keyDataJSON := se.SendNotifyToMerchant(bankOrderId)

	c.Data["json"] = keyDataJSON
	_ = c.ServeJSON()
}

func (c *SendNotify) SelfSendNotify() {
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))

	se := new(service.SendNotifyMerchantService)
	keyDataJSON := se.SelfSendNotify(bankOrderId)
	c.Data["json"] = keyDataJSON
	_ = c.ServeJSON()
}
