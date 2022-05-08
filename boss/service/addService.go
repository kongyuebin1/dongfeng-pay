package service

import (
	"boss/common"
	"boss/datas"
	"boss/models/accounts"
	"boss/models/agent"
	"boss/models/merchant"
	"boss/models/payfor"
	"boss/models/road"
	"boss/models/system"
	"boss/models/user"
	"boss/utils"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/rs/xid"
	"strconv"
	"strings"
)

type AddService struct {
}

func (c *AddService) AddMenu(oneMenu, userID string) *datas.BaseDataJSON {
	dataJSON := new(datas.BaseDataJSON)
	menuInfo := system.MenuInfo{
		MenuUid:    xid.New().String(),
		FirstMenu:  oneMenu,
		Status:     "active",
		Creater:    userID,
		CreateTime: utils.GetBasicDateTime(),
	}

	exist := system.FirstMenuIsExists(oneMenu)
	if !exist {
		menuInfo.MenuOrder = system.GetMenuLen() + 1
		flag := system.InsertMenu(menuInfo)
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
	return dataJSON
}

func (c *AddService) AddSecondMenu(firstMenuUid, secondRouter, secondMenu, userID string) *datas.KeyDataJSON {
	dataJSON := new(datas.KeyDataJSON)

	firstMenuInfo := system.GetMenuInfoByMenuUid(firstMenuUid)
	routerExists := system.SecondRouterExists(secondRouter)
	secondMenuExists := system.SecondMenuIsExists(secondMenu)

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
		sl := system.GetSecondMenuLenByFirstMenuUid(firstMenuUid)
		secondMenuInfo := system.SecondMenuInfo{
			MenuOrder:      sl + 1,
			FirstMenuUid:   firstMenuInfo.MenuUid,
			FirstMenu:      firstMenuInfo.FirstMenu,
			SecondMenuUid:  xid.New().String(),
			Status:         "active",
			SecondMenu:     secondMenu,
			SecondRouter:   secondRouter,
			Creater:        userID,
			CreateTime:     utils.GetBasicDateTime(),
			UpdateTime:     utils.GetBasicDateTime(),
			FirstMenuOrder: firstMenuInfo.MenuOrder,
		}
		if !system.InsertSecondMenu(secondMenuInfo) {
			dataJSON.Code = -1
			dataJSON.Msg = "添加二级菜单失败"
		} else {
			dataJSON.Code = 200
			dataJSON.Msg = "添加二级菜单成功"
		}
	}
	return dataJSON
}

func (c *AddService) AddPower(powerItem, powerID, firstMenuUid, secondMenuUid, userID string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = -1

	if powerItem == "" || len(powerItem) == 0 {
		keyDataJSON.Key = ".power-name-error"
		keyDataJSON.Msg = "*权限项名称不能为空"
		return keyDataJSON
	}
	if powerID == "" || len(powerID) == 0 {
		keyDataJSON.Key = ".power-id-error"
		keyDataJSON.Msg = "*权限项ID不能为空"
		return keyDataJSON
	}
	if system.PowerUidExists(powerID) {
		keyDataJSON.Key = ".power-id-error"
		keyDataJSON.Msg = "*权限项ID已经存在"
		return keyDataJSON
	}

	secondMenuInfo := system.GetSecondMenuInfoBySecondMenuUid(secondMenuUid)

	powerInfo := system.PowerInfo{
		SecondMenuUid: secondMenuUid,
		SecondMenu:    secondMenuInfo.SecondMenu,
		PowerId:       powerID, PowerItem: powerItem,
		Creater:      userID,
		Status:       "active",
		CreateTime:   utils.GetBasicDateTime(),
		UpdateTime:   utils.GetBasicDateTime(),
		FirstMenuUid: firstMenuUid,
	}

	keyDataJSON.Code = 200
	if !system.InsertPowerInfo(powerInfo) {
		keyDataJSON.Key = ".power-save-success"
		keyDataJSON.Msg = "添加权限项失败"
	} else {
		keyDataJSON.Key = ".power-save-success"
		keyDataJSON.Msg = "添加权限项成功"
	}

	return keyDataJSON
}

