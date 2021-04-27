/***************************************************
 ** @Desc : This file for ...收银台对接快一
 ** @Time : 2018-8-27 13:50
 ** @Author : Joker
 ** @File : home_action
 ** @Last Modified by : Joker
 ** @Last Modified time: 2018-08-29 17:59:48
 ** @Software: GoLand
****************************************************/
package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/rs/xid"
)

type HomeAction struct {
	beego.Controller
}

/*加载首页及数据*/
func (c *HomeAction) ShowHome() {
	//取值
	siteName, _ := beego.AppConfig.String("site.name")
	orderNo := xid.New().String()
	productName := "测试应用-支付功能体验(非商品消费)"

	//数据回显
	c.Data["siteName"] = siteName
	c.Data["pname"] = productName
	c.Data["orderNo"] = orderNo
	c.TplName = "index.html"
}

func (c *HomeAction) ErrorPage() {
	flash := beego.ReadFromRequest(&c.Controller)
	error := flash.Data["error"]
	c.Data["error"] = error
	c.TplName = "error.html"
}
