package fast

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/client/orm"
)

type RpAccount struct {
	Id            string `orm:"pk;column(id)"`
	CreateTime    string
	EditTime      string
	Version       int
	Remark        string
	AccountNo     string
	Balance       float64
	Unbalance     float64
	SecurityMoney float64
	Status        string
	TotalIncome   float64
	TodayIncome   float64
	SettAmount    float64
	UserNo        string
	AmountFrozen  float64
}

func (c *RpAccount) TableName() string {
	return "rp_account"
}

func GetAccontInfo(userNo string) *RpAccount {
	o := orm.NewOrm()

	rpAccount := new(RpAccount)
	if _, err := o.QueryTable("rp_account").Filter("user_no", userNo).All(rpAccount); err != nil {
		logs.Error("获取account信息失败：", err)
	}

	return rpAccount
}
