/***************************************************
 ** @Desc : 向下游返回支付结果
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
	"gateway/conf"
	"gateway/message"
	"gateway/models/notify"
	"gateway/utils"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-stomp/stomp"
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
	if !notify.NotifyInfoExistByBankOrderId(bankOrderId) {
		logs.Error("该订单不存在回调内容，bankOrderId= " + bankOrderId)
		return
	}

	notifyInfo := notify.GetNotifyInfoByBankOrderId(bankOrderId)
	if notifyInfo.Status == "success" {
		logs.Info(fmt.Sprintf("该订单= %s,已经回调", bankOrderId))
		return
	}
	notifyInfo.Times += 1
	notifyInfo.UpdateTime = utils.GetBasicDateTime()

	req := httplib.Post(notifyInfo.Url)
	response, err := req.String()

	if err == nil && ("success" == response || "SUCCESS" == response) {
		if strings.Contains(strings.ToLower(response), "success") {
			notifyInfo.Status = "success"
			if notify.UpdateNotifyInfo(notifyInfo) {
				logs.Info("订单回调成功， bankOrderId=", bankOrderId)
			} else {
				logs.Error("订单回调成功，但是更新数据库失败， bankOrderId=", bankOrderId)
			}
		} else {
			logs.Notice("订单已经回调，商户已经收到了回调通知，但是返回值错误: ", response)
		}
	} else {
		if notifyInfo.Times > LimitTimes {
			logs.Notice(fmt.Sprintf("该订单= %s，已经超过了回调次数", bankOrderId))
		} else {
			minute := GetOrderNotifyMinute(notifyInfo.Times)
			logs.Info(fmt.Sprintf("bankOrderId = %s, 进行第 %d 次回调，本次延时时间为：%d", notifyInfo.BankOrderId, notifyInfo.Times, minute))
			task := OrderNotifyTask{Delay: time.NewTimer(time.Duration(minute) * time.Minute),
				MerchantOrderId: notifyInfo.MerchantOrderId, BankOrderId: notifyInfo.BankOrderId, FirstNotifyTime: notifyInfo.CreateTime,
				NotifyTimes: notifyInfo.Times, LimitTimes: LimitTimes, Status: notifyInfo.Status}
			go OrderNotifyTimer(task)
			if !notify.UpdateNotifyInfo(notifyInfo) {
				logs.Error("订单回调失败，数据库更新失败:" + bankOrderId)
			}
		}
	}
}

func GetOrderNotifyMinute(times int) int {
	cur := 0
	switch times {
	case 0:
		cur = 0
		break
	case 1:
		cur = 1
		break
	case 2:
		cur = 2
		break
	case 3:
		cur = 5
		break
	case 4:
		cur = 15
		break
	case 5:
		cur = 30
		break
	default:
		cur = 45
		break
	}
	return cur
}

func OrderNotifyTimer(task OrderNotifyTask) {
	for {
		select {
		case <-task.Delay.C:
			SendOrderNotify(task.BankOrderId)
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
	params["create_time__gte"] = utils.GetDateTimeBeforeHours(48)
	notifyList := notify.GetNotifyInfosNotSuccess(params)
	for _, nf := range notifyList {
		minute := GetOrderNotifyMinute(nf.Times)
		task := OrderNotifyTask{Delay: time.NewTimer(time.Duration(minute) * time.Minute),
			MerchantOrderId: nf.MerchantOrderId, BankOrderId: nf.BankOrderId, FirstNotifyTime: nf.CreateTime,
			NotifyTimes: nf.Times, LimitTimes: LimitTimes, Status: nf.Status}
		go OrderNotifyTimer(task)
	}
}

//创建订单回调消费者
func CreateOrderNotifyConsumer() {
	CreateOrderDelayQueue()
	//启动定时任务
	conn := message.GetActiveMQConn()
	if conn == nil {
		logs.Error("启动消息队列消费者失败....")
		os.Exit(1)
	}

	logs.Notice("订单回调消息队列启动成功......")
	orderNotify, err := conn.Subscribe(conf.MqOrderNotify, stomp.AckClient)
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