func (c *AddService) AddRole(roleName, roleRemark, userID string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	if len(roleName) == 0 {
		keyDataJSON.Code = -1
		keyDataJSON.Key = ".role-name-error"
		keyDataJSON.Msg = "*角色名称不能为空"
		return keyDataJSON
	}

	if system.RoleNameExists(roleName) {
		keyDataJSON.Code = -1
		keyDataJSON.Key = ".role-name-error"
		keyDataJSON.Msg = "*角色名称已经存在"
		return keyDataJSON
	}

	roleInfo := system.RoleInfo{
		RoleName:   roleName,
		RoleUid:    xid.New().String(),
		Creater:    userID,
		Status:     "active",
		Remark:     roleRemark,
		CreateTime: utils.GetBasicDateTime(),
		UpdateTime: utils.GetBasicDateTime(),
	}

	if !system.InsertRole(roleInfo) {
		keyDataJSON.Code = -1
		keyDataJSON.Key = ".role-save-success"
		keyDataJSON.Msg = "添加角色失败"
		return keyDataJSON
	}

	keyDataJSON.Code = 200

	return keyDataJSON
}

func (c *AddService) SavePower(roleUid string, firstMenuUids, secondMenuUids, powerIds []string) *datas.BaseDataJSON {
	dataJSON := new(datas.BaseDataJSON)
	roleInfo := system.GetRoleByRoleUid(roleUid)
	if len(roleUid) == 0 || len(roleInfo.RoleUid) == 0 {
		dataJSON.Code = -1
		return dataJSON
	}

	roleInfo.UpdateTime = utils.GetBasicDateTime()
	roleInfo.ShowFirstUid = strings.Join(firstMenuUids, "||")
	roleInfo.ShowSecondUid = strings.Join(secondMenuUids, "||")
	roleInfo.ShowPowerUid = strings.Join(powerIds, "||")

	menuInfoList := system.GetMenuInfosByMenuUids(firstMenuUids)
	showFirstMenu := make([]string, 0)
	for _, m := range menuInfoList {
		showFirstMenu = append(showFirstMenu, m.FirstMenu)
	}
	roleInfo.ShowFirstMenu = strings.Join(showFirstMenu, "||")

	secondMenuInfoList := system.GetSecondMenuInfoBySecondMenuUids(secondMenuUids)
	showSecondMenu := make([]string, 0)
	for _, m := range secondMenuInfoList {
		showSecondMenu = append(showSecondMenu, m.SecondMenu)
	}
	roleInfo.ShowSecondMenu = strings.Join(showSecondMenu, "||")

	powerList := system.GetPowerByIds(powerIds)
	showPower := make([]string, 0)
	for _, p := range powerList {
		showPower = append(showPower, p.PowerItem)
	}
	roleInfo.ShowPower = strings.Join(showPower, "||")

	if !system.UpdateRoleInfo(roleInfo) {
		dataJSON.Code = -1
		dataJSON.Msg = "更新roleInfo失败"
	} else {
		dataJSON.Code = 200
		dataJSON.Msg = "更新roleInfo成功"
	}

	return dataJSON
}

func (c *AddService) AddOperator(loginAccount, loginPassword, role, status, remark string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
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
	} else if user.UserInfoExistByUserId(loginAccount) {
		keyDataJSON.Key = ".operator-name-error"
		keyDataJSON.Msg = "*账号已经存在"
	} else {
		if len(remark) == 0 {
			remark = loginAccount
		}
		roleInfo := system.GetRoleByRoleUid(role)
		userInfo := user.UserInfo{
			UserId: loginAccount,
			Passwd: utils.GetMD5Upper(loginPassword),
			Nick:   "壮壮", Remark: remark,
			Status:     status,
			Role:       role,
			RoleName:   roleInfo.RoleName,
			CreateTime: utils.GetBasicDateTime(),
			UpdateTime: utils.GetBasicDateTime(),
		}
		if !user.InsertUser(userInfo) {
			keyDataJSON.Code = 200
			keyDataJSON.Msg = "添加操作员失败"
		} else {
			keyDataJSON.Code = 200
			keyDataJSON.Msg = "添加操作员成功"
		}
	}
	return keyDataJSON
}

