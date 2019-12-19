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
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/rs/xid"
	"dongfeng-pay/service/common"
	"dongfeng-pay/service/models"
	"dongfeng-pay/service/utils"
	"strconv"
	"strings"
)

type AddController struct {
	BaseController
}

/*
* 添加一级菜单
 */
func (c *AddController) AddMenu() {
	oneMenu := c.GetString("oneMenu")

	dataJSON := new(BaseDataJSON)
	menuInfo := models.MenuInfo{MenuUid: xid.New().String(), FirstMenu: oneMenu, Status: "active",
		Creater: c.GetSession("userID").(string), CreateTime: utils.GetBasicDateTime()}

	exist := models.FirstMenuIsExists(oneMenu)
	if !exist {
		menuInfo.MenuOrder = models.GetMenuLen() + 1
		flag := models.InsertMenu(menuInfo)
		if !flag {
			dataJSON.Code = -1
			dataJSON.Msg = "添加菜单失败"
		} else {
			dataJSON.Code = 200
		}
	} else {
		dataJSON.Code = -1
		dataJSON.Msg = "一级菜单名已经存在"
	}
	c.GenerateJSON(dataJSON)
}

/*
* 添加二级菜单
 */
func (c *AddController) AddSecondMenu() {
	firstMenuUid := c.GetString("preMenuUid")
	secondMenu := c.GetString("secondMenu")
	secondRouter := c.GetString("secondRouter")

	dataJSON := new(KeyDataJSON)

	firstMenuInfo := models.GetMenuInfoByMenuUid(firstMenuUid)
	routerExists := models.SecondRouterExists(secondRouter)
	secondMenuExists := models.SecondMenuIsExists(secondMenu)

	if firstMenuInfo.MenuUid == "" {
		dataJSON.Code = -1
		dataJSON.Key = "pre-menu-error"
		dataJSON.Msg = "*一级菜单不存在"
	} else if routerExists {
		dataJSON.Code = -1
		dataJSON.Msg = "*该路由已存在"
		dataJSON.Key = "second-router-error"
	} else if secondMenuExists {
		dataJSON.Code = -1
		dataJSON.Key = "second-menu-error"
		dataJSON.Msg = "*该菜单名已经存在"
	} else {
		sl := models.GetSecondMenuLenByFirstMenuUid(firstMenuUid)
		secondMenuInfo := models.SecondMenuInfo{MenuOrder: sl + 1, FirstMenuUid: firstMenuInfo.MenuUid,
			FirstMenu: firstMenuInfo.FirstMenu, SecondMenuUid: xid.New().String(), Status: "active",
			SecondMenu: secondMenu, SecondRouter: secondRouter, Creater: c.GetSession("userID").(string),
			CreateTime: utils.GetBasicDateTime(), UpdateTime: utils.GetBasicDateTime(), FirstMenuOrder: firstMenuInfo.MenuOrder}
		if !models.InsertSecondMenu(secondMenuInfo) {
			dataJSON.Code = -1
			dataJSON.Msg = "添加二级菜单失败"
		} else {
			dataJSON.Code = 200
			dataJSON.Msg = "添加二级菜单成功"
		}
	}
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

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = -1
	if powerItem == "" || len(powerItem) == 0 {
		keyDataJSON.Key = ".power-name-error"
		keyDataJSON.Msg = "*权限项名称不能为空"
		c.GenerateJSON(keyDataJSON)
		return
	}
	if powerID == "" || len(powerID) == 0 {
		keyDataJSON.Key = ".power-id-error"
		keyDataJSON.Msg = "*权限项ID不能为空"
		c.GenerateJSON(keyDataJSON)
		return
	}
	if models.PowerUidExists(powerID) {
		keyDataJSON.Key = ".power-id-error"
		keyDataJSON.Msg = "*权限项ID已经存在"
		c.GenerateJSON(keyDataJSON)
		return
	}

	fmt.Println(powerID)

	secondMenuInfo := models.GetSecondMenuInfoBySecondMenuUid(secondMenuUid)

	powerInfo := models.PowerInfo{SecondMenuUid: secondMenuUid, SecondMenu: secondMenuInfo.SecondMenu,
		PowerId: powerID, PowerItem: powerItem, Creater: c.GetSession("userID").(string),
		Status: "active", CreateTime: utils.GetBasicDateTime(), UpdateTime: utils.GetBasicDateTime(),
		FirstMenuUid: firstMenuUid}

	keyDataJSON.Code = 200
	if !models.InsertPowerInfo(powerInfo) {
		keyDataJSON.Key = ".power-save-success"
		keyDataJSON.Msg = "添加权限项失败"
	} else {
		keyDataJSON.Key = ".power-save-success"
		keyDataJSON.Msg = "添加权限项成功"
	}
	c.GenerateJSON(keyDataJSON)
}

