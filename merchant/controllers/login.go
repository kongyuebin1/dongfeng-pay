/***************************************************
 ** @Desc : This file for 用户登录
 ** @Time : 19.11.29 13:52
 ** @Author : Joker
 ** @File : login
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.11.29 13:52
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dchest/captcha"
	"merchant/models"
	"merchant/sys"
	"merchant/sys/enum"
	"merchant/utils"
	"strconv"
	"strings"
)

var pubMethod = sys.PublicMethod{}
var encrypt = utils.Encrypt{}

type Login struct {
	beego.Controller
}

func (c *Login) UserLogin() {
	captchaCode := c.GetString("captchaCode")
	captchaId := c.GetString("captchaId")
	userName := strings.TrimSpace(c.GetString("userName"))
	password := c.GetString("Password")

	var (
		flag = enum.FailedFlag
		msg  = ""
		url  = "/"

		pwdMd5 string
		ran    string
		ranMd5 string

		verify bool
		u      models.MerchantInfo
	)

	us := c.GetSession(enum.UserSession)
	if us != nil {
		url = enum.DoMainUrl
		flag = enum.SuccessFlag
		goto stopRun
	}

	if userName == "" || password == "" {
		msg = "登录账号或密码不能为空!"
		goto stopRun
	}

	verify = captcha.VerifyString(captchaId, captchaCode)
	if !verify {
		url = strconv.Itoa(enum.FailedFlag)
		msg = "验证码不正确!"
		goto stopRun
	}

	u = models.GetMerchantByPhone(userName)
	if u.LoginPassword == "" {
		msg = "账户信息错误，请联系管理人员!"
		goto stopRun
	}

	if strings.Compare(enum.ACTIVE, u.Status) != 0 {
		msg = "登录账号或密码错误!"
		goto stopRun
	}

	//验证密码
	pwdMd5 = encrypt.EncodeMd5([]byte(password))
	if strings.Compare(strings.ToUpper(pwdMd5), u.LoginPassword) != 0 {
		msg = "登录账号或密码错误!"
		goto stopRun
	}

	c.SetSession(enum.UserSession, u)

	// 设置客户端用户信息有效保存时间
	ran = pubMethod.RandomString(46)
	ranMd5 = encrypt.EncodeMd5([]byte(ran))
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	url = enum.DoMainUrl
	flag = enum.SuccessFlag
	utils.LogNotice(fmt.Sprintf("【%s】用户登录成功，请求IP：%s", u.MerchantName, c.Ctx.Input.IP()))

stopRun:
	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, url)
	c.ServeJSON()
	c.StopRun()
}

func (c *Login) Index() {
	capt := struct {
		CaptchaId string
	}{
		captcha.NewLen(4),
	}
	c.Data["CaptchaId"] = capt.CaptchaId
	c.TplName = "login.html"
}

// 验证输入的验证码
func (c *Login) VerifyCaptcha() {
	captchaValue := c.Ctx.Input.Param(":value")
	captchaId := c.Ctx.Input.Param(":chaId")

	verify := captcha.VerifyString(captchaId, captchaValue)
	if verify {
		c.Data["json"] = pubMethod.JsonFormat(enum.SuccessFlag, "", "", "")
	} else {
		c.Data["json"] = pubMethod.JsonFormat(enum.FailedFlag, "", "验证码不匹配!", "")
	}
	c.ServeJSON()
	c.StopRun()
}

// 重绘验证码
func (c *Login) FlushCaptcha() {
	capt := struct {
		CaptchaId string
	}{
		captcha.NewLen(4),
	}
	c.Data["json"] = pubMethod.JsonFormat(enum.SuccessFlag, capt.CaptchaId, "验证码不匹配!", "")
	c.ServeJSON()
	c.StopRun()
}

// 退出登录
func (c *Login) LoginOut() {
	c.DelSession(enum.UserSession)

	c.Data["json"] = pubMethod.JsonFormat(enum.FailedFlag, "", "", "/")
	c.ServeJSON()
	c.StopRun()
}

// 对接文档
func (c *Login) PayDoc() {
	c.TplName = "api_doc/pay_doc.html"
}
