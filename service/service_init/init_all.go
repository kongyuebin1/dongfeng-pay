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
	_ "juhe/service/message_queue"
	"juhe/service/models"
	"juhe/service/controller"
)

func InitAll() {
	//初始化mysql
	models.Init()
	controller.Init()
}
