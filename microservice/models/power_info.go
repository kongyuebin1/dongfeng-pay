/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/8/28 17:59
 ** @Author : yuebin
 ** @File : power_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/28 17:59
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type PowerInfo struct {
	Id            int
	FirstMenuUid  string
	SecondMenuUid string
	SecondMenu    string
	PowerId       string
	PowerItem     string
	Creater       string
	Status        string
	CreateTime    string
	UpdateTime    string
}

const POWER_INFO = "power_info"

type PowerInfoSlice []PowerInfo

func (sm PowerInfoSlice) Len() int {
	return len(sm)
}

func (sm PowerInfoSlice) Swap(i, j int) {
	sm[i], sm[j] = sm[j], sm[i]
}

func (sm PowerInfoSlice) Less(i, j int) bool {
	return sm[i].SecondMenuUid < sm[j].SecondMenuUid
}

func PowerUidExists(powerUid string) bool {
	o := orm.NewOrm()
	exists := o.QueryTable(POWER_INFO).Filter("power_id", powerUid).Exist()
	return exists
}

func InsertPowerInfo(powerInfo PowerInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&powerInfo)
	if err != nil {
		logs.Error("insert power info fail: ", err)
		return false
	}
	return true
}

func GetPower() []PowerInfo {
	o := orm.NewOrm()
	var powerInfo []PowerInfo
	_, err := o.QueryTable(POWER_INFO).Limit(-1).All(&powerInfo)

	if err != nil {
		logs.Error("get power fail: ", err)
	}
	return powerInfo
}

func GetPowerById(powerId string) PowerInfo {
	o := orm.NewOrm()
	var powerInfo PowerInfo
	_, err := o.QueryTable(POWER_INFO).Filter("power_id", powerId).Limit(1).All(&powerInfo)
	if err != nil {
		logs.Error("get power by id fail: ", err)
	}
	return powerInfo
}

func GetPowerByIds(powerIds []string) []PowerInfo {
	var powerInfoList []PowerInfo
	for _, v := range powerIds {
		m := GetPowerById(v)
		powerInfoList = append(powerInfoList, m)
	}
	return powerInfoList
}

func GetPowerItemLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(POWER_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Limit(-1).Count()
	if err != nil {
		logs.Error("get power item len by map fail: ", err)
	}
	return int(cnt)
}

func GetPowerItemByMap(params map[string]string, displpay, offset int) []PowerInfo {
	o := orm.NewOrm()
	var powerItemList []PowerInfo
	qs := o.QueryTable(POWER_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	_, err := qs.Limit(displpay, offset).OrderBy("-update_time").All(&powerItemList)
	if err != nil {
		logs.Error("get power item by map fail: ", err)
	}
	return powerItemList
}

func DeletePowerItemByPowerID(powerID string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(POWER_INFO).Filter("power_id", powerID).Delete()
	if err != nil {
		logs.Error("delete power item by powerID fail: ", err)
		return false
	}
	return true
}

func DeletePowerBySecondUid(secondUid string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(POWER_INFO).Filter("second_menu_uid", secondUid).Delete()

	if err != nil {
		logs.Error("delete power by second menu uid fail: ", err)
		return false
	}
	return true
}
