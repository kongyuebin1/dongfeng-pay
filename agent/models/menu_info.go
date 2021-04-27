/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/8/21 9:33
 ** @Author : yuebin
 ** @File : menu_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/21 9:33
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type MenuInfo struct {
	Id         int
	MenuOrder  int
	MenuUid    string
	FirstMenu  string
	SecondMenu string
	Creater    string
	Status     string
	CreateTime string
	UpdateTime string
}

//实现排序的三个接口函数
type MenuInfoSlice []MenuInfo

func (m MenuInfoSlice) Len() int {
	return len(m)
}

func (m MenuInfoSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MenuInfoSlice) Less(i, j int) bool {
	return m[i].MenuOrder < m[j].MenuOrder //从小到大排序
}

const MENUINFO = "menu_info"

func InsertMenu(menuInfo MenuInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&menuInfo)
	if err != nil {
		logs.Error("insert new menu info fail：", err)
		return false
	}
	return true
}

func FirstMenuIsExists(firstMenu string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(MENUINFO).Filter("first_menu", firstMenu).Exist()
	return exist
}

func FirstMenuUidIsExists(firstMenUid string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(MENUINFO).Filter("menu_uid", firstMenUid).Exist()
	return exist
}

func MenuOrderIsExists(menuOrder int) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(MENUINFO).Filter("menu_order", menuOrder).Exist()
	return exist
}

func GetMenuLen() int {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(MENUINFO).Count()
	if err != nil {
		logs.Error("get menu info len length fail: ", err)
	}
	return int(cnt)
}

func GetMenuInfoByMenuUid(menuUid string) MenuInfo {
	o := orm.NewOrm()
	var menuInfo MenuInfo
	_, err := o.QueryTable(MENUINFO).Filter("menu_uid", menuUid).Limit(1).All(&menuInfo)
	if err != nil {
		logs.Error("get menu info by menuUid fail: ", err)
	}
	return menuInfo
}

func GetMenuInfosByMenuUids(menuUids []string) []MenuInfo {
	menuInfoList := make([]MenuInfo, 0)
	for _, v := range menuUids {
		m := GetMenuInfoByMenuUid(v)
		menuInfoList = append(menuInfoList, m)
	}
	return menuInfoList
}

func GetMenuInfoByMenuOrder(menuOrder int) MenuInfo {
	o := orm.NewOrm()
	var menuInfo MenuInfo
	_, err := o.QueryTable(MENUINFO).Filter("menu_order", menuOrder).Limit(1).All(&menuInfo)
	if err != nil {
		logs.Error("get menu info by menu order fail: ", err)
	}
	return menuInfo
}

func GetMenuAll() []MenuInfo {
	o := orm.NewOrm()
	var menuInfoList []MenuInfo
	_, err := o.QueryTable(MENUINFO).OrderBy("-update_time").All(&menuInfoList)
	if err != nil {
		logs.Error("get all menu list fail：", err)
	}
	return menuInfoList
}

func GetMenuOffset(displayCount, offset int) []MenuInfo {
	o := orm.NewOrm()
	var menuInfoList []MenuInfo
	_, err := o.QueryTable(MENUINFO).Limit(displayCount, offset).All(&menuInfoList)
	if err != nil {
		logs.Error("get menu offset fail: ", err)
	}
	return menuInfoList
}

func GetMenuOffsetByMap(params map[string]string, displayCount, offset int) []MenuInfo {
	o := orm.NewOrm()
	var menuInfoList []MenuInfo
	qs := o.QueryTable(MENUINFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&menuInfoList)
	if err != nil {
		logs.Error("get menu offset by map fail: ", err)
	}
	return menuInfoList
}

func GetMenuLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(MENUINFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Count()
	if err != nil {
		logs.Error("get menu len by map fail: ", err)
	}
	return int(cnt)
}

func UpdateMenuInfo(menuInfo MenuInfo) {
	o := orm.NewOrm()
	cnt, err := o.Update(&menuInfo)
	if err != nil {
		logs.Error("update menu info fail: ", err)
	}
	logs.Info("update menu info success, num: ", cnt)
}

func DeleteMenuInfo(menuUid string) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(MENUINFO).Filter("menu_uid", menuUid).Delete()
	if err != nil {
		logs.Error("delete menu info fail: ", err)
	}
	logs.Info("delete menu info num: ", cnt)
}
