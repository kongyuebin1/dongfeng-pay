/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/12/18 17:16
 ** @Author : yuebin
 ** @File : pay
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/12/18 17:16
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"strconv"
	"strings"
)

type PayController struct {
	web.Controller
}

func (c *PayController) Pay() {
	orderNo := strings.TrimSpace(c.GetString("orderid"))
	flash := web.NewFlash()
	if orderNo == "" {
		flash.Error("订单号为空")
		flash.Store(&c.Controller)
		c.Redirect("/error.html", 302)
		return
	}
	amount := strings.TrimSpace(c.GetString("amount"))
	if !c.judgeAmount(amount) {
		flash.Error("金额有误")
		flash.Store(&c.Controller)
		c.Redirect("/error.html", 302)
		return
	}
	isScan := strings.TrimSpace(c.GetString("SCAN"))
	isH5 := strings.TrimSpace(c.GetString("H5"))
	isKj := strings.TrimSpace(c.GetString("KJ"))
	if strings.Contains(isScan, "SCAN") {
		//扫码
		scanShop := new(ScanShopController)
		scanShop.Prepare()
		scanShop.Params["orderPrice"] = amount
		scanShop.Params["payWayCode"] = isScan
		scanShop.Params["orderNo"] = orderNo
		response := scanShop.Shop(c.Ctx.Request.Host)
		if response.Code == 200 {
			str := "/scan.html?" + "orderNo=" + orderNo + "&orderPrice=" + amount + "&qrCode=" + response.Qrcode + "&payWayCode=" + isScan
			c.Redirect(str, 302)
		} else {
			flash.Error(response.Msg)
			flash.Store(&c.Controller)
			c.Redirect("/error.html", 302)
		}
	} else if strings.Contains(isH5, "H5") {

	} else if strings.Contains(isKj, "FAST") {

	} else {
		flash.Error("不存在这样的支付类型")
		flash.Store(&c.Controller)
		c.Redirect("/error.html", 302)
		return
	}
}

func (c *PayController) judgeAmount(amount string) bool {
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		logs.Error("输入金额有误")
		return false
	}

	return true
}
