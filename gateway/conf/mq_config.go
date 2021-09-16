/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/6 11:37
 ** @Author : yuebin
 ** @File : mq_config
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/6 11:37
 ** @Software: GoLand
****************************************************/
package conf

import "net"

const (
	mqHost = "127.0.0.1"
	mqPort = "61613"

	MqOrderQuery    = "order_query"
	MQ_PAYFOR_QUERY = "payfor_query"
	MqOrderNotify   = "order_notify"
)

func GetMQAddress() string {
	return net.JoinHostPort(mqHost, mqPort)
}
