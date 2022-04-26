/***************************************************
 ** @Desc : c file for ...
 ** @Time : 2019/8/19 18:13
 ** @Author : yuebin
 ** @File : add
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/19 18:13
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"boss/common"
	"boss/datas"
	"boss/models/merchant"
	"boss/models/payfor"
	"boss/models/road"
	"boss/models/system"
	"boss/service"
	"boss/utils"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/rs/xid"
)

type AddController struct {
	BaseController
}

/*
* 添加一级菜单
 */
func (c *AddController) AddMenu() {
	oneMenu := c.GetString("oneMenu")
	userID := c.GetSession("userID").(string)
	se := new(service.AddService)
	dataJSON := se.AddMenu(oneMenu, userID)
	c.GenerateJSON(dataJSON)
}

/*
* 添加二级菜单
 */
func (c *AddController) AddSecondMenu() {
	firstMenuUid := c.GetString("preMenuUid")
	secondMenu := c.GetString("secondMenu")
	secondRouter := c.GetString("secondRouter")
	userID := c.GetSession("userID").(string)

	se := new(service.AddService)
	dataJSON := se.AddSecondMenu(firstMenuUid, secondRouter, secondMenu, userID)

	c.GenerateJSON(dataJSON)
}

/*
* 添加权限项的处理函数
 */
func (c *AddController) AddPower() {
	firstMenuUid := strings.TrimSpace(c.GetString("firstMenuUid"))
	secondMenuUid := strings.TrimSpace(c.GetString("secondMenuUid"))
	powerItem := strings.TrimSpace(c.GetString("powerItem"))
	powerID := strings.TrimSpace(c.GetString("powerID"))
	userID := c.GetSession("userID").(string)

	se := new(service.AddService)
	keyDataJSON := se.AddPower(powerItem, powerID, firstMenuUid, secondMenuUid, userID)

	c.GenerateJSON(keyDataJSON)
}

/*
* 添加权限角色
 */
func (c *AddController) AddRole() {
	roleName := strings.TrimSpace(c.GetString("roleNameAdd"))
	roleRemark := strings.TrimSpace(c.GetString("roleRemark"))
	userID := c.GetSession("userID").(string)
	se := new(service.AddService)
	keyDataJSON := se.AddRole(roleName, roleRemark, userID)
	c.GenerateJSON(keyDataJSON)
}

func (c *AddController) SavePower() {
	firstMenuUids := c.GetStrings("firstMenuUid[]")
	secondMenuUids := c.GetStrings("secondMenuUid[]")
	powerIds := c.GetStrings("powerId[]")
	roleUid := strings.TrimSpace(c.GetString("roleUid"))

	se := new(service.AddService)
	dataJSON := se.SavePower(roleUid, firstMenuUids, secondMenuUids, powerIds)
	c.GenerateJSON(dataJSON)
}

/*
* 添加操作员
 */
func (c *AddController) AddOperator() {
	loginAccount := strings.TrimSpace(c.GetString("operatorAccount"))
	loginPassword := strings.TrimSpace(c.GetString("operatorPassword"))
	role := strings.TrimSpace(c.GetString("operatorRole"))
	status := strings.TrimSpace(c.GetString("status"))
	remark := strings.TrimSpace(c.GetString("remark"))

	se := new(service.AddService)
	keyDataJSON := se.AddOperator(loginAccount, loginPassword, role, status, remark)

	c.GenerateJSON(keyDataJSON)
}

/*
* 添加银行卡
 */
func (c *AddController) AddBankCard() {
	userName := strings.TrimSpace(c.GetString("userName"))
	bankCode := strings.TrimSpace(c.GetString("bankCode"))
	accountName := strings.TrimSpace(c.GetString("accountName"))
	certificateType := strings.TrimSpace(c.GetString("certificateType"))
	phoneNo := strings.TrimSpace(c.GetString("phoneNo"))
	bankName := strings.TrimSpace(c.GetString("bankName"))
	bankAccountType := strings.TrimSpace(c.GetString("bankAccountType"))
	bankNo := strings.TrimSpace(c.GetString("bankNo"))
	identifyCard := strings.TrimSpace(c.GetString("certificateType"))
	certificateNo := strings.TrimSpace(c.GetString("certificateNo"))
	bankAddress := strings.TrimSpace(c.GetString("bankAddress"))
	uid := strings.TrimSpace(c.GetString("uid"))

	se := new(service.AddService)
	dataJSON := se.AddBankCard(userName, bankCode, accountName, certificateType, phoneNo,
		bankName, bankAccountType, bankNo, certificateNo, bankAddress, uid, identifyCard)

	c.GenerateJSON(dataJSON)
}

