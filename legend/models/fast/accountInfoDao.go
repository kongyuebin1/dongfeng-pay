package fast

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type AccountInfo struct {
	Id           string `orm:"pk;column(id)"`
	Status       string
	AccountUid   string
	AccountName  string
	Balance      float64
	SettleAmount float64
	LoanAmount   float64
	WaitAmount   float64
	FreezeAmount float64
	PayforAmount float64
	UpdateTime   string
	CreateTime   string
}

const ACCOUNTINFO = "account_info"

func (c *AccountInfo) TableName() string {
	return "account_info"
}

func GetAccountInfo(accountUid string) *AccountInfo {
	o := orm.NewOrm()

	account := new(AccountInfo)
	if _, err := o.QueryTable(ACCOUNTINFO).Filter("account_uid", accountUid).All(account); err != nil {
		logs.Error("获取account信息失败：", err)
	}

	return account
}