func (c *AddService) AddBankCard(userName, bankCode, accountName, certificateType,
	phoneNo, bankName, bankAccountType, bankNo, certificateNo, bankAddress, uid, identifyCard string) *datas.BaseDataJSON {
	dataJSON := new(datas.BaseDataJSON)

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
			bankCardInfo := system.GetBankCardByUid(uid)
			bankCardInfo = system.BankCardInfo{
				Id:              bankCardInfo.Id,
				UserName:        userName,
				BankName:        bankName,
				BankCode:        bankCode,
				BankAccountType: bankAccountType,
				AccountName:     accountName,
				BankNo:          bankNo,
				IdentifyCard:    identifyCard,
				CertificateNo:   certificateNo,
				PhoneNo:         phoneNo,
				BankAddress:     bankAddress,
				UpdateTime:      utils.GetBasicDateTime(),
				CreateTime:      bankCardInfo.CreateTime,
				Uid:             bankCardInfo.Uid,
			}
			if system.UpdateBankCard(bankCardInfo) {
				dataJSON.Code = 200
			}
		} else {
			bankCardInfo := system.BankCardInfo{
				Uid:             "3333" + xid.New().String(),
				UserName:        userName,
				BankName:        bankName,
				BankCode:        bankCode,
				BankAccountType: bankAccountType,
				AccountName:     accountName,
				BankNo:          bankNo,
				IdentifyCard:    identifyCard,
				CertificateNo:   certificateNo,
				PhoneNo:         phoneNo,
				BankAddress:     bankAddress,
				UpdateTime:      utils.GetBasicDateTime(),
				CreateTime:      utils.GetBasicDateTime(),
			}

			if system.InsertBankCardInfo(bankCardInfo) {
				dataJSON.Code = 200
			}
		}
	}
	return dataJSON
}

func (c *AddService) AddRoad(startHour, endHour, roadName, productUid,
	payType, basicRate, settleFee, roadTotalLimit, roadEverydayLimit,
	singleMinLimit, singleMaxLimit, roadUid, roadRemark, params string) *datas.BaseDataJSON {

	dataJSON := new(datas.BaseDataJSON)
	dataJSON.Code = -1

	startHourTmp, err1 := strconv.Atoi(startHour)
	endHourTmp, err2 := strconv.Atoi(endHour)

	if err1 != nil || err2 != nil {
		dataJSON.Msg = "开始时间或者结束时间设置有误"
		return dataJSON
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
			return dataJSON
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
			roadInfo := road.GetRoadInfoByRoadUid(roadUid)
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

			if road.UpdateRoadInfo(roadInfo) {
				dataJSON.Code = 200
			} else {
				dataJSON.Msg = "通道更新失败"
			}
		} else {
			//添加新的通道
			roadUid = "4444" + xid.New().String()
			roadInfo := road.RoadInfo{
				RoadName:       roadName,
				RoadUid:        roadUid,
				Remark:         roadRemark,
				ProductUid:     productUid,
				ProductName:    productName,
				PayType:        payType,
				BasicFee:       basicFee,
				SettleFee:      settleFeeTmp,
				TotalLimit:     totalLimit,
				TodayLimit:     todayLimit,
				SingleMinLimit: singleMinLimitTmp,
				Balance:        common.ZERO,
				SingleMaxLimit: singleMaxLimitTmp,
				StarHour:       startHourTmp,
				EndHour:        endHourTmp,
				Status:         "active",
				Params:         params,
				UpdateTime:     utils.GetBasicDateTime(),
				CreateTime:     utils.GetBasicDateTime(),
			}

			if road.InsertRoadInfo(roadInfo) {
				dataJSON.Code = 200
			} else {
				dataJSON.Msg = "添加新通道失败"
			}
		}
	}
	return dataJSON
}