/*
* 添加通道
 */
func (c *AddController) AddRoad() {
	roadUid := strings.TrimSpace(c.GetString("roadUid"))
	roadName := strings.TrimSpace(c.GetString("roadName"))
	roadRemark := strings.TrimSpace(c.GetString("roadRemark"))
	productUid := strings.TrimSpace(c.GetString("productName"))
	payType := strings.TrimSpace(c.GetString("payType"))
	basicRate := strings.TrimSpace(c.GetString("basicRate"))
	settleFee := strings.TrimSpace(c.GetString("settleFee"))
	roadTotalLimit := strings.TrimSpace(c.GetString("roadTotalLimit"))
	roadEverydayLimit := strings.TrimSpace(c.GetString("roadEverydayLimit"))
	singleMinLimit := strings.TrimSpace(c.GetString("singleMinLimit"))
	singleMaxLimit := strings.TrimSpace(c.GetString("singleMaxLimit"))
	startHour := strings.TrimSpace(c.GetString("startHour"))
	endHour := strings.TrimSpace(c.GetString("endHour"))
	params := strings.TrimSpace(c.GetString("params"))

	se := new(service.AddService)
	dataJSON := se.AddRoad(startHour, endHour, roadName, productUid, payType, basicRate, settleFee,
		roadTotalLimit, roadEverydayLimit, singleMinLimit, singleMaxLimit, roadUid, roadRemark, params)

	c.GenerateJSON(dataJSON)
}

func (c *AddController) AddRoadPool() {
	roadPoolName := strings.TrimSpace(c.GetString("roadPoolName"))
	roadPoolCode := strings.TrimSpace(c.GetString("roadPoolCode"))

	se := new(service.AddService)
	keyDataJSON := se.AddRoadPool(roadPoolName, roadPoolCode)

	c.GenerateJSON(keyDataJSON)
}

/*
* 添加或者更新通道池中的通道
 */
func (c *AddController) SaveRoadUid() {
	roadUids := c.GetStrings("roadUid[]")
	roadPoolCode := strings.TrimSpace(c.GetString("roadPoolCode"))

	se := new(service.AddService)
	dataJSON := se.SaveRoadUid(roadPoolCode, roadUids)

	c.GenerateJSON(dataJSON)
}

/*
* 添加代理信息
 */
func (c *AddController) AddAgent() {
	agentName := strings.TrimSpace(c.GetString("agentName"))
	agentPhone := strings.TrimSpace(c.GetString("agentPhone"))
	agentLoginPassword := strings.TrimSpace(c.GetString("agentLoginPassword"))
	agentVertifyPassword := strings.TrimSpace(c.GetString("agentVertifyPassword"))
	agentRemark := strings.TrimSpace(c.GetString("agentRemark"))
	status := strings.TrimSpace(c.GetString("status"))
	agentUid := strings.TrimSpace(c.GetString("agentUid"))

	se := new(service.AddService)
	keyDataJSON := se.AddAgent(agentName, agentPhone, agentLoginPassword, agentVertifyPassword, status, agentUid, agentRemark)

	c.GenerateJSON(keyDataJSON)
}

func (c *AddController) AddMerchant() {
	merchantName := strings.TrimSpace(c.GetString("merchantName"))
	phone := strings.TrimSpace(c.GetString("phone"))
	loginPassword := strings.TrimSpace(c.GetString("loginPassword"))
	verifyPassword := strings.TrimSpace(c.GetString("verifyPassword"))
	merchantStatus := strings.TrimSpace(c.GetString("merchantStatus"))
	remark := strings.TrimSpace(c.GetString("remark"))

	se := new(service.AddService)
	keyDataJSON := se.AddMerchant(merchantName, phone, loginPassword, verifyPassword, merchantStatus, remark)
	c.GenerateJSON(keyDataJSON)
}

/*
* 添加商戶支付配置參數
 */
func (c *AddController) AddMerchantDeploy() {
	merchantUid := strings.TrimSpace(c.GetString("merchantNo"))
	isAutoSettle := strings.TrimSpace(c.GetString("isAutoSettle"))
	isAutoPayfor := strings.TrimSpace(c.GetString("isAutoPayfor"))
	ipWhite := strings.TrimSpace(c.GetString("ipWhite"))
	payforRoadChoose := strings.TrimSpace(c.GetString("payforRoadChoose"))
	rollPayforRoadChoose := strings.TrimSpace(c.GetString("rollPayforRoadChoose"))
	payforFee := strings.TrimSpace(c.GetString("payforFee"))
	belongAgentName := strings.TrimSpace(c.GetString("belongAgentName"))
	belongAgentUid := strings.TrimSpace(c.GetString("belongAgentUid"))

	se := new(service.AddService)
	keyDataJSON := se.AddMerchantDeploy(merchantUid, isAutoSettle, isAutoPayfor, ipWhite, belongAgentName,
		belongAgentUid, payforRoadChoose, rollPayforRoadChoose, payforFee)

	c.GenerateJSON(keyDataJSON)
}

