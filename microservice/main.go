package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"microservice/notify"
	"microservice/order_settle"
	"microservice/pay_for"
	"microservice/query"
	_ "microservice/supplier"
)

func main() {
	RegisterLogs()
	//logs.SetLogger(logs.AdapterFile, `{"level": 7, "color":true, "filename":"jhmicro.log"}`)
	go notify.CreateOrderNotifyConsumer()
	go query.CreateSupplierOrderQueryCuConsumer()
	go pay_for.PayForInit()
	go query.CreatePayForQueryConsumer()
	go order_settle.OrderSettleInit()
	beego.Run()
}

/**
** 注册日志信息
 */
func RegisterLogs() {
	logs.SetLogger(logs.AdapterFile,
		`{
						"filename":"../logs/legend.log",
						"level":4,
						"maxlines":0,
						"maxsize":0,
						"daily":true,
						"maxdays":10,
						"color":true
					}`)

	f := &logs.PatternLogFormatter{
		Pattern:    "%F:%n|%w%t>> %m",
		WhenFormat: "2006-01-02",
	}

	logs.RegisterFormatter("pattern", f)
	_ = logs.SetGlobalFormatter("pattern")
}
