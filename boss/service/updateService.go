package service

import (
	"boss/common"
	"boss/datas"
	"boss/models"
	"boss/models/accounts"
	"boss/models/agent"
	"boss/models/merchant"
	"boss/models/payfor"
	"boss/models/road"
	"boss/models/system"
	"boss/models/user"
	"boss/utils"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type UpdateService struct {
}

func (c *UpdateService) UpMenu(menuUid string) *datas.BaseDataJSON {
	menuInfo := system.GetMenuInfoByMenuUid(menuUid)
	dataJSON := new(datas.BaseDataJSON)
	if menuInfo.MenuUid == "" {
		dataJSON.Msg = "更改排列顺序失败"
		dataJSON.Code = -1
	} else {
		exist := system.MenuOrderIsExists(menuInfo.MenuOrder - 1)
		if !exist {
			dataJSON.Msg = "已经是最高的顺序"
			dataJSON.Code = -1
		} else {
			//如果他前面有菜单，那么交换他们的menuOrder
			preMenuInfo := system.GetMenuInfoByMenuOrder(menuInfo.MenuOrder - 1)
			menuInfo.MenuOrder = menuInfo.MenuOrder - 1
			preMenuInfo.MenuOrder = preMenuInfo.MenuOrder + 1
			preMenuInfo.UpdateTime = utils.GetBasicDateTime()
			menuInfo.UpdateTime = utils.GetBasicDateTime()
			//更新菜单表
			system.UpdateMenuInfo(preMenuInfo)
			system.UpdateMenuInfo(menuInfo)
			//更新二级菜单表
			SortSecondMenuOrder(preMenuInfo)
			SortSecondMenuOrder(menuInfo)
			dataJSON.Code = 200
		}
	}
	return dataJSON
}

func (c *UpdateService) DownMenu(menuUid string) *datas.BaseDataJSON {
	menuInfo := system.GetMenuInfoByMenuUid(menuUid)
	dataJSON := new(datas.BaseDataJSON)
	if menuInfo.MenuUid == "" {
		dataJSON.Msg = "更改排列顺序失败"
		dataJSON.Code = -1
	} else {
		exist := system.MenuOrderIsExists(menuInfo.MenuOrder + 1)
		if !exist {
			dataJSON.Msg = "已经是最高的顺序"
			dataJSON.Code = -1
		} else {
			//如果他前面有菜单，那么交换他们的menuOrder
			lastMenuInfo := system.GetMenuInfoByMenuOrder(menuInfo.MenuOrder + 1)
			menuInfo.MenuOrder = menuInfo.MenuOrder + 1
			lastMenuInfo.MenuOrder = lastMenuInfo.MenuOrder - 1
			lastMenuInfo.UpdateTime = utils.GetBasicDateTime()
			menuInfo.UpdateTime = utils.GetBasicDateTime()
			//更新菜单表
			system.UpdateMenuInfo(lastMenuInfo)
			system.UpdateMenuInfo(menuInfo)
			//更新二级菜单表
			SortSecondMenuOrder(lastMenuInfo)
			SortSecondMenuOrder(menuInfo)
			dataJSON.Code = 200
		}
	}
	return dataJSON
}

func (c *UpdateService) UpSecondMenu(secondMenuUid string) *datas.BaseDataJSON {
	secondMenuInfo := system.GetSecondMenuInfoBySecondMenuUid(secondMenuUid)
	dataJSON := new(datas.BaseDataJSON)
	if secondMenuInfo.MenuOrder == 1 {
		dataJSON.Code = -1
	} else {
		preSecondMenuInfo := system.GetSecondMenuInfoByMenuOrder(secondMenuInfo.MenuOrder-1, secondMenuInfo.FirstMenuUid)
		preSecondMenuInfo.MenuOrder = preSecondMenuInfo.MenuOrder + 1
		preSecondMenuInfo.UpdateTime = utils.GetBasicDateTime()
		secondMenuInfo.MenuOrder = secondMenuInfo.MenuOrder - 1
		secondMenuInfo.UpdateTime = utils.GetBasicDateTime()
		//更新二级菜单项
		system.UpdateSecondMenu(preSecondMenuInfo)
		system.UpdateSecondMenu(secondMenuInfo)

		dataJSON.Code = 200
	}
	return dataJSON
}

func (c *UpdateService) DownSecondMenu(secondMenuUid string) *datas.BaseDataJSON {
	secondMenuInfo := system.GetSecondMenuInfoBySecondMenuUid(secondMenuUid)

	dataJSON := new(datas.BaseDataJSON)

	l := system.GetSecondMenuLenByFirstMenuUid(secondMenuInfo.FirstMenuUid)
	if l == secondMenuInfo.MenuOrder {
		dataJSON.Code = -1
	} else {
		lastSecondMenu := system.GetSecondMenuInfoByMenuOrder(secondMenuInfo.MenuOrder+1, secondMenuInfo.FirstMenuUid)
		lastSecondMenu.MenuOrder = lastSecondMenu.MenuOrder - 1
		lastSecondMenu.UpdateTime = utils.GetBasicDateTime()

		secondMenuInfo.MenuOrder = secondMenuInfo.MenuOrder + 1
		secondMenuInfo.UpdateTime = utils.GetBasicDateTime()

		system.UpdateSecondMenu(lastSecondMenu)
		system.UpdateSecondMenu(secondMenuInfo)

		dataJSON.Code = 200
	}
	return dataJSON
}

func (c *UpdateService) FreezeOperator(userId string) *datas.BaseDataJSON {
	dataJSON := new(datas.BaseDataJSON)

	if user.UpdateStauts(common.UNACTIVE, userId) {
		dataJSON.Code = 200
		dataJSON.Msg = "冻结成功"
	} else {
		dataJSON.Code = -1
		dataJSON.Msg = "冻结失败"
	}
	return dataJSON
}

func (c *UpdateService) UnfreezeOperator(userId string) *datas.BaseDataJSON {
	dataJSON := new(datas.BaseDataJSON)

	if user.UpdateStauts("active", userId) {
		dataJSON.Code = 200
		dataJSON.Msg = "解冻成功"
	} else {
		dataJSON.Code = -1
		dataJSON.Msg = "解冻失败"
	}

	return dataJSON
}

func (c *UpdateService) EditOperator(password, changePassword, role, userId, nick, remark string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)

	if (len(password) > 0 || len(changePassword) > 0) && password != changePassword {
		keyDataJSON.Code = -1
		keyDataJSON.Key = ".veritfy-operator-password-error"
		keyDataJSON.Msg = "*2次密码输入不一致"
		return keyDataJSON
	}

	if role == "" || role == "none" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = ".change-operator-role-error"
		keyDataJSON.Msg = "*角色不能为空"
		return keyDataJSON
	}

	userInfo := user.GetUserInfoByUserID(userId)
	if userInfo.UserId == "" {
		keyDataJSON.Code = -2
		keyDataJSON.Msg = "该用户不存在"
	} else {
		userInfo.UpdateTime = utils.GetBasicDateTime()
		userInfo.Remark = remark
		roleInfo := system.GetRoleByRoleUid(role)
		userInfo.RoleName = roleInfo.RoleName
		userInfo.Role = role
		if len(password) > 0 && len(changePassword) > 0 && password == changePassword {
			userInfo.Passwd = utils.GetMD5Upper(password)
		}
		userInfo.Nick = nick
		user.UpdateUserInfo(userInfo)
		keyDataJSON.Code = 200
	}
	return keyDataJSON

}

