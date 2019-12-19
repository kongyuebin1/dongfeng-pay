package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "juhe/jhboss/routers"
	_ "juhe/service/message_queue"
	_ "juhe/service/models"
	"juhe/service/service_init"
)

func main() {
	//设置日志打印
	logs.SetLogger(logs.AdapterFile, `{"filename":"jhboss.log", "level":7, "daily":true, "maxdays":10}`)
	service_init.InitAll()
	beego.Run()
}
