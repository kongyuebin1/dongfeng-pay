package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"legend/service"
)

type LoginController struct {
	web.Controller
}

/**
**展示登录页面
 */
func (c *LoginController) LoginPage() {
	c.TplName = "login.html"
}

/**
** 处理登录逻辑
 */
func (c *LoginController) Login() {

	userName := c.GetString("username")
	password := c.GetString("password")

	logs.Info("username："+userName, ";password: "+password)

	loginService := new(service.LoginService)

	loginJsonData := loginService.Login(userName, password)
	if loginJsonData.Code == 200 {
		_ = c.SetSession("userName", userName)
	}

	c.Data["json"] = loginJsonData

	err := c.ServeJSON()
	if err != nil {
		logs.Error("错误：", err)
	}
}

/**
** 更新登录密码
 */
func (c *LoginController) PersonPassword() {
	oldPassword := c.GetString("oldpass")
	newPassword := c.GetString("newpass")
	repeatPassword := c.GetString("repass")

	logs.Debug("用户跟换密码，旧密码：%s, 新密码：%s，确认密码：%s", oldPassword, newPassword, repeatPassword)

	userNname := c.GetSession("userName").(string)

	loginService := new(service.LoginService)
	loginJsonData := loginService.PersonPassword(newPassword, oldPassword, repeatPassword, userNname)

	c.Data["json"] = loginJsonData
	_ = c.ServeJSON()
}
