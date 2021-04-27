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
	"agent/models"
	"agent/sys/enum"
	"regexp"
	"strings"
)

type UserInfo struct {
	KeepSession
}

func (c *UserInfo) ShowModifyUserInfoUI() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	c.Data["userName"] = u.AgentName
	c.TplName = "modify_userInfo.html"
}

// 修改用户信息
func (c *UserInfo) ModifyUserInfo() {
	or_pwd := strings.TrimSpace(c.GetString("or_pwd"))
	new_pwd := strings.TrimSpace(c.GetString("new_pwd"))
	confirm_pwd := strings.TrimSpace(c.GetString("confirm_pwd"))

	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

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
	if strings.Compare(strings.ToUpper(pwdMd5), u.AgentPassword) != 0 {
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

	u.AgentPassword = strings.ToUpper(encrypt.EncodeMd5([]byte(new_pwd)))
	u.UpdateTime = pubMethod.GetNowTime()
	ud = models.UpdateAgentInfo(u)
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
func (c *UserInfo) ConfirmOriginPwd() {
	ori := strings.TrimSpace(c.GetString("c"))

	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	var (
		msg  = enum.FailedString
		flag = enum.FailedFlag
	)

	pwdMd5 := encrypt.EncodeMd5([]byte(ori))
	if strings.Compare(strings.ToUpper(pwdMd5), u.AgentPassword) != 0 {
		msg = "原始密码错误!"
	} else {
		flag = enum.SuccessFlag
	}

	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, "")
	c.ServeJSON()
	c.StopRun()
}

// 设置支付密码
func (c *UserInfo) SetPayPassword() {
	or_pwd := strings.TrimSpace(c.GetString("or_pwd"))
	new_pwd := strings.TrimSpace(c.GetString("new_pwd"))
	confirm_pwd := strings.TrimSpace(c.GetString("confirm_pwd"))

	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	var (
		msg  = enum.FailedString
		flag = enum.FailedFlag

		md     bool
		ud     bool
		pwdMd5 string
	)

	if new_pwd == "" || confirm_pwd == "" {
		msg = "密码不能为空!"
		goto stopRun
	}

	if u.PayPassword != "" {
		pwdMd5 = encrypt.EncodeMd5([]byte(or_pwd))
		if strings.Compare(strings.ToUpper(pwdMd5), u.AgentPassword) != 0 {
			msg = "原始密码错误!"
		}
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

	u.PayPassword = strings.ToUpper(encrypt.EncodeMd5([]byte(new_pwd)))
	u.UpdateTime = pubMethod.GetNowTime()
	ud = models.UpdateAgentInfo(u)
	if ud {
		msg = enum.SuccessString
		flag = enum.SuccessFlag

		// 重新写入缓存信息
		c.SetSession(enum.UserSession, u)
	}

stopRun:
	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, "")
	c.ServeJSON()
	c.StopRun()
}

// 验证原始支付密码
func (c *UserInfo) ConfirmOriginPayPwd() {
	ori := strings.TrimSpace(c.GetString("c"))

	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	var (
		msg  = enum.FailedString
		flag = enum.SuccessFlag
	)

	if u.PayPassword != "" {
		pwdMd5 := encrypt.EncodeMd5([]byte(ori))
		if strings.Compare(strings.ToUpper(pwdMd5), u.AgentPassword) != 0 {
			msg = "原始密码错误!"
			flag = enum.FailedFlag
		}
	}

	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, "")
	c.ServeJSON()
	c.StopRun()
}

// 展示用户信息
func (c *UserInfo) ShowUserInfoUI() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	c.Data["userName"] = u.AgentName
	c.TplName = "show_userInfo.html"
}

// 代理商列表
func (c *UserInfo) ShowMerchantUI() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	c.Data["payType"] = enum.GetPayType()
	c.Data["status"] = enum.GetOrderStatus()
	c.Data["userName"] = u.AgentName
	c.TplName = "merchant.html"
}

func (c *UserInfo) MerchantQueryAndListPage() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	mt := make(map[string]string)
	mt["belong_agent_uid"] = u.AgentUid
	mt["status"] = enum.ACTIVE

	// 该代理商下所有正常商户
	merchants := models.GetMerchantByParams(mt, -1, 0)

	type account struct {
		UId          string
		MerchantName string  // 商户名
		Mobile       string  // 手机号
		Balance      float64 // 账户余额
		SettleAmount float64 // 已结算金额
		WaitAmount   float64 // 待结算金额
		LoanAmount   float64 // 押款金额
		FreezeAmount float64 // 账户冻结金额
		PayAmount    float64 // 代付中金额
	}

	type deploy struct {
		ChannelName string  // 通道名
		PlatRate    float64 // 平台费率
		AgentRate   float64 // 代理商费率
	}

	mtd := make(map[string]string)
	mtd["status"] = enum.ACTIVE
	var (
		count    = 0       // 计算该代理商下的商户有多少个通道
		accounts []account // 每个商户的账户信息
	)
	for _, m := range merchants {
		mtd["merchant_uid"] = m.MerchantUid
		lens := models.GetMerchantDeployLenByMap(mtd)
		count += lens

		ac := models.GetAccountByUid(m.MerchantUid)
		accounts = append(accounts, account{
			UId:          m.MerchantUid,
			MerchantName: m.MerchantName,
			Mobile:       m.LoginAccount,
			Balance:      ac.Balance,
			SettleAmount: ac.SettleAmount,
			WaitAmount:   ac.WaitAmount,
			LoanAmount:   ac.LoanAmount,
			FreezeAmount: ac.FreezeAmount,
			PayAmount:    ac.PayforAmount,
		})
	}

	// 每个商户的通道信息
	deploys := make(map[string][]deploy)
	if count != 0 {
		for _, a := range accounts {
			mtd["merchant_uid"] = a.UId

			mdl := models.GetMerchantDeployListByMap(mtd, -1, 0)
			for _, m := range mdl {
				road := models.GetRoadInfoByRoadUid(m.SingleRoadUid)
				fee := road.BasicFee + m.SingleRoadPlatformRate + m.SingleRoadAgentRate
				if fee < 0.00000001 {
					fee = road.BasicFee + m.RollRoadPlatformRate + m.RollRoadAgentRate
				}
				afee := m.SingleRoadAgentRate
				if afee < 0.00000001 {
					afee = m.RollRoadAgentRate
				}
				deploys[a.UId] = append(deploys[a.UId], deploy{
					ChannelName: road.ProductName,
					PlatRate:    fee,
					AgentRate:   afee,
				})
			}
		}
	}

	// 数据回显
	out := make(map[string]interface{})
	out["ac"] = accounts
	out["dp"] = deploys
	out["count"] = count

	c.Data["json"] = out
	c.ServeJSON()
	c.StopRun()
}
