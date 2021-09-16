/***************************************************
 ** @Desc : 供下游订单状态查询和代付结果查询
 ** @Time : 2019/11/6 13:59
 ** @Author : yuebin
 ** @File : order_query
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/6 13:59
 ** @Software: GoLand
****************************************************/
package query

import (
	"encoding/json"
	"fmt"
	"gateway/models/merchant"
	"gateway/models/order"
	"gateway/utils"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"strings"
)

type MerchantQueryController struct {
	web.Controller
}

type OrderQueryFailData struct {
	PayKey     string `json:"payKey"`
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}

/*
** 改接口是为下游商户提供订单查询
 */
func (c *MerchantQueryController) OrderQuery() {
	orderNo := strings.TrimSpace(c.GetString("orderNo"))
	payKey := strings.TrimSpace(c.GetString("payKey"))
	sign := strings.TrimSpace(c.GetString("sign"))
	params := make(map[string]string)
	params["orderNo"] = orderNo
	params["payKey"] = payKey

	failData := new(OrderQueryFailData)
	failData.StatusCode = "01"
	failData.PayKey = payKey

	merchantInfo := merchant.GetMerchantByPaykey(payKey)
	if merchantInfo.MerchantUid == "" || len(merchantInfo.MerchantUid) == 0 {
		failData.Msg = "商户不存在，请核对payKey字段"
	}
	orderInfo := order.GetOrderByMerchantOrderId(orderNo)
	if orderInfo.BankOrderId == "" || len(orderInfo.BankOrderId) == 0 {
		failData.Msg = "不存在这样的订单，请核对orderNo字段"
	}
	keys := utils.SortMap(params)
	paySercet := merchantInfo.MerchantSecret
	tmpSign := utils.GetMD5Sign(params, keys, paySercet)
	if tmpSign != sign {
		failData.Msg = "签名错误"
	}
	if failData.Msg != "" {
		c.Data["json"] = failData
		_ = c.ServeJSON()
		return
	}
	p := make(map[string]string)
	p["orderNo"] = orderNo
	p["orderTime"] = strings.TrimSpace(strings.Replace("-", "", orderInfo.UpdateTime, -1))
	p["trxNo"] = orderInfo.BankOrderId
	p["tradeStatus"] = orderInfo.Status
	p["payKey"] = payKey
	p["orderPrice"] = fmt.Sprintf("%.2f", orderInfo.OrderAmount)
	p["factPrice"] = fmt.Sprintf("%.2f", orderInfo.FactAmount)
	p["statusCode"] = "00"
	keys = utils.SortMap(p)
	p["sign"] = utils.GetMD5Sign(p, keys, paySercet)
	s, err := json.Marshal(p)
	if err != nil {
		logs.Error("json marshal fail： ", err)
	}
	c.Data["json"] = s
}
