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
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"dongfeng-pay/service/common"
	"dongfeng-pay/service/controller"
	"dongfeng-pay/service/models"
	"dongfeng-pay/service/utils"
	"strconv"
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

	dataJSON := new(KeyDataJSON)
	dataJSON.Code = -1
	if !ok || userID == "" {
		dataJSON.Code = 404
		dataJSON.Msg = "请重新登录!"
	} else {
		userInfo := models.GetUserInfoByUserID(userID)
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
			c.DelSession("userID")
			//更新数据库的密码
			userInfo.Passwd = utils.GetMD5Upper(newPassword)
			models.UpdateUserInfoPassword(userInfo)
		}
	}
	c.GenerateJSON(dataJSON)
}

/*
* 更新菜单的排列顺序
 */
func (c *UpdateController) UpMenu() {
	menuUid := c.GetString("menuUid")
	menuInfo := models.GetMenuInfoByMenuUid(menuUid)
	dataJSON := new(BaseDataJSON)
	if menuInfo.MenuUid == "" {
		dataJSON.Msg = "更改排列顺序失败"
		dataJSON.Code = -1
	} else {
		exist := models.MenuOrderIsExists(menuInfo.MenuOrder - 1)
		if !exist {
			dataJSON.Msg = "已经是最高的顺序"
			dataJSON.Code = -1
		} else {
			//如果他前面有菜单，那么交换他们的menuOrder
			preMenuInfo := models.GetMenuInfoByMenuOrder(menuInfo.MenuOrder - 1)
			menuInfo.MenuOrder = menuInfo.MenuOrder - 1
			preMenuInfo.MenuOrder = preMenuInfo.MenuOrder + 1
			preMenuInfo.UpdateTime = utils.GetBasicDateTime()
			menuInfo.UpdateTime = utils.GetBasicDateTime()
			//更新菜单表
			models.UpdateMenuInfo(preMenuInfo)
			models.UpdateMenuInfo(menuInfo)
			//更新二级菜单表
			SortSecondMenuOrder(preMenuInfo)
			SortSecondMenuOrder(menuInfo)
			dataJSON.Code = 200
		}
	}
	c.GenerateJSON(dataJSON)
}
func (c *UpdateController) DownMenu() {
	menuUid := c.GetString("menuUid")
	menuInfo := models.GetMenuInfoByMenuUid(menuUid)
	dataJSON := new(BaseDataJSON)
	if menuInfo.MenuUid == "" {
		dataJSON.Msg = "更改排列顺序失败"
		dataJSON.Code = -1
	} else {
		exist := models.MenuOrderIsExists(menuInfo.MenuOrder + 1)
		if !exist {
			dataJSON.Msg = "已经是最高的顺序"
			dataJSON.Code = -1
		} else {
			//如果他前面有菜单，那么交换他们的menuOrder
			lastMenuInfo := models.GetMenuInfoByMenuOrder(menuInfo.MenuOrder + 1)
			menuInfo.MenuOrder = menuInfo.MenuOrder + 1
			lastMenuInfo.MenuOrder = lastMenuInfo.MenuOrder - 1
			lastMenuInfo.UpdateTime = utils.GetBasicDateTime()
			menuInfo.UpdateTime = utils.GetBasicDateTime()
			//更新菜单表
			models.UpdateMenuInfo(lastMenuInfo)
			models.UpdateMenuInfo(menuInfo)
			//更新二级菜单表
			SortSecondMenuOrder(lastMenuInfo)
			SortSecondMenuOrder(menuInfo)
			dataJSON.Code = 200
		}
	}
	c.GenerateJSON(dataJSON)
}

/*
* 提升二级菜单的顺序号
 */
