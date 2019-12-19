/***************************************************
 ** @Desc : c file for ...
 ** @Time : 2019/8/13 18:09
 ** @Author : yuebin
 ** @File : base_controller
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/13 18:09
 ** @Software: GoLand
****************************************************/
package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) GenerateJSON(dataJSON interface{}) {
	c.Data["json"] = dataJSON
	c.ServeJSON()
}

func (c *BaseController) Prepare() {
	userID, ok := c.GetSession("userID").(string)
	if !ok || userID == "" {
		//用户没有登录，或者登录到期了，则跳转登录主页面
		dataJSON := new(BaseDataJSON)
		dataJSON.Code = 404
		dataJSON.Msg = "登录已经过期!"
		c.Data["json"] = dataJSON
		c.ServeJSON()
	} else {
		//重新赋值给session
		c.SetSession("userID", userID)
	}
}
