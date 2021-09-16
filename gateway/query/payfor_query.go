/***************************************************
 ** @Desc : 处理代付查询功能
 ** @Time : 2019/12/3 15:07
 ** @Author : yuebin
 ** @File : pay_for_query
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/12/3 15:07
 ** @Software: GoLand
****************************************************/
package query

import (
	"fmt"
	"gateway/conf"
	"gateway/message"
	"gateway/models/payfor"
	"gateway/models/road"
	"gateway/pay_for"
	"gateway/supplier/third_party"
	"gateway/utils"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-stomp/stomp"
	"os"
	"time"
)

type PayForQueryTask struct {
	Delay           *time.Timer
	MerchantOrderId string
	BankOrderId     string
	FirstNotifyTime string
	QueryTimes      int
	LimitTimes      int
	Status          string
}

const (
	PayForLimitTimes    = 12 //最多查询次数
	PayForQueryInterval = 5  //时间间隔为5分钟
)

func PayForQueryTimer(task PayForQueryTask) {
	for {
		select {
		case <-task.Delay.C:
			PayForSupplier(task)
			task.Delay.Stop()
			return
			//70分钟没有执行该协程，那么退出协程
		case <-time.After(time.Minute * 70):
			return
		}
	}
}

func PayForSupplier(task PayForQueryTask) {
	logs.Info(fmt.Sprintf("执行代付查询任务：%+v", task))
	payFor := payfor.GetPayForByBankOrderId(task.BankOrderId)
	roadInfo := road.GetRoadInfoByRoadUid(payFor.RoadUid)
	supplier := third_party.GetPaySupplierByCode(roadInfo.ProductUid)
	if supplier == nil {
		logs.Error("代付查询返回supplier为空")
		return
	}
	res, _ := supplier.PayForQuery(payFor)
	if res == conf.PAYFOR_SUCCESS {
		//代付成功了
		pay_for.PayForSuccess(payFor)
	} else if res == conf.PAYFOR_FAIL {
		//代付失败
		pay_for.PayForFail(payFor)
	} else if res == conf.PAYFOR_BANKING {
		//银行处理中，那么就继续执行查询，直到次数超过最大次数
		if task.QueryTimes <= task.LimitTimes {
			task.QueryTimes += 1
			task.Delay = time.NewTimer(time.Duration(PayForQueryInterval) * time.Minute)
			go PayForQueryTimer(task)
		} else {
			logs.Info(fmt.Sprintf("该代付订单已经超过最大查询次数，bankOrderId = %s", task.BankOrderId))
		}
	}
}

func payForQueryConsumer(bankOrderId string) {
	exist := payfor.IsExistPayForByBankOrderId(bankOrderId)
	if !exist {
		logs.Error(fmt.Sprintf("代付记录不存在，bankOrderId = %s", bankOrderId))
		return
	}

	payFor := payfor.GetPayForByBankOrderId(bankOrderId)

	if payFor.Status != conf.PAYFOR_BANKING {
		logs.Info(fmt.Sprintf("代付状态不是银行处理中，不需要去查询，bankOrderId = %s", bankOrderId))
		return
	}

	payForQueryTask := PayForQueryTask{Delay: time.NewTimer(time.Duration(PayForQueryInterval) * time.Minute), MerchantOrderId: payFor.MerchantOrderId,
		BankOrderId: payFor.BankOrderId, FirstNotifyTime: utils.GetBasicDateTime(), QueryTimes: 1, LimitTimes: PayForLimitTimes, Status: payFor.Status}

	go PayForQueryTimer(payForQueryTask)
}

/*
* 创建代付查询的消费者
 */
func CreatePayForQueryConsumer() {
	//启动定时任务
	conn := message.GetActiveMQConn()
	if conn == nil {
		logs.Error("启动消息队列消费者失败....")
		os.Exit(1)
	}

	logs.Notice("代付查询消费启动成功......")

	payForQuery, err := conn.Subscribe(conf.MQ_PAYFOR_QUERY, stomp.AckClient)
	if err != nil {
		logs.Error("订阅代付查询失败......")
		os.Exit(1)
	}

	for {
		select {
		case v := <-payForQuery.C:
			if v != nil {
				bankOrderId := string(v.Body)
				go payForQueryConsumer(bankOrderId)
				//应答，重要
				err := conn.Ack(v)
				if err != nil {
					logs.Error("消息应答失败！")
				}
			}
		}
	}
}
