package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "dongfeng-pay/jhboss/routers"
	_ "dongfeng-pay/service/message_queue"
	_ "dongfeng-pay/service/models"
	"dongfeng-pay/service/service_init"
)

func main() {
	//设置日志打印
	logs.SetLogger(logs.AdapterFile, `{"filename":"jhboss.log", "level":7, "daily":true, "maxdays":10}`)
	service_init.InitAll()
	beego.Run()
}