/*
* 添加权限角色
 */
func (this *AddController) AddRole() {
	roleName := strings.TrimSpace(this.GetString("roleNameAdd"))
	roleRemark := strings.TrimSpace(this.GetString("roleRemark"))

	keyDataJSON := new(KeyDataJSON)
	if len(roleName) == 0 {
		keyDataJSON.Code = -1
		keyDataJSON.Key = ".role-name-error"
		keyDataJSON.Msg = "*角色名称不能为空"
		this.GenerateJSON(keyDataJSON)
		return
	}

	if models.RoleNameExists(roleName) {
		keyDataJSON.Code = -1
		keyDataJSON.Key = ".role-name-error"
		keyDataJSON.Msg = "*角色名称已经存在"
		this.GenerateJSON(keyDataJSON)
		return
	}

	roleInfo := models.RoleInfo{RoleName: roleName, RoleUid: xid.New().String(),
		Creater: this.GetSession("userID").(string), Status: "active", Remark: roleRemark,
		CreateTime: utils.GetBasicDateTime(), UpdateTime: utils.GetBasicDateTime()}

	if !models.InsertRole(roleInfo) {
		keyDataJSON.Code = -1
		keyDataJSON.Key = ".role-save-success"
		keyDataJSON.Msg = "添加角色失败"
		this.GenerateJSON(keyDataJSON)
		return
	}

	keyDataJSON.Code = 200
	this.GenerateJSON(keyDataJSON)
}

func (this *AddController) SavePower() {
	firstMenuUids := this.GetStrings("firstMenuUid[]")
	secondMenuUids := this.GetStrings("secondMenuUid[]")
	powerIds := this.GetStrings("powerId[]")
	roleUid := strings.TrimSpace(this.GetString("roleUid"))

	dataJSON := new(BaseDataJSON)
	roleInfo := models.GetRoleByRoleUid(roleUid)
	if len(roleUid) == 0 || len(roleInfo.RoleUid) == 0 {
		dataJSON.Code = -1
		this.GenerateJSON(dataJSON)
	}

	roleInfo.UpdateTime = utils.GetBasicDateTime()
	roleInfo.ShowFirstUid = strings.Join(firstMenuUids, "||")
	roleInfo.ShowSecondUid = strings.Join(secondMenuUids, "||")
	roleInfo.ShowPowerUid = strings.Join(powerIds, "||")

	menuInfoList := models.GetMenuInfosByMenuUids(firstMenuUids)
	showFirstMenu := make([]string, 0)
	for _, m := range menuInfoList {
		showFirstMenu = append(showFirstMenu, m.FirstMenu)
	}
	roleInfo.ShowFirstMenu = strings.Join(showFirstMenu, "||")

	secondMenuInfoList := models.GetSecondMenuInfoBySecondMenuUids(secondMenuUids)
	showSecondMenu := make([]string, 0)
	for _, m := range secondMenuInfoList {
		showSecondMenu = append(showSecondMenu, m.SecondMenu)
	}
	roleInfo.ShowSecondMenu = strings.Join(showSecondMenu, "||")

	powerList := models.GetPowerByIds(powerIds)
	showPower := make([]string, 0)
	for _, p := range powerList {
		showPower = append(showPower, p.PowerItem)
	}
	roleInfo.ShowPower = strings.Join(showPower, "||")

	if !models.UpdateRoleInfo(roleInfo) {
		dataJSON.Code = -1
		dataJSON.Msg = "更新roleInfo失败"
	} else {
		dataJSON.Code = 200
		dataJSON.Msg = "更新roleInfo成功"
	}
	this.GenerateJSON(dataJSON)
}

/*
* 添加操作员
 */
func (this *AddController) AddOperator() {
	loginAccount := strings.TrimSpace(this.GetString("operatorAccount"))
	loginPassword := strings.TrimSpace(this.GetString("operatorPassword"))
	role := strings.TrimSpace(this.GetString("operatorRole"))
	status := strings.TrimSpace(this.GetString("status"))
	remark := strings.TrimSpace(this.GetString("remark"))

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = -1
	if len(loginAccount) == 0 {
		keyDataJSON.Key = ".operator-name-error"
		keyDataJSON.Msg = "*登录账号不能为空"
	} else if len(loginPassword) == 0 {
		keyDataJSON.Key = ".operator-password-error"
		keyDataJSON.Msg = "*初始密码不能为空"
	} else if len(role) == 0 || role == "none" {
		keyDataJSON.Key = ".operator-role-error"
		keyDataJSON.Msg = "请选择角色"
	} else if models.UserInfoExistByUserId(loginAccount) {
		keyDataJSON.Key = ".operator-name-error"
		keyDataJSON.Msg = "*账号已经存在"
	} else {
		if len(remark) == 0 {
			remark = loginAccount
		}
		roleInfo := models.GetRoleByRoleUid(role)
		userInfo := models.UserInfo{UserId: loginAccount, Passwd: utils.GetMD5Upper(loginPassword), Nick: "壮壮", Remark: remark,
			Status: status, Role: role, RoleName: roleInfo.RoleName, CreateTime: utils.GetBasicDateTime(), UpdateTime: utils.GetBasicDateTime()}
		if !models.InsertUser(userInfo) {
			keyDataJSON.Code = 200
			keyDataJSON.Msg = "添加操作员失败"
		} else {
			keyDataJSON.Code = 200
			keyDataJSON.Msg = "添加操作员成功"
		}
	}
	this.GenerateJSON(keyDataJSON)
}

