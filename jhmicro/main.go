package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"juhe/jhmicro/notify"
	"juhe/jhmicro/order_settle"
	"juhe/jhmicro/pay_for"
	"juhe/jhmicro/query"
	"juhe/service/service_init"
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
