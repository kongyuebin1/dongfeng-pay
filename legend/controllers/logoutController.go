package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/server/web"
)

type LogoutController struct {
	web.Controller
}

func (c *LogoutController) Logout() {
	if err := c.DelSession("username"); err != nil {
		logs.Error("用户退出登录出错，错误信息：", err)
	}
	c.Redirect("/login.html", 302)
}

/**
** 切换用户登录
 */
func (c *LogoutController) SwitchLogin() {
	err := c.DelSession("username")
	if err != nil {
		logs.Error("切换账号失败，错误信息：", err)
	}

	c.Redirect("/login.html", 302)
}
