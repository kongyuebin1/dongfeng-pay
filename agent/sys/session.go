/***************************************************
 ** @Desc : This file for session配置
 ** @Time : 19.11.30 17:44
 ** @Author : Joker
 ** @File : session
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.11.30 17:44
 ** @Software: GoLand
****************************************************/
package sys

import (
	"agent/sys/enum"
	beego "github.com/beego/beego/v2/server/web"
)

func InitSession() {
	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true

	//beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionName = enum.LocalSessionName
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = enum.SessionExpireTime
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = enum.SessionExpireTime

	//is, _ := pubMethod.PathExists(enum.SessionPath)
	//if !is {
	//	_ = os.Mkdir(enum.SessionPath, os.ModePerm)
	//}
	beego.BConfig.WebConfig.Session.SessionProviderConfig = enum.SessionPath
}
