/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/22 23:02
 ** @Author : yuebin
 ** @File : order_query
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/22 23:02
 ** @Software: GoLand
****************************************************/
package query

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/go-stomp/stomp"
	"dongfeng-pay/service/common"
	"dongfeng-pay/service/controller"
	"dongfeng-pay/service/message_queue"
	"dongfeng-pay/service/models"
	"os"
	"time"
)

type OrderQueryTask struct {
	BankOrderId     string
	OrderQueryTimer *time.Timer
	Times           int
}

const (
	DelayTime  = 5 //延时时间为5分钟查询一次
	LimitTimes = 5 //最多查询5次
)

/*
** 该接口是查询上游的订单
 */
func solveSupplierOrderQuery(task OrderQueryTask) {
	bankOrderId := task.BankOrderId
	orderInfo := models.GetOrderByBankOrderId(bankOrderId)
	if orderInfo.BankOrderId == "" || len(orderInfo.BankOrderId) == 0 {
		logs.Error("不存在这样的订单，订单查询结束")
		return
	}
	if orderInfo.Status != "" && orderInfo.Status != "wait" {
		logs.Error(fmt.Sprintf("该订单=%s，已经处理完毕，", bankOrderId))
		return
	}
	supplierCode := orderInfo.PayProductCode
	supplier := controller.GetPaySupplierByCode(supplierCode)
	flag := supplier.PayQuery(orderInfo)
	if flag {
		logs.Info("订单查询成功， bankOrderId：", bankOrderId)
	} else {
		if task.Times <= LimitTimes {
			task.Times += 1
			task.OrderQueryTimer = time.NewTimer(time.Duration(5) * time.Minute)
			DelayOrderQueryQueue(task)
		} else {
			logs.Notice(fmt.Sprintf("订单id=%s, 已经查询超过次数"))
		}
	}
}

/*
* 延时队列
 */
func DelayOrderQueryQueue(task OrderQueryTask) {
	for {
		select {
		case <-task.OrderQueryTimer.C:
			logs.Info(fmt.Sprintf("订单id=%s,执行第：%d 次查询", task.BankOrderId, task.Times))
			solveSupplierOrderQuery(task)
			return
		case <-time.After(time.Duration(2*DelayTime) * time.Minute):
			return
		}
	}
}

/*
** 启动消息订单查询的消息队列消费者
 */
func CreateSupplierOrderQueryCuConsumer() {
	conn := message_queue.GetActiveMQConn()
	if conn == nil {
		logs.Error("supplier order query consumer fail")
		os.Exit(1)
	}
	logs.Notice("启动订单查询的消费者成功.....")
	orderQuerySub, _ := conn.Subscribe(common.MqOrderQuery, stomp.AckClient)

	for {
		select {
		case v := <-orderQuerySub.C:
			if v != nil {
				bankOrderId := string(v.Body)
				logs.Info("消费者正在处理订单查询： " + bankOrderId)
				task := OrderQueryTask{BankOrderId: bankOrderId, OrderQueryTimer: time.NewTimer(time.Second * 1), Times: 1}
				DelayOrderQueryQueue(task)
				//应答，重要
				err := conn.Ack(v)
				if err != nil {
					logs.Error("消息应答失败！")
				}
			}
		}
	}
}
