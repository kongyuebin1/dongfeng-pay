/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/8/9 13:48
 ** @Author : yuebin
 ** @File : init
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/9 13:48
 ** @Software: GoLand
****************************************************/
package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbHost := beego.AppConfig.String("mysql::dbhost")
	dbUser := beego.AppConfig.String("mysql::dbuser")
	dbPassword := beego.AppConfig.String("mysql::dbpasswd")
	dbBase := beego.AppConfig.String("mysql::dbbase")
	dbPort := beego.AppConfig.String("mysql::dbport")

	link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbBase)

	fmt.Println(link)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", link, 30, 30)
	orm.RegisterModel()
}
