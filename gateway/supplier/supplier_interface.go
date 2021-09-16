/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/28 9:39
 ** @Author : yuebin
 ** @File : supplier_interface
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/28 9:39
 ** @Software: GoLand
****************************************************/
package supplier

import (
	"gateway/models/merchant"
	"gateway/models/order"
	"gateway/models/payfor"
	"gateway/models/road"
)

//定义扫码支付的返回值
type ScanData struct {
	Supplier   string //上游的通道供应商
	PayType    string //支付类型
	OrderNo    string //下游商户请求订单号
	BankNo     string //本系统的请求订单号
	OrderPrice string //订单金额
	FactPrice  string //实际的展示在客户面前的金额
	Status     string //状态码 '00' 成功
	PayUrl     string //支付二维码链接地址
	Msg        string //附加的信息
}

type PayInterface interface {
	Scan(order.OrderInfo, road.RoadInfo, merchant.MerchantInfo) ScanData
	H5(order.OrderInfo, road.RoadInfo, merchant.MerchantInfo) ScanData
	Fast(order.OrderInfo, road.RoadInfo, merchant.MerchantInfo) bool
	Syt(order.OrderInfo, road.RoadInfo, merchant.MerchantInfo) ScanData
	Web(order.OrderInfo, road.RoadInfo, merchant.MerchantInfo) bool
	PayNotify()
	PayQuery(order.OrderInfo) bool
	PayFor(payfor.PayforInfo) string
	PayForNotify() string
	PayForQuery(payfor.PayforInfo) (string, string)
	BalanceQuery(road.RoadInfo) float64
}