func (c *AddController) AddMerchantPayType() {
	merchantNo := strings.TrimSpace(c.GetString("merchantNo"))
	payType := strings.TrimSpace(c.GetString("payType"))
	singleRoad := strings.TrimSpace(c.GetString("singleRoad"))
	singleRoadPlatformFee := strings.TrimSpace(c.GetString("singleRoadPlatformFee"))
	singleRoadAgentFee := strings.TrimSpace(c.GetString("singleRoadAgentFee"))
	rollPoolRoad := strings.TrimSpace(c.GetString("rollPoolRoad"))
	rollRoadPlatformFee := strings.TrimSpace(c.GetString("rollRoadPlatformFee"))
	rollRoadAgentFee := strings.TrimSpace(c.GetString("rollRoadAgentFee"))
	isLoan := strings.TrimSpace(c.GetString("isLoan"))
	loanRate := strings.TrimSpace(c.GetString("loanRate"))
	loanDays := strings.TrimSpace(c.GetString("loanDays"))
	unfreezeTimeHour := strings.TrimSpace(c.GetString("unfreezeTimeHour"))

	keyDataJSON := new(datas.KeyDataJSON)
	if payType == "" || payType == "none" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "操作失败，请选择支付类型"
		c.GenerateJSON(keyDataJSON)
		return
	}

	if singleRoad == "" && rollPoolRoad == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "单通道、轮询通道至少要有一个不为空！"
	}

	if singleRoad != "" && singleRoadPlatformFee == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "单通道平台利润率不能为0"
	}

	if rollPoolRoad != "" && rollRoadPlatformFee == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "轮询通道平台利润率不能为0"
	}

	if keyDataJSON.Code == -1 {
		c.GenerateJSON(keyDataJSON)
		return
	}

	//将字符串转变为float64或者int类型
	a, err := strconv.ParseFloat(singleRoadPlatformFee, 64)
	if err != nil {
		a = 0.0
	}
	b, err := strconv.ParseFloat(singleRoadAgentFee, 64)
	if err != nil {
		b = 0.0
	}
	cs, err := strconv.ParseFloat(rollRoadPlatformFee, 64)
	if err != nil {
		cs = 0.0
	}
	d, err := strconv.ParseFloat(rollRoadAgentFee, 64)
	if err != nil {
		d = 0.0
	}
	e, err := strconv.ParseFloat(loanRate, 64)
	if err != nil {
		e = 0.0
	}
	i, err := strconv.Atoi(loanDays)
	if err != nil {
		i = 0
	}
	j, err := strconv.Atoi(unfreezeTimeHour)
	if err != nil {
		j = 0
	}

	var merchantDeployInfo merchant.MerchantDeployInfo
	merchantDeployInfo.MerchantUid = merchantNo
	merchantDeployInfo.PayType = payType
	merchantDeployInfo.SingleRoadName = singleRoad
	merchantDeployInfo.SingleRoadPlatformRate = a
	merchantDeployInfo.SingleRoadAgentRate = b
	merchantDeployInfo.RollRoadPlatformRate = cs
	merchantDeployInfo.RollRoadAgentRate = d
	merchantDeployInfo.IsLoan = isLoan
	merchantDeployInfo.LoanRate = e
	merchantDeployInfo.LoanDays = i
	merchantDeployInfo.UnfreezeHour = j
	merchantDeployInfo.RollRoadName = rollPoolRoad
	roadInfo := road.GetRoadInfoByName(singleRoad)
	rollPoolInfo := road.GetRoadPoolByName(rollPoolRoad)
	merchantDeployInfo.SingleRoadUid = roadInfo.RoadUid
	merchantDeployInfo.RollRoadCode = rollPoolInfo.RoadPoolCode

	//如果该用户的改支付类型已经存在,那么进行更新，否则进行添加
	if merchant.IsExistByUidAndPayType(merchantNo, payType) {
		if singleRoad == "" && rollPoolRoad == "" {
			//表示需要删除该支付类型的通道
			if merchant.DeleteMerchantDeployByUidAndPayType(merchantNo, payType) {
				keyDataJSON.Code = 200
				keyDataJSON.Msg = "删除该支付类型通道成功"
			} else {
				keyDataJSON.Code = -1
				keyDataJSON.Msg = "删除该支付类型通道失败"
			}
		} else {
			tmpInfo := merchant.GetMerchantDeployByUidAndPayType(merchantNo, payType)
			merchantDeployInfo.Id = tmpInfo.Id
			merchantDeployInfo.Status = tmpInfo.Status
			merchantDeployInfo.UpdateTime = utils.GetBasicDateTime()
			merchantDeployInfo.CreateTime = tmpInfo.CreateTime
			if merchant.UpdateMerchantDeploy(merchantDeployInfo) {
				keyDataJSON.Code = 200
				keyDataJSON.Msg = "更新成功"
			} else {
				keyDataJSON.Code = -1
				keyDataJSON.Msg = "更新失败"
			}
		}
	} else {
		if singleRoad == "" && rollPoolRoad == "" {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "参数不能为空"
		} else {
			merchantDeployInfo.CreateTime = utils.GetBasicDateTime()
			merchantDeployInfo.UpdateTime = utils.GetBasicDateTime()
			merchantDeployInfo.Status = common.ACTIVE
			if merchant.InsertMerchantDeployInfo(merchantDeployInfo) {
				keyDataJSON.Code = 200
				keyDataJSON.Msg = "添加支付类型成功"
			} else {
				keyDataJSON.Code = -1
				keyDataJSON.Msg = "添加支付类型失败"
			}
		}
	}
	c.GenerateJSON(keyDataJSON)
}