func (c *AddService) AddRoadPool(roadPoolName, roadPoolCode string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = -1

	if len(roadPoolName) == 0 {
		keyDataJSON.Msg = "*通道池名称不能为空"
	} else if len(roadPoolCode) == 0 {
		keyDataJSON.Msg = "*通道池编号不能为空"
	}

	roadPoolInfo := road.RoadPoolInfo{
		Status:       "active",
		RoadPoolName: roadPoolName,
		RoadPoolCode: roadPoolCode,
		UpdateTime:   utils.GetBasicDateTime(),
		CreateTime:   utils.GetBasicDateTime(),
	}

	if road.InsertRoadPool(roadPoolInfo) {
		keyDataJSON.Code = 200
		keyDataJSON.Msg = "添加通道池成功"
	} else {
		keyDataJSON.Msg = "添加通道池失败"
	}

	return keyDataJSON
}

func (c *AddService) SaveRoadUid(roadPoolCode string, roadUids []string) *datas.BaseDataJSON {
	dataJSON := new(datas.BaseDataJSON)
	dataJSON.Code = -1
	roadPoolInfo := road.GetRoadPoolByRoadPoolCode(roadPoolCode)
	if roadPoolInfo.RoadPoolCode == "" {
		return dataJSON
	}
	var uids []string
	for _, uid := range roadUids {
		//去掉空格
		if len(uid) > 0 && road.RoadInfoExistByRoadUid(uid) {
			uids = append(uids, uid)
		}
	}
	if len(uids) > 0 {
		roadUid := strings.Join(uids, "||")
		roadPoolInfo.RoadUidPool = roadUid
	}
	roadPoolInfo.UpdateTime = utils.GetBasicDateTime()
	if road.UpdateRoadPool(roadPoolInfo) {
		dataJSON.Code = 200
	}
	return dataJSON
}

func (c *AddService) AddAgent(agentName, agentPhone, agentLoginPassword,
	agentVertifyPassword, status, agentUid, agentRemark string) *datas.KeyDataJSON {

	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = 200

	if agentName == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-name-error"
		keyDataJSON.Msg = "代理名不能为空"
	} else if agent.IsEixstByAgentName(agentName) {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-name-error"
		keyDataJSON.Msg = "已存在该代理名称"
	} else if agentPhone == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#agent-phone-error"
		keyDataJSON.Msg = "代理注册手机号不能为空"
	} else if agent.IsEixstByAgentPhone(agentPhone) {
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
		return keyDataJSON
	}

	if status == "" {
		status = "active"
	}

	if agentUid == "" {

		agentUid = "9999" + xid.New().String()

		agentInfo := agent.AgentInfo{
			Status:        status,
			AgentName:     agentName,
			AgentPhone:    agentPhone,
			AgentPassword: utils.GetMD5Upper(agentLoginPassword),
			AgentUid:      agentUid,
			UpdateTime:    utils.GetBasicDateTime(),
			CreateTime:    utils.GetBasicDateTime(),
			AgentRemark:   agentRemark,
		}

		if !agent.InsertAgentInfo(agentInfo) {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "添加代理商失败"
		}
	}

	//创建新的账户
	account := accounts.GetAccountByUid(agentUid)
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
		if accounts.InsetAcount(account) {
			keyDataJSON.Code = 200
			keyDataJSON.Msg = "插入成功"
		} else {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "掺入失败"
		}
	}
	return keyDataJSON
}
func (c *AddService) AddMerchant(merchantName, phone, loginPassword,
	verifyPassword, merchantStatus, remark string) *datas.KeyDataJSON {

	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = 200
	if merchantName == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#merchant-name-error"
		keyDataJSON.Msg = "商户名称为空"
	} else if merchant.IsExistByMerchantName(merchantName) {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#merchant-name-error"
		keyDataJSON.Msg = "商户名已经存在"
	} else if phone == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Key = "#merchant-phone-error"
		keyDataJSON.Msg = "手机号为空"
	} else if merchant.IsExistByMerchantPhone(phone) {
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
		return keyDataJSON
	}
	merchantUid := "8888" + xid.New().String()
	merchantKey := "kkkk" + xid.New().String()    //商户key
	merchantSecret := "ssss" + xid.New().String() //商户密钥
	merchantInfo := merchant.MerchantInfo{
		MerchantName:   merchantName,
		MerchantUid:    merchantUid,
		LoginAccount:   phone,
		MerchantKey:    merchantKey,
		MerchantSecret: merchantSecret,
		LoginPassword:  utils.GetMD5Upper(loginPassword),
		Status:         merchantStatus,
		Remark:         remark,
		UpdateTime:     utils.GetBasicDateTime(),
		CreateTime:     utils.GetBasicDateTime(),
	}

	if merchant.InsertMerchantInfo(merchantInfo) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "插入失败"
	}
	//创建新的账户
	account := accounts.GetAccountByUid(merchantUid)
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
		if accounts.InsetAcount(account) {
			keyDataJSON.Code = 200
			keyDataJSON.Msg = "插入成功"
		} else {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "掺入失败"
		}
	}
	return keyDataJSON
}

