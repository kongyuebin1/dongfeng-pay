package fast

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/client/orm"
)

type RpUserBankAccount struct {
	Id              string `orm:"pk;column(id)"`
	CreateTime      string
	EditTime        string
	Version         string
	Remark          string
	Status          string
	UserNo          string
	BankName        string
	BankCode        string
	BankAccountName string
	BankAccountNo   string
	CardType        string
	CardNo          string
	MobileNo        string
	IsDefault       string
	Province        string
	City            string
	Areas           string
	Street          string
	BankAccountType string
}

func (c *RpUserBankAccount) TableName() string {
	return "rp_user_bank_account"
}

func GetBankInfoByUserNo(userNo string) *RpUserBankAccount {
	o := orm.NewOrm()
	userBankAccount := new(RpUserBankAccount)
	if _, err := o.QueryTable("rp_user_bank_account").
		Filter("user_no", userNo).
		All(userBankAccount); err != nil {
		logs.Error("获取用户银行卡信息失败：", err)
	}
	return userBankAccount
}
