/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/22 13:07
 ** @Author : yuebin
 ** @File : merchant_load_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/22 13:07
 ** @Software: GoLand
****************************************************/
package merchant

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type MerchantLoadInfo struct {
	Id          int
	Status      string
	MerchantUid string
	RoadUid     string
	LoadDate    string
	LoadAmount  float64
	UpdateTime  string
	CreateTime  string
}

const MERCHANT_LOAD_INFO = "merchant_load_info"

func GetMerchantLoadInfoByMap(params map[string]string) []MerchantLoadInfo {
	o := orm.NewOrm()
	qs := o.QueryTable(MERCHANT_LOAD_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	var merchantLoadList []MerchantLoadInfo
	_, err := qs.Limit(-11).All(&merchantLoadList)
	if err != nil {
		logs.Error("get merchant load info fail: ", err)
	}
	return merchantLoadList
}

func IsExistMerchantLoadByParams(params map[string]string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(MERCHANT_LOAD_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	return qs.Exist()
}
