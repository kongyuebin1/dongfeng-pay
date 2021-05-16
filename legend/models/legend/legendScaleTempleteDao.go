package legend

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type ScaleTemplate struct {
	Id           int `orm:"pk;column(id)"`
	MerchantUid  string
	TemplateName string
	UserUid      string
	UserWarn     string
	MoneyType    string
	PresentType  string
	UpdateTime   string
	CreateTime   string
}

const SCALETMPLETE = "legend_scale_template"

func (c *ScaleTemplate) TableName() string {
	return SCALETMPLETE
}

func InsertScaleTemplate(scaleTemplate *ScaleTemplate) bool {
	o := orm.NewOrm()

	if _, err := o.Insert(scaleTemplate); err != nil {
		logs.Error("insert scale template err: ", err)
		return false
	}

	return true
}

func IsExistsScaleTemplateByName(name string) bool {
	o := orm.NewOrm()
	return o.QueryTable(SCALETMPLETE).Filter("template_name", name).Exist()
}

func GetScaleTemplateList(offset, limit int) []ScaleTemplate {
	o := orm.NewOrm()

	var scaleTemplates []ScaleTemplate
	if _, err := o.QueryTable(SCALETMPLETE).Limit(limit, offset).OrderBy("-create_time").All(&scaleTemplates); err != nil {
		logs.Error("get scale template list err : ", err)
	}

	return scaleTemplates
}

func GetScaleTemplateAll() int {
	o := orm.NewOrm()
	count, err := o.QueryTable(SCALETMPLETE).Count()
	if err != nil {
		logs.Error("get scale template all err：", err)
	}
	return int(count)
}

func GetScaleTemplateByName(name string) *ScaleTemplate {
	o := orm.NewOrm()
	scaleTemplate := new(ScaleTemplate)
	if _, err := o.QueryTable(SCALETMPLETE).Filter("template_name", name).Limit(1).All(scaleTemplate); err != nil {
		logs.Error("get scale template by name err：", err)
	}

	return scaleTemplate
}

func GetScaleTemplateByNameAndMerchantUid(name, merchantUid string) *ScaleTemplate {
	o := orm.NewOrm()
	scaleTemplate := new(ScaleTemplate)
	if _, err := o.QueryTable(SCALETMPLETE).Filter("template_name", name).
		Filter("merchant_uid", merchantUid).Limit(1).All(scaleTemplate); err != nil {
		logs.Error("get scale template by name and merchantUid err：", err)
	}

	return scaleTemplate
}

func DeleteScaleTemplate(templateName string) bool {
	o := orm.NewOrm()
	if _, err := o.QueryTable(SCALETMPLETE).Filter("template_name", templateName).Delete(); err != nil {
		logs.Error("delete template err：", err)
		return false
	}

	return true
}

func UpdateScaleTemplate(template *ScaleTemplate) bool {
	o := orm.NewOrm()
	if _, err := o.Update(template); err != nil {
		logs.Error("update scale template err: ", err)
		return false
	}

	return true
}
