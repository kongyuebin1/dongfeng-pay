/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/8/26 9:33
 ** @Author : yuebin
 ** @File : second_menu_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/26 9:33
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

const SECOND_MENU_INFO = "second_menu_info"

type SecondMenuInfo struct {
	Id             int
	FirstMenuOrder int
	FirstMenuUid   string
	FirstMenu      string
	MenuOrder      int
	SecondMenuUid  string
	SecondMenu     string
	SecondRouter   string
	Creater        string
	Status         string
	CreateTime     string
	UpdateTime     string
}

type SecondMenuSlice []SecondMenuInfo

func (sm SecondMenuSlice) Len() int {
	return len(sm)
}

func (sm SecondMenuSlice) Swap(i, j int) {
	sm[i], sm[j] = sm[j], sm[i]
}

func (sm SecondMenuSlice) Less(i, j int) bool {
	if sm[i].FirstMenuOrder == sm[j].FirstMenuOrder {
		return sm[i].MenuOrder < sm[j].MenuOrder
	}
	return sm[i].FirstMenuOrder < sm[j].FirstMenuOrder
}

func GetSecondMenuLen() int {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(SECOND_MENU_INFO).Count()
	if err != nil {
		logs.Error("get second meun len fail: ", err)
	}
	return int(cnt)
}

func GetSecondMenuInfoByMenuOrder(menuOrder int, firstMenuUid string) SecondMenuInfo {
	o := orm.NewOrm()
	var secondMenuInfo SecondMenuInfo
	_, err := o.QueryTable(SECOND_MENU_INFO).Filter("first_menu_uid", firstMenuUid).Filter("menu_order", menuOrder).Limit(1).All(&secondMenuInfo)
	if err != nil {
		logs.Error("get second menu info by menu order fail: ", err)
	}
	return secondMenuInfo
}

func GetSecondMenuLenByFirstMenuUid(firstMenuUid string) int {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(SECOND_MENU_INFO).Filter("first_menu_uid", firstMenuUid).Count()
	if err != nil {
		logs.Error("get second menu len by first menu uid fail: ", err)
	}
	return int(cnt)
}

func GetSecondMenuList() []SecondMenuInfo {
	o := orm.NewOrm()
	var secondMenuList []SecondMenuInfo
	_, err := o.QueryTable(SECOND_MENU_INFO).Limit(-1).OrderBy("-update_time").All(&secondMenuList)
	if err != nil {
		logs.Error("get second menu list fail: ", err)
	}
	return secondMenuList
}

func GetSecondMenuInfoBySecondMenuUid(secondMenuUid string) SecondMenuInfo {
	o := orm.NewOrm()
	var secondMenuInfo SecondMenuInfo
	_, err := o.QueryTable(SECOND_MENU_INFO).Filter("second_menu_uid", secondMenuUid).Limit(1).All(&secondMenuInfo)
	if err != nil {
		logs.Error("get scond menu info by second menu uid fail: ", err)
	}
	return secondMenuInfo
}

func GetSecondMenuInfoBySecondMenuUids(secondMenuUids []string) []SecondMenuInfo {
	secondMenuInfoList := make([]SecondMenuInfo, 0)
	for _, v := range secondMenuUids {
		sm := GetSecondMenuInfoBySecondMenuUid(v)
		secondMenuInfoList = append(secondMenuInfoList, sm)
	}
	return secondMenuInfoList
}

func GetSecondMenuListByFirstMenuUid(firstMenuUid string) []SecondMenuInfo {
	o := orm.NewOrm()
	var secondMenuList []SecondMenuInfo
	_, err := o.QueryTable(SECOND_MENU_INFO).Filter("first_menu_uid", firstMenuUid).Limit(-1).OrderBy("-update_time").All(&secondMenuList)
	if err != nil {
		logs.Error("get second menu list by first menu uid fail: ", err)
	}
	return secondMenuList
}

func GetSecondMenuLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(SECOND_MENU_INFO)
	for k, v := range params {
		qs = qs.Filter(k, v)
	}
	cnt, err := qs.Limit(-1).Count()
	if err != nil {
		logs.Error("get second menu len by map fail: ", err)
	}
	return int(cnt)
}

func GetSecondMenuByMap(params map[string]string, displayCount, offset int) []SecondMenuInfo {
	o := orm.NewOrm()
	var secondMenuList []SecondMenuInfo
	qs := o.QueryTable(SECOND_MENU_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&secondMenuList)
	if err != nil {
		logs.Error("get second menu by map fail: ", err)
	}
	return secondMenuList
}
func InsertSecondMenu(secondMenuInfo SecondMenuInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&secondMenuInfo)
	if err != nil {
		logs.Error("insert second menu fail: ", err)
		return false
	}
	return true
}

func SecondMenuIsExists(seconfMenu string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(SECOND_MENU_INFO).Filter("second_menu", seconfMenu).Exist()
	return exist
}

func SecondMenuUidIsExists(secondMenuUid string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(SECOND_MENU_INFO).Filter("second_menu_uid", secondMenuUid).Exist()
	return exist
}

func SecondRouterExists(secondRouter string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(SECOND_MENU_INFO).Filter("second_router", secondRouter).Exist()
	return exist
}

func DeleteSecondMenuByFirstMenuUid(firstMenuUid string) bool {
	o := orm.NewOrm()
	num, err := o.QueryTable(SECOND_MENU_INFO).Filter("first_menu_uid", firstMenuUid).Delete()
	if err != nil {
		logs.Error("delete second menu by first menu uid fail: ", err)
		return false
	}
	logs.Info("delete second menu by first menu uid success, num: ", num)
	return true
}

func DeleteSecondMenuBySecondMenuUid(secondMenuUid string) bool {
	o := orm.NewOrm()
	num, err := o.QueryTable(SECOND_MENU_INFO).Filter("second_menu_uid", secondMenuUid).Delete()
	if err != nil {
		logs.Error("delete second menu by second menu uid fail: ", err)
		return false
	}
	logs.Info("delete second menu by second menu uid success, num: ", num)
	return true
}

func UpdateSecondMenuOrderBySecondUid(secondUid string, order int) {
	o := orm.NewOrm()
	_, err := o.QueryTable(SECOND_MENU_INFO).Filter("second_menu_uid", secondUid).Update(orm.Params{"menu_order": order})
	if err != nil {
		logs.Error("update second menu order by second menu uid fail: ", err)
	}
}

func UpdateSecondMenu(secondMenu SecondMenuInfo) {
	o := orm.NewOrm()
	_, err := o.Update(&secondMenu)
	if err != nil {
		logs.Error("update second menu for first order fail: ", err)
	}
}

func SecondMenuExistByMenuOrder(menuOrder int) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(SECOND_MENU_INFO).Filter("menu_order", menuOrder).Exist()
	return exist
}
