/***************************************************
 ** @Desc : 快付支付的实现逻辑
 ** @Time : 2019/10/28 14:12
 ** @Author : yuebin
 ** @File : kuaifu
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/28 14:12
 ** @Software: GoLand
****************************************************/
package third_party

import (
	"fmt"
	"gateway/conf"
	"gateway/models/merchant"
	"gateway/models/order"
	"gateway/models/payfor"
	"gateway/models/road"
	"gateway/service"
	"gateway/supplier"
	"gateway/utils"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/rs/xid"
	"github.com/widuu/gojson"
	"strconv"
	"strings"
)

type KuaiFuImpl struct {
	beego.Controller
}

const (
	HOST             = "localhost"
	KF_SCAN_HOST     = "http://" + HOST + "/gateway/scanPay/payService"
	KF_PAYFOR_HOST   = "http://" + HOST + "/gateway/remittance/pay"
	KF_BALANCE_QUERY = "http://" + HOST + "/gateway/remittance/getBalance"
	KF_ORDER_QUERY   = "http://" + HOST + "/gateway/scanPay/orderQuery"
	KF_PAYFOR_QUERY  = "http://" + HOST + "/gateway/remittance/query"
	KF_PAY_KEY       = "xxxxxxx"
	KF_PAY_SECRET    = "xxxxxx"
)

func (c *KuaiFuImpl) Scan(orderInfo order.OrderInfo, roadInfo road.RoadInfo, merchantInfo merchant.MerchantInfo) supplier.ScanData {
	payWayCode := ""
	switch orderInfo.PayTypeCode {
	case "ALI_SCAN":
		payWayCode = "SCAN_ALIPAY"
	case "WEIXIN_SCAN":
		payWayCode = "SCAN_WEIXIN"
	case "QQ_SCAN":
		payWayCode = "SCAN_QQ"
	case "UNION_SCAN":
		payWayCode = "SCAN_YL"
	case "BAIDU_SCAN":
	case "JD_SCAN":
	}
	//将金额转为带有2位小数点的float
	order := fmt.Sprintf("%0.2f", orderInfo.OrderAmount)
	params := make(map[string]string)
	params["orderNo"] = orderInfo.BankOrderId
	params["productName"] = orderInfo.ShopName
	params["orderPeriod"] = orderInfo.OrderPeriod
	params["orderPrice"] = order
	params["payWayCode"] = payWayCode
	params["osType"] = orderInfo.OsType
	params["notifyUrl"] = "KF"
	params["payKey"] = KF_PAY_KEY
	//params["field1"] = "field1"

	keys := utils.SortMap(params)
	sign := utils.GetMD5Sign(params, keys, KF_PAY_SECRET)
	params["sign"] = sign

	req := httplib.Post(KF_SCAN_HOST)
	for k, v := range params {
		req.Param(k, v)
	}
	var scanData supplier.ScanData
	scanData.Supplier = orderInfo.PayProductCode
	scanData.PayType = orderInfo.PayTypeCode
	scanData.OrderNo = orderInfo.MerchantOrderId
	scanData.BankNo = orderInfo.BankOrderId
	scanData.OrderPrice = params["orderPrice"]
	response, err := req.String()
	if err != nil {
		logs.Error("KF 请求失败：", err)
		scanData.Status = "01"
		scanData.Msg = gojson.Json(response).Get("statusMsg").Tostring()
		return scanData
	}
	statusCode := gojson.Json(response).Get("statusCode").Tostring()
	if statusCode != "00" {
		logs.Error("KF生成扫码地址失败")
		scanData.Status = "01"
		scanData.Msg = "生成扫码地址失败"
		return scanData
	}
	payUrl := gojson.Json(response).Get("payURL").Tostring()
	scanData.Status = "00"
	scanData.PayUrl = payUrl
	scanData.Msg = "请求成功"
	return scanData
}

func (c *KuaiFuImpl) H5(orderInfo order.OrderInfo, roadInfo road.RoadInfo, merchantInfo merchant.MerchantInfo) supplier.ScanData {
	var scanData supplier.ScanData
	scanData.Status = "01"
	return scanData
}

