/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/9/28 16:47
 ** @Author : yuebin
 ** @File : merchant_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/9/28 16:47
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
)

type MerchantInfo struct {
	Id                   int
	Status               string
	BelongAgentUid       string
	BelongAgentName      string
	MerchantName         string
	MerchantUid          string
	MerchantKey          string
	MerchantSecret       string
	LoginPassword        string
	LoginAccount         string
	AutoSettle           string
	AutoPayFor           string
	WhiteIps             string
	Remark               string
	SinglePayForRoadUid  string
	SinglePayForRoadName string
	RollPayForRoadCode   string
	RollPayForRoadName   string
	PayforFee            float64
	UpdateTime           string
	CreateTime           string
}

const MERCHANT_INFO = "merchant_info"

func IsExistByMerchantName(merchantName string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(MERCHANT_INFO).Filter("merchant_name", merchantName).Exist()

	return exist
}

func IsExistByMerchantUid(uid string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(MERCHANT_INFO).Filter("merchant_uid", uid).Exist()

	return exist
}

func IsExistMerchantByAgentUid(uid string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(MERCHANT_INFO).Filter("belong_agent_uid", uid).Exist()

	return exist
}

func IsExistByMerchantPhone(phone string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(MERCHANT_INFO).Filter("LoginAccount", phone).Exist()

	return exist
}

func GetMerchantByPhone(phone string) (m MerchantInfo) {
	o := orm.NewOrm()
	_, e := o.QueryTable(MERCHANT_INFO).Filter("LoginAccount", phone).Limit(1).All(&m)
	if e != nil {
		logs.Error("GetMerchantByPhone merchant fail: ", e)
	}
	return m
}

func InsertMerchantInfo(merchantInfo MerchantInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&merchantInfo)
	if err != nil {
		logs.Error("insert merchant fail: ", err)
		return false
	}
	return true
}

func GetMerchantLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(MERCHANT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Count()
	if err != nil {
		logs.Error("get merchant len by map fail: ", err)
	}
	return int(cnt)
}

func GetMerchantListByMap(params map[string]string, displayCount, offset int) []MerchantInfo {
	o := orm.NewOrm()
	qs := o.QueryTable(MERCHANT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	var merchantList []MerchantInfo
	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&merchantList)
	if err != nil {
		logs.Error("get merchant list by map fail: ", err)
	}
	return merchantList
}

func GetAllMerchant() []MerchantInfo {
	o := orm.NewOrm()
	var merchantList []MerchantInfo

	_, err := o.QueryTable(MERCHANT_INFO).Limit(-1).All(&merchantList)
	if err != nil {
		logs.Error("get all merchant fail：", err)
	}

	return merchantList
}

func GetMerchantByParams(params map[string]string, displayCount, offset int) []MerchantInfo {
	o := orm.NewOrm()
	var merchantList []MerchantInfo
	qs := o.QueryTable(MERCHANT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	qs.Limit(displayCount, offset).All(&merchantList)

	return merchantList
}

func GetMerchantLenByParams(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(MERCHANT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	cnt, err := qs.Limit(-1).Count()

	if err != nil {
		logs.Error("get merchant len by params fail: ", err)
	}
	return int(cnt)
}

func GetMerchantByUid(merchantUid string) MerchantInfo {
	o := orm.NewOrm()
	var merchantInfo MerchantInfo
	_, err := o.QueryTable(MERCHANT_INFO).Filter("merchant_uid", merchantUid).Limit(1).All(&merchantInfo)
	if err != nil {
		logs.Error("get merchant info fail: ", err)
	}
	return merchantInfo
}

func GetMerchantByPaykey(payKey string) MerchantInfo {
	o := orm.NewOrm()
	var merchantInfo MerchantInfo
	_, err := o.QueryTable(MERCHANT_INFO).Filter("merchant_key", payKey).Limit(1).All(&merchantInfo)
	if err != nil {
		logs.Error("get merchant by merchantKey fail: ", err)
	}
	return merchantInfo
}

func UpdateMerchant(merchantInfo MerchantInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&merchantInfo)

	if err != nil {
		logs.Error("update merchant fail: ", err)
		return false
	}

	return true
}

func DeleteMerchantByUid(merchantUid string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(MERCHANT_INFO).Filter("merchant_uid", merchantUid).Delete()
	if err != nil {
		logs.Error("delete merchant fail: ", err)
		return false
	}
	return true
}
