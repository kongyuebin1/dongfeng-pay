package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"legend/models/fast"
	"legend/models/legend"
	"os"
)

/**
** 链接数据库，注册已经存在的数据表，进行orm映射操作
 */
func init() {
	initLegend()
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

	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxIdleConns("default", 30)

	orm.RegisterModel(new(fast.MerchantInfo))
	orm.RegisterModel(new(fast.MerchantDeployInfo))
	orm.RegisterModel(new(fast.BankCardInfo))
	orm.RegisterModel(new(fast.AccountInfo))
	orm.RegisterModel(new(fast.OrderInfo))

	orm.RegisterModel(new(legend.AnyMoney))
	orm.RegisterModel(new(legend.FixMoney))
	orm.RegisterModel(new(legend.FixPresent))
	orm.RegisterModel(new(legend.ScalePresent))
	orm.RegisterModel(new(legend.ScaleTemplate))
	orm.RegisterModel(new(legend.Group))

	logs.Info("init legend success ......")

}
