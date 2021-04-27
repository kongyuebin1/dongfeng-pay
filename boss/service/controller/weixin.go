/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/28 16:38
 ** @Author : yuebin
 ** @File : weixin
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/28 16:38
 ** @Software: GoLand
****************************************************/
package controller

import (
	"dongfeng/service/models"
)

type WeiXinImpl struct {
}

func (c *WeiXinImpl) Scan(orderInfo models.OrderInfo, roadInfo models.RoadInfo, merchantInfo models.MerchantInfo) ScanData {
	var scanData ScanData
	scanData.Status = "01"
	return scanData
}

func (c *WeiXinImpl) H5(orderInfo models.OrderInfo, roadInfo models.RoadInfo, merchantInfo models.MerchantInfo) ScanData {
	var scanData ScanData
	scanData.Status = "01"
	return scanData
}

func (c *WeiXinImpl) Syt(orderInfo models.OrderInfo, roadInfo models.RoadInfo, merchantInfo models.MerchantInfo) ScanData {
	var scanData ScanData
	scanData.Status = "01"
	return scanData
}

func (c *WeiXinImpl) Fast(orderInfo models.OrderInfo, roadInfo models.RoadInfo, merchantInfo models.MerchantInfo) bool {
	var scanData ScanData
	scanData.Status = "01"
	return true
}

func (c *WeiXinImpl) Web(orderInfo models.OrderInfo, roadInfo models.RoadInfo, merchantInfo models.MerchantInfo) bool {
	var scanData ScanData
	scanData.Status = "01"
	return true
}

func (c *WeiXinImpl) PayNotify() {
}

func (c *WeiXinImpl) PayQuery(orderInfo models.OrderInfo) bool {
	return true
}

func (c *WeiXinImpl) PayFor(payFor models.PayforInfo) string {
	return ""
}

func (c *WeiXinImpl) PayForNotify() string {
	return ""
}

func (c *WeiXinImpl) PayForQuery(payFor models.PayforInfo) (string, string) {
	return "", ""
}

func (c *WeiXinImpl) BalanceQuery(roadInfo models.RoadInfo) float64 {
	return 0.00
}
