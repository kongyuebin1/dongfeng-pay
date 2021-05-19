package legend

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type ScalePresent struct {
	Id           int `orm:"pk;column(id)"`
	Uid          string
	TemplateName string
	Money        float64
	PresentScale float64
	BaseDao
}

const SCALEPRESENT = "legend_scale_present"

func (c *ScalePresent) TableName() string {
	return SCALEPRESENT
}

func InsertScalePresent(scalePresent *ScalePresent) bool {
	o := orm.NewOrm()
	if _, err := o.Insert(scalePresent); err != nil {
		logs.Error("insert scale present err：", err)
		return false
	}

	return true
}

func GetScalePresentsByName(name string) []ScalePresent {
	o := orm.NewOrm()
	var scalePresents []ScalePresent
	if _, err := o.QueryTable(SCALEPRESENT).Filter("template_name", name).Limit(-1).All(&scalePresents); err != nil {
		logs.Error("get scale present err：", err)
	}

	return scalePresents
}

func GetScalePresentByUid(uid string) *ScalePresent {
	o := orm.NewOrm()
	scalePresent := new(ScalePresent)
	if _, err := o.QueryTable(SCALEPRESENT).Filter("uid", uid).Limit(1).All(scalePresent); err != nil {
		logs.Error("get scale present err：", err)
	}

	return scalePresent
}

func DeleteScalePresent(templateName string) bool {
	o := orm.NewOrm()
	if _, err := o.QueryTable(SCALEPRESENT).Filter("template_name", templateName).Delete(); err != nil {
		logs.Error("delete scale present err：", err)
		return false
	}

	return true
}

func DeleteScalePresentByUid(uid string) bool {
	o := orm.NewOrm()
	if _, err := o.QueryTable(SCALEPRESENT).Filter("uid", uid).Delete(); err != nil {
		logs.Error("delete scale present by uid err：", err)
		return false
	}

	return true
}

func UpdateScalePresent(present *ScalePresent) bool {
	o := orm.NewOrm()
	if _, err := o.Update(present); err != nil {
		logs.Error("update scale present err: ", err)
		return false
	}

	return true
}