func (c *UpdateController) UpSecondMenu() {
	secondMenuUid := c.GetString("secondMenuUid")
	secondMenuInfo := models.GetSecondMenuInfoBySecondMenuUid(secondMenuUid)
	dataJSON := new(BaseDataJSON)
	if secondMenuInfo.MenuOrder == 1 {
		dataJSON.Code = -1
	} else {
		preSecondMenuInfo := models.GetSecondMenuInfoByMenuOrder(secondMenuInfo.MenuOrder-1, secondMenuInfo.FirstMenuUid)
		preSecondMenuInfo.MenuOrder = preSecondMenuInfo.MenuOrder + 1
		preSecondMenuInfo.UpdateTime = utils.GetBasicDateTime()
		secondMenuInfo.MenuOrder = secondMenuInfo.MenuOrder - 1
		secondMenuInfo.UpdateTime = utils.GetBasicDateTime()
		//更新二级菜单项
		models.UpdateSecondMenu(preSecondMenuInfo)
		models.UpdateSecondMenu(secondMenuInfo)

		dataJSON.Code = 200
	}
	c.GenerateJSON(dataJSON)
}

/*
* 降低二级菜单的顺序号
 */
func (c *UpdateController) DownSecondMenu() {
	secondMenuUid := c.GetString("secondMenuUid")
	secondMenuInfo := models.GetSecondMenuInfoBySecondMenuUid(secondMenuUid)

	dataJSON := new(BaseDataJSON)

	l := models.GetSecondMenuLenByFirstMenuUid(secondMenuInfo.FirstMenuUid)
	if l == secondMenuInfo.MenuOrder {
		dataJSON.Code = -1
	} else {
		lastSecondMenu := models.GetSecondMenuInfoByMenuOrder(secondMenuInfo.MenuOrder+1, secondMenuInfo.FirstMenuUid)
		lastSecondMenu.MenuOrder = lastSecondMenu.MenuOrder - 1
		lastSecondMenu.UpdateTime = utils.GetBasicDateTime()

		secondMenuInfo.MenuOrder = secondMenuInfo.MenuOrder + 1
		secondMenuInfo.UpdateTime = utils.GetBasicDateTime()

		models.UpdateSecondMenu(lastSecondMenu)
		models.UpdateSecondMenu(secondMenuInfo)

		dataJSON.Code = 200
	}
	c.GenerateJSON(dataJSON)
}

func (c *UpdateController) FreezeOperator() {
	userId := strings.TrimSpace(c.GetString("operatorName"))

	dataJSON := new(BaseDataJSON)

	if models.UpdateStauts("unactive", userId) {
		dataJSON.Code = 200
		dataJSON.Msg = "冻结成功"
	} else {
		dataJSON.Code = -1
		dataJSON.Msg = "冻结失败"
	}

	c.GenerateJSON(dataJSON)
}

func (c *UpdateController) UnfreezeOperator() {
	userId := strings.TrimSpace(c.GetString("operatorName"))

	dataJSON := new(BaseDataJSON)

	if models.UpdateStauts("active", userId) {
		dataJSON.Code = 200
		dataJSON.Msg = "解冻成功"
	} else {
		dataJSON.Code = -1
		dataJSON.Msg = "解冻失败"
	}
	c.GenerateJSON(dataJSON)
}

