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
	"gateway/conf"
	"gateway/models/accounts"
	"gateway/models/agent"
	"gateway/models/merchant"
	"gateway/models/notify"
	"gateway/models/order"
	"gateway/models/payfor"
	"gateway/models/road"
	"gateway/models/system"
	"gateway/models/user"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbHost := conf.DB_HOST
	dbUser := conf.DB_USER
	dbPassword := conf.DB_PASSWORD
	dbBase := conf.DB_BASE
	dbPort := conf.DB_PORT

	link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbBase)

	logs.Info("mysql init.....", link)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", link)
	orm.RegisterModel(new(user.UserInfo),
		new(system.MenuInfo),
		new(system.SecondMenuInfo),
		new(system.PowerInfo),
		new(system.RoleInfo),
		new(system.BankCardInfo),
		new(road.RoadInfo),
		new(road.RoadPoolInfo),
		new(agent.AgentInfo),
		new(merchant.MerchantInfo),
		new(merchant.MerchantDeployInfo),
		new(accounts.AccountInfo),
		new(accounts.AccountHistoryInfo),
		new(order.OrderInfo),
		new(order.OrderProfitInfo),
		new(order.OrderSettleInfo),
		new(notify.NotifyInfo),
		new(merchant.MerchantLoadInfo),
		new(payfor.PayforInfo))
}
