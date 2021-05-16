package fast

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type BankCardInfo struct {
	Id              string `orm:"pk;column(id)"`
	Uid             string
	UserName        string
	BankName        string
	BankCode        string
	BankAccountType string
	AccountName     string
	BankNo          string
	IdentifyCard    string
	CertificateNo   string
	PhoneNo         string
	BankAddress     string
	CreateTime      string
	UpdateTime      string
}

const BANKCARDINFO = "bank_card_info"

func (c *BankCardInfo) TableName() string {
	return BANKCARDINFO
}

func GetBankCardInfoByUserNo(merchantNo string) *BankCardInfo {
	o := orm.NewOrm()
	bankCardInfo := new(BankCardInfo)
	if _, err := o.QueryTable(BANKCARDINFO).Filter("user_name", merchantNo).Limit(1).All(bankCardInfo); err != nil {

		logs.Error("获取用户银行卡信息失败：", err)
	}
	return bankCardInfo
}
