/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/8/9 14:02
 ** @Author : yuebin
 ** @File : user_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/9 14:02
 ** @Software: GoLand
****************************************************/
package user

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

const (
	USERINFO = "user_info"
)

type UserInfo struct {
	Id         int
	UserId     string
	Passwd     string
	Nick       string
	Remark     string
	Ip         string
	Status     string
	Role       string
	RoleName   string
	CreateTime string
	UpdateTime string
}

func GetUserInfoByUserID(userID string) UserInfo {
	o := orm.NewOrm()
	var userInfo UserInfo
	err := o.QueryTable(USERINFO).Exclude("status", "delete").Filter("user_id", userID).One(&userInfo)
	if err != nil {
		logs.Error("get user info fail: ", err)
	}
	return userInfo
}

func GetOperatorByMap(params map[string]string, displayCount, offset int) []UserInfo {
	o := orm.NewOrm()
	var userInfo []UserInfo
	qs := o.QueryTable(USERINFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Exclude("status", "delete").Limit(displayCount, offset).OrderBy("-update_time").All(&userInfo)

	if err != nil {
		logs.Error("get operator by map fail: ", err)
	}
	return userInfo
}

func GetOperatorLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(USERINFO)
	for k, v := range params {
		qs = qs.Filter(k, v)
	}
	cnt, err := qs.Exclude("status", "delete").Count()
	if err != nil {
		logs.Error("get operator len by map fail: ", err)
	}
	return int(cnt)
}

func UpdateUserInfoIP(userInfo UserInfo) {
	o := orm.NewOrm()
	num, err := o.QueryTable(USERINFO).Exclude("status", "delete").Filter("user_id", userInfo.UserId).Update(orm.Params{"ip": userInfo.Ip})
	if err != nil {
		logs.Error("%s update user info ip fail: %v", userInfo.UserId, err)
	} else {
		logs.Info("%s update user info ip success, num: %d", userInfo.UserId, num)
	}
}

func UpdateUserInfoPassword(userInfo UserInfo) {
	o := orm.NewOrm()
	num, err := o.QueryTable(USERINFO).Exclude("status", "delete").Filter("user_id", userInfo.UserId).Update(orm.Params{"passwd": userInfo.Passwd})
	if err != nil {
		logs.Error("%s update user info password fail: %v", userInfo.UserId, err)
	} else {
		logs.Info("%s update user info password success, update num: %d", userInfo.UserId, num)
	}
}

func UpdateUserInfo(userInfo UserInfo) {
	o := orm.NewOrm()
	if num, err := o.Update(&userInfo); err != nil {
		logs.Error("update user info fail: ", err)
	} else {
		logs.Info("update user info success, num: ", num)
	}
}

func UpdateStauts(status, userId string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(USERINFO).Filter("user_id", userId).Update(orm.Params{"status": status})

	if err != nil {
		logs.Error("update status fail: ", err)
		return false
	}
	return true
}

func UserInfoExistByUserId(userId string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(USERINFO).Exclude("status", "delete").Filter("user_id", userId).Exist()
	return exist
}

func NickIsExist(nick string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(USERINFO).Exclude("status", "delete").Filter("nick", nick).Exist()
	return exist
}

func InsertUser(userInfo UserInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&userInfo)
	if err != nil {
		logs.Error("insert user fail: ", err)
		return false
	}
	return true
}

func DeleteUserByUserId(userId string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(USERINFO).Exclude("status", "delete").Filter("user_id", userId).Update(orm.Params{"status": "delete"})

	if err != nil {
		logs.Error("delete user by userId fail: ", err)
		return false
	}
	return true
}