/*
* 添加银行卡
 */
func (this *AddController) AddBankCard() {
	userName := strings.TrimSpace(this.GetString("userName"))
	bankCode := strings.TrimSpace(this.GetString("bankCode"))
	accountName := strings.TrimSpace(this.GetString("accountName"))
	certificateType := strings.TrimSpace(this.GetString("certificateType"))
	phoneNo := strings.TrimSpace(this.GetString("phoneNo"))
	bankName := strings.TrimSpace(this.GetString("bankName"))
	bankAccountType := strings.TrimSpace(this.GetString("bankAccountType"))
	bankNo := strings.TrimSpace(this.GetString("bankNo"))
	identifyCard := strings.TrimSpace(this.GetString("certificateType"))
	certificateNo := strings.TrimSpace(this.GetString("certificateNo"))
	bankAddress := strings.TrimSpace(this.GetString("bankAddress"))
	uid := strings.TrimSpace(this.GetString("uid"))

	dataJSON := new(BaseDataJSON)

	dataJSON.Code = -1
	if len(userName) == 0 {
		dataJSON.Msg = "用户名不能为空"
	} else if len(bankCode) == 0 {
		dataJSON.Msg = "银行编码不能为空"
	} else if len(accountName) == 0 {
		dataJSON.Msg = "银行开户名不能为空"
	} else if len(certificateType) == 0 {
		dataJSON.Msg = "证件种类不能为空"
	} else if len(phoneNo) == 0 {
		dataJSON.Msg = "手机号不能为空"
	} else if len(bankName) == 0 {
		dataJSON.Msg = "银行名称不能为空"
	} else if len(bankAccountType) == 0 {
		dataJSON.Msg = "银行账户类型不能为空"
	} else if len(bankNo) == 0 {
		dataJSON.Msg = "银行账号不能为空"
	} else if len(certificateNo) == 0 {
		dataJSON.Msg = "身份证号不能为空"
	} else if len(bankAddress) == 0 {
		dataJSON.Msg = "银行地址不能为空"
	} else {

	}
	if dataJSON.Msg != "" {
		logs.Error("添加银行卡校验失败")
	} else {
		if len(uid) > 0 {
			bankCardInfo := models.GetBankCardByUid(uid)
			bankCardInfo = models.BankCardInfo{
				Id: bankCardInfo.Id, UserName: userName, BankName: bankName,
				BankCode: bankCode, BankAccountType: bankAccountType,
				AccountName: accountName, BankNo: bankNo, IdentifyCard: identifyCard,
				CertificateNo: certificateNo, PhoneNo: phoneNo,
				BankAddress: bankAddress, UpdateTime: utils.GetBasicDateTime(),
				CreateTime: bankCardInfo.CreateTime, Uid: bankCardInfo.Uid,
			}
			if models.UpdateBankCard(bankCardInfo) {
				dataJSON.Code = 200
			}
		} else {
			bankCardInfo := models.BankCardInfo{Uid: "3333" + xid.New().String(), UserName: userName, BankName: bankName,
				BankCode: bankCode, BankAccountType: bankAccountType, AccountName: accountName, BankNo: bankNo,
				IdentifyCard: identifyCard, CertificateNo: certificateNo, PhoneNo: phoneNo, BankAddress: bankAddress,
				UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

			if models.InsertBankCardInfo(bankCardInfo) {
				dataJSON.Code = 200
			}
		}
	}

	this.GenerateJSON(dataJSON)
}

/*
* 添加通道
 */
func (this *AddController) AddRoad() {
	roadUid := strings.TrimSpace(this.GetString("roadUid"))
	roadName := strings.TrimSpace(this.GetString("roadName"))
	roadRemark := strings.TrimSpace(this.GetString("roadRemark"))
	productUid := strings.TrimSpace(this.GetString("productName"))
	payType := strings.TrimSpace(this.GetString("payType"))
	basicRate := strings.TrimSpace(this.GetString("basicRate"))
	settleFee := strings.TrimSpace(this.GetString("settleFee"))
	roadTotalLimit := strings.TrimSpace(this.GetString("roadTotalLimit"))
	roadEverydayLimit := strings.TrimSpace(this.GetString("roadEverydayLimit"))
	singleMinLimit := strings.TrimSpace(this.GetString("singleMinLimit"))
	singleMaxLimit := strings.TrimSpace(this.GetString("singleMaxLimit"))
	startHour := strings.TrimSpace(this.GetString("startHour"))
	endHour := strings.TrimSpace(this.GetString("endHour"))
	params := strings.TrimSpace(this.GetString("params"))

	dataJSON := new(BaseDataJSON)
	dataJSON.Code = -1

	startHourTmp, err1 := strconv.Atoi(startHour)
	endHourTmp, err2 := strconv.Atoi(endHour)

	if err1 != nil || err2 != nil {
		dataJSON.Msg = "开始时间或者结束时间设置有误"
		this.GenerateJSON(dataJSON)
		return
	}

	valid := validation.Validation{}
	if v := valid.Required(roadName, "roadName"); !v.Ok {
		dataJSON.Msg = "通道名称不能为空"
	} else if v := valid.Required(productUid, "productUid"); !v.Ok {
		dataJSON.Msg = "产品名称不能为空"
	} else if v := valid.Required(payType, "payType"); !v.Ok {
		dataJSON.Msg = "支付类型不能为空"
	} else if v := valid.Required(basicRate, ""); !v.Ok {
		dataJSON.Msg = "成本费率不能为空"
	} else if v := valid.Range(startHourTmp, 0, 23, ""); !v.Ok {
		dataJSON.Msg = "开始时间设置有误"
	} else if v := valid.Range(endHourTmp, 0, 23, ""); !v.Ok {
		dataJSON.Msg = "结束时间设置有误"
	} else {
		basicFee, err := strconv.ParseFloat(basicRate, 64)
		if err != nil {
			dataJSON.Msg = "成本汇率设置不符合规范"
		}
		settleFeeTmp, err := strconv.ParseFloat(settleFee, 64)
		if err != nil {
			dataJSON.Msg = "代付手续费设置不符合规范"
		}
		totalLimit, err := strconv.ParseFloat(roadTotalLimit, 64)
		if err != nil {
			dataJSON.Msg = "通道总额度设置不符合规范"
		}
		todayLimit, err := strconv.ParseFloat(roadEverydayLimit, 64)
		if err != nil {
			dataJSON.Msg = "每天额度设置不符合规范"
		}
		singleMinLimitTmp, err := strconv.ParseFloat(singleMinLimit, 64)
		if err != nil {
			dataJSON.Msg = "单笔最小金额设置不符合规范"
		}
		singleMaxLimitTmp, err := strconv.ParseFloat(singleMaxLimit, 64)
		if err != nil {
			dataJSON.Msg = "单笔最大金额设置不符合规范"
		}
		if len(dataJSON.Msg) > 0 {
			this.GenerateJSON(dataJSON)
			return
		}
		productName := ""
		supplierMap := common.GetSupplierMap()
		for k, v := range supplierMap {
			if k == productUid {
				productName = v
			}
		}

		if len(roadUid) > 0 {
			//更新通道
			roadInfo := models.GetRoadInfoByRoadUid(roadUid)
			roadInfo.RoadName = roadName
			roadInfo.Remark = roadRemark
			roadInfo.ProductUid = productUid
			roadInfo.ProductName = productName
			roadInfo.PayType = payType
			roadInfo.BasicFee = basicFee
			roadInfo.SettleFee = settleFeeTmp
			roadInfo.TotalLimit = totalLimit
			roadInfo.TodayLimit = todayLimit
			roadInfo.SingleMaxLimit = singleMaxLimitTmp
			roadInfo.SingleMinLimit = singleMinLimitTmp
			roadInfo.StarHour = startHourTmp
			roadInfo.EndHour = endHourTmp
			roadInfo.Params = params

			if models.UpdateRoadInfo(roadInfo) {
				dataJSON.Code = 200
			} else {
				dataJSON.Msg = "通道更新失败"
			}
		} else {
			//添加新的通道
			roadUid = "4444" + xid.New().String()
			roadInfo := models.RoadInfo{RoadName: roadName, RoadUid: roadUid, Remark: roadRemark,
				ProductUid: productUid, ProductName: productName, PayType: payType, BasicFee: basicFee, SettleFee: settleFeeTmp,
				TotalLimit: totalLimit, TodayLimit: todayLimit, SingleMinLimit: singleMinLimitTmp, Balance: common.ZERO,
				SingleMaxLimit: singleMaxLimitTmp, StarHour: startHourTmp, EndHour: endHourTmp, Status: "active",
				Params: params, UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime(),
			}

			if models.InsertRoadInfo(roadInfo) {
				dataJSON.Code = 200
			} else {
				dataJSON.Msg = "添加新通道失败"
			}
		}
	}

	this.GenerateJSON(dataJSON)
}

func (this *AddController) AddRoadPool() {
	roadPoolName := strings.TrimSpace(this.GetString("roadPoolName"))
	roadPoolCode := strings.TrimSpace(this.GetString("roadPoolCode"))

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = -1

	if len(roadPoolName) == 0 {
		keyDataJSON.Msg = "*通道池名称不能为空"
	} else if len(roadPoolCode) == 0 {
		keyDataJSON.Msg = "*通道池编号不能为空"
	}

	roadPoolInfo := models.RoadPoolInfo{Status: "active", RoadPoolName: roadPoolName, RoadPoolCode: roadPoolCode,
		UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

	if models.InsertRoadPool(roadPoolInfo) {
		keyDataJSON.Code = 200
		keyDataJSON.Msg = "添加通道池成功"
	} else {
		keyDataJSON.Msg = "添加通道池失败"
	}

	this.GenerateJSON(keyDataJSON)
}

/*
* 添加或者更新通道池中的通道
 */
func (this *AddController) SaveRoadUid() {
	roadUids := this.GetStrings("roadUid[]")
	roadPoolCode := strings.TrimSpace(this.GetString("roadPoolCode"))

	dataJSON := new(BaseDataJSON)
	dataJSON.Code = -1
	roadPoolInfo := models.GetRoadPoolByRoadPoolCode(roadPoolCode)
	if roadPoolInfo.RoadPoolCode == "" {
		this.GenerateJSON(dataJSON)
		return
	}
	var uids []string
	for _, uid := range roadUids {
		//去掉空格
		if len(uid) > 0 && models.RoadInfoExistByRoadUid(uid) {
			uids = append(uids, uid)
		}
	}
	if len(uids) > 0 {
		roadUid := strings.Join(uids, "||")
		roadPoolInfo.RoadUidPool = roadUid
	}
	roadPoolInfo.UpdateTime = utils.GetBasicDateTime()
	if models.UpdateRoadPool(roadPoolInfo) {
		dataJSON.Code = 200
	}
	this.GenerateJSON(dataJSON)
}

/*
* 添加代理信息
 */
func (this *AddController) AddAgent() {
	agentName := strings.TrimSpace(this.GetString("agentName"))
	agentPhone := strings.TrimSpace(this.GetString("agentPhone"))
	agentLoginPassword := strings.TrimSpace(this.GetString("agentLoginPassword"))
	agentVertifyPassword := strings.TrimSpace(this.GetString("agentVertifyPassword"))
	agentRemark := strings.TrimSpace(this.GetString("agentRemark"))
	status := strings.TrimSpace(this.GetString("status"))
	agentUid := strings.TrimSpace(this.GetString("agentUid"))

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = 200

	if agentName == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-name-error"
		keyDataJSON.Msg = "代理名不能为空"
	} else if models.IsEixstByAgentName(agentName) {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-name-error"
		keyDataJSON.Msg = "已存在该代理名称"
	} else if agentPhone == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-phone-error"
		keyDataJSON.Msg = "代理注册手机号不能为空"
	} else if models.IsEixstByAgentPhone(agentPhone) {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-phone-error"
		keyDataJSON.Msg = "代理商手机号已被注册"
	} else if agentLoginPassword == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-login-password-error"
		keyDataJSON.Msg = "密码不能为空"
	} else if agentLoginPassword != agentVertifyPassword {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-vertify-password-error"
		keyDataJSON.Msg = "二次密码输入不一致"
	}

	if keyDataJSON.Code == -1 {
		this.GenerateJSON(keyDataJSON)
		return
	}

	if status == "" {
		status = "active"
	}

	if agentUid == "" {

		agentUid = "9999" + xid.New().String()

		agentInfo := models.AgentInfo{Status: status, AgentName: agentName, AgentPhone: agentPhone,
			AgentPassword: utils.GetMD5Upper(agentLoginPassword), AgentUid: agentUid, UpdateTime: utils.GetBasicDateTime(),
			CreateTime: utils.GetBasicDateTime(), AgentRemark: agentRemark}

		if !models.InsertAgentInfo(agentInfo) {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "添加代理商失败"
		}
	}

	//创建新的账户
	account := models.GetAccountByUid(agentUid)
	if account.AccountUid == "" {
		account.Status = "active"
		account.AccountUid = agentUid
		account.AccountName = agentName
		account.Balance = 0.0
		account.LoanAmount = 0.0
		account.FreezeAmount = 0.0
		account.PayforAmount = 0.0
		account.SettleAmount = 0.0
		account.WaitAmount = 0.0
		account.UpdateTime = utils.GetBasicDateTime()
		account.CreateTime = utils.GetBasicDateTime()
		if models.InsetAcount(account) {
			keyDataJSON.Code = 200
			keyDataJSON.Msg = "插入成功"
		} else {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "掺入失败"
		}
	}

	this.GenerateJSON(keyDataJSON)
}

