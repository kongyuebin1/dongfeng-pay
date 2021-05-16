package fast

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type MerchantDeployInfo struct {
	Id             string `orm:"pk;column(id)"`
	Status         string
	MerchantUid    string
	PayType        string
	SingleRoadUid  string
	SingleRoadName string
}

const MERCHANTDEPLOYINFO = "merchant_deploy_info"

func (c *MerchantDeployInfo) TableName() string {
	return MERCHANTDEPLOYINFO
}

func GetUserPayConfigByUserNo(userNo string) *MerchantDeployInfo {
	o := orm.NewOrm()
	userPayConfig := new(MerchantDeployInfo)
	_, err := o.QueryTable(MERCHANTDEPLOYINFO).Filter("user_no", userNo).All(userPayConfig)
	if err != nil {
		logs.Error("获取用户支付配置错误：", err)
	}

	return userPayConfig
}
