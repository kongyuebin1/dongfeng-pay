package routers

import (
	"github.com/beego/beego/v2/server/web"
	"legend/controllers"
	"legend/filter"
)

func init() {
	pageInit()
	logicInit()
}

/**
** 初始化展示页面路由器
 */
func pageInit() {

	//web.Router("/favicon.ico", &controllers.ShowPageController{}, "*:FaviconPage")
	web.Router("/", &controllers.IndexController{}, "*:Index")
	web.Router("/index.html", &controllers.IndexController{}, "*:Index")
	web.Router("/login.html", &controllers.LoginController{}, "*:LoginPage")
	web.Router("/welcome.html", &controllers.ShowPageController{}, "*:WelcomePage")
	web.Router("/merchantKey.html", &controllers.ShowPageController{}, "*:MerchantKeyPage")
	web.Router("/orderList.html", &controllers.ShowPageController{}, "*:OrderListPage")
	web.Router("/scaleTemplate.html", &controllers.ShowPageController{}, "*:ScaleTemplatePage")
	web.Router("/templateAdd.html", &controllers.ShowPageController{}, "*:TemplateAdd")
	web.Router("/templateEdit.html", &controllers.ShowPageController{}, "*:TemplateEdit")
	web.Router("/groupList.html", &controllers.ShowPageController{}, "*:GroupListPage")
	web.Router("/areaList.html", &controllers.ShowPageController{}, "*:AreaListPage")
	web.Router("/imitateOrder.html", &controllers.ShowPageController{}, "*:ImitateChargePage")
	web.Router("/settleList.html", &controllers.ShowPageController{}, "*:SettleListPage")
	web.Router("/everydayChargeCount.html", &controllers.ShowPageController{}, "*:EverydayChargeCountPage")
	web.Router("/groupChargeCount.html", &controllers.ShowPageController{}, "*:GroupChargeCountPage")
	web.Router("/areaChargeCount.html", &controllers.ShowPageController{}, "*:AreaChargePage")
	web.Router("/person.html", &controllers.ShowPageController{}, "*:PersonPage")
	web.Router("areaAddOrEdit.html", &controllers.ShowPageController{}, "*:AreaAddOrEdit")
}

/**
** 业务逻辑路由
 */
func logicInit() {
	web.Router("/login", &controllers.LoginController{}, "*:Login")
	web.Router("/logout.html", &controllers.LogoutController{}, "*:Logout")
	web.Router("/switch/login", &controllers.LogoutController{}, "*:SwitchLogin")
	web.Router("/person/password", &controllers.LoginController{}, "*:PersonPassword")
	web.Router("/add/template", &controllers.TemplateController{}, "*:TemplateAdd")
	web.Router("/template/list", &controllers.TemplateController{}, "*:TemplateList")
	web.Router("/delete/template", &controllers.TemplateController{}, "*:TemplateDelete")
	web.Router("/template/info", &controllers.TemplateController{}, "*:TemplateAllInfo")
	web.Router("/group/list", &controllers.GroupController{}, "*:ListGroup")
	web.Router("/add/group", &controllers.GroupController{}, "*:AddGroup")
	web.Router("/delete/group", &controllers.GroupController{}, "*:DeleteGroup")
	web.Router("/edit/group", &controllers.GroupController{}, "*:EditGroup")
	web.InsertFilter("/*", web.BeforeRouter, filter.LoginFilter)
}
