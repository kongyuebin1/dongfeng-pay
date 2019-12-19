/***************************************************
 ** @Desc : 处理代付问题
 ** @Time : 2019/11/29 14:07
 ** @Author : yuebin
 ** @File : payfor
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/29 14:07
 ** @Software: GoLand
****************************************************/
package pay_for

import (
	"github.com/astaxie/beego/logs"
	"juhe/service/common"
	"juhe/service/controller"
	"time"
)

func PayForInit() {
	payForIntervalTimer := time.NewTimer(time.Duration(common.PAYFOR_INTERVAL * time.Second))
	for {
		select {
		case <-payForIntervalTimer.C:
			logs.Info("代付小程序开始执行任务......")
			payForIntervalTimer = time.NewTimer(time.Duration(common.PAYFOR_INTERVAL * time.Minute))
			controller.SolvePayForConfirm()
			controller.SolvePayFor()
		case <-time.After(time.Duration(10 * time.Minute)):
			logs.Notice("代付小程序已经10分钟没有执行了.........")
		}
	}
}
