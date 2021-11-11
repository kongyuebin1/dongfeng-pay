/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/7 11:52
 ** @Author : yuebin
 ** @File : merchant_deploy_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/7 11:52
 ** @Software: GoLand
****************************************************/
package merchant

import (
	"github.com/beego/beego/v2/client/orm"
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

func GetMerchantDeployByUidAndPayType(uid, payType string) MerchantDeployInfo {
	o := orm.NewOrm()
	var merchantDeployInfo MerchantDeployInfo
	_, err := o.QueryTable(MERCHANT_DEPLOY_INFO).Filter("merchant_uid", uid).Filter("pay_type", payType).Limit(1).All(&merchantDeployInfo)
	if err != nil {
		logs.Error("get merchant deploy by uid and paytype fail:", err)
	}
	return merchantDeployInfo
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
