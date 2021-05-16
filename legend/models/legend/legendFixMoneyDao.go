package legend

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type FixMoney struct {
	Id           int `orm:"pk;column(id)"`
	Uid          string
	TemplateName string
	Price        float64
	GoodsName    string
	GoodsNo      string
	BuyTimes     int
	UpdateTime   string
	CreateTime   string
}

const FIXMONEY = "legend_fix_money"

func (c *FixMoney) TableName() string {
	return FIXMONEY
}

func InsertFixMoney(fixMoney *FixMoney) bool {
	o := orm.NewOrm()
	if _, err := o.Insert(fixMoney); err != nil {
		logs.Error("insert fix money err: ", err)
		return false
	}

	return true
}

func GetFixMoneyByName(name string) []FixMoney {
	o := orm.NewOrm()
	var fixMoneys []FixMoney
	if _, err := o.QueryTable(FIXMONEY).Filter("template_name", name).Limit(-1).All(&fixMoneys); err != nil {
		logs.Error("get fix money err：", err)
	}

	return fixMoneys
}

func GetFixMoneyByUid(uid string) *FixMoney {
	o := orm.NewOrm()
	fixMoney := new(FixMoney)
	if _, err := o.QueryTable(FIXMONEY).Filter("uid", uid).Limit(1).All(fixMoney); err != nil {
		logs.Error("get fix Money by uid err：", err)
	}

	return fixMoney
}

func DeleteFixMoney(templateName string) bool {
	o := orm.NewOrm()
	if _, err := o.QueryTable(FIXMONEY).Filter("template_name", templateName).Delete(); err != nil {
		logs.Error("delete fix money err：", err)
		return false
	}

	return true
}

func DeleteFixMoneyByUid(uid string) bool {
	o := orm.NewOrm()
	if _, err := o.QueryTable(FIXMONEY).Filter("uid", uid).Delete(); err != nil {
		logs.Error("delete fix money err：", err)
		return false
	}

	return true
}

func UpdateFixMoney(fixMoney *FixMoney) bool {
	o := orm.NewOrm()
	if _, err := o.Update(fixMoney); err != nil {
		logs.Error("update fix money err: ", err)
		return false
	}

	return true
}
