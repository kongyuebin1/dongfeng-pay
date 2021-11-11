/***************************************************
 ** @Desc : c file for ...
 ** @Time : 2019/8/16 9:49
 ** @Author : yuebin
 ** @File : update
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/16 9:49
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"boss/datas"
	"boss/models/user"
	"boss/service"
	"boss/utils"
	"github.com/beego/beego/v2/adapter/validation"
	"strings"
)

type UpdateController struct {
	BaseController
}

/*
*更新密码
 */
func (c *UpdateController) UpdatePassword() {
	oldPassword := c.GetString("oldPassword")
	newPassword := c.GetString("newPassword")
	twicePassword := c.GetString("twicePassword")

	userID, ok := c.GetSession("userID").(string)

	dataJSON := new(datas.KeyDataJSON)
	dataJSON.Code = -1
	if !ok || userID == "" {
		dataJSON.Code = 404
		dataJSON.Msg = "请重新登录!"
	} else {
		userInfo := user.GetUserInfoByUserID(userID)
		valid := validation.Validation{}
		if userInfo.Passwd != utils.GetMD5Upper(oldPassword) {
			dataJSON.Key = ".old-error"
			dataJSON.Msg = "输入密码不正确"
		} else if v := valid.Min(len(newPassword), 8, ".new-error"); !v.Ok {
			dataJSON.Key = v.Error.Key
			dataJSON.Msg = "新密码长度必须大于等于8个字符!"
		} else if v := valid.Max(len(newPassword), 16, ".new-error"); !v.Ok {
			dataJSON.Key = v.Error.Key
			dataJSON.Msg = "新密码长度不能大于16个字符!"
		} else if v := valid.AlphaNumeric(newPassword, ".new-error"); !v.Ok {
			dataJSON.Key = v.Error.Key
			dataJSON.Msg = "新密码必须有数字和字母组成!"
		} else if newPassword != twicePassword {
			dataJSON.Key = ".twice-error"
			dataJSON.Msg = "两次密码不一致!"
		} else {
			dataJSON.Code = 200
			dataJSON.Msg = "密码修改成功!"
			//删除原先的session状态
			_ = c.DelSession("userID")
			//更新数据库的密码
			userInfo.Passwd = utils.GetMD5Upper(newPassword)
			user.UpdateUserInfoPassword(userInfo)
		}
	}
	c.GenerateJSON(dataJSON)
}

/*
* 更新菜单的排列顺序
 */
func (c *UpdateController) UpMenu() {
	menuUid := c.GetString("menuUid")
	se := new(service.UpdateService)
	dataJSON := se.UpMenu(menuUid)

	c.GenerateJSON(dataJSON)
}
func (c *UpdateController) DownMenu() {
	menuUid := c.GetString("menuUid")
	se := new(service.UpdateService)
	dataJSON := se.DownMenu(menuUid)

	c.GenerateJSON(dataJSON)
}

/*
* 提升二级菜单的顺序号
 */
func (c *UpdateController) UpSecondMenu() {
	secondMenuUid := c.GetString("secondMenuUid")
	se := new(service.UpdateService)
	dataJSON := se.UpSecondMenu(secondMenuUid)

	c.GenerateJSON(dataJSON)
}

/*
* 降低二级菜单的顺序号
 */
func (c *UpdateController) DownSecondMenu() {
	secondMenuUid := c.GetString("secondMenuUid")
	se := new(service.UpdateService)
	dataJSON := se.DownSecondMenu(secondMenuUid)
	c.GenerateJSON(dataJSON)
}

func (c *UpdateController) FreezeOperator() {
	userId := strings.TrimSpace(c.GetString("operatorName"))
	se := new(service.UpdateService)
	dataJSON := se.FreezeOperator(userId)

	c.GenerateJSON(dataJSON)
}

func (c *UpdateController) UnfreezeOperator() {
	userId := strings.TrimSpace(c.GetString("operatorName"))

	se := new(service.UpdateService)
	dataJSON := se.UnfreezeOperator(userId)

	c.GenerateJSON(dataJSON)
}

func (c *UpdateController) EditOperator() {
	userId := strings.TrimSpace(c.GetString("userId"))
	password := strings.TrimSpace(c.GetString("password"))
	changePassword := strings.TrimSpace(c.GetString("changePassword"))
	role := strings.TrimSpace(c.GetString("role"))
	nick := strings.TrimSpace(c.GetString("nick"))
	remark := strings.TrimSpace(c.GetString("remark"))

	se := new(service.UpdateService)
	keyDataJSON := se.EditOperator(password, changePassword, role, userId, nick, remark)

	c.GenerateJSON(keyDataJSON)
}

/*
* 更新通道的状态
 */
func (c *UpdateController) UpdateRoadStatus() {
	roadUid := strings.TrimSpace(c.GetString("roadUid"))

	se := new(service.UpdateService)
	dataJSON := se.UpdateRoadStatus(roadUid)

	c.GenerateJSON(dataJSON)
}

/*
* 冻结商户
 */
func (c *UpdateController) UpdateMerchantStatus() {
	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))
	se := new(service.UpdateService)
	keyDataJSON := se.UpdateMerchantStatus(merchantUid)

	c.GenerateJSON(keyDataJSON)
}

/*
* 更新账户的状态
 */
func (c *UpdateController) UpdateAccountStatus() {
	accountUid := strings.TrimSpace(c.GetString("accountUid"))

	se := new(service.UpdateService)
	dataJSON := se.UpdateAccountStatus(accountUid)

	c.GenerateJSON(dataJSON)
}
func (c *UpdateController) OperatorAccount() {
	accountUid := strings.TrimSpace(c.GetString("accountUid"))
	accountOperator := strings.TrimSpace(c.GetString("accountOperator"))
	amount := strings.TrimSpace(c.GetString("amount"))

	se := new(service.UpdateService)
	accountDataJSON := se.OperatorAccount(accountOperator, amount, accountUid)

	c.GenerateJSON(accountDataJSON)
}

func (c *UpdateController) UpdateAgentStatus() {
	agentUid := strings.TrimSpace(c.GetString("agentUid"))
	se := new(service.UpdateService)
	c.GenerateJSON(se.UpdateAgentStatus(agentUid))
}

func (c *UpdateController) ResetAgentPassword() {
	agentUid := strings.TrimSpace(c.GetString("agentUid"))
	newPassword := strings.TrimSpace(c.GetString("newPassword"))
	newVertifyPassword := strings.TrimSpace(c.GetString("newVertifyPassword"))

	se := new(service.UpdateService)

	c.GenerateJSON(se.ResetAgentPassword(agentUid, newPassword, newVertifyPassword))
}

/*
* 手动选择了打款通道
 */
func (c *UpdateController) ChoosePayForRoad() {
	roadName := strings.TrimSpace(c.GetString("roadName"))
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))
	remark := strings.TrimSpace(c.GetString("remark"))
	confirmType := strings.TrimSpace(c.GetString("confirmType"))

	se := new(service.UpdateService)

	c.GenerateJSON(se.ChoosePayForRoad(confirmType, roadName, bankOrderId, remark))
}

/*
* 处理打款结果的处理
 */
func (c *UpdateController) ResultPayFor() {
	resultType := strings.TrimSpace(c.GetString("resultType"))
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))

	se := new(service.UpdateService)

	c.GenerateJSON(se.ResultPayFor(resultType, bankOrderId))
}

func (c *UpdateController) UpdateOrderStatus() {
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))
	solveType := strings.TrimSpace(c.GetString("solveType"))

	se := new(service.UpdateService)

	c.GenerateJSON(se.UpdateOrderStatus(bankOrderId, solveType))
}
