package legend

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type Group struct {
	Id         int `orm:"pk;column(id)"`
	GroupName  string
	Uid        string
	UpdateTime string
	CreateTime string
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