func (c *UpdateController) EditOperator() {
	userId := strings.TrimSpace(c.GetString("userId"))
	password := strings.TrimSpace(c.GetString("password"))
	changePassword := strings.TrimSpace(c.GetString("changePassword"))
	role := strings.TrimSpace(c.GetString("role"))
	nick := strings.TrimSpace(c.GetString("nick"))
	remark := strings.TrimSpace(c.GetString("remark"))

	keyDataJSON := new(KeyDataJSON)

	if (len(password) > 0 || len(changePassword) > 0) && password != changePassword {
		keyDataJSON.Code = -1
		keyDataJSON.Key = ".veritfy-operator-password-error"
		keyDataJSON.Msg = "*2次密码输入不一致"
		c.GenerateJSON(keyDataJSON)
	}

	if role == "" || role == "none" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = ".change-operator-role-error"
		keyDataJSON.Msg = "*角色不能为空"
		c.GenerateJSON(keyDataJSON)
	}

	userInfo := models.GetUserInfoByUserID(userId)
	if userInfo.UserId == "" {
		keyDataJSON.Code = -2
		keyDataJSON.Msg = "该用户不存在"
	} else {
		userInfo.UpdateTime = utils.GetBasicDateTime()
		userInfo.Remark = remark
		roleInfo := models.GetRoleByRoleUid(role)
		userInfo.RoleName = roleInfo.RoleName
		userInfo.Role = role
		if len(password) > 0 && len(changePassword) > 0 && password == changePassword {
			userInfo.Passwd = utils.GetMD5Upper(password)
		}
		userInfo.Nick = nick
		models.UpdateUserInfo(userInfo)
		keyDataJSON.Code = 200
	}

	c.GenerateJSON(keyDataJSON)
}

/*
* 更新通道的状态
 */
func (c *UpdateController) UpdateRoadStatus() {
	roadUid := strings.TrimSpace(c.GetString("roadUid"))

	dataJSON := new(BaseDataJSON)
	dataJSON.Code = 200

	roadInfo := models.GetRoadInfoByRoadUid(roadUid)
	if roadInfo.Status == "active" {
		roadInfo.Status = "unactive"
	} else {
		roadInfo.Status = "active"
	}
	if models.UpdateRoadInfo(roadInfo) {
		dataJSON.Code = 200
	} else {
		dataJSON.Code = -1
	}
	c.GenerateJSON(dataJSON)
}

/*
* 冻结商户
 */
func (c *UpdateController) UpdateMerchantStatus() {
	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))
	keyDataJSON := new(KeyDataJSON)
	if merchantUid == "" {
		keyDataJSON.Code = -1
		c.GenerateJSON(keyDataJSON)
		return
	}

	merchantInfo := models.GetMerchantByUid(merchantUid)

	if merchantInfo.MerchantUid == "" {
		keyDataJSON.Code = -1
		c.GenerateJSON(keyDataJSON)
		return
	}

	if merchantInfo.Status == "active" {
		merchantInfo.Status = "unactive"
	} else {
		merchantInfo.Status = "active"
	}
	merchantInfo.UpdateTime = utils.GetBasicDateTime()

	if models.UpdateMerchant(merchantInfo) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
	}
	c.GenerateJSON(keyDataJSON)
}

/*
* 更新账户的状态
 */
func (c *UpdateController) UpdateAccountStatus() {
	accountUid := strings.TrimSpace(c.GetString("accountUid"))

	accountInfo := models.GetAccountByUid(accountUid)
	if accountInfo.Status == "active" {
		accountInfo.Status = "unactive"
	} else {
		accountInfo.Status = "active"
	}
	accountInfo.UpdateTime = utils.GetBasicDateTime()

	dataJSON := new(BaseDataJSON)
	if models.UpdateAccount(accountInfo) {
		dataJSON.Code = 200
		dataJSON.Msg = "更新账户状态成功"
	} else {
		dataJSON.Code = -1
		dataJSON.Msg = "更新账户状态失败"
	}
	c.GenerateJSON(dataJSON)
}
func (c *UpdateController) OperatorAccount() {
	accountUid := strings.TrimSpace(c.GetString("accountUid"))
	accountOperator := strings.TrimSpace(c.GetString("accountOperator"))
	amount := strings.TrimSpace(c.GetString("amount"))

	accountDataJSON := new(AccountDataJSON)
	switch accountOperator {
	case common.PLUS_AMOUNT:
	case common.SUB_AMOUNT:
	case common.FREEZE_AMOUNT:
	case common.UNFREEZE_AMOUNT:
	default:
		accountDataJSON.Code = -1
	}
	a, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		accountDataJSON.Msg = "处理金额输入有误"
	}
	if accountDataJSON.Code == -1 {
		c.GenerateJSON(accountDataJSON)
		return
	}
	msg, flag := models.OperatorAccount(accountUid, accountOperator, a)
	if flag {
		accountDataJSON.Code = 200
		accountDataJSON.Msg = "处理成功，请检查对应账户信息"
		accountDataJSON.AccountList = append(accountDataJSON.AccountList, models.GetAccountByUid(accountUid))
	} else {
		accountDataJSON.Code = -1
		accountDataJSON.Msg = msg
	}

	c.GenerateJSON(accountDataJSON)
}

