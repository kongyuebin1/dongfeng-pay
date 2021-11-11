package main

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "merchant/models"
	_ "merchant/routers"
	"merchant/sys"
	"merchant/utils"
)

func init() {

	// 初始化Session
	sys.InitSession()
	utils.InitLogs()
	// 如果是开发模式，则显示命令信息
	s, _ := beego.AppConfig.String("runmode")
	isDev := !(s != "dev")
	if isDev {
		orm.Debug = isDev
	}
}

func main() {
	beego.Run()
}