func (this *AddController) AddMerchant() {
	merchantName := strings.TrimSpace(this.GetString("merchantName"))
	phone := strings.TrimSpace(this.GetString("phone"))
	loginPassword := strings.TrimSpace(this.GetString("loginPassword"))
	verifyPassword := strings.TrimSpace(this.GetString("verifyPassword"))
	merchantStatus := strings.TrimSpace(this.GetString("merchantStatus"))
	remark := strings.TrimSpace(this.GetString("remark"))

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = 200
	if merchantName == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#merchant-name-error"
		keyDataJSON.Msg = "商户名称为空"
	} else if models.IsExistByMerchantName(merchantName) {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#merchant-name-error"
		keyDataJSON.Msg = "商户名已经存在"
	} else if phone == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#merchant-phone-error"
		keyDataJSON.Msg = "手机号为空"
	} else if models.IsExistByMerchantPhone(phone) {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#merchant-phone-error"
		keyDataJSON.Msg = "该手机号已经注册"
	} else if loginPassword == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#merchant-login-password-error"
		keyDataJSON.Msg = "登录密码为空"
	} else if verifyPassword == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#merchant-verify-password-error"
		keyDataJSON.Msg = "密码确认为空"
	} else if loginPassword != verifyPassword {
		keyDataJSON.Key = "#merchant-verify-password-error"
		keyDataJSON.Msg = "两次密码输入不正确"
	} else if merchantStatus == "" {
		merchantStatus = "active"
	}
	if keyDataJSON.Code == -1 {
		this.GenerateJSON(keyDataJSON)
		return
	}
	merchantUid := "8888" + xid.New().String()
	merchantKey := "kkkk" + xid.New().String()    //商户key
	merchantSecret := "ssss" + xid.New().String() //商户密钥
	merchantInfo := models.MerchantInfo{MerchantName: merchantName, MerchantUid: merchantUid,
		LoginAccount: phone, MerchantKey: merchantKey, MerchantSecret: merchantSecret,
		LoginPassword: utils.GetMD5Upper(loginPassword), Status: merchantStatus, Remark: remark,
		UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

	if models.InsertMerchantInfo(merchantInfo) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "插入失败"
	}
	//创建新的账户
	account := models.GetAccountByUid(merchantUid)
	if account.AccountUid == "" {
		account.Status = "active"
		account.AccountUid = merchantUid
		account.AccountName = merchantName
		account.Balance = 0.0
		account.LoanAmount = 0.0
		account.FreezeAmount = 0.0
		account.PayforAmount = 0.0
		account.SettleAmount = 0.0
		account.WaitAmount = 0.0
		account.UpdateTime = utils.GetBasicDateTime()
		account.CreateTime = utils.GetBasicDateTime()
		if models.InsetAcount(account) {
			keyDataJSON.Code = 200
			keyDataJSON.Msg = "插入成功"
		} else {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "掺入失败"
		}
	}
	this.GenerateJSON(keyDataJSON)
}

