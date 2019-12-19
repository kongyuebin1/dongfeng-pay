package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"juhe/jhmerchant/controllers"
)

func init() {
	//生产登录验证码
	beego.Handler("/img.do/*.png", captcha.Server(130, 40))

	beego.Include(
		&controllers.Login{},
		&controllers.Index{},
		&controllers.TradeRecord{},
		&controllers.UserInfo{},
		&controllers.Withdraw{},
		&controllers.DealExcel{},
		&controllers.MultiWithdraw{},
		&controllers.History{},
	)
}
