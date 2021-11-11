/***************************************************
 ** @Desc : c file for ...
 ** @Time : 2019/8/21 16:51
 ** @Author : yuebin
 ** @File : delete
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/21 16:51
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"boss/service"
	"strings"
)

type DeleteController struct {
	BaseController
}

func (c *DeleteController) Finish() {
	se := new(service.DeleteService)
	se.Finish()
}

func (c *DeleteController) DeleteMenu() {
	menuUid := c.GetString("menuUid")
	se := new(service.DeleteService)
	dataJSON := se.DeleteMenu(menuUid, c.GetSession("userID").(string))

	c.Data["json"] = dataJSON
	_ = c.ServeJSONP()
}

func (c *DeleteController) DeleteSecondMenu() {
	secondMenuUid := strings.TrimSpace(c.GetString("secondMenuUid"))
	se := new(service.DeleteService)
	dataJSON := se.DeleteSecondMenu(secondMenuUid)

	c.Data["json"] = dataJSON
	_ = c.ServeJSON()
}

/*
* 删除权限项
 */
func (c *DeleteController) DeletePowerItem() {
	powerID := strings.TrimSpace(c.GetString("powerID"))
	se := new(service.DeleteService)
	dataJSON := se.DeletePowerItem(powerID)
	c.GenerateJSON(dataJSON)
}

/*
* 删除角色
 */
func (c *DeleteController) DeleteRole() {
	roleUid := strings.TrimSpace(c.GetString("roleUid"))
	se := new(service.DeleteService)
	dataJSON := se.DeleteRole(roleUid)
	c.GenerateJSON(dataJSON)
}

/*
* 删除操作员
 */
func (c *DeleteController) DeleteOperator() {
	userId := strings.TrimSpace(c.GetString("userId"))
	se := new(service.DeleteService)
	dataJSON := se.DeleteOperator(userId)

	c.GenerateJSON(dataJSON)
}

func (c *DeleteController) DeleteBankCardRecord() {
	uid := strings.TrimSpace(c.GetString("uid"))
	se := new(service.DeleteService)

	dataJSON := se.DeleteBankCardRecord(uid)

	c.GenerateJSON(dataJSON)
}

/*
* 删除通道操作
 */
func (c *DeleteController) DeleteRoad() {
	roadUid := strings.TrimSpace(c.GetString("roadUid"))

	se := new(service.DeleteService)
	dataJSON := se.DeleteRoad(roadUid)

	c.GenerateJSON(dataJSON)
}

/*
* 删除通道池
 */
func (c *DeleteController) DeleteRoadPool() {
	roadPoolCode := strings.TrimSpace(c.GetString("roadPoolCode"))

	se := new(service.DeleteService)
	dataJSON := se.DeleteRoadPool(roadPoolCode)

	c.GenerateJSON(dataJSON)
}

/*
* 删除商户
 */
func (c *DeleteController) DeleteMerchant() {
	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))
	se := new(service.DeleteService)
	keyDataJSON := se.DeleteMerchant(merchantUid)

	c.GenerateJSON(keyDataJSON)
}

/*
* 删除账户
 */
func (c *DeleteController) DeleteAccount() {
	accountUid := strings.TrimSpace(c.GetString("accountUid"))
	se := new(service.DeleteService)
	dataJSON := se.DeleteAccount(accountUid)

	c.GenerateJSON(dataJSON)
}

func (c *DeleteController) DeleteAgent() {
	agentUid := strings.TrimSpace(c.GetString("agentUid"))

	se := new(service.DeleteService)
	keyDataJSON := se.DeleteAgent(agentUid)

	c.GenerateJSON(keyDataJSON)
}

func (c *DeleteController) DeleteAgentRelation() {

	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))

	se := new(service.DeleteService)

	keyDataJSON := se.DeleteAgentRelation(merchantUid)

	c.GenerateJSON(keyDataJSON)
}
