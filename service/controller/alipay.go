/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/28 16:38
 ** @Author : yuebin
 ** @File : alipay
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/28 16:38
 ** @Software: GoLand
****************************************************/
package controller

import (
	"juhe/service/models"
)

type AlipayImpl struct {
}

func (c *AlipayImpl) Scan(orderInfo models.OrderInfo, roadInfo models.RoadInfo, merchantInfo models.MerchantInfo) ScanData {
	var scanData ScanData
	scanData.Status = "01"
	return scanData
}

func (c *AlipayImpl) H5(orderInfo models.OrderInfo, roadInfo models.RoadInfo, merchantInfo models.MerchantInfo) ScanData {
	var scanData ScanData
	scanData.Status = "01"
	return scanData
}

func (c *AlipayImpl) Syt(orderInfo models.OrderInfo, roadInfo models.RoadInfo, merchantInfo models.MerchantInfo) ScanData {
	var scanData ScanData
	scanData.Status = "01"
	return scanData
}

func (c *AlipayImpl) Fast(orderInfo models.OrderInfo, roadInfo models.RoadInfo, merchantInfo models.MerchantInfo) bool {
	var scanData ScanData
	scanData.Status = "01"
	return true
}

func (c *AlipayImpl) Web(orderInfo models.OrderInfo, roadInfo models.RoadInfo, merchantInfo models.MerchantInfo) bool {
	var scanData ScanData
	scanData.Status = "01"
	return true
}

func (c *AlipayImpl) PayNotify() {
}

func (c *AlipayImpl) PayQuery(orderInfo models.OrderInfo) bool {
	return true
}

func (c *AlipayImpl) PayFor(info models.PayforInfo) string {
	return ""
}

func (c *AlipayImpl) PayForNotify() string {
	return ""
}

func (c *AlipayImpl) PayForQuery(payFor models.PayforInfo) (string, string) {
	return "", ""
}

func (c *AlipayImpl) BalanceQuery(roadInfo models.RoadInfo) float64 {
	return 0.00
}
