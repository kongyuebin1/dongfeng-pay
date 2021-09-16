/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/21 15:53
 ** @Author : yuebin
 ** @File : send_message
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/21 15:53
 ** @Software: GoLand
****************************************************/
package message

import (
	"github.com/beego/beego/v2/core/logs"
	"os"
)

func SendMessage(topic, message string) {

	conn := GetActiveMQConn()

	if conn == nil {
		logs.Error("send message get Active mq fail")
		os.Exit(1)
	}

	err := conn.Send(topic, "text/plain", []byte(message))

	if err != nil {
		logs.Error("发送消息给activeMQ失败, message=", message)
	} else {
		logs.Info("发送消息给activeMQ成功，message=", message)
	}
}
