package legend

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type FixPresent struct {
	Id           int `orm:"pk;column(id)"`
	Uid          string
	TemplateName string
	Money        float64
	PresentMoney float64
	UpdateTime   string
	CreateTime   string
}

const FIXPRESENT = "legend_fix_present"

func (c *FixPresent) TableName() string {
	return FIXPRESENT
}

func InsertFixPresent(fixPresnet *FixPresent) bool {
	o := orm.NewOrm()
	if _, err := o.Insert(fixPresnet); err != nil {
		logs.Error("insert fix present err：", err)
		return false
	}

	return true
}

func GetFixPresentsByName(name string) []FixPresent {
	o := orm.NewOrm()
	var fixPresents []FixPresent
	if _, err := o.QueryTable(FIXPRESENT).Filter("template_name", name).Limit(-1).All(&fixPresents); err != nil {
		logs.Error("get fix presents err：", err)
	}

	return fixPresents
}

func GetFixPresentByUid(uid string) *FixPresent {
	o := orm.NewOrm()
	fixPresent := new(FixPresent)
	if _, err := o.QueryTable(FIXPRESENT).Filter("uid", uid).Limit(1).All(fixPresent); err != nil {
		logs.Error("get fix present err：", err)
	}

	return fixPresent
}

func DeleteFixPresent(templateName string) bool {
	o := orm.NewOrm()
	if _, err := o.QueryTable(FIXPRESENT).Filter("template_name", templateName).Delete(); err != nil {
		logs.Error("delete fix present err: ", err)
		return false
	}

	return true
}

func DeleteFixPresentByUid(uid string) bool {
	o := orm.NewOrm()
	if _, err := o.QueryTable(FIXPRESENT).Filter("uid", uid).Delete(); err != nil {
		logs.Error("delete fix present by uid err：", err)
		return false
	}

	return true
}

func UpdatePresentFixMoney(present *FixPresent) bool {
	o := orm.NewOrm()
	if _, err := o.Update(present); err != nil {
		logs.Error("update fix present err: ", err)
		return false
	}

	return true
}
