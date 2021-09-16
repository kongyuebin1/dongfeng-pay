/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/20 13:13
 ** @Author : yuebin
 ** @File : notify_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/20 13:13
 ** @Software: GoLand
****************************************************/
package notify

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type NotifyInfo struct {
	Id              int
	Type            string //订单-order，代付-payfor
	BankOrderId     string
	MerchantOrderId string
	Status          string
	Times           int
	Url             string
	Response        string
	UpdateTime      string
	CreateTime      string
}

const NOTIFYINFO = "notify_info"

func InsertNotifyInfo(notifyInfo NotifyInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&notifyInfo)
	if err != nil {
		logs.Error("insert notify fail：", err)
		return false
	}
	return true
}

func NotifyInfoExistByBankOrderId(bankOrderId string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(NOTIFYINFO).Filter("bank_order_id", bankOrderId).Exist()
	return exist
}

func GetNotifyInfoByBankOrderId(bankOrderId string) NotifyInfo {
	o := orm.NewOrm()
	var notifyInfo NotifyInfo
	_, err := o.QueryTable(NOTIFYINFO).Filter("bank_order_id", bankOrderId).All(&notifyInfo)
	if err != nil {
		logs.Error("get notify info by bankOrderId fail: ", err)
	}

	return notifyInfo
}

func GetNotifyInfosNotSuccess(params map[string]interface{}) []NotifyInfo {
	o := orm.NewOrm()
	var notifyInfoList []NotifyInfo
	qs := o.QueryTable(NOTIFYINFO)
	for k, v := range params {
		qs = qs.Filter(k, v)
	}
	qs = qs.Exclude("status", "success")
	_, err := qs.Limit(-1).All(&notifyInfoList)

	if err != nil {
		logs.Error("get notifyinfos fail: ", err)
	}

	return notifyInfoList
}

func UpdateNotifyInfo(notifyInfo NotifyInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&notifyInfo)
	if err != nil {
		logs.Error("update notify info fail: ", err)
		return false
	}
	return true
}
