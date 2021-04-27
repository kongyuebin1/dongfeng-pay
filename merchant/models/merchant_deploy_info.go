/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/7 11:52
 ** @Author : yuebin
 ** @File : merchant_deploy_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/7 11:52
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
)

type MerchantDeployInfo struct {
	Id                     int
	Status                 string
	MerchantUid            string
	PayType                string
	SingleRoadUid          string
	SingleRoadName         string
	SingleRoadPlatformRate float64
	SingleRoadAgentRate    float64
	RollRoadCode           string
	RollRoadName           string
	RollRoadPlatformRate   float64
	RollRoadAgentRate      float64
	IsLoan                 string
	LoanRate               float64
	LoanDays               int
	UnfreezeHour           int
	WaitUnfreezeAmount     float64
	LoanAmount             float64
	UpdateTime             string
	CreateTime             string
}

const MERCHANT_DEPLOY_INFO = "merchant_deploy_info"

func InsertMerchantDeployInfo(merchantDeployInfo MerchantDeployInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&merchantDeployInfo)
	if err != nil {
		logs.Error("insert merchant deploy info fail: ", err)
		return false
	}
	return true
}

func IsExistByUidAndPayType(uid, payType string) bool {
	o := orm.NewOrm()
	isEixst := o.QueryTable(MERCHANT_DEPLOY_INFO).Filter("merchant_uid", uid).Filter("pay_type", payType).Exist()
	return isEixst
}

func GetMerchantDeployByUidAndPayType(uid, payType string) MerchantDeployInfo {
	o := orm.NewOrm()
	var merchantDeployInfo MerchantDeployInfo
	_, err := o.QueryTable(MERCHANT_DEPLOY_INFO).Filter("merchant_uid", uid).Filter("pay_type", payType).Limit(1).All(&merchantDeployInfo)
	if err != nil {
		logs.Error("get merchant deploy by uid and paytype fail:", err)
	}
	return merchantDeployInfo
}

func GetMerchantDeployByUid(uid string) (ms []MerchantDeployInfo) {
	o := orm.NewOrm()
	_, err := o.QueryTable(MERCHANT_DEPLOY_INFO).Filter("merchant_uid", uid).All(&ms)
	if err != nil {
		logs.Error("get merchant deploy by uid fail:", err)
	}
	return ms
}

func GetMerchantDeployByHour(hour int) []MerchantDeployInfo {
	o := orm.NewOrm()
	var merchantDeployList []MerchantDeployInfo
	_, err := o.QueryTable(MERCHANT_DEPLOY_INFO).Filter("unfreeze_hour", hour).Filter("status", "active").Limit(-1).All(&merchantDeployList)
	if err != nil {
		logs.Error("get merchant deploy list fail: ", err)
	}

	return merchantDeployList
}
func DeleteMerchantDeployByUidAndPayType(uid, payType string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(MERCHANT_DEPLOY_INFO).Filter("merchant_uid", uid).Filter("pay_type", payType).Delete()
	if err != nil {
		logs.Error("delete merchant deploy by uid and payType fail: ", err)
		return false
	}
	return true
}

func UpdateMerchantDeploy(merchantDeploy MerchantDeployInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&merchantDeploy)
	if err != nil {
		logs.Error("update merchant deploy fail: ", err)
		return false
	}
	return true
}

func GetMerchantDeployLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(MERCHANT_DEPLOY_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Count()
	if err != nil {
		logs.Error("get merchant deploy len by map fail: ", err)
	}
	return int(cnt)
}

func GetMerchantDeployListByMap(params map[string]string, displayCount, offset int) (md []MerchantDeployInfo) {
	o := orm.NewOrm()
	qs := o.QueryTable(MERCHANT_DEPLOY_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&md)
	if err != nil {
		logs.Error("get merchant deploy list by map fail: ", err)
	}
	return md
}
