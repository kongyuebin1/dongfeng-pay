/***************************************************
 ** @Desc : 模拟商户扫码支付请求
 ** @Time : 2019/10/26 9:48
 ** @Author : yuebin
 ** @File : scan
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/26 9:48
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/skip2/go-qrcode"
	"github.com/widuu/gojson"
	"dongfeng-pay/service/utils"
	"strings"
)

type ScanShopController struct {
	beego.Controller
	Params map[string]string
}

type DataJSON struct {
	Code int
	Msg  string
}

type ResponseJSON struct {
	Code    int
	Msg     string
	OrderNo string
	Url     string
	Qrcode  string
}

const (
	HOST       = "http://localhost:10081"
	SCAN_HOST  = HOST + "/gateway/scan"
	H5_HOST    = HOST + "/gateway/h5"
	SYT_HOST   = HOST + "/gateway/syt"
	FAST_HOST  = HOST + "/gateway/fast"
	NOTIFY_URL = HOST + "/shop/notify"
	RETURN_URL = HOST + "/shop/return"
	PAY_KEY    = "kkkkbmrb9gijhrt0th4naoag"
	PAY_SERCET = "ssssbmrb9gijhrt0th4naob0"
)

func (c *ScanShopController) Prepare() {
	c.Params = make(map[string]string)
	//c.Params["orderNo"] = xid.New().String()
	c.Params["productName"] = "测试"
	c.Params["orderPeriod"] = "1"
	c.Params["osType"] = "1"
	c.Params["notifyUrl"] = NOTIFY_URL
	c.Params["returnUrl"] = RETURN_URL
	c.Params["payKey"] = PAY_KEY
}

func (c *ScanShopController) Shop(requestHost string) *ResponseJSON {

	responseJSON := new(ResponseJSON)

	reqUrl := SCAN_HOST

	keys := utils.SortMap(c.Params)
	sign := utils.GetMD5Sign(c.Params, keys, PAY_SERCET)
	c.Params["sign"] = sign
	req := httplib.Post(reqUrl)
	for k, v := range c.Params {
		req.Param(k, v)
	}
	response, err := req.String()
	if err != nil {
		logs.Error("扫码请求失败")
		responseJSON.Code = -1
		responseJSON.Msg = response + " ;" + err.Error()
	} else {
		statusCode := gojson.Json(response).Get("statusCode").Tostring()
		if statusCode != "00" {
			msg := gojson.Json(response).Get("msg").Tostring()
			responseJSON.Code = -1
			responseJSON.Msg = msg
		} else {
			responseJSON.Code = 200
			payUrl := gojson.Json(response).Get("payURL").Tostring()
			orderNo := gojson.Json(response).Get("orderNo").Tostring()
			qrCodePathName := "./static/img/" + orderNo + ".png"
			qrCode := "/static/img/" + orderNo + ".png"
			GenerateQrcode(payUrl, qrCodePathName)
			responseJSON.OrderNo = orderNo
			responseJSON.Url = payUrl
			responseJSON.Qrcode = "http://" + requestHost + qrCode
		}
	}

	return responseJSON
}

func GenerateQrcode(codeUrl, qrcodePathName string) {
	err := qrcode.WriteFile(codeUrl, qrcode.Medium, 256, qrcodePathName)
	if err != nil {
		logs.Error("generate qrCode fail: ", err)
	}
}

func (c *ScanShopController) ScanRender() {
	orderNo := strings.TrimSpace(c.GetString("orderNo"))
	orderPrice := strings.TrimSpace(c.GetString("orderPrice"))
	qrCode := strings.TrimSpace(c.GetString("qrCode"))
	payWayCode := strings.TrimSpace(c.GetString("payWayCode"))
	if strings.Contains(payWayCode, "UNION") {
		c.Data["payTypeName"] = "云闪付app"
		c.Data["openApp"] = "云闪付app [扫一扫]"
	}
	c.Data["qrCode"] = qrCode
	c.Data["orderNo"] = orderNo
	c.Data["price"] = orderPrice
	c.TplName = "pay/scan.html"
}