func (c *UpdateService) UpdateRoadStatus(roadUid string) *datas.BaseDataJSON {
	dataJSON := new(datas.BaseDataJSON)
	dataJSON.Code = 200

	roadInfo := road.GetRoadInfoByRoadUid(roadUid)
	if roadInfo.Status == "active" {
		roadInfo.Status = common.UNACTIVE
	} else {
		roadInfo.Status = "active"
	}
	if road.UpdateRoadInfo(roadInfo) {
		dataJSON.Code = 200
	} else {
		dataJSON.Code = -1
	}
	return dataJSON
}

func (c *UpdateService) UpdateMerchantStatus(merchantUid string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	if merchantUid == "" {
		keyDataJSON.Code = -1
		return keyDataJSON
	}

	merchantInfo := merchant.GetMerchantByUid(merchantUid)

	if merchantInfo.MerchantUid == "" {
		keyDataJSON.Code = -1
		return keyDataJSON
	}

	if merchantInfo.Status == common.ACTIVE {
		merchantInfo.Status = common.UNACTIVE
	} else {
		merchantInfo.Status = common.ACTIVE
	}
	merchantInfo.UpdateTime = utils.GetBasicDateTime()

	if merchant.UpdateMerchant(merchantInfo) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
	}
	return keyDataJSON
}

func (c *UpdateService) UpdateAccountStatus(accountUid string) *datas.BaseDataJSON {
	accountInfo := accounts.GetAccountByUid(accountUid)
	if accountInfo.Status == common.ACTIVE {
		accountInfo.Status = common.UNACTIVE
	} else {
		accountInfo.Status = common.ACTIVE
	}
	accountInfo.UpdateTime = utils.GetBasicDateTime()

	dataJSON := new(datas.BaseDataJSON)
	if accounts.UpdateAccount(accountInfo) {
		dataJSON.Code = 200
		dataJSON.Msg = "更新账户状态成功"
	} else {
		dataJSON.Code = -1
		dataJSON.Msg = "更新账户状态失败"
	}
	return dataJSON
}

