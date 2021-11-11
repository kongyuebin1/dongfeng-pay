/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/19 14:56
 ** @Author : yuebin
 ** @File : account_history_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/19 14:56
 ** @Software: GoLand
****************************************************/
package accounts

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type AccountHistoryInfo struct {
	Id          int
	AccountUid  string
	AccountName string
	Type        string
	Amount      float64
	Balance     float64
	UpdateTime  string
	CreateTime  string
}

const ACCOUNT_HISTORY_INFO = "account_history_info"

func GetAccountHistoryLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(ACCOUNT_HISTORY_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Limit(-1).Count()
	if err != nil {
		logs.Error("get account history len by map fail: ", err)
	}
	return int(cnt)
}

func GetAccountHistoryByMap(params map[string]string, displayCount, offset int) []AccountHistoryInfo {
	o := orm.NewOrm()
	qs := o.QueryTable(ACCOUNT_HISTORY_INFO)
	var accountHistoryList []AccountHistoryInfo
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&accountHistoryList)
	if err != nil {
		logs.Error("get account history by map fail: ", err)
	}
	return accountHistoryList
}
