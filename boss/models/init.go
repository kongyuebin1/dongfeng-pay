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
	"boss/models/accounts"
	"boss/models/agent"
	"boss/models/merchant"
	"boss/models/notify"
	"boss/models/order"
	"boss/models/payfor"
	"boss/models/road"
	"boss/models/system"
	"boss/models/user"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbHost, _ := web.AppConfig.String("mysql::dbhost")
	dbUser, _ := web.AppConfig.String("mysql::dbuser")
	dbPassword, _ := web.AppConfig.String("mysql::dbpasswd")
	dbBase, _ := web.AppConfig.String("mysql::dbbase")
	dbPort, _ := web.AppConfig.String("mysql::dbport")

	link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbBase)

	logs.Info("mysql init.....", link)

	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", link)
	orm.RegisterModel(new(user.UserInfo), new(system.MenuInfo), new(system.SecondMenuInfo),
		new(system.PowerInfo), new(system.RoleInfo), new(system.BankCardInfo), new(road.RoadInfo),
		new(road.RoadPoolInfo), new(agent.AgentInfo), new(merchant.MerchantInfo), new(merchant.MerchantDeployInfo),
		new(accounts.AccountInfo), new(accounts.AccountHistoryInfo), new(order.OrderInfo), new(order.OrderProfitInfo),
		new(order.OrderSettleInfo), new(notify.NotifyInfo), new(merchant.MerchantLoadInfo),
		new(payfor.PayforInfo))
}
