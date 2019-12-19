/***************************************************
 ** @Desc : This file for ...日志信息
 ** @Time : 2018-12-22 17:34:38
 ** @Author : Joker
 ** @File : DefaultTest.go
 ** @Last Modified by : Joker
 ** @Last Modified time: 2019-11-29 10:54:13
 ** @Software: GoLand
****************************************************/

package utils

import (
	"dongfeng-pay/jhmerchant/sys"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger

// fileLogs 生产环境下日志
var fileLogs *logs.BeeLogger

//运行方式
var runmode string

var pubMethod = sys.PublicMethod{}

func InitLogs() {
	//日志输出调用的文件名和文件行号
	logs.EnableFuncCallDepth(true)
	beego.SetLogFuncCall(true)

	consoleLogs = logs.NewLogger(1)
	_ = consoleLogs.SetLogger(logs.AdapterConsole)
	consoleLogs.Async() //异步
	fileLogs = logs.NewLogger(10000)

	//读取配置信息
	filepath := beego.AppConfig.String("logs::filepath")
	level := beego.AppConfig.String("logs::level")
	separate := beego.AppConfig.String("logs::separate")
	maxdays := beego.AppConfig.String("logs::maxdays")
	config := `{"filename":"` + filepath + `",
		"separate":` + separate + `,
		"level":` + level + `,
		"daily":true,
		"maxdays":` + maxdays + `}`
	_ = fileLogs.SetLogger(logs.AdapterMultiFile, config)

	fileLogs.Async(1e3) //异步,设置缓冲 chan 的大小
	runmode = strings.TrimSpace(strings.ToLower(beego.AppConfig.String("runmode")))
	if runmode == "" {
		runmode = "dev"
	}
	LogNotice("商户日志初始化成功!")
}

//根据错误/异常打印不同日志
func LogEmergency(v ...interface{}) {
	log("emergency", v)
}
func LogAlert(v ...interface{}) {
	log("alert", v)
}
func LogCritical(v ...interface{}) {
	log("critical", v)
}
func LogError(v ...interface{}) {
	log("error", v)
}
func LogWarning(v ...interface{}) {
	log("warning", v)
}
func LogNotice(v ...interface{}) {
	log("notice", v)
}
func LogInfo(v ...interface{}) {
	log("info", v)
}
func LogDebug(v ...interface{}) {
	log("debug", v)
}

func LogTrace(v ...interface{}) {
	log("trace", v)
}

//Log 输出日志
func log(level string, v ...interface{}) {
	format := "%s"
	if level == "" {
		level = "debug"
	}
	//若是开发者模式,则将日志同时输出到控制台
	if runmode == "dev" {
		switch level {
		case "emergency":
			consoleLogs.Emergency(format, v...)
		case "alert":
			consoleLogs.Alert(format, v...)
		case "critical":
			consoleLogs.Critical(format, v...)
		case "warning":
			consoleLogs.Warning(format, v...)
		case "error":
			consoleLogs.Error(format, v...)
		case "notice":
			consoleLogs.Notice(format, v...)
		case "info":
			consoleLogs.Info(format, v...)
		case "debug":
			consoleLogs.Debug(format, v...)
		case "trace":
			consoleLogs.Trace(format, v...)
		default:
			consoleLogs.Debug(format, v...)
		}
	}

	switch level {
	case "emergency":
		fileLogs.Emergency(format, v...)
	case "alert":
		fileLogs.Alert(format, v...)
	case "critical":
		fileLogs.Critical(format, v...)
	case "error":
		fileLogs.Error(format, v...)
	case "warning":
		fileLogs.Warning(format, v...)
	case "notice":
		fileLogs.Notice(format, v...)
	case "info":
		fileLogs.Info(format, v...)
	case "debug":
		fileLogs.Debug(format, v...)
	case "trace":
		fileLogs.Trace(format, v...)
	default:
		fileLogs.Debug(format, v...)
	}
}