func (c *KuaiFuImpl) Syt(orderInfo order.OrderInfo, roadInfo road.RoadInfo, merchantInfo merchant.MerchantInfo) supplier.ScanData {
	var scanData supplier.ScanData
	scanData.Status = "01"
	return scanData
}

func (c *KuaiFuImpl) Fast(orderInfo order.OrderInfo, roadInfo road.RoadInfo, merchantInfo merchant.MerchantInfo) bool {
	var scanData supplier.ScanData
	scanData.Status = "01"
	return true
}

func (c *KuaiFuImpl) Web(orderInfo order.OrderInfo, roadInfo road.RoadInfo, merchantInfo merchant.MerchantInfo) bool {
	var scanData supplier.ScanData
	scanData.Status = "01"
	return true
}

//支付回调
func (c *KuaiFuImpl) PayNotify() {
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
		logs.Error("快付回调签名异常，回调失败")
		c.StopRun()
	}
	//实际支付金额
	factAmount, err := strconv.ParseFloat(params["orderPrice"], 64)
	if err != nil {
		logs.Error("快付回调实际金额有误， factAmount=", params["orderPrice"])
		c.StopRun()
	}
	orderInfo.FactAmount = factAmount
	orderInfo.BankTransId = params["trxNo"]
	tradeStatus := params["tradeStatus"]
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

