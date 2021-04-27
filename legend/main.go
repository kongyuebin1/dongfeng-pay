package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"legend/controllers/expection"
	_ "legend/models"
	_ "legend/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	RegisterLogs()
	RegisterErrorHandler()
	beego.Run()
}

/**
**注册错误页面
 */
func RegisterErrorHandler() {
	beego.ErrorHandler("404", expection.PageNotFind)
	beego.ErrorHandler("503", expection.SystemExpection)
}

/**
** 注册日志信息
 */
func RegisterLogs() {
	logs.SetLogger(logs.AdapterFile,
		`{
						"filename":"../../logs/legend.log",
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
