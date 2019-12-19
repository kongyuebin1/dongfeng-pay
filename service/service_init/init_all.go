/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/19 17:48
 ** @Author : yuebin
 ** @File : init_all
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/19 17:48
 ** @Software: GoLand
****************************************************/
package service_init

import (
	_ "dongfeng-pay/service/message_queue"
	"dongfeng-pay/service/models"
	"dongfeng-pay/service/controller"
)

func InitAll() {
	//初始化mysql
	models.Init()
	controller.Init()
}
