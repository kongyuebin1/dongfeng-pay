package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"legend/controllers/base"
	"legend/request"
	"legend/service"
	"legend/utils"
	"strings"
)

type TemplateController struct {
	base.BasicController
}

func (c *TemplateController) TemplateAdd() {

	addTemplate := new(request.AddTemplateReq)
	if err := c.ParseForm(addTemplate); err != nil {
		logs.Error("错误：", err)
	}

	addTemplate.FixPrices = utils.StringToFloats(c.GetString("fixPrices"))
	addTemplate.PresentFixMoneys = utils.StringToFloats(c.GetString("presentFixMoneys"))
	addTemplate.PresentFixPresentMoneys = utils.StringToFloats(c.GetString("presentFixPresentMoneys"))
	addTemplate.PresentScaleMoneys = utils.StringToFloats(c.GetString("presentScaleMoneys"))
	addTemplate.PresentScales = utils.StringToFloats(c.GetString("presentScales"))

	addTemplate.FixUids = strings.Split(c.GetString("fixUids"), ",")
	addTemplate.GoodsNames = strings.Split(c.GetString("goodsNames"), ",")
	addTemplate.GoodsNos = strings.Split(c.GetString("goodsNos"), ",")

	addTemplate.PresentFixUids = strings.Split(c.GetString("presentFixUids"), ",")
	addTemplate.PresentScaleUids = strings.Split(c.GetString("presentScaleUids"), ",")
	addTemplate.Limits = utils.StringToInt(c.GetString("limits"))

	se := new(service.TemplateService)
	merchantUid := c.Data["merchantUid"].(string)
	t := c.GetString("type")
	if t == "edit" {
		c.Data["json"] = se.UpdateTemplate(addTemplate, merchantUid)
	} else {
		c.Data["json"] = se.AddTemplate(addTemplate, merchantUid)
	}

	_ = c.ServeJSON()
}

func (c *TemplateController) TemplateList() {

	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")

	se := new(service.TemplateService)
	list := se.GetTemplateList(page, limit)

	c.Data["json"] = list
	_ = c.ServeJSON()
}

func (c *TemplateController) TemplateDelete() {
	templateName := c.GetString("TemplateName")

	logs.Debug("template TemplateName ：", templateName)

	se := new(service.TemplateService)
	baseResp := se.DeleteTemplate(templateName)

	c.Data["json"] = baseResp
	_ = c.ServeJSON()
}

func (c *TemplateController) TemplateAllInfo() {
	templateName := c.GetString("scaleTemplateName")
	logs.Debug("获取到的scaleTemplateName：", templateName)

	se := new(service.TemplateService)
	allInfo := se.AllTemplateInfo(templateName)

	logs.Debug("scale template all info：", fmt.Sprintf("%+v", allInfo))

	c.Data["json"] = allInfo

	_ = c.ServeJSON()
}
