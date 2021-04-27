package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"legend/models/fast"
	"os"
)

/**
** 链接数据库，注册已经存在的数据表，进行orm映射操作
 */
func init() {
	//initFastPay()
	initLegend()
}

/**
** 初始化快付支付系统的mysql数据库
 */
func initFastPay() {
	dbType, _ := web.AppConfig.String("dbtype")
	mysqlHost, _ := web.AppConfig.String("fast::host")
	mysqlPort, _ := web.AppConfig.String("fast::port")
	mysqlUserName, _ := web.AppConfig.String("fast::username")
	mysqlPassword, _ := web.AppConfig.String("fast::password")
	mysqlDbName, _ := web.AppConfig.String("fast::dbname")

	logs.Info("host:%s, port:%s, usreName:%s, password:%s, dbname:%s, dbType:%s", mysqlHost, mysqlPort,
		mysqlUserName, mysqlPassword, mysqlDbName, dbType)

	pStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",
		mysqlUserName, mysqlPassword, mysqlHost, mysqlPort, mysqlDbName)

	if err := orm.RegisterDataBase("default", dbType, pStr); err != nil {
		logs.Error("init fast fail：%s", err)
		os.Exit(1)
	}
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxIdleConns("default", 30)

	orm.RegisterModel(new(fast.RpUserInfo))
	orm.RegisterModel(new(fast.RpUserPayConfig))
	orm.RegisterModel(new(fast.RpUserBankAccount))
	orm.RegisterModel(new(fast.RpAccount))

	logs.Info("init fast success ......")
}

/**
** 初始化传奇支付系统的mysql数据库
 */
func initLegend() {
	dbType, _ := web.AppConfig.String("dbtype")
	mysqlHost, _ := web.AppConfig.String("legend::host")
	mysqlPort, _ := web.AppConfig.String("legend::port")
	mysqlUserName, _ := web.AppConfig.String("legend::username")
	mysqlPassword, _ := web.AppConfig.String("legend::password")
	mysqlDbName, _ := web.AppConfig.String("legend::dbname")

	logs.Info("host:%s, port:%s, usreName:%s, password:%s, dbname:%s", mysqlHost, mysqlPort,
		mysqlUserName, mysqlPassword, mysqlDbName)

	pStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",
		mysqlUserName, mysqlPassword, mysqlHost, mysqlPort, mysqlDbName)

	if err := orm.RegisterDataBase("default", dbType, pStr); err != nil {
		logs.Error("init legend fail：%s", err)
		os.Exit(1)
	}

	logs.Info("init legend success ......")

	orm.RegisterModel()
}
