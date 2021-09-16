/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/9/8 12:09
 ** @Author : yuebin
 ** @File : road_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/9/8 12:09
 ** @Software: GoLand
****************************************************/
package road

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type RoadInfo struct {
	Id             int
	Status         string
	RoadName       string
	RoadUid        string
	Remark         string
	ProductName    string
	ProductUid     string
	PayType        string
	BasicFee       float64
	SettleFee      float64
	TotalLimit     float64
	TodayLimit     float64
	SingleMinLimit float64
	SingleMaxLimit float64
	StarHour       int
	EndHour        int
	Params         string
	TodayIncome    float64
	TotalIncome    float64
	TodayProfit    float64
	TotalProfit    float64
	Balance        float64
	RequestAll     int
	RequestSuccess int
	UpdateTime     string
	CreateTime     string
}

const ROAD_INFO = "road_info"

func GetRoadInfoByRoadUid(roadUid string) RoadInfo {
	o := orm.NewOrm()
	var roadInfo RoadInfo
	_, err := o.QueryTable(ROAD_INFO).Exclude("status", "delete").Filter("road_uid", roadUid).Limit(1).All(&roadInfo)
	if err != nil {
		logs.Error("get road info by road uid fail: ", err)
	}
	return roadInfo
}

func GetRoadInfosByRoadUids(roadUids []string) []RoadInfo {
	o := orm.NewOrm()
	var roadInfoList []RoadInfo
	_, err := o.QueryTable(ROAD_INFO).Filter("road_uid__in", roadUids).OrderBy("update_time").All(&roadInfoList)
	if err != nil {
		logs.Error("get roadInfos by roadUids fail: ", err)
	}
	return roadInfoList
}

func GetRoadInfoByName(roadName string) RoadInfo {
	o := orm.NewOrm()
	var roadInfo RoadInfo
	_, err := o.QueryTable(ROAD_INFO).Exclude("status", "delete").Filter("road_name", roadName).Limit(1).All(&roadInfo)
	if err != nil {
		logs.Error("get road info by name fail: ", err)
	}
	return roadInfo
}

func GetRoadLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(ROAD_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Exclude("status", "delete").Limit(-1).Count()
	if err != nil {
		logs.Error("get road len by map fail: ", err)
	}
	return int(cnt)
}

func GetRoadInfoByMap(params map[string]string, displayCount, offset int) []RoadInfo {
	o := orm.NewOrm()
	var roadInfoList []RoadInfo
	qs := o.QueryTable(ROAD_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	_, err := qs.Exclude("status", "delete").Limit(displayCount, offset).OrderBy("-update_time").All(&roadInfoList)
	if err != nil {
		logs.Error("get road info by map fail: ", err)
	}
	return roadInfoList
}

func GetAllRoad(params map[string]string) []RoadInfo {
	o := orm.NewOrm()
	var roadInfoList []RoadInfo
	qs := o.QueryTable(ROAD_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(-1).All(&roadInfoList)
	if err != nil {
		logs.Error("get all road fail: ", err)
	}
	return roadInfoList
}

func InsertRoadInfo(roadInfo RoadInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&roadInfo)

	if err != nil {
		logs.Error("insert road info fail: ", err)
		return false
	}
	return true
}

func RoadInfoExistByRoadUid(roadUid string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(ROAD_INFO).Filter("status", "active").Filter("road_uid", roadUid).Exist()

	return exist
}

func UpdateRoadInfo(roadInfo RoadInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&roadInfo)
	if err != nil {
		logs.Error("update road info fail: ", err)
		return false
	}
	return true
}

func DeleteRoadByRoadUid(roadUid string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(ROAD_INFO).Filter("road_uid", roadUid).Delete()
	if err != nil {
		logs.Error("delete road by road uid fail: ", err)
		return false
	}
	return true
}
