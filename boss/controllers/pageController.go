/***************************************************
 ** @Desc : c file for ...
 ** @Time : 2019/10/23 15:20
 ** @Author : yuebin
 ** @File : page_controller
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/23 15:20
 ** @Software: GoLand
****************************************************/
package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type PageController struct {
	beego.Controller
}

func (c *PageController) Index() {
	c.TplName = "index.html"
}

func (c *PageController) LoginPage() {
	c.TplName = "login.html"
}

func (c *PageController) AccountPage() {
	c.TplName = "account.html"
}

func (c *PageController) AccountHistoryPage() {
	c.TplName = "account_history.html"
}

func (c *PageController) BankCardPage() {
	c.TplName = "bank_card.html"
}

func (c *PageController) CreateAgentPage() {
	c.TplName = "create_agent.html"
}

func (c *PageController) EditRolePage() {
	c.TplName = "edit_role.html"
}

func (c *PageController) FirstMenuPage() {
	c.TplName = "first_menu.html"
}

func (c *PageController) MainPage() {
	c.TplName = "main.html"
}

func (c *PageController) MenuPage() {
	c.TplName = "menu.html"
}

func (c *PageController) MerchantPage() {
	c.TplName = "merchant.html"
}

func (c *PageController) OperatorPage() {
	c.TplName = "operator.html"
}

func (c *PageController) PowerPage() {
	c.TplName = "power.html"
}

func (c *PageController) RoadPage() {
	c.TplName = "road.html"
}

func (c *PageController) RoadPoolPage() {
	c.TplName = "road_pool.html"
}

func (c *PageController) RoadProfitPage() {
	c.TplName = "road_profit.html"
}

func (c *PageController) RolePage() {
	c.TplName = "role.html"
}

func (c *PageController) SecondMenuPage() {
	c.TplName = "second_menu.html"
}

func (c *PageController) OrderInfoPage() {
	c.TplName = "order_info.html"
}

func (c *PageController) OrderProfitPage() {
	c.TplName = "order_profit.html"
}

func (c *PageController) MerchantPayforPage() {
	c.TplName = "merchant_payfor.html"
}

func (c *PageController) SelfPayforPage() {
	c.TplName = "self_payfor.html"
}

func (c *PageController) PayforRecordPage() {
	c.TplName = "payfor_record.html"
}

func (c *PageController) ConfirmPage() {
	c.TplName = "confirm.html"
}

func (c *PageController) SelfNotifyPage() {
	c.TplName = "self_notify.html"
}

func (c *PageController) SelfPlusSubPage() {
	c.TplName = "self_plus_sub.html"
}

func (c *PageController) AgentToMerchantPage() {
	c.TplName = "agent_to_merchant.html"
}

func (c *PageController) PlatFormProfitPage() {
	c.TplName = "platform_profit.html"
}

func (c *PageController) AgentProfitPage() {
	c.TplName = "agent_profit.html"
}