/*
* 添加商戶支付配置參數
 */
func (this *AddController) AddMerchantDeploy() {
	//merchantName := strings.TrimSpace(this.GetString("merchantName"))
	merchantUid := strings.TrimSpace(this.GetString("merchantNo"))
	isAutoSettle := strings.TrimSpace(this.GetString("isAutoSettle"))
	isAutoPayfor := strings.TrimSpace(this.GetString("isAutoPayfor"))
	ipWhite := strings.TrimSpace(this.GetString("ipWhite"))
	payforRoadChoose := strings.TrimSpace(this.GetString("payforRoadChoose"))
	rollPayforRoadChoose := strings.TrimSpace(this.GetString("rollPayforRoadChoose"))
	payforFee := strings.TrimSpace(this.GetString("payforFee"))
	belongAgentName := strings.TrimSpace(this.GetString("belongAgentName"))
	belongAgentUid := strings.TrimSpace(this.GetString("belongAgentUid"))

	keyDataJSON := new(KeyDataJSON)
	merchantInfo := models.GetMerchantByUid(merchantUid)
	merchantInfo.AutoSettle = isAutoSettle
	merchantInfo.AutoPayFor = isAutoPayfor
	merchantInfo.WhiteIps = ipWhite
	merchantInfo.BelongAgentName = belongAgentName
	merchantInfo.BelongAgentUid = belongAgentUid

	if payforRoadChoose != "" {
		roadInfo := models.GetRoadInfoByName(payforRoadChoose)
		merchantInfo.SinglePayForRoadName = payforRoadChoose
		merchantInfo.SinglePayForRoadUid = roadInfo.RoadUid
	}
	if rollPayforRoadChoose != "" {
		rollPoolInfo := models.GetRoadPoolByName(rollPayforRoadChoose)
		merchantInfo.RollPayForRoadName = rollPayforRoadChoose
		merchantInfo.RollPayForRoadCode = rollPoolInfo.RoadPoolCode
	}
	tmp, err := strconv.ParseFloat(payforFee, 64)
	if err != nil {
		logs.Error("手续费由字符串转为float64失败")
		tmp = common.PAYFOR_FEE
	}
	merchantInfo.PayforFee = tmp
	if models.UpdateMerchant(merchantInfo) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
	}
	this.GenerateJSON(keyDataJSON)
}