func (c *KuaiFuImpl) PayQuery(orderInfo order.OrderInfo) bool {
	if orderInfo.Status != "wait" && orderInfo.Status != "" {
		logs.Error("订单已经被处理，不需要查询，bankOrderId：", orderInfo.BankOrderId)
		return false
	}
	params := make(map[string]string)
	params["orderNo"] = orderInfo.BankOrderId
	params["payKey"] = KF_PAY_KEY
	paySecret := KF_PAY_SECRET
	keys := utils.SortMap(params)
	params["sign"] = utils.GetMD5Sign(params, keys, paySecret)
	req := httplib.Get(KF_ORDER_QUERY)
	for k, v := range params {
		req.Param(k, v)
	}
	response, err := req.String()
	if err != nil {
		logs.Error("快付订单查询失败,bankOrderId: ", orderInfo.BankOrderId)
		logs.Error("err: ", err)
		return false
	}
	statusCode := gojson.Json(response).Get("statusCode").Tostring()
	if statusCode != "00" {
		logs.Error("快付订单查询返回失败，bankOrderId：", orderInfo.BankOrderId)
		logs.Error("err: ", response)
		return false
	}
	//获取用户的实际支付金额
	orderPrice := gojson.Json(response).Get("orderPrice").Tostring()
	factAmount, err := strconv.ParseFloat(orderPrice, 64)
	if err != nil {
		logs.Error("快速查询得到的实际金额错误, orderPrice=", orderPrice)
	}

	//orderInfo.FactAmount = orderInfo.OrderAmount
	tradeStatus := gojson.Json(response).Get("tradeStatus").Tostring()
	trxNo := gojson.Json(response).Get("trxNo").Tostring()
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

func (c *KuaiFuImpl) PayFor(payFor payfor.PayforInfo) string {
	params := make(map[string]string)
	params["merchantKey"] = KF_PAY_KEY
	params["realname"] = payFor.BankAccountName
	params["cardNo"] = payFor.BankAccountNo
	params["bankCode"] = payFor.BankCode
	if payFor.BankAccountType == conf.PRIVATE {
		params["accType"] = "01"
	} else {
		params["accType"] = "02"
	}
	params["province"] = payFor.BankAccountAddress
	params["city"] = payFor.BankAccountAddress
	params["bankAccountAddress"] = payFor.BankAccountAddress
	//将float64转为字符串
	params["amount"] = strconv.FormatFloat(payFor.PayforAmount, 'f', 2, 64)
	params["moblieNo"] = payFor.PhoneNo
	params["merchantOrderId"] = payFor.BankOrderId
	keys := utils.SortMap(params)
	sign := utils.GetMD5Sign(params, keys, KF_PAY_SECRET)
	params["sign"] = sign
	req := httplib.Post(KF_PAYFOR_HOST)
	for k, v := range params {
		req.Param(k, v)
	}
	response, err := req.String()
	if err != nil {
		logs.Error("快付代付返回错误结果： ", response)
	} else {
		json := gojson.Json(response)
		resultCode := json.Get("resultCode").Tostring()
		resultMsg := json.Get("resultMsg").Tostring()
		if resultCode != "00" {
			logs.Error("快付代付返回错误信息：", resultMsg)
			return "fail"
		}
		settStatus := json.Get("settStatus").Tostring()
		if settStatus == "00" {
			logs.Info(fmt.Sprintf("代付uid=%s，已经成功发送给了上游处理", payFor.PayforUid))
		} else if settStatus == "01" {
			logs.Info(fmt.Sprintf("代付uid=%s，发送失败", payFor.PayforUid))
		}
	}
	return "success"
}

func (c *KuaiFuImpl) PayForNotify() string {
	return ""
}

func (c *KuaiFuImpl) PayForQuery(payFor payfor.PayforInfo) (string, string) {
	params := make(map[string]string)
	params["merchantKey"] = KF_PAY_KEY
	params["timestamp"] = utils.GetNowTimesTamp()
	params["merchantOrderId"] = payFor.BankOrderId
	keys := utils.SortMap(params)
	sign := utils.GetMD5Sign(params, keys, KF_PAY_SECRET)
	params["sign"] = sign
	req := httplib.Get(KF_PAYFOR_QUERY)
	for k, v := range params {
		req.Param(k, v)
	}
	response, err := req.String()
	if err != nil {
		logs.Error("快付代付查询失败：", err)
		return conf.PAYFOR_SOLVING, "查询失败"
	}

	payFor.ResponseContent = response
	payFor.ResponseTime = utils.GetBasicDateTime()
	payFor.UpdateTime = utils.GetBasicDateTime()
	if !payfor.UpdatePayFor(payFor) {
		logs.Error("更新快付代付订单状态失败")
	}

	resultCode := gojson.Json(response).Get("resultCode").Tostring()
	resultMsg := gojson.Json(response).Get("resultMsg").Tostring()

	if resultCode != "00" {
		logs.Error("快付代付查询返回错误：", resultMsg)
		return conf.PAYFOR_SOLVING, resultMsg
	}

	logs.Info("快付代付查询返回结果：", resultMsg)

	merchantOrderId := gojson.Json(response).Get("merchantOrderId").Tostring()
	if merchantOrderId != payFor.BankOrderId {
		logs.Error("快付代付返回结果，订单id不一致: ", merchantOrderId)
		return conf.PAYFOR_SOLVING, "快付代付返回结果，订单id不一致"
	}

	settStatus := gojson.Json(response).Get("settStatus").Tostring()

	if settStatus == "00" {
		return conf.PAYFOR_SUCCESS, "代付成功"
	} else if settStatus == "01" {
		return conf.PAYFOR_FAIL, "代付失败"
	} else {
		return conf.PAYFOR_BANKING, "银行处理中"
	}
}

func (c *KuaiFuImpl) BalanceQuery(roadInfo road.RoadInfo) float64 {
	params := make(map[string]string)
	params["merchantKey"] = KF_PAY_KEY
	params["timestamp"] = utils.GetNowTimesTamp()
	params["merchantOrderId"] = xid.New().String()
	keys := utils.SortMap(params)
	sign := utils.GetMD5Sign(params, keys, KF_PAY_SECRET)
	params["sign"] = sign
	req := httplib.Get(KF_BALANCE_QUERY)
	for k, v := range params {
		req.Param(k, v)
	}

	response, err := req.String()
	if err != nil {
		logs.Error("快付余额查询失败,err: ", err)
		return 0.00
	}
	logs.Debug("快付余额查询返回：", response)

	resultCode := gojson.Json(response).Get("resultCode").Tostring()
	resultMsg := gojson.Json(response).Get("resultMsg").Tostring()
	logs.Notice("快付返回信息：", resultMsg)

	if resultCode != "00" {
		return 0.00
	}

	balance := gojson.Json(response).Get("balance").Tostring()
	availableAmount := gojson.Json(response).Get("availableAmount").Tostring()

	logs.Info(fmt.Sprintf("快付余额=%s，可用金额=%s", balance, availableAmount))

	f, err := strconv.ParseFloat(availableAmount, 64)
	return f
}
