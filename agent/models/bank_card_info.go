/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/9/6 10:19
 ** @Author : yuebin
 ** @File : bank_card_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/9/6 10:19
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type BankCardInfo struct {
	Id              int
	Uid             string
	UserName        string
	BankName        string
	BankCode        string
	BankAccountType string
	AccountName     string
	BankNo          string
	IdentifyCard    string
	CertificateNo   string
	PhoneNo         string
	BankAddress     string
	UpdateTime      string
	CreateTime      string
}

const BANK_CARD_INFO = "bank_card_info"

func InsertBankCardInfo(bankCardInfo BankCardInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&bankCardInfo)

	if err != nil {
		logs.Error("insert bank card info fail: ", err)
		return false
	}
	return true
}

func GetBankCardLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(BANK_CARD_INFO)
	for k, v := range params {
		qs = qs.Filter(k, v)
	}
	cnt, err := qs.Limit(-1).Count()
	if err != nil {
		logs.Error("get bank card len by map fail: ", err)
	}
	return int(cnt)
}

func GetBankCardByMap(params map[string]string, displayCount, offset int) []BankCardInfo {
	o := orm.NewOrm()
	var bankCardList []BankCardInfo
	qs := o.QueryTable(BANK_CARD_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&bankCardList)
	if err != nil {
		logs.Error("get bank card by map fail: ", err)
	}
	return bankCardList
}

func GetBankCardByUid(uid string) BankCardInfo {
	o := orm.NewOrm()
	var bankCardInfo BankCardInfo
	_, err := o.QueryTable(bankCardInfo).Filter("uid", uid).Limit(1).All(&bankCardInfo)
	if err != nil {
		logs.Error("get bank card by uid fail: ", err)
	}

	return bankCardInfo
}

func DeleteBankCardByUid(uid string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(BANK_CARD_INFO).Filter("uid", uid).Delete()

	if err != nil {
		logs.Error("delete bank card by uid fail: ", err)
		return false
	}
	return true
}

func UpdateBankCard(bankCard BankCardInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&bankCard)
	if err != nil {
		logs.Error("update bank card fail: ", err)
		return false
	}
	return true
}
