/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/8/29 14:43
 ** @Author : yuebin
 ** @File : role_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/29 14:43
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type RoleInfo struct {
	Id             int
	RoleName       string
	RoleUid        string
	ShowFirstMenu  string
	ShowFirstUid   string
	ShowSecondMenu string
	ShowSecondUid  string
	ShowPower      string
	ShowPowerUid   string
	Creater        string
	Status         string
	Remark         string
	CreateTime     string
	UpdateTime     string
}

const ROLE_INFO = "role_info"

func GetRoleLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(ROLE_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Count()
	if err != nil {
		logs.Error("get role len by map fail: ", err)
	}
	return int(cnt)
}

func GetRole() []RoleInfo {
	o := orm.NewOrm()
	var roleInfo []RoleInfo
	_, err := o.QueryTable(ROLE_INFO).Limit(-1).OrderBy("-update_time").All(&roleInfo)
	if err != nil {
		logs.Error("get all role fail: ", err)
	}
	return roleInfo
}

func GetRoleByMap(params map[string]string, display, offset int) []RoleInfo {
	o := orm.NewOrm()
	var roleInfo []RoleInfo
	qs := o.QueryTable(ROLE_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(display, offset).OrderBy("-update_time").All(&roleInfo)
	if err != nil {
		logs.Error("get role by map fail: ", err)
	}
	return roleInfo
}

func GetRoleByRoleUid(roleUid string) RoleInfo {
	o := orm.NewOrm()
	var roleInfo RoleInfo
	_, err := o.QueryTable(ROLE_INFO).Filter("role_uid", roleUid).Limit(1).All(&roleInfo)

	if err != nil {
		logs.Error("get role by role uid fail: ", err)
	}
	return roleInfo
}

func RoleNameExists(roleName string) bool {
	o := orm.NewOrm()
	exists := o.QueryTable(ROLE_INFO).Filter("role_name", roleName).Exist()
	return exists
}

func InsertRole(roleInfo RoleInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&roleInfo)
	if err != nil {
		logs.Error("insert role fail: ", err)
		return false
	}
	return true
}

func DeleteRoleByRoleUid(roleUid string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(ROLE_INFO).Filter("role_uid", roleUid).Delete()
	if err != nil {
		logs.Error("delete role by role uid fail: ", err)
		return false
	}
	return true
}

func UpdateRoleInfo(roleInfo RoleInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&roleInfo)

	if err != nil {
		logs.Error("update role info fail: ", err)
		return false
	}
	return true
}
