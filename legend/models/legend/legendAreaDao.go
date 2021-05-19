package legend

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type Area struct {
	Id           int `orm:"pk;column(id)"`
	AreaName     string
	Uid          string
	GroupName    string
	TemplateName string
	NotifyUrl    string
	AttachParams string
	BaseDao
}

const AREA = "legend_area"

func (c *Area) TableName() string {
	return AREA
}

func InsertArea(area *Area) bool {
	o := orm.NewOrm()
	if _, err := o.Insert(area); err != nil {
		logs.Error("insert area err: ", err)
		return false
	}
	return true
}

func GetAreaByName(name string) *Area {
	o := orm.NewOrm()
	area := new(Area)
	if _, err := o.QueryTable(AREA).Filter("area_name", name).Limit(1).All(area); err != nil {
		logs.Error("get area by name err：", err)
	}

	return area
}

func GetAreaAllCount() int {
	o := orm.NewOrm()
	count, err := o.QueryTable(AREA).Count()
	if err != nil {
		logs.Error("get area all count err：", err)
	}

	return int(count)
}

func GetAreaList(offset, limit int) []Area {
	o := orm.NewOrm()
	var areas []Area
	if _, err := o.QueryTable(AREA).Limit(limit, offset).OrderBy("-create_time").All(&areas); err != nil {
		logs.Error(" get area list err：", err)
	}

	return areas
}

func GetAreaByUid(uid string) *Area {
	o := orm.NewOrm()
	area := new(Area)
	if _, err := o.QueryTable(AREA).Filter("uid", uid).Limit(1).All(area); err != nil {
		logs.Error(" get area by uid err : ", err)
	}

	return area
}

func UpdateArea(area *Area) bool {
	o := orm.NewOrm()
	if _, err := o.Update(area); err != nil {
		logs.Error("update area err：", err)
		return false
	}

	return true
}

func DeleteAreaByUid(uid string) bool {
	o := orm.NewOrm()
	if _, err := o.QueryTable(AREA).Filter("uid", uid).Delete(); err != nil {
		logs.Error(" delete area by uid err：", err)
		return false
	}

	return true
}