func (c *AddService) AddMerchantDeploy(merchantUid, isAutoSettle, isAutoPayfor, ipWhite, belongAgentName,
	belongAgentUid, payforRoadChoose, rollPayforRoadChoose, payforFee string) *datas.KeyDataJSON {

	keyDataJSON := new(datas.KeyDataJSON)
	merchantInfo := merchant.GetMerchantByUid(merchantUid)
	merchantInfo.AutoSettle = isAutoSettle
	merchantInfo.AutoPayFor = isAutoPayfor
	merchantInfo.WhiteIps = ipWhite
	merchantInfo.BelongAgentName = belongAgentName
	merchantInfo.BelongAgentUid = belongAgentUid

	if payforRoadChoose != "" {
		roadInfo := road.GetRoadInfoByName(payforRoadChoose)
		merchantInfo.SinglePayForRoadName = payforRoadChoose
		merchantInfo.SinglePayForRoadUid = roadInfo.RoadUid
	}
	if rollPayforRoadChoose != "" {
		rollPoolInfo := road.GetRoadPoolByName(rollPayforRoadChoose)
		merchantInfo.RollPayForRoadName = rollPayforRoadChoose
		merchantInfo.RollPayForRoadCode = rollPoolInfo.RoadPoolCode
	}
	tmp, err := strconv.ParseFloat(payforFee, 64)
	if err != nil {
		logs.Error("手续费由字符串转为float64失败")
		tmp = common.PAYFOR_FEE
	}
	merchantInfo.PayforFee = tmp
	if merchant.UpdateMerchant(merchantInfo) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
	}
	return keyDataJSON
}

func (c *AddService) AddMerchantPayType(payType, singleRoad, rollPoolRoad,
	singleRoadPlatformFee, rollRoadPlatformFee, singleRoadAgentFee, rollRoadAgentFee,
	loanRate, loanDays, unfreezeTimeHour, merchantNo, isLoan string) *datas.KeyDataJSON {

	keyDataJSON := new(datas.KeyDataJSON)
	if payType == "" || payType == "none" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "操作失败，请选择支付类型"
		return keyDataJSON
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
		return keyDataJSON
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
	return keyDataJSON
}

