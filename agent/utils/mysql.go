/***************************************************
 ** @Desc : This file for 配置数据库连接
 ** @Time : 2018-12-22 13:55:26
 ** @Author : Joker
 ** @File : init_database.go
 ** @Last Modified by : Joker
 ** @Last Modified time:2018-12-22 13:55:26
 ** @Software: GoLand
****************************************************/
package utils

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

//初始化数据连接
func InitDatabase() bool {
	//读取配置文件，设置数据库参数
	dbType, _ := beego.AppConfig.String("db_type")
	dbAlias, _ := beego.AppConfig.String(dbType + "::db_alias")
	dbName, _ := beego.AppConfig.String(dbType + "::db_name")
	dbUser, _ := beego.AppConfig.String(dbType + "::db_user")
	dbPwd, _ := beego.AppConfig.String(dbType + "::db_pwd")
	dbHost, _ := beego.AppConfig.String(dbType + "::db_host")
	dbPort, _ := beego.AppConfig.String(dbType + "::db_port")

	var err error
	switch dbType {
	case "sqlite3":
		err = orm.RegisterDataBase(dbAlias, dbType, dbName)
	case "mysql":
		dbCharset, _ := beego.AppConfig.String(dbType + "::db_charset")
		err = orm.RegisterDriver(dbType, orm.DRMySQL)
		err = orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
			dbPort+")/"+dbName+"?charset="+dbCharset)
	}

	if err != nil {
		return false
	}
	return true
}