func (c *UpdateController) UpdateAgentStatus() {
	agentUid := strings.TrimSpace(c.GetString("agentUid"))
	agentInfo := models.GetAgentInfoByAgentUid(agentUid)

	keyDataJSON := new(KeyDataJSON)

	if agentInfo.AgentUid == "" {
		keyDataJSON.Code = -1
		c.GenerateJSON(keyDataJSON)
	}

	if agentInfo.Status == "active" {
		agentInfo.Status = "unactive"
	} else {
		agentInfo.Status = "active"
	}
	agentInfo.UpdateTime = utils.GetBasicDateTime()
	if models.UpdateAgentInfo(agentInfo) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
	}
	c.GenerateJSON(keyDataJSON)
}

func (c *UpdateController) ResetAgentPassword() {
	agentUid := strings.TrimSpace(c.GetString("agentUid"))
	newPassword := strings.TrimSpace(c.GetString("newPassword"))
	newVertifyPassword := strings.TrimSpace(c.GetString("newVertifyPassword"))

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = 200
	if agentUid == "" {
		keyDataJSON.Code = -2
	} else if newPassword == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-login-password-error-reset"
		keyDataJSON.Msg = " *新密码不能为空"
	} else if newVertifyPassword != newPassword {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-vertify-password-error-reset"
		keyDataJSON.Msg = " *两次密码输入不一致"
	}

	if keyDataJSON.Code != 200 {
		c.GenerateJSON(keyDataJSON)
		return
	}

	agentInfo := models.GetAgentInfoByAgentUid(agentUid)
	agentInfo.UpdateTime = utils.GetBasicDateTime()
	agentInfo.AgentPassword = utils.GetMD5Upper(newPassword)
	if !models.UpdateAgentInfo(agentInfo) {
		keyDataJSON.Code = -1
	}
	c.GenerateJSON(keyDataJSON)
}

/*
* 手动选择了打款通道
 */
func (c *UpdateController) ChoosePayForRoad() {
	roadName := strings.TrimSpace(c.GetString("roadName"))
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))
	remark := strings.TrimSpace(c.GetString("remark"))
	confirmType := strings.TrimSpace(c.GetString("confirmType"))

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = 200

	if confirmType == common.PAYFOR_ROAD && roadName == "" {
		keyDataJSON.Msg = "打款通道不能为空"
		keyDataJSON.Code = -1
		c.GenerateJSON(keyDataJSON)
		return
	}

	payForInfo := models.GetPayForByBankOrderId(bankOrderId)
	roadInfo := models.GetRoadInfoByName(roadName)

	if payForInfo.Status != common.PAYFOR_COMFRIM {
		keyDataJSON.Msg = "结算状态错误，请刷新后确认"
	} else {
		payForInfo.UpdateTime = utils.GetBasicDateTime()
		payForInfo.GiveType = confirmType
		if confirmType == common.PAYFOR_REFUSE {
			//拒绝打款
			payForInfo.Status = common.PAYFOR_FAIL
		} else {
			payForInfo.Status = common.PAYFOR_SOLVING
		}
		payForInfo.RoadUid = roadInfo.RoadUid
		payForInfo.RoadName = roadInfo.RoadName
		payForInfo.Remark = remark

		if !models.ForUpdatePayFor(payForInfo) {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "更新代付记录失败"
		}
	}

	c.GenerateJSON(keyDataJSON)
}

