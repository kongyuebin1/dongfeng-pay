package main

import (
	_ "gateway/models"
	_ "gateway/supplier"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//启动订单查询消费者
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
