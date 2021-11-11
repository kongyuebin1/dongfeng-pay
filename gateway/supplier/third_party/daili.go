/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/28 16:38
 ** @Author : yuebin
 ** @File : alipay
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/28 16:38
 ** @Software: GoLand
****************************************************/
package third_party

import (
	"gateway/models/merchant"
	"gateway/models/order"
	"gateway/models/payfor"
	"gateway/models/road"
	"gateway/service"
	"gateway/supplier"
	"gateway/utils"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/rs/xid"
	"github.com/widuu/gojson"
	"strconv"
	"strings"
)

type DaiLiImpl struct {
	web.Controller
}

const NOTITY_URL = "http://localhost:12306/accept/notify"
const URL = "http://zhaoyin.lfwin.com/payapi/pay/jspay3"

func (c *DaiLiImpl) Scan(orderInfo order.OrderInfo, roadInfo road.RoadInfo, merchantInfo merchant.MerchantInfo) supplier.ScanData {
	// 从boss后台获取数据
	service := gojson.Json(roadInfo.Params).Get("service").Tostring()
	apiKey := gojson.Json(roadInfo.Params).Get("apikey").Tostring()
	signKey := gojson.Json(roadInfo.Params).Get("signkey").Tostring()

	params := make(map[string]string)
	params["service"] = service
	params["apikey"] = apiKey
	params["money"] = strconv.FormatFloat(orderInfo.OrderAmount, 'f', 2, 32)
	params["nonce_str"] = xid.New().String()
	params["mch_orderid"] = orderInfo.BankOrderId
	params["notify_url"] = NOTITY_URL

	waitStr := utils.MapToString(utils.SortMapByKeys(params))
	waitStr = waitStr + "&signkey=" + signKey
	sign := utils.GetMD5LOWER(waitStr)
	params["sign"] = sign

	request := URL + "?" + utils.MapToString(params)

	logs.Info("代丽请求字符串 = " + request)

	var scanData supplier.ScanData
	scanData.Status = "00"
	response, err := httplib.Post(request).String()
	if err != nil {
		logs.Error("代丽支付请求失败：" + err.Error())
		scanData.Status = "-1"
		scanData.Msg = "请求失败：" + err.Error()
	} else {
		/*logs.Info("代丽支付返回 = " + response)
		status := gojson.Json(response).Get("status").Tostring()
		message := gojson.Json(response).Get("message").Tostring()
		if "10000" != status {
			scanData.Status = "-1"
			scanData.Msg = message
		} else {*/
		codeUrl := gojson.Json(response).Get("url").Tostring()
		codeUrl = "http://www.baidu.com"
		scanData.PayUrl = codeUrl
		scanData.OrderNo = orderInfo.BankOrderId
		scanData.OrderPrice = strconv.FormatFloat(orderInfo.OrderAmount, 'f', 2, 64)
		//}
	}

	return scanData
}

func (c *DaiLiImpl) H5(orderInfo order.OrderInfo, roadInfo road.RoadInfo, merchantInfo merchant.MerchantInfo) supplier.ScanData {
	var scanData supplier.ScanData
	scanData.Status = "01"
	return scanData
}

func (c *DaiLiImpl) Syt(orderInfo order.OrderInfo, roadInfo road.RoadInfo, merchantInfo merchant.MerchantInfo) supplier.ScanData {
	var scanData supplier.ScanData
	scanData.Status = "01"
	return scanData
}

func (c *DaiLiImpl) Fast(orderInfo order.OrderInfo, roadInfo road.RoadInfo, merchantInfo merchant.MerchantInfo) bool {
	var scanData supplier.ScanData
	scanData.Status = "01"
	return true
}

func (c *DaiLiImpl) Web(orderInfo order.OrderInfo, roadInfo road.RoadInfo, merchantInfo merchant.MerchantInfo) bool {
	var scanData supplier.ScanData
	scanData.Status = "01"
	return true
}