/*
* 处理打款结果的处理
 */
func (c *UpdateController) ResultPayFor() {
	resultType := strings.TrimSpace(c.GetString("resultType"))
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = 200

	if resultType == "" || bankOrderId == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "提交的数据有误"
		c.GenerateJSON(keyDataJSON)
		return
	}

	payFor := models.GetPayForByBankOrderId(bankOrderId)

	if payFor.Type == common.SELF_HELP {
		//如果是管理员在后台提现，不用做任何的商户减款,只需要更新代付订单状态
		payFor.UpdateTime = utils.GetBasicDateTime()
		payFor.Status = resultType

		if !models.ForUpdatePayFor(payFor) {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "系统处理失败"
		}
		c.GenerateJSON(keyDataJSON)
		return
	}

	if payFor.Status == common.PAYFOR_FAIL || payFor.Status == common.PAYFOR_SUCCESS {
		logs.Error(fmt.Sprintf("该代付订单=%s，状态有误....", bankOrderId))
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "订单状态有误，请刷新重新判断"
		c.GenerateJSON(keyDataJSON)
		return
	}

	if resultType == common.PAYFOR_FAIL {
		//处理代付失败的逻辑，减去相应的代付冻结金额
		if !controller.PayForFail(payFor) {
			logs.Error(fmt.Sprintf("商户uid=%s,处理代付失败逻辑出错", payFor.MerchantUid))
			keyDataJSON.Msg = "代付失败逻辑，处理失败"
			keyDataJSON.Code = -1
		}
	} else if resultType == common.PAYFOR_SUCCESS {
		//代付成功，减去相应的代付冻结金额，并且余额减掉，可用金额减掉
		if !controller.PayForSuccess(payFor) {
			logs.Error(fmt.Sprintf("商户uid=%s,处理代付成功逻辑出错", payFor.MerchantUid))
			keyDataJSON.Msg = "代付成功逻辑，处理失败"
			keyDataJSON.Code = -1
		}
	}

	if keyDataJSON.Code == 200 {
		keyDataJSON.Msg = "处理成功"
	}

	c.GenerateJSON(keyDataJSON)
}

func (c *UpdateController) UpdateOrderStatus() {
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))
	solveType := strings.TrimSpace(c.GetString("solveType"))

	keyDataJSON := new(KeyDataJSON)
	orderInfo := models.GetOrderByBankOrderId(bankOrderId)
	if orderInfo.BankOrderId == "" {
		logs.Error("该订单不存在,bankOrderId=", bankOrderId)
		keyDataJSON.Code = -1
	} else {
		paySolve := new(controller.PaySolveController)
		flag := false
		switch solveType {
		case common.SUCCESS:
			flag = paySolve.SolvePaySuccess(bankOrderId, orderInfo.FactAmount, common.SUCCESS)
		case common.FAIL:
			flag = paySolve.SolvePayFail(orderInfo, common.FAIL)
		case common.FREEZE_AMOUNT:
			//将这笔订单进行冻结
			flag = paySolve.SolveOrderFreeze(bankOrderId)
		case common.UNFREEZE_AMOUNT:
			//将这笔订单金额解冻
			flag = paySolve.SolveOrderUnfreeze(bankOrderId)
		case common.REFUND:
			if orderInfo.Status == common.SUCCESS {
				flag = paySolve.SolveRefund(bankOrderId)
			}
		case common.ORDERROLL:
			if orderInfo.Status == common.SUCCESS {
				flag = paySolve.SolveOrderRoll(bankOrderId)
			}
		default:
			logs.Error("不存在这样的处理类型")
		}
		if flag {
			keyDataJSON.Code = 200
		} else {
			keyDataJSON.Code = -1
		}
	}

	c.GenerateJSON(keyDataJSON)
}
