package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "juhe/jhmerchant/routers"
	"juhe/jhmerchant/sys"
	"juhe/jhmerchant/utils"
	"juhe/service/service_init"
)

func init() {
	// 初始化日志
	utils.InitLogs()

	// 初始化数据库
	service_init.InitAll()

	// 初始化Session
	sys.InitSession()

	// 如果是开发模式，则显示命令信息
	isDev := !(beego.AppConfig.String("runmode") != "dev")
	if isDev {
		orm.Debug = isDev
	}
}

func main() {
	beego.Run()
}
