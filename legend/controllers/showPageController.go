package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"legend/controllers/base"
	"legend/service"
	"legend/utils"
)

type ShowPageController struct {
	base.BasicController
}

func (c *ShowPageController) FaviconPage() {
	c.TplName = "favicon.png"
}

/**
** 展示后台第一个页面
 */
func (c *ShowPageController) WelcomePage() {

	userName := c.GetSession("userName").(string)

	fmt.Println(userName)

	accountService := new(service.AccountService)
	accountInfo := accountService.GetAccountInfo(userName)

	logs.Debug("account信息：", accountInfo)

	c.Data["balance"] = accountInfo.Balance
	c.Data["unBalance"] = accountInfo.FreezeAmount
	c.Data["settleAmount"] = accountInfo.SettleAmount
	// 获取今天充值金额
	todayIncome := accountService.GetTodayIncome()
	logs.Debug("获取到当天充值金额：", todayIncome)
	c.Data["todayAmount"] = todayIncome

	c.TplName = "welcome.html"
}

/**
** 展示商户密钥
 */
func (c *ShowPageController) MerchantKeyPage() {
	userName := c.GetSession("userName").(string)

	merchantService := new(service.MerchantService)
	userInfo, bankInfo := merchantService.GetMerchantBankInfo(userName)

	c.Data["currentTime"] = utils.GetNowTime()
	c.Data["userName"] = userName
	c.Data["userInfo"] = userInfo
	c.Data["bankInfo"] = bankInfo
	c.TplName = "merchant-key.html"
}

/**
** 比例模板
 */
func (c *ShowPageController) ScaleTemplatePage() {
	c.TplName = "scale-template.html"
}

/**
** 增加模板
 */
func (c *ShowPageController) TemplateAdd() {
	c.TplName = "template-add.html"
}

func (c *ShowPageController) TemplateEdit() {
	templateName := c.GetString("templateName")
	fmt.Println(templateName)
	c.Data["scaleTemplateName"] = templateName
	c.TplName = "template-edit.html"
}

/**
** 分组列表
 */
func (c *ShowPageController) GroupListPage() {
	c.TplName = "group-list.html"
}

/**
** 分区列表
 */
func (c *ShowPageController) AreaListPage() {
	c.TplName = "area-list.html"
}

func (c *ShowPageController) AreaEdit() {
	uid := c.GetString("uid")
	c.Data["uid"] = uid
	c.TplName = "area-edit.html"
}

/**
** 充值订单
 */
func (c *ShowPageController) OrderListPage() {
	c.TplName = "order-list.html"
}

/**
** 模拟充值
 */
func (c *ShowPageController) ImitateChargePage() {
	c.TplName = "imitate-order.html"
}

/**
** 结算管理
 */
func (c *ShowPageController) SettleListPage() {
	c.TplName = "settle-list.html"
}

/**
**每日充值统计
 */
func (c *ShowPageController) EverydayChargeCountPage() {
	c.TplName = "everyday-charge-count.html"
}

/**
** 分组充值统计
 */
func (c *ShowPageController) GroupChargeCountPage() {
	c.TplName = "group-charge-count.html"
}

/**
** 分区充值统计
 */
func (c *ShowPageController) AreaChargePage() {
	c.TplName = "area-charge-count.html"
}

/**
** 创建分区和编辑分区
 */
func (c *ShowPageController) AreaAddOrEdit() {
	c.TplName = "area-add.html"
}

/**
** 个人页面
 */
func (c *ShowPageController) PersonPage() {
	userName, ok := c.GetSession("userName").(string)
	if !ok {
		c.Abort("404")
	} else {
		merchantService := new(service.MerchantService)
		userInfo := merchantService.MerchantInfo(userName)
		c.Data["userName"] = userInfo.MerchantName
	}

	c.TplName = "person.html"
}
