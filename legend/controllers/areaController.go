package controllers

import (
	"legend/controllers/base"
	"legend/request"
	"legend/service"
)

type AreaController struct {
	base.BasicController
}

func (c *AreaController) getRequestPrams() *request.AreaReq {
	req := new(request.AreaReq)
	req.AreaName = c.GetString("areaName")
	req.GroupName = c.GetString("groupName")
	req.TemplateName = c.GetString("templateName")
	req.NotifyUrl = c.GetString("notifyUrl")
	req.AttachParams = c.GetString("attachParams")
	return req
}

func (c *AreaController) AreaAdd() {

	req := c.getRequestPrams()

	se := new(service.AreaService)
	area := se.AddArea(req)

	c.Data["json"] = area
	_ = c.ServeJSON()
}

func (c *AreaController) AreaEdit() {
	req := c.getRequestPrams()
	uid := c.GetString("uid")

	se := new(service.AreaService)
	resp := se.EditArea(req, uid)

	c.Data["json"] = resp

	_ = c.ServeJSON()
}

func (c *AreaController) AreaList() {
	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")

	se := new(service.AreaService)
	list := se.AreaList(page, limit)

	c.Data["json"] = list
	_ = c.ServeJSON()
}

func (c *AreaController) AreaDelete() {
	uid := c.GetString("uid")

	se := new(service.AreaService)
	resp := se.DeleteArea(uid)
	c.Data["json"] = resp

	_ = c.ServeJSON()
}

func (c *AreaController) AreaGet() {
	uid := c.GetString("uid")

	se := new(service.AreaService)
	resp := se.GetArea(uid)

	c.Data["json"] = resp
	_ = c.ServeJSON()
}
