/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/25 14:32
 ** @Author : yuebin
 ** @File : payfor_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/25 14:32
 ** @Software: GoLand
****************************************************/
package payfor

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type PayforInfo struct {
	Id                 int
	PayforUid          string
	MerchantUid        string
	MerchantName       string
	MerchantOrderId    string
	BankOrderId        string
	BankTransId        string
	RoadUid            string
	RoadName           string
	RollPoolCode       string
	RollPoolName       string
	PayforFee          float64
	PayforAmount       float64
	PayforTotalAmount  float64
	BankCode           string
	BankName           string
	BankAccountName    string
	BankAccountNo      string
	BankAccountType    string
	Country            string
	City               string
	Ares               string
	BankAccountAddress string
	PhoneNo            string
	GiveType           string
	Type               string
	NotifyUrl          string
	Status             string
	IsSend             string
	RequestTime        string
	ResponseTime       string
	ResponseContent    string
	Remark             string
	CreateTime         string
	UpdateTime         string
}

const PAYFORINFO = "payfor_info"

func InsertPayfor(payFor PayforInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&payFor)
	if err != nil {
		logs.Error("insert payfor fail: ", err)
		return false
	}
	return true
}

func IsExistPayForByBankOrderId(bankOrderId string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(PAYFORINFO).Filter("bank_order_id", bankOrderId).Exist()

	return exist
}

func IsExistPayForByMerchantOrderId(merchantOrderId string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(PAYFORINFO).Filter("merchant_order_id", merchantOrderId).Exist()

	return exist
}

func GetPayForByBankOrderId(bankOrderId string) PayforInfo {
	o := orm.NewOrm()
	var payFor PayforInfo
	_, err := o.QueryTable(PAYFORINFO).Filter("bank_order_id", bankOrderId).Limit(1).All(&payFor)

	if err != nil {
		logs.Error("get pay for by bank_order_id fail: ", err)
	}

	return payFor
}

func GetPayForByMerchantOrderId(merchantOrderId string) PayforInfo {
	o := orm.NewOrm()
	var payFor PayforInfo

	_, err := o.QueryTable(PAYFORINFO).Filter("merchant_order_id", merchantOrderId).Limit(1).All(&payFor)

	if err != nil {
		logs.Error("fail: ", err)
	}

	return payFor
}

func UpdatePayFor(payFor PayforInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&payFor)

	if err != nil {
		logs.Error("update pay for fail：", err)
		return false
	}

	return true
}
