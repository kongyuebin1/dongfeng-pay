package legend

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type Group struct {
	Id        int `orm:"pk;column(id)"`
	GroupName string
	Uid       string
	BaseDao
}

const GROUP = "legend_group"

func (c *Group) TableName() string {
	return GROUP
}

func InsertGroup(group *Group) bool {
	o := orm.NewOrm()
	if _, err := o.Insert(group); err != nil {
		logs.Error("insert group err：", err)
		return false
	}

	return true
}

func GetGroupAllCont() int {
	o := orm.NewOrm()
	count, err := o.QueryTable(GROUP).Limit(-1).Count()
	if err != nil {
		logs.Error(" get group all count err: ", err)
	}

	return int(count)
}

func GetGroupList(offset, limit int) []Group {
	o := orm.NewOrm()

	var groups []Group
	if _, err := o.QueryTable(GROUP).Limit(limit, offset).OrderBy("-create_time").All(&groups); err != nil {
		logs.Error("get scale template list err : ", err)
	}

	return groups
}

func GetGroupByName(name string) *Group {
	o := orm.NewOrm()
	group := new(Group)
	if _, err := o.QueryTable(GROUP).Filter("group_name", name).Limit(1).All(group); err != nil {
		logs.Error(" get group by name err：", err)
	}

	return group
}

func GetGroupByUid(uid string) *Group {
	o := orm.NewOrm()
	group := new(Group)
	if _, err := o.QueryTable(GROUP).Filter("uid", uid).Limit(1).All(group); err != nil {
		logs.Error("get group by uid err: ", err)
	}

	return group
}

func DeleteGroupByUid(uid string) bool {
	o := orm.NewOrm()
	if _, err := o.QueryTable(GROUP).Filter("uid", uid).Delete(); err != nil {
		logs.Error("delete group by uid err: ", err)
		return false
	}

	return true
}

func UpdateGroup(group *Group) bool {
	o := orm.NewOrm()
	if _, err := o.Update(group); err != nil {
		logs.Error("update group err：", err)
		return false
	}

	return true
}
