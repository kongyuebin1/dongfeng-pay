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
	"gateway/response"
	"gateway/service"
	"gateway/supplier/third_party"
	"gateway/utils"
	"strings"
)

type ScanController struct {
	BaseGateway
}

//处理错误的返回
func (c *ScanController) SolveFailJSON(p *response.PayBaseResp) {
	scanFailJSON := new(response.ScanFailData)
	scanFailJSON.StatusCode = "01"
	scanFailJSON.PayKey = p.Params["payKey"]
	scanFailJSON.Msg = p.Msg
	c.Data["json"] = scanFailJSON
	_ = c.ServeJSON()
	c.StopRun()
}

//处理扫码的请求
func (c *ScanController) Scan() {

	p := c.PayPrepare()

	if p.Code == -1 {
		c.SolveFailJSON(p)
	}
	//签名验证
	p.Params["returnUrl"] = strings.TrimSpace(c.GetString("returnUrl"))
	paySecret := p.MerchantInfo.MerchantSecret
	if !utils.Md5Verify(p.Params, paySecret) {
		p.Code = -1
		p.Msg = "签名异常"
		c.SolveFailJSON(p)
	}
	//选择通道
	p = service.ChooseRoad(p)
	if p.Code == -1 {
		c.SolveFailJSON(p)
	}
	//生成订单记录
	orderInfo, _ := service.GenerateRecord(p)
	if p.Code == -1 {
		c.SolveFailJSON(p)
	}
	//获取到对应的上游
	supplierCode := p.RoadInfo.ProductUid
	supplier := third_party.GetPaySupplierByCode(supplierCode)
	scanData := supplier.Scan(orderInfo, p.RoadInfo, p.MerchantInfo)
	if scanData.Status == "00" {
		scanSuccessData := service.GenerateSuccessData(scanData, p)
		c.Data["json"] = scanSuccessData
		_ = c.ServeJSON()
	} else {
		p.Msg = scanData.Msg
		c.SolveFailJSON(p)
	}
}