func (c *AddService) AddPayFor(merchantUid, bankUid, payForAmount, bankNo, accountName,
	phone, merchantName, bankName, bankAddress string) *datas.KeyDataJSON {

	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = -1

	if merchantUid == "" {
		keyDataJSON.Msg = "请选择需要下发的商户"
		return keyDataJSON
	}

	if bankUid == "" {
		keyDataJSON.Msg = "请选择发下银行卡"
		return keyDataJSON
	}

	money, err := strconv.ParseFloat(payForAmount, 64)
	if err != nil {
		logs.Error("add pay for fail： ", err)
		keyDataJSON.Msg = "下发金额输入不正确"
		return keyDataJSON
	}

	accountInfo := accounts.GetAccountByUid(merchantUid)
	if accountInfo.SettleAmount < money+common.PAYFOR_FEE {
		keyDataJSON.Msg = "用户可用金额不够"
		return keyDataJSON
	}

	bankInfo := system.GetBankCardByUid(bankUid)

	if bankInfo.BankNo != bankNo || bankInfo.AccountName != accountName || bankInfo.PhoneNo != phone {
		keyDataJSON.Msg = "银行卡信息有误，请连接管理员"
		return keyDataJSON
	}

	payFor := payfor.PayforInfo{
		PayforUid:    "pppp" + xid.New().String(),
		MerchantUid:  merchantUid,
		MerchantName: merchantName, PhoneNo: phone,
		MerchantOrderId: xid.New().String(),
		BankOrderId:     "4444" + xid.New().String(),
		PayforFee:       common.PAYFOR_FEE, Type: common.SELF_MERCHANT,
		PayforAmount:      money,
		PayforTotalAmount: money + common.PAYFOR_FEE,
		BankCode:          bankInfo.BankCode,
		BankName:          bankName, IsSend: common.NO,
		BankAccountName:    bankInfo.AccountName,
		BankAccountNo:      bankInfo.BankNo,
		BankAccountType:    bankInfo.BankAccountType,
		BankAccountAddress: bankAddress,
		Status:             common.PAYFOR_COMFRIM,
		RequestTime:        utils.GetBasicDateTime(),
		CreateTime:         utils.GetBasicDateTime(),
		UpdateTime:         utils.GetBasicDateTime(),
	}

	if payfor.InsertPayfor(payFor) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "代付下发提交失败"
	}

	return keyDataJSON
}

func (c *AddService) AddSelfPayFor(bankUid, payForAmount, accountName,
	bankNo, phone, bankName, bankAddress string) *datas.KeyDataJSON {

	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = -1

	if bankUid == "" {
		keyDataJSON.Msg = "银行卡uid不能为空，请联系技术人员"
		return keyDataJSON
	}
	money, err := strconv.ParseFloat(payForAmount, 64)
	if err != nil {
		logs.Error("self payfor money fail: ", err)
		keyDataJSON.Msg = "输入金额有误，请仔细检查"
		return keyDataJSON
	}

	bankInfo := system.GetBankCardByUid(bankUid)

	//需要对前端传入的数据做校验，不能完全相信前端的数据
	if bankInfo.AccountName != accountName || bankInfo.BankNo != bankNo || bankInfo.PhoneNo != phone {
		keyDataJSON.Msg = "前端页面数据有篡改，请注意资金安全"
		return keyDataJSON
	}

	selfPayFor := payfor.PayforInfo{
		PayforUid:          "pppp" + xid.New().String(),
		BankOrderId:        "4444" + xid.New().String(),
		PayforFee:          common.ZERO,
		Type:               common.SELF_HELP,
		PayforAmount:       money,
		PayforTotalAmount:  money + common.ZERO,
		BankCode:           bankInfo.BankCode,
		BankName:           bankName,
		IsSend:             common.NO,
		BankAccountName:    bankInfo.AccountName,
		BankAccountNo:      bankInfo.BankNo,
		BankAccountType:    bankInfo.BankAccountType,
		BankAccountAddress: bankAddress,
		Status:             common.PAYFOR_COMFRIM,
		CreateTime:         utils.GetBasicDateTime(),
		UpdateTime:         utils.GetBasicDateTime(),
	}

	if payfor.InsertPayfor(selfPayFor) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Msg = "数据处理失败，请重新提交"
	}
	return keyDataJSON
}
