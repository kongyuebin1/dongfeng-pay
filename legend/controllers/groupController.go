package controllers

import (
	"legend/controllers/base"
	"legend/service"
)

type GroupController struct {
	base.BasicController
}

func (c *GroupController) AddGroup() {
	groupName := c.GetString("groupName")

	se := new(service.GroupService)
	resp := se.GroupAdd(groupName)

	c.Data["json"] = resp
	_ = c.ServeJSON()
}

func (c *GroupController) ListGroup() {
	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")

	se := new(service.GroupService)
	list := se.GroupList(page, limit)

	c.Data["json"] = list
	_ = c.ServeJSON()
}

func (c *GroupController) DeleteGroup() {
	uid := c.GetString("uid")

	se := new(service.GroupService)
	resp := se.GroupDelete(uid)

	c.Data["json"] = resp
	_ = c.ServeJSON()
}

func (c *GroupController) EditGroup() {

	uid := c.GetString("uid")
	groupName := c.GetString("groupName")

	se := new(service.GroupService)
	resp := se.GroupEdit(uid, groupName)

	c.Data["json"] = resp
	_ = c.ServeJSON()
}
