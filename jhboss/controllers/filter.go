/***************************************************
 ** @Desc : 过滤功能
 ** @Time : 2019/8/8 16:10
 ** @Author : yuebin
 ** @File : filter
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/8 16:10
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

type FilterController struct {
	beego.Controller
}

var FilterLogin = func(ctx *context.Context) {
	userID, ok := ctx.Input.Session("userID").(string)
	if !ok || userID == "" {
		if !strings.Contains(ctx.Request.RequestURI, "/login.html") &&
			!strings.Contains(ctx.Request.RequestURI, "/getVerifyImg") &&
			!strings.Contains(ctx.Request.RequestURI, "/favicon.ico") &&
			!ctx.Input.IsAjax() {
			ctx.Redirect(302, "/login.html")
		}
	} else {
		if strings.Contains(ctx.Request.RequestURI, "/login.html") {
			ctx.Redirect(302, "/")
		}
	}
}

//jsonp请求过来的函数
func (c *FilterController) Filter() {
	userID, ok := c.GetSession("userID").(string)

	dataJSON := new(struct {
		Code int
	})

	if !ok || userID == "" {
		dataJSON.Code = 404
	} else {
		dataJSON.Code = 200
		c.SetSession("userID", userID)
	}
	fmt.Println(dataJSON)
	c.Data["json"] = dataJSON
	c.ServeJSON()
}
