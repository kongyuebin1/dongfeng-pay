package fast

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/client/orm"
)

type RpUserInfo struct {
	Id                    string `orm:"pk;column(id)"`
	CreateTime            string
	Status                string
	UserNo                string
	UserName              string
	AccountNo             string
	Mobile                string
	Password              string
	PayPwd                string
	LastSmsVerifyCodeTime string
	Email                 string
	Ips                   string
}

func (c *RpUserInfo) TableName() string {
	return "rp_user_info"

}

func tableName() string {
	return "rp_user_info"
}

func GetUserInfoByUserName(userName string) *RpUserInfo {

	o := orm.NewOrm()
	userInfo := new(RpUserInfo)

	_, err := o.QueryTable(tableName()).Filter("mobile", userName).All(userInfo)

	if err != nil {
		logs.Error("根据用户名从数据获取用户信息失败：", err)
	}

	return userInfo
}

/**
** 更新用户信息
 */
func UpdateUserInfo(userInfo *RpUserInfo) bool {
	o := orm.NewOrm()

	if _, err := o.Update(userInfo); err != nil {
		logs.Error("更新用户信息失败，错误：%s", err)
		return false
	}

	return true
}
