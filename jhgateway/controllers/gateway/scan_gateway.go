/***************************************************
 ** @Desc : 下游请求扫码支付的处理逻辑
 ** @Time : 2019/10/24 11:15
 ** @Author : yuebin
 ** @File : gateway
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/24 11:15
 ** @Software: GoLand
****************************************************/
package gateway

import (
	"juhe/service/controller"
	"juhe/service/utils"
	"strings"
)

type ScanController struct {
	BaseGateway
}

type ScanSuccessData struct {
	OrderNo    string `json:"orderNo"`
	Sign       string `json:"sign"`
	OrderPrice string `json:"orderPrice"`
	PayKey     string `json:"payKey"`
	PayUrl     string `json:"payURL"`
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}

type ScanFailData struct {
	PayKey     string `json:"payKey"`
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}

//处理错误的返回
func (c *ScanController) SolveFailJSON() {
	scanFailJSON := new(ScanFailData)
	scanFailJSON.StatusCode = "01"
	scanFailJSON.PayKey = c.Params["payKey"]
	scanFailJSON.Msg = c.Msg
	c.Data["json"] = scanFailJSON
	c.ServeJSON()
	c.StopRun()
}

//处理扫码的请求
func (c *ScanController) Scan() {

	c.PayPrepare()

	if c.Code == -1 {
		c.SolveFailJSON()
	}
	//签名验证
	c.Params["returnUrl"] = strings.TrimSpace(c.GetString("returnUrl"))
	paySecret := c.MerchantInfo.MerchantSecret
	if !utils.Md5Verify(c.Params, paySecret) {
		c.Code = -1
		c.Msg = "签名异常"
		c.SolveFailJSON()
	}
	//选择通道
	c.ChooseRoad()
	if c.Code == -1 {
		c.SolveFailJSON()
	}
	//升级订单记录
	orderInfo, _ := c.GenerateRecord()
	if c.Code == -1 {
		c.SolveFailJSON()
	}
	//获取到对应的上游
	supplierCode := c.RoadInfo.ProductUid
	supplier := controller.GetPaySupplierByCode(supplierCode)
	scanData := supplier.Scan(orderInfo, c.RoadInfo, c.MerchantInfo)
	if scanData.Status == "00" {
		scanSuccessData := c.GenerateSuccessData(scanData)
		c.Data["json"] = scanSuccessData
		c.ServeJSON()
	} else {
		c.SolveFailJSON()
	}
}
