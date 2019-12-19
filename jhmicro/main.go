package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"dongfeng-pay/jhmicro/notify"
	"dongfeng-pay/jhmicro/order_settle"
	"dongfeng-pay/jhmicro/pay_for"
	"dongfeng-pay/jhmicro/query"
	"dongfeng-pay/service/service_init"
)

func main() {
	logs.SetLogger(logs.AdapterFile, `{"level": 7, "color":true, "filename":"jhmicro.log"}`)
	service_init.InitAll()
	go notify.CreateOrderNotifyConsumer()
	go query.CreateSupplierOrderQueryCuConsumer()
	go pay_for.PayForInit()
	go query.CreatePayForQueryConsumer()
	go order_settle.OrderSettleInit()
	beego.Run()
}
