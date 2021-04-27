/***************************************************
 ** @Desc : 处理网关模块的一些需要操作数据库的功能
 ** @Time : 2019/12/7 16:40
 ** @Author : yuebin
 ** @File : gateway_solve
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/12/7 16:40
 ** @Software: GoLand
****************************************************/
package controller

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"dongfeng/service/models"
)

/*
* 插入支付订单记录和订单利润记录，保证一致性
 */
func InsertOrderAndOrderProfit(orderInfo models.OrderInfo, orderProfitInfo models.OrderProfitInfo) bool {
	o := orm.NewOrm()
	o.Begin()

	defer func(interface{}) {
		if err := recover(); err != nil {
			o.Rollback()
		}
	}(o)

	if _, err := o.Insert(&orderInfo); err != nil {
		logs.Error("insert orderInfo fail: ", err)
		o.Rollback()
		return false
	}
	if _, err := o.Insert(&orderProfitInfo); err != nil {
		logs.Error("insert orderProfit fail: ", err)
		o.Rollback()
		return false
	}

	if err := o.Commit(); err != nil {
		logs.Error("insert order and orderProfit fail：", err)
	} else {
		logs.Info("插入order和orderProfit记录成功")
	}
	return true
}
