/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/9/9 16:35
 ** @Author : yuebin
 ** @File : road_pool_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/9/9 16:35
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type RoadPoolInfo struct {
	Id           int
	Status       string
	RoadPoolName string
	RoadPoolCode string
	RoadUidPool  string
	UpdateTime   string
	CreateTime   string
}

const ROAD_POOL_INFO = "road_pool_info"

func InsertRoadPool(roadPool RoadPoolInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&roadPool)
	if err != nil {
		logs.Error("insert road pool fail: ", err)
		return false
	}
	return true
}

func GetRoadPoolLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(ROAD_POOL_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Limit(-1).Count()
	if err != nil {
		logs.Error("get road pool len by map fail: ", err)
	}
	return int(cnt)
}

func GetRoadPoolByMap(params map[string]string, displayCount, offset int) []RoadPoolInfo {
	o := orm.NewOrm()
	var roadPoolList []RoadPoolInfo
	qs := o.QueryTable(ROAD_POOL_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&roadPoolList)
	if err != nil {
		logs.Error("get road pool by map fail: ", err)
	}
	return roadPoolList
}

func GetRoadPoolByRoadPoolCode(roadPoolCode string) RoadPoolInfo {
	o := orm.NewOrm()
	var roadPoolInfo RoadPoolInfo
	_, err := o.QueryTable(ROAD_POOL_INFO).Filter("road_pool_code", roadPoolCode).Limit(1).All(&roadPoolInfo)

	if err != nil {
		logs.Error("get road pool info by road pool code fail: ", err)
	}

	return roadPoolInfo
}

func GetAllRollPool(params map[string]string) []RoadPoolInfo {
	o := orm.NewOrm()
	var roadPoolList []RoadPoolInfo
	qs := o.QueryTable(ROAD_POOL_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(-1).All(&roadPoolList)
	if err != nil {
		logs.Error("get all roll pool fail: ", err)
	}
	return roadPoolList
}

func GetRoadPoolByName(roadPoolName string) RoadPoolInfo {
	o := orm.NewOrm()
	var roadPoolInfo RoadPoolInfo
	_, err := o.QueryTable(ROAD_POOL_INFO).Filter("road_pool_name", roadPoolName).Limit(1).All(&roadPoolInfo)
	if err != nil {
		logs.Error("get road pool by name fail: ", err)
	}
	return roadPoolInfo
}

func DeleteRoadPoolByCode(roadPoolCode string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(ROAD_POOL_INFO).Filter("road_pool_code", roadPoolCode).Delete()
	if err != nil {
		logs.Error("delete road pool by code fail: ", err)
		return false
	}
	return true
}

func UpdateRoadPool(roadPool RoadPoolInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&roadPool)

	if err != nil {
		logs.Error("update road pool fail: ", err)
		return false
	}
	return true
}
