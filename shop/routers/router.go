package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"shop/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeAction{}, "*:ShowHome") //初始化首页
	beego.Router("/pay.html", &controllers.PayController{}, "*:Pay")
	beego.Router("/pay_requst.html", &controllers.ScanShopController{})
	beego.Router("/scan.html", &controllers.ScanShopController{}, "*:ScanRender")
	beego.Router("/error.html", &controllers.HomeAction{}, "*:ErrorPage")
}
