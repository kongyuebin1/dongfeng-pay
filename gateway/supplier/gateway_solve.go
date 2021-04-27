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
	"context"
	"gateway/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

/*
* 插入支付订单记录和订单利润记录，保证一致性
 */
func InsertOrderAndOrderProfit(orderInfo models.OrderInfo, orderProfitInfo models.OrderProfitInfo) bool {
	o := orm.NewOrm()
	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		if _, err := txOrm.Insert(&orderInfo); err != nil {
			logs.Error("insert orderInfo fail: ", err)
			return err
		}
		if _, err := txOrm.Insert(&orderProfitInfo); err != nil {
			logs.Error("insert orderProfit fail: ", err)
			return err
		}

		return nil

	}); err != nil {
		return false
	}
	return true
}
