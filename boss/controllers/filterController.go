/***************************************************
 ** @Desc : 过滤功能
 ** @Time : 2019/8/8 16:10
 ** @Author : yuebin
 ** @File : filter
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/8 16:10
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type FilterController struct {
	web.Controller
}

//jsonp请求过来的函数
func (c *FilterController) Filter() {
	userID, ok := c.GetSession("userID").(string)

	dataJSON := new(struct {
		Code int
	})

	if !ok || userID == "" {
		dataJSON.Code = 404
	} else {
		dataJSON.Code = 200
		_ = c.SetSession("userID", userID)
	}
	fmt.Println(dataJSON)
	c.Data["json"] = dataJSON
	_ = c.ServeJSON()
}
