package controllers

import (
	"legend/controllers/base"
)

type IndexController struct {
	base.BasicController
}

/**
** 用户登录后跳转的页面，也是后台的整个主题框架
 */
func (c *IndexController) Index() {
	c.TplName = "index.html"
}
