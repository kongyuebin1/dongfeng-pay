/***************************************************
 ** @Desc : This file for 用户信息控制
 ** @Time : 19.12.3 10:38
 ** @Author : Joker
 ** @File : user_info
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.3 10:38
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"juhe/jhmerchant/sys/enum"
	"juhe/service/models"
	"regexp"
	"strings"
)

type UserInfo struct {
	KeepSession
}

// @router /user_info/show_modify_ui
func (c *UserInfo) ShowModifyUserInfoUI() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	c.Data["userName"] = u.MerchantName
	c.TplName = "modify_userInfo.html"
}

// 修改用户信息
// @router /user_info/modify_userInfo/?:params [post]
func (c *UserInfo) ModifyUserInfo() {
	or_pwd := strings.TrimSpace(c.GetString("or_pwd"))
	new_pwd := strings.TrimSpace(c.GetString("new_pwd"))
	confirm_pwd := strings.TrimSpace(c.GetString("confirm_pwd"))

	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	var (
		msg  = enum.FailedString
		flag = enum.FailedFlag

		md     bool
		ud     bool
		pwdMd5 string
	)

	if or_pwd == "" ||
		new_pwd == "" ||
		confirm_pwd == "" {
		msg = "密码不能为空!"
		goto stopRun
	}

	pwdMd5 = encrypt.EncodeMd5([]byte(or_pwd))
	if strings.Compare(strings.ToUpper(pwdMd5), u.LoginPassword) != 0 {
		msg = "原始密码错误!"
	}

	md, _ = regexp.MatchString(enum.PasswordReg, new_pwd)
	if !md {
		msg = "密码只能输入6-20个以字母开头、可带数字、“_”、“.”的字串!"
		goto stopRun
	}

	md, _ = regexp.MatchString(enum.PasswordReg, confirm_pwd)
	if !md {
		msg = "密码只能输入6-20个以字母开头、可带数字、“_”、“.”的字串!"
		goto stopRun
	}

	if strings.Compare(new_pwd, confirm_pwd) != 0 {
		msg = "两次密码不匹配!"
		goto stopRun
	}

	u.LoginPassword = strings.ToUpper(encrypt.EncodeMd5([]byte(new_pwd)))
	ud = models.UpdateMerchant(u)
	if ud {
		msg = enum.SuccessString
		flag = enum.SuccessFlag

		// 退出重新登录
		c.DelSession(enum.UserSession)
	}

stopRun:
	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, "")
	c.ServeJSON()
	c.StopRun()
}

// 验证原始密码
// @router /user_info/confirm_pwd/?:params [post]
func (c *UserInfo) ConfirmOriginPwd() {
	ori := strings.TrimSpace(c.GetString("c"))

	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	var (
		msg  = enum.FailedString
		flag = enum.FailedFlag
	)

	pwdMd5 := encrypt.EncodeMd5([]byte(ori))
	if strings.Compare(strings.ToUpper(pwdMd5), u.LoginPassword) != 0 {
		msg = "原始密码错误!"
	} else {
		flag = enum.SuccessFlag
	}

	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, "")
	c.ServeJSON()
	c.StopRun()
}

// 展示用户信息
// @router /user_info/show_ui
func (c *UserInfo) ShowUserInfoUI() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	c.Data["userName"] = u.MerchantName
	c.Data["mobile"] = u.LoginAccount
	c.Data["email"] = u.LoginAccount
	c.Data["riskDay"] = "1"
	//c.Data["key"] = uPayConfig.PayKey
	//c.Data["secret"] = uPayConfig.PaySecret
	c.TplName = "show_userInfo.html"
}
