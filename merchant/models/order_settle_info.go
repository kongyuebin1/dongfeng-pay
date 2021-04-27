/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/30 11:41
 ** @Author : yuebin
 ** @File : order_settle_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/30 11:41
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type OrderSettleInfo struct {
	Id               int
	PayProductCode   string
	PayProductName   string
	PayTypeCode      string
	RoadUid          string
	PayTypeName      string
	MerchantUid      string
	MerchantName     string
	MerchantOrderId  string
	BankOrderId      string
	SettleAmount     float64
	IsAllowSettle    string
	IsCompleteSettle string
	UpdateTime       string
	CreateTime       string
}

const ORDER_SETTLE_INFO = "order_settle_info"

func GetOrderSettleListByParams(params map[string]string) []OrderSettleInfo {
	o := orm.NewOrm()
	qs := o.QueryTable(ORDER_SETTLE_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	var orderSettleList []OrderSettleInfo
	if _, err := qs.Limit(-1).All(&orderSettleList); err != nil {
		logs.Error("get order settle list fail: ", err)
	}

	return orderSettleList
}
