/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/26 16:56
 ** @Author : yuebin
 ** @File : error_gateway
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/26 16:56
 ** @Software: GoLand
****************************************************/
package gateway

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ErrorGatewayController struct {
	beego.Controller
}

func (c *ErrorGatewayController) ErrorParams() {
	beego.ReadFromRequest(&c.Controller)
	c.TplName = "err/params.html"
}