func (this *AddController) AddMerchantPayType() {
	merchantNo := strings.TrimSpace(this.GetString("merchantNo"))
	payType := strings.TrimSpace(this.GetString("payType"))
	singleRoad := strings.TrimSpace(this.GetString("singleRoad"))
	singleRoadPlatformFee := strings.TrimSpace(this.GetString("singleRoadPlatformFee"))
	singleRoadAgentFee := strings.TrimSpace(this.GetString("singleRoadAgentFee"))
	rollPoolRoad := strings.TrimSpace(this.GetString("rollPoolRoad"))
	rollRoadPlatformFee := strings.TrimSpace(this.GetString("rollRoadPlatformFee"))
	rollRoadAgentFee := strings.TrimSpace(this.GetString("rollRoadAgentFee"))
	isLoan := strings.TrimSpace(this.GetString("isLoan"))
	loanRate := strings.TrimSpace(this.GetString("loanRate"))
	loanDays := strings.TrimSpace(this.GetString("loanDays"))
	unfreezeTimeHour := strings.TrimSpace(this.GetString("unfreezeTimeHour"))

	keyDataJSON := new(KeyDataJSON)
	if payType == "" || payType == "none" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "操作失败，请选择支付类型"
		this.GenerateJSON(keyDataJSON)
		return
	}
	if singleRoad == "" && (singleRoadPlatformFee != "" || singleRoadAgentFee != "") {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "单通道选项不能为空"
	} else if rollPoolRoad == "" && (rollRoadPlatformFee != "" || rollRoadAgentFee != "") {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "轮询通道选项不能为空"
	}

	if keyDataJSON.Code == -1 {
		this.GenerateJSON(keyDataJSON)
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
	c, err := strconv.ParseFloat(rollRoadPlatformFee, 64)
	if err != nil {
		c = 0.0
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

	var merchantDeployInfo models.MerchantDeployInfo
	merchantDeployInfo.MerchantUid = merchantNo
	merchantDeployInfo.PayType = payType
	merchantDeployInfo.SingleRoadName = singleRoad
	merchantDeployInfo.SingleRoadPlatformRate = a
	merchantDeployInfo.SingleRoadAgentRate = b
	merchantDeployInfo.RollRoadPlatformRate = c
	merchantDeployInfo.RollRoadAgentRate = d
	merchantDeployInfo.IsLoan = isLoan
	merchantDeployInfo.LoanRate = e
	merchantDeployInfo.LoanDays = i
	merchantDeployInfo.UnfreezeHour = j
	merchantDeployInfo.RollRoadName = rollPoolRoad
	roadInfo := models.GetRoadInfoByName(singleRoad)
	rollPoolInfo := models.GetRoadPoolByName(rollPoolRoad)
	merchantDeployInfo.SingleRoadUid = roadInfo.RoadUid
	merchantDeployInfo.RollRoadCode = rollPoolInfo.RoadPoolCode

	//如果该用户的改支付类型已经存在,那么进行更新，否则进行添加
	if models.IsExistByUidAndPayType(merchantNo, payType) {
		if singleRoad == "" && rollPoolRoad == "" {
			//表示需要删除该支付类型的通道
			if models.DeleteMerchantDeployByUidAndPayType(merchantNo, payType) {
				keyDataJSON.Code = 200
				keyDataJSON.Msg = "删除该支付类型通道成功"
			} else {
				keyDataJSON.Code = -1
				keyDataJSON.Msg = "删除该支付类型通道失败"
			}
		} else {
			tmpInfo := models.GetMerchantDeployByUidAndPayType(merchantNo, payType)
			merchantDeployInfo.Id = tmpInfo.Id
			merchantDeployInfo.Status = tmpInfo.Status
			merchantDeployInfo.UpdateTime = utils.GetBasicDateTime()
			if models.UpdateMerchantDeploy(merchantDeployInfo) {
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
			if models.InsertMerchantDeployInfo(merchantDeployInfo) {
				keyDataJSON.Code = 200
				keyDataJSON.Msg = "添加支付类型成功"
			} else {
				keyDataJSON.Code = -1
				keyDataJSON.Msg = "添加支付类型失败"
			}
		}
	}
	this.GenerateJSON(keyDataJSON)
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
	//cardType := strings.TrimSpace(c.GetString("cardType"))
	bankAddress := strings.TrimSpace(c.GetString("bankAddress"))
	phone := strings.TrimSpace(c.GetString("phone"))
	payForAmount := strings.TrimSpace(c.GetString("payForAmount"))

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = -1

	if merchantUid == "" {
		keyDataJSON.Msg = "请选择需要下发的商户"
		c.GenerateJSON(keyDataJSON)
		return
	}

	if bankUid == "" {
		keyDataJSON.Msg = "请选择发下银行卡"
		c.GenerateJSON(keyDataJSON)
		return
	}

	money, err := strconv.ParseFloat(payForAmount, 64)
	if err != nil {
		logs.Error("add pay for fail： ", err)
		keyDataJSON.Msg = "下发金额输入不正确"
		c.GenerateJSON(keyDataJSON)
		return
	}

	accountInfo := models.GetAccountByUid(merchantUid)
	if accountInfo.SettleAmount < money+common.PAYFOR_FEE {
		keyDataJSON.Msg = "用户可用金额不够"
		c.GenerateJSON(keyDataJSON)
		return
	}

	bankInfo := models.GetBankCardByUid(bankUid)

	if bankInfo.BankNo != bankNo || bankInfo.AccountName != accountName || bankInfo.PhoneNo != phone {
		keyDataJSON.Msg = "银行卡信息有误，请连接管理员"
		c.GenerateJSON(keyDataJSON)
		return
	}

	payFor := models.PayforInfo{PayforUid: "pppp" + xid.New().String(), MerchantUid: merchantUid, MerchantName: merchantName, PhoneNo: phone,
		MerchantOrderId: xid.New().String(), BankOrderId: "4444" + xid.New().String(), PayforFee: common.PAYFOR_FEE, Type: common.SELF_MERCHANT,
		PayforAmount: money, PayforTotalAmount: money + common.PAYFOR_FEE, BankCode: bankInfo.BankCode, BankName: bankName, IsSend: common.NO,
		BankAccountName: bankInfo.AccountName, BankAccountNo: bankInfo.BankNo, BankAccountType: bankInfo.BankAccountType, BankAccountAddress: bankAddress,
		Status: common.PAYFOR_COMFRIM, CreateTime: utils.GetBasicDateTime(), UpdateTime: utils.GetBasicDateTime()}

	if models.InsertPayfor(payFor) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "代付下发提交失败"
	}

	c.GenerateJSON(keyDataJSON)
}

func (c *AddController) AddSelfPayFor() {
	bankUid := strings.TrimSpace(c.GetString("bankUid"))
	bankName := strings.TrimSpace(c.GetString("bankName"))
	accountName := strings.TrimSpace(c.GetString("accountName"))
	bankNo := strings.TrimSpace(c.GetString("bankNo"))
	//cardType := strings.TrimSpace(c.GetString("cardType"))
	bankAddress := strings.TrimSpace(c.GetString("bankAddress"))
	phone := strings.TrimSpace(c.GetString("phone"))
	payForAmount := strings.TrimSpace(c.GetString("payForAmount"))

	keyDataJSON := new(KeyDataJSON)
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

	bankInfo := models.GetBankCardByUid(bankUid)

	//需要对前端传入的数据做校验，不能完全相信前端的数据
	if bankInfo.AccountName != accountName || bankInfo.BankNo != bankNo || bankInfo.PhoneNo != phone {
		keyDataJSON.Msg = "前端页面数据有篡改，请注意资金安全"
		c.GenerateJSON(keyDataJSON)
		return
	}

	selfPayFor := models.PayforInfo{PayforUid: "pppp" + xid.New().String(), BankOrderId: "4444" + xid.New().String(), PayforFee: common.ZERO, Type: common.SELF_HELP,
		PayforAmount: money, PayforTotalAmount: money + common.ZERO, BankCode: bankInfo.BankCode, BankName: bankName, IsSend: common.NO,
		BankAccountName: bankInfo.AccountName, BankAccountNo: bankInfo.BankNo, BankAccountType: bankInfo.BankAccountType, BankAccountAddress: bankAddress,
		Status: common.PAYFOR_COMFRIM, CreateTime: utils.GetBasicDateTime(), UpdateTime: utils.GetBasicDateTime()}

	if models.InsertPayfor(selfPayFor) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Msg = "数据处理失败，请重新提交"
	}

	c.GenerateJSON(keyDataJSON)
}
