/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/23 15:17
 ** @Author : yuebin
 ** @File : router_pages
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/23 15:17
 ** @Software: GoLand
****************************************************/
package routers

import (
	"github.com/astaxie/beego"
	"juhe/jhboss/controllers"
)

func init() {
	beego.Router("/", &controllers.PageController{}, "*:Index")
	beego.Router("/index.html", &controllers.PageController{}, "*:Index")
	beego.Router("/login.html", &controllers.PageController{}, "*:LoginPage")
	beego.Router("/account.html", &controllers.PageController{}, "*:AccountPage")
	beego.Router("/account_history.html", &controllers.PageController{}, "*:AccountHistoryPage")
	beego.Router("/bank_card.html", &controllers.PageController{}, "*:BankCardPage")
	beego.Router("/create_agent.html", &controllers.PageController{}, "*:CreateAgentPage")
	beego.Router("/edit_role.html", &controllers.PageController{}, "*:EditRolePage")
	beego.Router("/first_menu.html", &controllers.PageController{}, "*:FirstMenuPage")
	beego.Router("/main.html", &controllers.PageController{}, "*:MainPage")
	beego.Router("/menu.html", &controllers.PageController{}, "*:MenuPage")
	beego.Router("/merchant.html", &controllers.PageController{}, "*:MerchantPage")
	beego.Router("/operator.html", &controllers.PageController{}, "*:OperatorPage")
	beego.Router("/power.html", &controllers.PageController{}, "*:PowerPage")
	beego.Router("/road.html", &controllers.PageController{}, "*:RoadPage")
	beego.Router("/road_pool.html", &controllers.PageController{}, "*:RoadPoolPage")
	beego.Router("/road_profit.html", &controllers.PageController{}, "*:RoadProfitPage")
	beego.Router("/role.html", &controllers.PageController{}, "*:RolePage")
	beego.Router("/second_menu.html", &controllers.PageController{}, "*:SecondMenuPage")
	beego.Router("/order_info.html", &controllers.PageController{}, "*:OrderInfoPage")
	beego.Router("/order_profit.html", &controllers.PageController{}, "*:OrderProfitPage")
	beego.Router("/merchant_payfor.html", &controllers.PageController{}, "*:MerchantPayforPage")
	beego.Router("/self_payfor.html", &controllers.PageController{}, "*:SelfPayforPage")
	beego.Router("/payfor_record.html", &controllers.PageController{}, "*:PayforRecordPage")
	beego.Router("/confirm.html", &controllers.PageController{}, "*:ConfirmPage")
	beego.Router("/self_notify.html", &controllers.PageController{}, "*:SelfNotifyPage")
	beego.Router("/self_plus_sub.html", &controllers.PageController{}, "*:SelfPlusSubPage")
	beego.Router("/agent_to_merchant.html", &controllers.PageController{}, "*:AgentToMerchantPage")
	beego.Router("/platform_profit.html", &controllers.PageController{}, "*:PlatFormProfitPage")
	beego.Router("/agent_profit.html", &controllers.PageController{}, "*:AgentProfitPage")
}
