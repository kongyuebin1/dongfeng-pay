/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/20 1:35
 ** @Author : yuebin
 ** @File : order_notify
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/20 1:35
 ** @Software: GoLand
****************************************************/
package notify

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/go-stomp/stomp"
	"juhe/service/common"
	"juhe/service/message_queue"
	"juhe/service/models"
	"juhe/service/utils"
	"os"
	"strings"
	"time"
)

type OrderNotifyTask struct {
	Delay           *time.Timer
	MerchantOrderId string
	BankOrderId     string
	FirstNotifyTime string
	NotifyTimes     int
	LimitTimes      int
	Status          string //success-通知成功，其余的为待通知或者通知未完成
}

const (
	LimitTimes = 5 //最多通知5次
)

//给商户发送订单结果
func SendOrderNotify(bankOrderId string) {
	if !models.NotifyInfoExistByBankOrderId(bankOrderId) {
		logs.Error("该订单不存在回调内容，bankOrderId=", bankOrderId)
		return
	}

	notifyInfo := models.GetNotifyInfoByBankOrderId(bankOrderId)
	if notifyInfo.Status == "success" {
		logs.Info(fmt.Sprintf("该订单=%s,已经回调", bankOrderId))
		return
	}
	notifyInfo.Times += 1
	notifyInfo.UpdateTime = utils.GetBasicDateTime()

	req := httplib.Post(notifyInfo.Url)
	response, err := req.String()

	if err == nil {
		if strings.Contains(strings.ToLower(response), "success") {
			notifyInfo.Status = "success"
			if models.UpdateNotifyInfo(notifyInfo) {
				logs.Info("订单回调成功， bankOrderId=", bankOrderId)
			} else {
				logs.Error("订单回调成功，但是更新数据库失败， bankOrderId=", bankOrderId)
			}
		} else {
			logs.Notice("订单已经回调，商户已经收到了回调通知，但是返回值错误: ", response)
		}
	} else {
		if notifyInfo.Times > LimitTimes {
			logs.Notice(fmt.Sprintf("该订单=%s，已经超过了回调次数", bankOrderId))
		} else {
			minute := GetOrderNotifyMinute(notifyInfo.Times)
			task := OrderNotifyTask{Delay: time.NewTimer(time.Duration(minute) * time.Minute),
				MerchantOrderId: notifyInfo.MerchantOrderId, BankOrderId: notifyInfo.BankOrderId, FirstNotifyTime: notifyInfo.CreateTime,
				NotifyTimes: notifyInfo.Times, LimitTimes: LimitTimes, Status: notifyInfo.Status}
			logs.Info(fmt.Sprintf("订单bankOrderId=%s，已经是第：%d,回调", bankOrderId, notifyInfo.Times))
			go OrderNotifyTimer(task)
			if !models.UpdateNotifyInfo(notifyInfo) {
				logs.Error("订单回调失败，数据库更新失败:", bankOrderId)
			}
		}
	}
}

func GetOrderNotifyMinute(times int) int {
	cur := 1
	switch times {
	case 0:
		cur = 1
	case 2:
		cur = 2
	case 3:
		cur = 5
	case 4:
		cur = 15
	case 5:
		cur = 30
	}
	return cur
}

func OrderNotifyTimer(task OrderNotifyTask) {
	for {
		select {
		case <-task.Delay.C:
			SendOrderNotify(task.BankOrderId)
			task.Delay.Stop()
			return
			//70分钟没有执行该协程，那么退出协程
		case <-time.After(time.Minute * 70):
			logs.Notice("订单回调延时执行，70分钟没有执行")
			return
		}
	}
}

//读取一小时之内，未发送成功，并且还没有到达回调限制次数的记录读取，存入延迟队列
func CreateOrderDelayQueue() {
	params := make(map[string]interface{})
	params["times__lte"] = LimitTimes
	params["create_time__gte"] = utils.GetDateTimeBeforeHours(1)
	notifyList := models.GetNotifyInfosNotSuccess(params)
	for _, notify := range notifyList {
		minute := GetOrderNotifyMinute(notify.Times)
		task := OrderNotifyTask{Delay: time.NewTimer(time.Duration(minute) * time.Minute),
			MerchantOrderId: notify.MerchantOrderId, BankOrderId: notify.BankOrderId, FirstNotifyTime: notify.CreateTime,
			NotifyTimes: notify.Times, LimitTimes: LimitTimes, Status: notify.Status}
		go OrderNotifyTimer(task)
	}
}

//创建订单回调消费者
func CreateOrderNotifyConsumer() {
	CreateOrderDelayQueue()
	//启动定时任务
	conn := message_queue.GetActiveMQConn()
	if conn == nil {
		logs.Error("启动消息队列消费者失败....")
		os.Exit(1)
	}

	logs.Notice("订单回调消息队列启动成功......")
	orderNotify, err := conn.Subscribe(common.MqOrderNotify, stomp.AckClient)
	if err != nil {
		logs.Error("订阅订单回调失败......")
		os.Exit(1)
	}
	for {
		select {
		case v := <-orderNotify.C:
			if v != nil {
				bankOrderId := string(v.Body)
				go SendOrderNotify(bankOrderId)
				//应答，重要
				err := conn.Ack(v)
				if err != nil {
					logs.Error("消息应答失败！")
				}
			}
		}
	}
}
