package controllers

import (
	"boss/datas"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) GenerateJSON(dataJSON interface{}) {
	c.Data["json"] = dataJSON
	_ = c.ServeJSON()
}

func (c *BaseController) Prepare() {
	userID, ok := c.GetSession("userID").(string)
	if !ok || userID == "" {
		//用户没有登录，或者登录到期了，则跳转登录主页面
		dataJSON := new(datas.BaseDataJSON)
		dataJSON.Code = 404
		dataJSON.Msg = "登录已经过期!"
		c.Data["json"] = dataJSON
		_ = c.ServeJSON()
	} else {
		//重新赋值给session
		_ = c.SetSession("userID", userID)
	}
}
