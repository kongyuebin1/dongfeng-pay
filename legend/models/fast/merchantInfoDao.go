package fast

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type MerchantInfo struct {
	Id                   string `orm:"pk;column(id)"`
	Status               string
	BelongAgentUid       string
	BelongAgentName      string
	MerchantName         string
	MerchantUid          string
	MerchantKey          string
	MerchantSecret       string
	LoginAccount         string
	LoginPassword        string
	AutoSettle           string
	AutoPayFor           string
	WhiteIps             string
	Remark               string
	SinglePayForRoadUid  string
	SinglePayForRoadName string
	RollPayForRoadCode   string
	RollPayForRoadName   string
	PayforFee            string
	CreateTime           string
	UpdateTime           string
}

func (c *MerchantInfo) TableName() string {
	return "merchant_info"

}

func tableName() string {
	return "merchant_info"
}

func GetMerchantInfoByUserName(userName string) *MerchantInfo {

	o := orm.NewOrm()
	userInfo := new(MerchantInfo)

	_, err := o.QueryTable(tableName()).Filter("login_account", userName).All(userInfo)

	if err != nil {
		logs.Error("根据用户名从数据获取用户信息失败：", err)
	}

	return userInfo
}

/**
** 更新用户信息
 */
func UpdateMerchantInfo(merchantInfo *MerchantInfo) bool {
	o := orm.NewOrm()

	if _, err := o.Update(merchantInfo); err != nil {
		logs.Error("更新用户信息失败，错误：%s", err)
		return false
	}

	return true
}
