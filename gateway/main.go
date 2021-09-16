package main

import (
	_ "gateway/message"
	_ "gateway/models"
	"gateway/notify"
	"gateway/query"
	_ "gateway/routers"
	"gateway/service"
	_ "gateway/supplier/third_party"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	RegisterLogs()
	web.BConfig.WebConfig.Session.SessionOn = true

	go notify.CreateOrderNotifyConsumer()
	//go pay_for.PayForInit()
	go query.CreatePayForQueryConsumer()
	go service.OrderSettleInit()
	go query.CreateSupplierOrderQueryCuConsumer()

	web.Run()
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
