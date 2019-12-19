package main

import (
	"github.com/astaxie/beego"
	_ "juhe/jhgateway/routers"
	"juhe/service/service_init"
)

func main() {
	//启动订单查询消费者
	//go gateway.CreateSupplierOrderQueryCuConsumer()
	service_init.InitAll()
	beego.Run()
}
