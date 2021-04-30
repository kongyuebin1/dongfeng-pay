package base

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/server/web"
	"legend/models/fast"
)

/**
** 基础controller插件，重写一些公共的方法
 */
type BasicController struct {
	web.Controller
}

func (c *BasicController) Prepare() {

	userName, ok := c.GetSession("userName").(string)
	if ok {
		logs.Info("该用户已经登录， userName：", userName)
		userInfo := fast.GetMerchantInfoByUserName(userName)
		if userInfo.LoginAccount != "" {
			c.Data["nickName"] = userInfo.MerchantName
		}
	} else {
		c.Data["nickName"] = "史蒂芬-库里"
	}
}
