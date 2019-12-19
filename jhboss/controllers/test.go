/***************************************************
 ** @Desc : c file for ...
 ** @Time : 2019/9/20 14:38
 ** @Author : yuebin
 ** @File : test
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/9/20 14:38
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"github.com/astaxie/beego/httplib"
)

const HOST = "https://gw.open.icbc.com.cn/ui/b2c/pay/transfer/V2"

func (c *BaseController) Test() {
	//sign := c.GetString("sign")
	msg_id := c.GetString("msg_id")
	app_id := c.GetString("app_id")
	sign_type := c.GetString("sign_type")
	timestamp := c.GetString("timestamp")
	//host := HOST + "&sign=" + sign + "&msg_id=" + msg_id + "&app_id=" + app_id + "&sign_type=" + sign_type + "&timestamp=" + timestamp
	biz_content := c.GetString("biz_content")
	clientType := c.GetString("clientType")
	interfaceVersion := c.GetString("interfaceVersion")
	interfaceName := c.GetString("interfaceName")
	notify_url := c.GetString("notify_url")
	ca := c.GetString("ca")
	req := httplib.Post(HOST)
	//req.Header("Content‐Type", "application/x‐www‐form‐urlencoded")
	req.Header("charset", "GBK")
	req.Param("charset", "UTF-8")
	req.Param("format", "json")
	req.Param("sign", "ERERERERERERERE")
	req.Param("msg_id", msg_id)
	req.Param("app_id", app_id)
	req.Param("sign_type", sign_type)
	req.Param("timestamp", timestamp)
	req.Param("biz_content", biz_content)
	req.Param("clientType", clientType)
	req.Param("interfaceVersion", interfaceVersion)
	req.Param("interfaceName", interfaceName)
	req.Param("notify_url", notify_url)
	req.Param("ca", ca)

	res, _ := req.String()
	c.Ctx.WriteString(res)
	c.ServeJSON()
}
