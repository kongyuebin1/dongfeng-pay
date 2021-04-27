/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/6 11:43
 ** @Author : yuebin
 ** @File : active_mq
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/6 11:43
 ** @Software: GoLand
****************************************************/
package message_queue

import (
	"boss/common"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-stomp/stomp"
	"os"
	"time"
)

//解决第一个问题的代码
var activeConn *stomp.Conn

var options = []func(*stomp.Conn) error{
	//设置读写超时，超时时间为1个小时
	stomp.ConnOpt.HeartBeat(7200*time.Second, 7200*time.Second),
	stomp.ConnOpt.HeartBeatError(360 * time.Second),
}

func init() {
	address := common.GetMQAddress()

	conn, err := stomp.Dial("tcp", address, options...)
	if err != nil {
		logs.Error("链接active mq 失败：", err.Error())
		os.Exit(1)
	}

	activeConn = conn
}

func GetActiveMQConn() *stomp.Conn {
	return activeConn
}


