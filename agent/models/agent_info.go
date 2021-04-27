/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/9/19 14:41
 ** @Author : yuebin
 ** @File : agent_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/9/19 14:41
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type AgentInfo struct {
	Id            int
	Status        string
	AgentName     string
	AgentPassword string
	PayPassword   string
	AgentRemark   string
	AgentUid      string
	AgentPhone    string
	UpdateTime    string
	CreateTime    string
}

const AGENT_INFO = "agent_info"

func IsEixstByAgentName(agentName string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(AGENT_INFO).Filter("agent_name", agentName).Exist()

	return exist
}

func IsExistByAgentUid(uid string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(AGENT_INFO).Filter("agent_uid", uid).Exist()

	return exist
}

func IsEixstByAgentPhone(agentPhone string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(AGENT_INFO).Filter("agent_phone", agentPhone).Exist()
	return exist
}

func InsertAgentInfo(agentInfo AgentInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&agentInfo)
	if err != nil {
		logs.Error("insert agent info fail: ", err)
		return false
	}

	return true
}

func GetAgentInfoByAgentUid(agentUid string) AgentInfo {
	o := orm.NewOrm()
	var agentInfo AgentInfo
	_, err := o.QueryTable(AGENT_INFO).Filter("agent_uid", agentUid).Limit(1).All(&agentInfo)

	if err != nil {
		logs.Error("get agent info by agentUid fail: ", err)
	}

	return agentInfo
}

func GetAgentInfoByPhone(phone string) AgentInfo {
	o := orm.NewOrm()
	var agentInfo AgentInfo
	_, err := o.QueryTable(AGENT_INFO).Filter("agent_phone", phone).Limit(1).All(&agentInfo)

	if err != nil {
		logs.Error("get agent info by phone fail: ", err)
	}

	return agentInfo
}

func GetAgentInfoLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(AGENT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Limit(-1).Count()
	if err != nil {
		logs.Error("get agentinfo len by map fail: ", err)
	}

	return int(cnt)
}

func GetAgentInfoByMap(params map[string]string, displayCount, offset int) []AgentInfo {
	o := orm.NewOrm()
	var agentInfoList []AgentInfo

	qs := o.QueryTable(AGENT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&agentInfoList)

	if err != nil {
		logs.Error("get agentInfo by map fail: ", err)
	}

	return agentInfoList
}

func GetAllAgentByMap(parmas map[string]string) []AgentInfo {
	o := orm.NewOrm()
	var agentList []AgentInfo

	qs := o.QueryTable(AGENT_INFO)
	for k, v := range parmas {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	_, err := qs.Limit(-1).All(&agentList)
	if err != nil {
		logs.Error("get all agent by map fail: ", err)
	}

	return agentList
}

func UpdateAgentInfo(agentInfo AgentInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&agentInfo)

	if err != nil {
		logs.Error("update agentinfo fail: ", err)
		return false
	}

	return true
}

func DeleteAgentByAgentUid(agentUid string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(AGENT_INFO).Filter("agent_uid", agentUid).Delete()
	if err != nil {
		logs.Error("delete agent by agent uid fail: ", err)
		return false
	}
	return true
}
