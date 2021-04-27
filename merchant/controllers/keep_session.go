/***************************************************
 ** @Desc : This file for 保持会话
 ** @Time : 19.11.29 13:55
 ** @Author : Joker
 ** @File : keep_session
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.11.29 13:55
 ** @Software: GoLand
****************************************************/
package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"merchant/sys/enum"
)

type KeepSession struct {
	beego.Controller
}

// 生成随机md5值，客户端和服务端各保存一份
// 客户端每次请求都将重写md5值
// 若用户在30分钟内没有操作行为或连续操作时间超过3小时，则自动退出
func (c *KeepSession) Prepare() {
	// 以免session值不是string类型而panic
	defer func() {
		if err := recover(); err != nil {
			c.DelSession(enum.UserSession)
			c.Ctx.Redirect(302, "/")
		}
	}()

	us := c.GetSession(enum.UserSession)
	uc := c.GetSession(enum.UserCookie)
	if us == nil || uc == nil {
		c.DelSession(enum.UserSession)
		c.Ctx.Redirect(302, "/")
	}

	if uc.(string) == "" {
		c.DelSession(enum.UserSession)
		c.Ctx.Redirect(302, "/")
	}

	_, b := c.Ctx.GetSecureCookie(uc.(string), enum.UserCookie)
	//utils.LogNotice(fmt.Sprintf("客户端cookie：%s，服务端cookie：%s", cookie, uc.(string)))
	if !b {
		c.DelSession(enum.UserSession)
		c.Ctx.Redirect(302, "/")
	}
}
