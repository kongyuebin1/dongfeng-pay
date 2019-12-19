/***************************************************
 ** @Desc : 将待结算的订单金额，加入账户可用金额中
 ** @Time : 2019/11/21 23:43
 ** @Author : yuebin
 ** @File : settle
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/21 23:43
 ** @Software: GoLand
****************************************************/
package order_settle

import (
	"github.com/astaxie/beego/logs"
	"juhe/service/controller"
	"time"
)

const (
	SettleInterval = 5  //隔多少分钟进行结算
	OneMinute      = 15 //每隔15分钟，进行扫码，看有没有隔天押款金额
)

func OrderSettleInit() {
	//每隔5分钟，巡查有没有可以进行结算的订单
	go func() {
		settleTimer := time.NewTimer(time.Duration(SettleInterval) * time.Minute)
		oneMinuteTimer := time.NewTimer(time.Duration(OneMinute) * time.Minute)
		for {
			select {
			case <-settleTimer.C:
				settleTimer = time.NewTimer(time.Duration(SettleInterval) * time.Minute)
				logs.Info("开始对商户进行支付订单结算>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
				controller.OrderSettle()
			case <-oneMinuteTimer.C:
				oneMinuteTimer = time.NewTimer(time.Duration(OneMinute) * time.Minute)
				logs.Info("开始执行商户的解款操作>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
				controller.MerchantLoadSolve()
			}
		}
	}()
}
