package fast

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/client/orm"
)

type RpUserPayConfig struct {
	Id        string `orm:"pk;column(id)"`
	UserNo    string
	UserName  string
	Status    string
	PayKey    string
	PaySecret string
}

func (c *RpUserPayConfig) TableName() string {
	return "rp_user_pay_config"
}

func getTableName() string {
	return "rp_user_pay_config"
}

func GetUserPayConfigByUserNo(userNo string) *RpUserPayConfig {
	o := orm.NewOrm()
	userPayConfig := new(RpUserPayConfig)
	_, err := o.QueryTable(getTableName()).Filter("user_no", userNo).All(userPayConfig)
	if err != nil {
		logs.Error("获取用户支付配置错误：", err)
	}

	return userPayConfig
}