func (c *DaiLiImpl) PayNotify() {
	params := make(map[string]string)
	orderNo := strings.TrimSpace(c.GetString("orderNo"))
	orderInfo := order.GetOrderByBankOrderId(orderNo)
	if orderInfo.BankOrderId == "" || len(orderInfo.BankOrderId) == 0 {
		logs.Error("快付回调的订单号不存在，订单号=", orderNo)
		c.StopRun()
	}
	roadInfo := road.GetRoadInfoByRoadUid(orderInfo.RoadUid)
	if roadInfo.RoadUid == "" || len(roadInfo.RoadUid) == 0 {
		logs.Error("支付通道已经关系或者删除，不进行回调")
		c.StopRun()
	}
	merchantUid := orderInfo.MerchantUid
	merchantInfo := merchant.GetMerchantByUid(merchantUid)
	if merchantInfo.MerchantUid == "" || len(merchantInfo.MerchantUid) == 0 {
		logs.Error("快付回调失败，该商户不存在或者已经删除，商户uid=", merchantUid)
		c.StopRun()
	}
	paySecret := merchantInfo.MerchantSecret
	params["orderNo"] = orderNo
	params["orderPrice"] = strings.TrimSpace(c.GetString("orderPrice"))
	params["orderTime"] = strings.TrimSpace(c.GetString("orderTime"))
	params["trxNo"] = strings.TrimSpace(c.GetString("trxNo"))
	params["statusCode"] = strings.TrimSpace(c.GetString("statusCode"))
	params["tradeStatus"] = strings.TrimSpace(c.GetString("tradeStatus"))
	params["field1"] = strings.TrimSpace(c.GetString("field1"))
	params["payKey"] = strings.TrimSpace(c.GetString("payKey"))
	//对参数进行验签
	keys := utils.SortMap(params)
	tmpSign := utils.GetMD5Sign(params, keys, paySecret)
	sign := strings.TrimSpace(c.GetString("sign"))
	if tmpSign != sign {
		logs.Error("代丽回调签名异常，回调失败")
		//c.StopRun()
	}
	//实际支付金额
	factAmount, err := strconv.ParseFloat(params["orderPrice"], 64)
	if err != nil {
		orderInfo.FactAmount = 0
	}
	orderInfo.FactAmount = factAmount
	orderInfo.BankTransId = params["trxNo"]
	tradeStatus := params["tradeStatus"]

	//paySolveController := new(service.PaySolveController)
	if tradeStatus == "FAILED" {
		if !service.SolvePayFail(orderInfo.BankOrderId, "") {
			logs.Error("solve order fail fail")
		}
	} else if tradeStatus == "CANCELED" {
		if !service.SolvePayFail(orderInfo.BankOrderId, "") {
			logs.Error("solve order cancel fail")
		}
	} else if tradeStatus == "WAITING_PAYMENT" {
		logs.Notice("快付回调，该订单还处于等待支付，订单id=", orderNo)
	} else if tradeStatus == "SUCCESS" {
		//订单支付成功，需要搞很多事情 TODO
		service.SolvePaySuccess(orderInfo.BankOrderId, orderInfo.FactAmount, c.GetString("trxNo"))
	}
	c.Ctx.WriteString("success")
}

func (c *DaiLiImpl) PayQuery(orderInfo order.OrderInfo) bool {

	tradeStatus := "SUCCESS"
	trxNo := orderInfo.BankOrderId
	factAmount := 100.00
	if tradeStatus == "SUCCESS" {
		//调用支付成功的接口，做加款更新操作，需要把实际支付金额传入
		if !service.SolvePaySuccess(orderInfo.BankOrderId, factAmount, trxNo) {
			return false
		}
	} else if tradeStatus == "FAILED" {
		if !service.SolvePayFail(orderInfo.BankOrderId, "") {
			return false
		}
	} else {
		logs.Info("订单状态处于：" + tradeStatus + "；bankOrderId：" + orderInfo.BankOrderId)
	}
	return true
}

func (c *DaiLiImpl) PayFor(info payfor.PayforInfo) string {
	return ""
}

func (c *DaiLiImpl) PayForNotify() string {
	return ""
}

func (c *DaiLiImpl) PayForQuery(payFor payfor.PayforInfo) (string, string) {
	return "", ""
}

func (c *DaiLiImpl) BalanceQuery(roadInfo road.RoadInfo) float64 {
	return 0.00
}
