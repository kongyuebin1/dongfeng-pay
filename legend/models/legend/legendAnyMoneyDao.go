package legend

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type AnyMoney struct {
	Id             int `orm:"pk;column(id)"`
	TemplateName   string
	GameMoneyName  string
	GameMoneyScale int
	LimitLow       float64
	UpdateTime     string
	CreateTime     string
}

const ANYMONEY = "legend_any_money"

func (c *AnyMoney) TableName() string {
	return ANYMONEY
}

func InsertAnyMoney(anyMoney *AnyMoney) bool {
	o := orm.NewOrm()
	if _, err := o.Insert(anyMoney); err != nil {
		logs.Error("insert any money info err: ", err)
		return false
	}

	return true
}

func GetAnyMoneyByName(name string) *AnyMoney {
	o := orm.NewOrm()
	anyMoney := new(AnyMoney)
	if _, err := o.QueryTable(ANYMONEY).Filter("template_name", name).Limit(1).All(anyMoney); err != nil {
		logs.Error("get any money err：", err)
	}
	return anyMoney
}

func DeleteAnyMoney(templateName string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(ANYMONEY).Filter("template_name", templateName).Delete()
	if err != nil {
		logs.Error("delete any money err：", err)
		return false
	}

	return true
}

func UpdateAnyMoney(anyMoney *AnyMoney) bool {
	o := orm.NewOrm()
	if _, err := o.Update(anyMoney); err != nil {
		logs.Error("update any money err: ", err)
		return false
	}

	return true
}
