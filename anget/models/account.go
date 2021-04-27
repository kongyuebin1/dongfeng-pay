/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/16 11:11
 ** @Author : yuebin
 ** @File : account
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/16 11:11
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type AccountInfo struct {
	Id           int
	Status       string
	AccountUid   string
	AccountName  string
	Balance      float64 //账户总余额
	SettleAmount float64 //已经结算的金额
	LoanAmount   float64 //账户押款金额
	FreezeAmount float64 //账户冻结金额
	WaitAmount   float64 //待结算资金
	PayforAmount float64 //代付在途金额
	//AbleBalance  float64 //账户可用金额
	UpdateTime string
	CreateTime string
}

const ACCOUNT_INFO = "account_info"

func InsetAcount(account AccountInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&account)
	if err != nil {
		logs.Error("insert account fail: ", err)
		return false
	}
	return true
}

func GetAccountByUid(accountUid string) AccountInfo {
	o := orm.NewOrm()
	var account AccountInfo
	_, err := o.QueryTable(ACCOUNT_INFO).Filter("account_uid", accountUid).Limit(1).All(&account)
	if err != nil {
		logs.Error("get account by uid fail: ", err)
	}

	return account
}

func GetAccountLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(ACCOUNT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Limit(-1).OrderBy("-update_time").Count()
	if err != nil {
		logs.Error("get account len by map fail: ", err)
	}
	return int(cnt)
}

func GetAccountByMap(params map[string]string, displayCount, offset int) []AccountInfo {
	o := orm.NewOrm()
	var accountList []AccountInfo
	qs := o.QueryTable(ACCOUNT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&accountList)
	if err != nil {
		logs.Error("get account by map fail: ", err)
	}
	return accountList
}

func GetAllAccount() []AccountInfo {
	o := orm.NewOrm()
	var accountList []AccountInfo

	_, err := o.QueryTable(ACCOUNT_INFO).Limit(-1).All(&accountList)

	if err != nil {
		logs.Error("get all account fail: ", err)
	}

	return accountList
}

func UpdateAccount(account AccountInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&account)
	if err != nil {
		logs.Error("update account fail: ", err)
		return false
	}
	return true
}

func DeleteAccountByUid(accountUid string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(ACCOUNT_INFO).Filter("account_uid", accountUid).Delete()
	if err != nil {
		logs.Error("delete account fail: ", err)
		return false
	}
	return true
}
