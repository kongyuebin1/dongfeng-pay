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
	"juhe/service/conf"
)

func Init() {
	dbHost := conf.DB_HOST
	dbUser := conf.DB_USER
	dbPassword := conf.DB_PASSWORD
	dbBase := conf.DB_BASE
	dbPort := conf.DB_PORT

	link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbBase)

	beego.Info("mysql init.....", link)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", link, 30, 30)
	orm.RegisterModel(new(UserInfo), new(MenuInfo), new(SecondMenuInfo),
		new(PowerInfo), new(RoleInfo), new(BankCardInfo), new(RoadInfo),
		new(RoadPoolInfo), new(AgentInfo), new(MerchantInfo), new(MerchantDeployInfo),
		new(AccountInfo), new(AccountHistoryInfo), new(OrderInfo), new(OrderProfitInfo),
		new(OrderSettleInfo), new(NotifyInfo), new(MerchantLoadInfo),
		new(PayforInfo))
}
