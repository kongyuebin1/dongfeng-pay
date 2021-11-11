/***************************************************
 ** @Desc : This file for 账户变动
 ** @Time : 19.12.10 10:42
 ** @Author : Joker
 ** @File : account_history
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.10 10:42
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"merchant/models"
	"merchant/sys/enum"
	"strconv"
	"strings"
)

type History struct {
	KeepSession
}

// 账户资产变动列表
func (c *History) ShowHistoryListUI() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	_ = c.SetSession(enum.UserCookie, ranMd5)

	c.Data["payType"] = enum.GetHistoryStatus()
	c.Data["userName"] = u.MerchantName
	c.TplName = "history_record.html"
}

func (c *History) HistoryQueryAndListPage() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	// 分页参数
	page, _ := strconv.Atoi(c.GetString("page"))
	limit, _ := strconv.Atoi(c.GetString("limit"))
	if limit == 0 {
		limit = 15
	}

	// 查询参数
	in := make(map[string]string)
	start := strings.TrimSpace(c.GetString("start"))
	end := strings.TrimSpace(c.GetString("end"))
	status := strings.TrimSpace(c.GetString("status"))

	in["type"] = status
	in["account_uid"] = u.MerchantUid

	if start != "" {
		in["create_time__gte"] = start
	}
	if end != "" {
		in["create_time__lte"] = end
	}

	// 计算分页数
	count := models.GetAccountHistoryLenByMap(in)
	totalPage := count / limit // 计算总页数
	if count%limit != 0 {      // 不满一页的数据按一页计算
		totalPage++
	}

	// 数据获取
	var list []models.AccountHistoryInfo
	if page <= totalPage {
		list = models.GetAccountHistoryByMap(in, limit, (page-1)*limit)
	}

	// 数据回显
	out := make(map[string]interface{})
	out["limit"] = limit // 分页数据
	out["page"] = page
	out["totalPage"] = totalPage
	out["root"] = list // 显示数据

	c.Data["json"] = out
	c.ServeJSON()
	c.StopRun()
}