/*
*后台提交的下发记录
 */
func (c *AddController) AddPayFor() {
	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))
	merchantName := strings.TrimSpace(c.GetString("merchantName"))
	bankName := strings.TrimSpace(c.GetString("bankName"))
	accountName := strings.TrimSpace(c.GetString("accountName"))
	bankUid := strings.TrimSpace(c.GetString("bankUid"))
	bankNo := strings.TrimSpace(c.GetString("bankNo"))
	bankAddress := strings.TrimSpace(c.GetString("bankAddress"))
	phone := strings.TrimSpace(c.GetString("phone"))
	payForAmount := strings.TrimSpace(c.GetString("payForAmount"))

	se := new(service.AddService)
	keyDataJSON := se.AddPayFor(merchantUid, bankUid, payForAmount, bankNo,
		accountName, phone, merchantName, bankName, bankAddress)

	c.GenerateJSON(keyDataJSON)
}

func (c *AddController) AddSelfPayFor() {
	bankUid := strings.TrimSpace(c.GetString("bankUid"))
	bankName := strings.TrimSpace(c.GetString("bankName"))
	accountName := strings.TrimSpace(c.GetString("accountName"))
	bankNo := strings.TrimSpace(c.GetString("bankNo"))
	bankAddress := strings.TrimSpace(c.GetString("bankAddress"))
	phone := strings.TrimSpace(c.GetString("phone"))
	payForAmount := strings.TrimSpace(c.GetString("payForAmount"))

	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = -1

	if bankUid == "" {
		keyDataJSON.Msg = "银行卡uid不能为空，请联系技术人员"
		c.GenerateJSON(keyDataJSON)
		return
	}
	money, err := strconv.ParseFloat(payForAmount, 64)
	if err != nil {
		logs.Error("self payfor money fail: ", err)
		keyDataJSON.Msg = "输入金额有误，请仔细检查"
		c.GenerateJSON(keyDataJSON)
		return
	}

	bankInfo := system.GetBankCardByUid(bankUid)

	//需要对前端传入的数据做校验，不能完全相信前端的数据
	if bankInfo.AccountName != accountName || bankInfo.BankNo != bankNo || bankInfo.PhoneNo != phone {
		keyDataJSON.Msg = "前端页面数据有篡改，请注意资金安全"
		c.GenerateJSON(keyDataJSON)
		return
	}

	selfPayFor := payfor.PayforInfo{PayforUid: "pppp" + xid.New().String(), BankOrderId: "4444" + xid.New().String(), PayforFee: common.ZERO, Type: common.SELF_HELP,
		PayforAmount: money, PayforTotalAmount: money + common.ZERO, BankCode: bankInfo.BankCode, BankName: bankName, IsSend: common.NO,
		BankAccountName: bankInfo.AccountName, BankAccountNo: bankInfo.BankNo, BankAccountType: bankInfo.BankAccountType, BankAccountAddress: bankAddress,
		Status: common.PAYFOR_COMFRIM, CreateTime: utils.GetBasicDateTime(), UpdateTime: utils.GetBasicDateTime()}

	if payfor.InsertPayfor(selfPayFor) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Msg = "数据处理失败，请重新提交"
	}

	c.GenerateJSON(keyDataJSON)
}