func (c *UpdateService) OperatorAccount(accountOperator, amount, accountUid string) *datas.AccountDataJSON {
	accountDataJSON := new(datas.AccountDataJSON)
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
		return accountDataJSON
	}
	msg, flag := models.OperatorAccount(accountUid, accountOperator, a)
	if flag {
		accountDataJSON.Code = 200
		accountDataJSON.Msg = "处理成功，请检查对应账户信息"
		accountDataJSON.AccountList = append(accountDataJSON.AccountList, accounts.GetAccountByUid(accountUid))
	} else {
		accountDataJSON.Code = -1
		accountDataJSON.Msg = msg
	}
	return accountDataJSON
}

func (c *UpdateService) UpdateAgentStatus(agentUid string) *datas.KeyDataJSON {

	keyDataJSON := new(datas.KeyDataJSON)
	agentInfo := agent.GetAgentInfoByAgentUid(agentUid)
	if agentInfo.AgentUid == "" {
		keyDataJSON.Code = -1
		return keyDataJSON
	}

	if agentInfo.Status == common.ACTIVE {
		agentInfo.Status = common.UNACTIVE
	} else {
		agentInfo.Status = "active"
	}
	agentInfo.UpdateTime = utils.GetBasicDateTime()
	if agent.UpdateAgentInfo(agentInfo) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
	}
	return keyDataJSON
}

func (c *UpdateService) ResetAgentPassword(agentUid, newPassword, newVertifyPassword string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
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
		return keyDataJSON
	}

	agentInfo := agent.GetAgentInfoByAgentUid(agentUid)
	agentInfo.UpdateTime = utils.GetBasicDateTime()
	agentInfo.AgentPassword = utils.GetMD5Upper(newPassword)
	if !agent.UpdateAgentInfo(agentInfo) {
		keyDataJSON.Code = -1
	}
	return keyDataJSON
}

func (c *UpdateService) ChoosePayForRoad(confirmType, roadName, bankOrderId, remark string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = 200

	if confirmType == common.PAYFOR_ROAD && roadName == "" {
		keyDataJSON.Msg = "打款通道不能为空"
		keyDataJSON.Code = -1
		return keyDataJSON
	}

	payForInfo := payfor.GetPayForByBankOrderId(bankOrderId)
	roadInfo := road.GetRoadInfoByName(roadName)

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

		if !payfor.ForUpdatePayFor(payForInfo) {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "更新代付记录失败"
		}
	}
	return keyDataJSON
}

func (c *UpdateService) UpdateOrderStatus(bankOrderId, solveType string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	updateOrderUrl, _ := web.AppConfig.String("gateway::host")
	res, err := httplib.Get(updateOrderUrl + "gateway/update/order" + "?bankOrderId=" + bankOrderId + "&solveType=" + solveType).String()
	if err != nil {
		logs.Error("update order status err = ", err)
		keyDataJSON.Code = -1
	} else {
		keyDataJSON.Code = 200
		keyDataJSON.Msg = res
	}

	return keyDataJSON
}

func (c *UpdateService) ResultPayFor(resultType, bankOrderId string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = 200

	if resultType == "" || bankOrderId == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "提交的数据有误"
		return keyDataJSON
	}

	payFor := payfor.GetPayForByBankOrderId(bankOrderId)

	if payFor.Type == common.SELF_HELP {
		//如果是管理员在后台提现，不用做任何的商户减款,只需要更新代付订单状态
		payFor.UpdateTime = utils.GetBasicDateTime()
		payFor.Status = resultType

		if !payfor.ForUpdatePayFor(payFor) {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "系统处理失败"
		}
		return keyDataJSON
	}

	if payFor.Status == common.PAYFOR_FAIL || payFor.Status == common.PAYFOR_SUCCESS {
		logs.Error(fmt.Sprintf("该代付订单=%s，状态有误....", bankOrderId))
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "订单状态有误，请刷新重新判断"
		return keyDataJSON
	}

	u, _ := web.AppConfig.String("gateway::host")
	u = u + "/solve/payfor/result?" + "resultType=" + resultType + "&bankOrderId=" + bankOrderId
	s, err := httplib.Get(u).String()
	if err != nil || s == common.FAIL {
		logs.Error("手动处理代付结果请求gateway系统失败：", err)
		keyDataJSON.Msg = "处理失败"
		keyDataJSON.Code = -1
	} else {
		keyDataJSON.Code = 200
		keyDataJSON.Msg = "处理成功"
	}

	return keyDataJSON
}
