/***************************************************
 ** @Desc : c file for ...
 ** @Time : 2019/8/21 11:18
 ** @Author : yuebin
 ** @File : get
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/21 11:18
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"boss/common"
	"boss/models"
	controller "boss/supplier"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type GetController struct {
	BaseController
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	JumpPage     int
	Offset       int
}

/*
* 处理分页的函数
 */
func (c *GetController) GetCutPage(l int) {
	c.DisplayCount, _ = c.GetInt("displayCount")
	c.CurrentPage, _ = c.GetInt("currentPage")
	c.TotalPage, _ = c.GetInt("totalPage")
	c.JumpPage, _ = c.GetInt("jumpPage")

	if c.CurrentPage == 0 {
		c.CurrentPage = 1
	}
	if c.DisplayCount == 0 {
		c.DisplayCount = 20
	}
	if c.JumpPage > 0 {
		c.CurrentPage = c.JumpPage
	}

	if l > 0 {
		c.TotalPage = l / c.DisplayCount
		if l%c.DisplayCount > 0 {
			c.TotalPage += 1
		}
	} else {
		c.TotalPage = 0
		c.CurrentPage = 0
	}
	//假如当前页超过了总页数
	if c.CurrentPage > c.TotalPage {
		c.CurrentPage = c.TotalPage
	}
	//计算出偏移量
	c.Offset = (c.CurrentPage - 1) * c.DisplayCount
}

func (c *GetController) GetMenu() {

	firstMenuSearch := strings.TrimSpace(c.GetString("firstMenuSearch"))

	params := make(map[string]string)
	params["first_menu__icontains"] = firstMenuSearch
	c.GetCutPage(models.GetMenuLenByMap(params))

	menuDataJSON := new(MenuDataJSON)
	menuDataJSON.DisplayCount = c.DisplayCount
	menuDataJSON.CurrentPage = c.CurrentPage
	menuDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		menuDataJSON.Code = -1
		menuDataJSON.MenuList = make([]models.MenuInfo, 0)
		c.GenerateJSON(menuDataJSON)
		return
	}

	menuInfoList := models.GetMenuOffsetByMap(params, c.DisplayCount, c.Offset)
	sort.Sort(models.MenuInfoSlice(menuInfoList))
	for i, m := range menuInfoList {
		secondMenuInfoList := models.GetSecondMenuListByFirstMenuUid(m.MenuUid)
		smenus := ""
		for j := 0; j < len(secondMenuInfoList); j++ {
			smenus += secondMenuInfoList[j].SecondMenu
			if j != (len(secondMenuInfoList) - 1) {
				smenus += "||"
			}
		}
		m.SecondMenu = smenus
		menuInfoList[i] = m
	}
	menuDataJSON.Code = 200
	menuDataJSON.MenuList = menuInfoList
	menuDataJSON.StartIndex = c.Offset

	if len(menuInfoList) == 0 {
		menuDataJSON.Msg = "获取菜单列表失败"
	}

	c.GenerateJSON(menuDataJSON)
}

func (c *GetController) GetFirstMenu() {
	menuDataJSON := new(MenuDataJSON)
	menuList := models.GetMenuAll()

	if len(menuList) == 0 {
		menuDataJSON.Code = -1
	} else {
		menuDataJSON.Code = 200
	}
	sort.Sort(models.MenuInfoSlice(menuList))
	menuDataJSON.MenuList = menuList
	c.GenerateJSON(menuDataJSON)
}

/*
*获取所有的二级菜单
 */
func (c *GetController) GetSecondMenu() {

	firstMenuSearch := strings.TrimSpace(c.GetString("firstMenuSerach"))
	secondMenuSearch := strings.TrimSpace(c.GetString("secondMenuSerach"))

	params := make(map[string]string)
	params["first_menu__icontains"] = firstMenuSearch
	params["second_menu__icontains"] = secondMenuSearch
	l := models.GetSecondMenuLenByMap(params)

	c.GetCutPage(l)

	secondMenuDataJSON := new(SecondMenuDataJSON)
	secondMenuDataJSON.DisplayCount = c.DisplayCount

	secondMenuDataJSON.Code = 200
	secondMenuDataJSON.CurrentPage = c.CurrentPage
	secondMenuDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		secondMenuDataJSON.SecondMenuList = make([]models.SecondMenuInfo, 0)
		c.GenerateJSON(secondMenuDataJSON)
		return
	}

	secondMenuList := models.GetSecondMenuByMap(params, c.DisplayCount, c.Offset)
	sort.Sort(models.SecondMenuSlice(secondMenuList))
	secondMenuDataJSON.SecondMenuList = secondMenuList
	secondMenuDataJSON.StartIndex = c.Offset

	if len(secondMenuList) == 0 {
		secondMenuDataJSON.Msg = "获取二级菜单失败"
	}

	c.GenerateJSON(secondMenuDataJSON)
}

func (c *GetController) GetSecondMenus() {
	firstMenuUid := strings.TrimSpace(c.GetString("firMenuUid"))

	secondMenuDataJSON := new(SecondMenuDataJSON)

	secondMenuList := models.GetSecondMenuListByFirstMenuUid(firstMenuUid)

	secondMenuDataJSON.Code = 200
	secondMenuDataJSON.SecondMenuList = secondMenuList
	c.GenerateJSON(secondMenuDataJSON)
}

func (c *GetController) GetOneMenu() {
	menuUid := c.GetString("menuUid")

	dataJSON := new(MenuDataJSON)
	menuInfo := models.GetMenuInfoByMenuUid(menuUid)
	if menuInfo.MenuUid == "" {
		dataJSON.Code = -1
		dataJSON.Msg = "该菜单项不存在"
	} else {
		dataJSON.MenuList = make([]models.MenuInfo, 0)
		dataJSON.MenuList = append(dataJSON.MenuList, menuInfo)
		dataJSON.Code = 200
	}
	c.Data["json"] = dataJSON
	c.ServeJSONP()
}

func (c *GetController) GetPowerItem() {
	powerID := c.GetString("powerID")
	powerItem := c.GetString("powerItem")

	params := make(map[string]string)
	params["power_uid__icontains"] = powerID
	params["power_item_icontains"] = powerItem

	l := models.GetPowerItemLenByMap(params)

	c.GetCutPage(l)

	powerItemDataJSON := new(PowerItemDataJSON)
	powerItemDataJSON.DisplayCount = c.DisplayCount
	powerItemDataJSON.Code = 200
	powerItemDataJSON.CurrentPage = c.CurrentPage
	powerItemDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		powerItemDataJSON.PowerItemList = make([]models.PowerInfo, 0)
		c.GenerateJSON(powerItemDataJSON)
		return
	}

	powerItemDataJSON.StartIndex = c.Offset
	powerItemList := models.GetPowerItemByMap(params, c.DisplayCount, c.Offset)
	sort.Sort(models.PowerInfoSlice(powerItemList))
	powerItemDataJSON.PowerItemList = powerItemList

	c.GenerateJSON(powerItemDataJSON)
}

func (c *GetController) GetRole() {
	roleName := strings.TrimSpace(c.GetString("roleName"))

	params := make(map[string]string)
	params["role_name__icontains"] = roleName

	l := models.GetRoleLenByMap(params)

	c.GetCutPage(l)

	roleInfoDataJSON := new(RoleInfoDataJSON)
	roleInfoDataJSON.DisplayCount = c.DisplayCount
	roleInfoDataJSON.Code = 200
	roleInfoDataJSON.CurrentPage = c.CurrentPage
	roleInfoDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		roleInfoDataJSON.RoleInfoList = make([]models.RoleInfo, 0)
		c.GenerateJSON(roleInfoDataJSON)
		return
	}
	roleInfoDataJSON.StartIndex = c.Offset
	roleInfoDataJSON.RoleInfoList = models.GetRoleByMap(params, c.DisplayCount, c.Offset)
	c.GenerateJSON(roleInfoDataJSON)
}

func (c *GetController) GetAllRole() {
	roleInfoDataJSON := new(RoleInfoDataJSON)
	roleInfoList := models.GetRole()
	fmt.Println(roleInfoList)
	if len(roleInfoList) == 0 {
		roleInfoDataJSON.Code = -1
	} else {
		roleInfoDataJSON.Code = 200
		roleInfoDataJSON.RoleInfoList = roleInfoList
	}
	c.GenerateJSON(roleInfoDataJSON)
}

func (c *GetController) GetDeployTree() {
	roleUid := strings.TrimSpace(c.GetString("roleUid"))
	roleInfo := models.GetRoleByRoleUid(roleUid)

	allFirstMenu := models.GetMenuAll()
	sort.Sort(models.MenuInfoSlice(allFirstMenu))
	allSecondMenu := models.GetSecondMenuList()
	sort.Sort(models.SecondMenuSlice(allSecondMenu))
	allPower := models.GetPower()

	deployTreeJSON := new(DeployTreeJSON)
	deployTreeJSON.Code = 200
	deployTreeJSON.AllFirstMenu = allFirstMenu
	deployTreeJSON.AllSecondMenu = allSecondMenu
	deployTreeJSON.AllPower = allPower
	deployTreeJSON.ShowFirstMenuUid = make(map[string]bool)
	for _, v := range strings.Split(roleInfo.ShowFirstUid, "||") {
		deployTreeJSON.ShowFirstMenuUid[v] = true
	}
	deployTreeJSON.ShowSecondMenuUid = make(map[string]bool)
	for _, v := range strings.Split(roleInfo.ShowSecondUid, "||") {
		deployTreeJSON.ShowSecondMenuUid[v] = true
	}
	deployTreeJSON.ShowPowerUid = make(map[string]bool)
	for _, v := range strings.Split(roleInfo.ShowPowerUid, "||") {
		deployTreeJSON.ShowPowerUid[v] = true
	}

	c.GenerateJSON(deployTreeJSON)
}

/*
* 获取操作员列表
 */
func (c *GetController) GetOperator() {
	operatorName := strings.TrimSpace(c.GetString("operatorName"))

	params := make(map[string]string)
	params["user_id__icontains"] = operatorName

	l := models.GetOperatorLenByMap(params)
	c.GetCutPage(l)
	operatorDataJSON := new(OperatorDataJSON)
	operatorDataJSON.DisplayCount = c.DisplayCount
	operatorDataJSON.Code = 200
	operatorDataJSON.CurrentPage = c.CurrentPage
	operatorDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		operatorDataJSON.OperatorList = make([]models.UserInfo, 0)
		c.GenerateJSON(operatorDataJSON)
		return
	}

	operatorDataJSON.StartIndex = c.Offset
	operatorDataJSON.OperatorList = models.GetOperatorByMap(params, c.DisplayCount, c.Offset)
	c.GenerateJSON(operatorDataJSON)
}

func (c *GetController) GetOneOperator() {
	userId := strings.TrimSpace(c.GetString("userId"))

	userInfo := models.GetUserInfoByUserID(userId)

	operatorDataJSON := new(OperatorDataJSON)
	operatorDataJSON.OperatorList = make([]models.UserInfo, 0)
	operatorDataJSON.OperatorList = append(operatorDataJSON.OperatorList, userInfo)

	operatorDataJSON.Code = 200

	c.GenerateJSON(operatorDataJSON)
}

func (c *GetController) GetEditOperator() {
	userId := strings.TrimSpace(c.GetString("userId"))

	editOperatorDataJSON := new(EditOperatorDataJSON)
	userInfo := models.GetUserInfoByUserID(userId)
	fmt.Println(userInfo)
	editOperatorDataJSON.OperatorList = append(editOperatorDataJSON.OperatorList, userInfo)
	editOperatorDataJSON.RoleList = models.GetRole()
	editOperatorDataJSON.Code = 200

	c.GenerateJSON(editOperatorDataJSON)
}

func (c *GetController) GetBankCard() {
	accountNameSearch := strings.TrimSpace(c.GetString("accountNameSearch"))
	params := make(map[string]string)
	params["account_name__icontains"] = accountNameSearch

	l := models.GetBankCardLenByMap(params)
	c.GetCutPage(l)

	bankCardDataJSON := new(BankCardDataJSON)
	bankCardDataJSON.DisplayCount = c.DisplayCount
	bankCardDataJSON.Code = 200
	bankCardDataJSON.CurrentPage = c.CurrentPage
	bankCardDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		bankCardDataJSON.BankCardInfoList = make([]models.BankCardInfo, 0)
		c.GenerateJSON(bankCardDataJSON)
		return
	}

	bankCardDataJSON.StartIndex = c.Offset
	bankCardDataJSON.BankCardInfoList = models.GetBankCardByMap(params, c.DisplayCount, c.Offset)
	c.GenerateJSON(bankCardDataJSON)
}

func (c *GetController) GetOneBankCard() {
	uid := strings.TrimSpace(c.GetString("uid"))
	bankCardInfo := models.GetBankCardByUid(uid)

	bankCardDataJSON := new(BankCardDataJSON)
	bankCardDataJSON.Code = -1

	if bankCardInfo.Uid != "" {
		bankCardDataJSON.BankCardInfoList = append(bankCardDataJSON.BankCardInfoList, bankCardInfo)
		bankCardDataJSON.Code = 200
	}

	c.GenerateJSON(bankCardDataJSON)
}

/*
* 获取通道
 */
func (c *GetController) GetRoad() {
	roadName := strings.TrimSpace(c.GetString("roadName"))
	productName := strings.TrimSpace(c.GetString("productName"))
	roadUid := strings.TrimSpace(c.GetString("roadUid"))
	payType := strings.TrimSpace(c.GetString("payType"))
	roadPoolCode := strings.TrimSpace(c.GetString("roadPoolCode"))

	params := make(map[string]string)
	params["road_name__icontains"] = roadName
	params["product_name__icontains"] = productName
	params["road_uid"] = roadUid
	params["pay_type"] = payType

	l := models.GetRoadLenByMap(params)
	c.GetCutPage(l)

	roadDataJSON := new(RoadDataJSON)
	roadDataJSON.DisplayCount = c.DisplayCount
	roadDataJSON.Code = 200
	roadDataJSON.CurrentPage = c.CurrentPage
	roadDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		roadDataJSON.RoadInfoList = make([]models.RoadInfo, 0)
		c.GenerateJSON(roadDataJSON)
		return
	}

	roadDataJSON.StartIndex = c.Offset
	roadDataJSON.RoadInfoList = models.GetRoadInfoByMap(params, c.DisplayCount, c.Offset)
	roadDataJSON.RoadPool = models.GetRoadPoolByRoadPoolCode(roadPoolCode)
	c.GenerateJSON(roadDataJSON)
}

func (c *GetController) GetAllRoad() {
	roadName := strings.TrimSpace(c.GetString("roadName"))
	params := make(map[string]string)
	params["road_name__icontains"] = roadName

	roadDataJSON := new(RoadDataJSON)
	roadInfoList := models.GetAllRoad(params)

	roadDataJSON.Code = 200
	roadDataJSON.RoadInfoList = roadInfoList
	c.GenerateJSON(roadDataJSON)
}

/*
* 获取单个通道
 */
func (c *GetController) GetOneRoad() {
	roadUid := strings.TrimSpace(c.GetString("roadUid"))

	roadInfo := models.GetRoadInfoByRoadUid(roadUid)
	roadDataJSON := new(RoadDataJSON)
	roadDataJSON.Code = -1

	if roadInfo.RoadUid != "" {
		roadDataJSON.RoadInfoList = append(roadDataJSON.RoadInfoList, roadInfo)
		roadDataJSON.Code = 200
	} else {
		roadDataJSON.RoadInfoList = make([]models.RoadInfo, 0)
	}

	c.GenerateJSON(roadDataJSON)
}

func (c *GetController) GetRoadPool() {
	roadPoolName := strings.TrimSpace(c.GetString("roadPoolName"))
	roadPoolCode := strings.TrimSpace(c.GetString("roadPoolCode"))

	params := make(map[string]string)
	params["road_pool_name__icontains"] = roadPoolName
	params["road_pool_code__icontains"] = roadPoolCode

	l := models.GetRoadPoolLenByMap(params)
	c.GetCutPage(l)

	roadPoolDataJSON := new(RoadPoolDataJSON)
	roadPoolDataJSON.DisplayCount = c.DisplayCount
	roadPoolDataJSON.Code = 200
	roadPoolDataJSON.CurrentPage = c.CurrentPage
	roadPoolDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		roadPoolDataJSON.RoadPoolInfoList = make([]models.RoadPoolInfo, 0)
		c.GenerateJSON(roadPoolDataJSON)
		return
	}

	roadPoolDataJSON.StartIndex = c.Offset
	roadPoolDataJSON.RoadPoolInfoList = models.GetRoadPoolByMap(params, c.DisplayCount, c.Offset)
	c.GenerateJSON(roadPoolDataJSON)
}

func (c *GetController) GetAllRollPool() {
	rollPoolName := strings.TrimSpace(c.GetString("rollPoolName"))
	params := make(map[string]string)
	params["road_pool_name__icontains"] = rollPoolName

	roadPoolDataJSON := new(RoadPoolDataJSON)
	roadPoolDataJSON.Code = 200
	roadPoolDataJSON.RoadPoolInfoList = models.GetAllRollPool(params)
	c.GenerateJSON(roadPoolDataJSON)
}

func (c *GetController) GetMerchant() {
	merchantName := strings.TrimSpace(c.GetString("merchantName"))
	merchantNo := strings.TrimSpace(c.GetString("merchantNo"))

	params := make(map[string]string)
	params["merchant_name__icontains"] = merchantName
	params["merchant_uid__icontains"] = merchantNo

	l := models.GetMerchantLenByMap(params)
	c.GetCutPage(l)

	merchantDataJSON := new(MerchantDataJSON)
	merchantDataJSON.DisplayCount = c.DisplayCount
	merchantDataJSON.Code = 200
	merchantDataJSON.CurrentPage = c.CurrentPage
	merchantDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		merchantDataJSON.MerchantList = make([]models.MerchantInfo, 0)
		c.GenerateJSON(merchantDataJSON)
		return
	}

	merchantDataJSON.StartIndex = c.Offset
	merchantDataJSON.MerchantList = models.GetMerchantListByMap(params, c.DisplayCount, c.Offset)
	c.GenerateJSON(merchantDataJSON)
}

func (c *GetController) GetAllMerchant() {
	merchantDataJSON := new(MerchantDataJSON)
	merchantDataJSON.Code = 200
	merchantDataJSON.MerchantList = models.GetAllMerchant()
	c.GenerateJSON(merchantDataJSON)
}

func (c *GetController) GetOneMerchant() {
	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))
	merchantDataJSON := new(MerchantDataJSON)

	if merchantUid == "" {
		merchantDataJSON.Code = -1
		c.GenerateJSON(merchantDataJSON)
		return
	}

	merchantInfo := models.GetMerchantByUid(merchantUid)

	merchantDataJSON.Code = 200
	merchantDataJSON.MerchantList = append(merchantDataJSON.MerchantList, merchantInfo)
	c.GenerateJSON(merchantDataJSON)
}

func (c *GetController) GetOneMerchantDeploy() {
	merchantNo := strings.TrimSpace(c.GetString("merchantNo"))
	payType := strings.TrimSpace(c.GetString("payType"))

	merchantDeployDataJSON := new(MerchantDeployDataJSON)

	merchantDeployInfo := models.GetMerchantDeployByUidAndPayType(merchantNo, payType)

	if merchantDeployInfo.Status == "active" {
		merchantDeployDataJSON.Code = 200
		merchantDeployDataJSON.MerchantDeploy = merchantDeployInfo
	} else {
		merchantDeployDataJSON.Code = -1
		merchantDeployDataJSON.MerchantDeploy = merchantDeployInfo
	}

	c.GenerateJSON(merchantDeployDataJSON)
}

func (c *GetController) GetAllAccount() {
	accountDataJSON := new(AccountDataJSON)
	accountDataJSON.Code = 200

	accountDataJSON.AccountList = models.GetAllAccount()

	c.GenerateJSON(accountDataJSON)
}

func (c *GetController) GetAccount() {
	accountName := strings.TrimSpace(c.GetString("accountName"))
	accountUid := strings.TrimSpace(c.GetString("accountNo"))

	params := make(map[string]string)
	params["account_name__icontains"] = accountName
	params["account_uid_icontains"] = accountUid

	l := models.GetAccountLenByMap(params)
	c.GetCutPage(l)

	accountDataJSON := new(AccountDataJSON)
	accountDataJSON.DisplayCount = c.DisplayCount
	accountDataJSON.Code = 200
	accountDataJSON.CurrentPage = c.CurrentPage
	accountDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		accountDataJSON.AccountList = make([]models.AccountInfo, 0)
		c.GenerateJSON(accountDataJSON)
		return
	}

	accountDataJSON.StartIndex = c.Offset
	accountDataJSON.AccountList = models.GetAccountByMap(params, c.DisplayCount, c.Offset)
	c.GenerateJSON(accountDataJSON)
}

func (c *GetController) GetOneAccount() {
	//从http的body中获取accountUid字段，并且这个字段是string类型
	accountUid := strings.TrimSpace(c.GetString("accountUid"))
	//new一个accountDataJSON结构体对象，用来做jsonp返回
	accountDataJSON := new(AccountDataJSON)
	//用accountuid作为过滤字段，从数据库中读取一条信息
	accountInfo := models.GetAccountByUid(accountUid)
	//code初始值为200
	accountDataJSON.Code = 200
	//将从数据库读出来的数据插入到accountList数组中
	accountDataJSON.AccountList = append(accountDataJSON.AccountList, accountInfo)
	//返回jsonp格式的数据给前端
	c.GenerateJSON(accountDataJSON)
}

func (c *GetController) GetAccountHistory() {
	accountName := strings.TrimSpace(c.GetString("accountHistoryName"))
	accountUid := strings.TrimSpace(c.GetString("accountHistoryNo"))
	operatorType := strings.TrimSpace(c.GetString("operatorType"))
	startTime := c.GetString("startTime")
	endTime := c.GetString("endTime")

	switch operatorType {
	case "plus-amount":
		operatorType = common.PLUS_AMOUNT
	case "sub-amount":
		operatorType = common.SUB_AMOUNT
	case "freeze-amount":
		operatorType = common.FREEZE_AMOUNT
	case "unfreeze-amount":
		operatorType = common.UNFREEZE_AMOUNT
	}
	params := make(map[string]string)
	params["account_name__icontains"] = accountName
	params["account_uid__icontains"] = accountUid
	params["type"] = operatorType
	params["create_time__gte"] = startTime
	params["create_time__lte"] = endTime

	l := models.GetAccountHistoryLenByMap(params)
	c.GetCutPage(l)

	accountHistoryDataJSON := new(AccountHistoryDataJSON)
	accountHistoryDataJSON.DisplayCount = c.DisplayCount
	accountHistoryDataJSON.Code = 200
	accountHistoryDataJSON.CurrentPage = c.CurrentPage
	accountHistoryDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		accountHistoryDataJSON.AccountHistoryList = make([]models.AccountHistoryInfo, 0)
		c.GenerateJSON(accountHistoryDataJSON)
		return
	}

	accountHistoryDataJSON.StartIndex = c.Offset
	accountHistoryDataJSON.AccountHistoryList = models.GetAccountHistoryByMap(params, c.DisplayCount, c.Offset)
	c.GenerateJSON(accountHistoryDataJSON)
}

func (c *GetController) GetAgent() {
	agentName := strings.TrimSpace(c.GetString("agentName"))
	params := make(map[string]string)
	params["agnet_name__icontains"] = agentName

	l := models.GetAgentInfoLenByMap(params)
	c.GetCutPage(l)

	agentDataJSON := new(AgentDataJSON)
	agentDataJSON.DisplayCount = c.DisplayCount
	agentDataJSON.Code = 200
	agentDataJSON.CurrentPage = c.CurrentPage
	agentDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		agentDataJSON.AgentList = make([]models.AgentInfo, 0)
		c.GenerateJSON(agentDataJSON)
		return
	}

	agentDataJSON.StartIndex = c.Offset
	agentDataJSON.AgentList = models.GetAgentInfoByMap(params, c.DisplayCount, c.Offset)
	c.GenerateJSON(agentDataJSON)
}

func (c *GetController) GetAllAgent() {
	agentName := strings.TrimSpace(c.GetString("agentName"))
	params := make(map[string]string)
	params["agent_name__icontains"] = agentName

	agentDataJSON := new(AgentDataJSON)
	agentDataJSON.Code = 200
	agentDataJSON.AgentList = models.GetAllAgentByMap(params)

	c.GenerateJSON(agentDataJSON)
}

func (c *GetController) GetProduct() {
	supplierCode2Name := common.GetSupplierMap()
	productDataJSON := new(ProductDataJSON)
	productDataJSON.Code = 200
	productDataJSON.ProductMap = supplierCode2Name
	c.GenerateJSON(productDataJSON)
}

func (c *GetController) GetAgentToMerchant() {
	agentUid := strings.TrimSpace(c.GetString("agentUid"))
	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))

	params := make(map[string]string)
	params["belong_agent_uid"] = agentUid
	params["merchant_uid"] = merchantUid

	l := models.GetMerchantLenByParams(params)
	c.GetCutPage(l)

	merchantDataJSON := new(MerchantDataJSON)
	merchantDataJSON.DisplayCount = c.DisplayCount
	merchantDataJSON.Code = 200
	merchantDataJSON.CurrentPage = c.CurrentPage
	merchantDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		merchantDataJSON.MerchantList = make([]models.MerchantInfo, 0)
	} else {
		merchantDataJSON.MerchantList = models.GetMerchantByParams(params, c.DisplayCount, c.Offset)
	}

	c.GenerateJSON(merchantDataJSON)
}

/*
* 获取订单数据
 */
func (c *GetController) GetOrder() {
	startTime := strings.TrimSpace(c.GetString("startTime"))
	endTime := strings.TrimSpace(c.GetString("endTime"))
	merchantName := strings.TrimSpace(c.GetString("merchantName"))
	orderNo := strings.TrimSpace(c.GetString("merchantOrderId"))
	//bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))
	status := strings.TrimSpace(c.GetString("orderStatus"))
	supplierUid := strings.TrimSpace(c.GetString("supplierUid"))
	payWayCode := strings.TrimSpace(c.GetString("payWayCode"))
	freeStatus := strings.TrimSpace(c.GetString("freeStatus"))

	params := make(map[string]string)
	params["create_time__gte"] = startTime
	params["create_time__lte"] = endTime
	params["merchant_name__icontains"] = merchantName
	params["merchant_order_id"] = orderNo
	//params["bank_order_id"] = bankOrderId
	params["status"] = status
	params["pay_product_code"] = supplierUid
	params["pay_type_code"] = payWayCode
	switch freeStatus {
	case "free":
		params["free"] = "yes"
	case "unfree":
		params["unfree"] = "yes"
	case "refund":
		params["refund"] = "yes"
	}

	l := models.GetOrderLenByMap(params)
	c.GetCutPage(l)

	orderDataJSON := new(OrderDataJSON)
	orderDataJSON.DisplayCount = c.DisplayCount
	orderDataJSON.Code = 200
	orderDataJSON.CurrentPage = c.CurrentPage
	orderDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		orderDataJSON.OrderList = make([]models.OrderInfo, 0)
		c.GenerateJSON(orderDataJSON)
		return
	}

	orderDataJSON.StartIndex = c.Offset
	orderDataJSON.OrderList = models.GetOrderByMap(params, c.DisplayCount, c.Offset)
	orderDataJSON.SuccessRate = models.GetSuccessRateByMap(params)
	orderDataJSON.AllAmount = models.GetAllAmountByMap(params)
	c.GenerateJSON(orderDataJSON)
}

func (c *GetController) GetOneOrder() {
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))
	orderDataJSON := new(OrderDataJSON)
	orderInfo := models.GetOneOrder(bankOrderId)

	orderDataJSON.Code = 200
	orderDataJSON.OrderList = append(orderDataJSON.OrderList, orderInfo)
	notifyInfo := models.GetNotifyInfoByBankOrderId(bankOrderId)
	if notifyInfo.Url == "" || len(notifyInfo.Url) == 0 {
		orderDataJSON.NotifyUrl = orderInfo.NotifyUrl
	} else {
		orderDataJSON.NotifyUrl = notifyInfo.Url
	}
	c.GenerateJSON(orderDataJSON)
}

func (c *GetController) GetOrderProfit() {
	startTime := strings.TrimSpace(c.GetString("startTime"))
	endTime := strings.TrimSpace(c.GetString("endTime"))
	merchantName := strings.TrimSpace(c.GetString("merchantName"))
	agentName := strings.TrimSpace(c.GetString("agentName"))
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))
	status := strings.TrimSpace(c.GetString("orderStatus"))
	supplierUid := strings.TrimSpace(c.GetString("supplierUid"))
	payWayCode := strings.TrimSpace(c.GetString("payWayCode"))

	params := make(map[string]string)
	params["create_time__gte"] = startTime
	params["create_time__lte"] = endTime
	params["merchant_name__icontains"] = merchantName
	params["agent_name__icontains"] = agentName
	params["bank_order_id"] = bankOrderId
	params["status"] = status
	params["pay_product_code"] = supplierUid
	params["pay_type_code"] = payWayCode

	l := models.GetOrderProfitLenByMap(params)
	c.GetCutPage(l)

	listDataJSON := new(ListDataJSON)
	listDataJSON.DisplayCount = c.DisplayCount
	listDataJSON.Code = 200
	listDataJSON.CurrentPage = c.CurrentPage
	listDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		listDataJSON.List = make([]models.OrderProfitInfo, 0)
		c.GenerateJSON(listDataJSON)
		return
	}

	listDataJSON.StartIndex = c.Offset
	listDataJSON.List = models.GetOrderProfitByMap(params, c.DisplayCount, c.Offset)
	supplierAll := 0.0
	platformAll := 0.0
	agentAll := 0.0
	allAmount := 0.0
	for _, v := range listDataJSON.List {
		if v.Status != "success" {
			continue
		}
		allAmount += v.FactAmount
		supplierAll += v.SupplierProfit
		platformAll += v.PlatformProfit
		agentAll += v.AgentProfit
	}

	listDataJSON.SupplierProfit, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", supplierAll), 3)
	listDataJSON.PlatformProfit, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", platformAll), 3)
	listDataJSON.AgentProfit, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", agentAll), 3)
	listDataJSON.AllAmount, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", allAmount), 3)
	c.GenerateJSON(listDataJSON)
}

func (c *GetController) GetPayFor() {
	startTime := strings.TrimSpace(c.GetString("startTime"))
	endTime := strings.TrimSpace(c.GetString("endTime"))
	merchantOrderId := strings.TrimSpace(c.GetString("merchantOrderId"))
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))
	status := strings.TrimSpace(c.GetString("status"))

	params := make(map[string]string)
	params["create_time__lte"] = endTime
	params["create_time__gte"] = startTime
	params["merchant_order_id"] = merchantOrderId
	params["bank_order_id"] = bankOrderId
	params["status"] = status

	l := models.GetPayForLenByMap(params)
	c.GetCutPage(l)

	listDataJSON := new(PayForDataJSON)
	listDataJSON.DisplayCount = c.DisplayCount
	listDataJSON.Code = 200
	listDataJSON.CurrentPage = c.CurrentPage
	listDataJSON.TotalPage = c.TotalPage

	if c.Offset < 0 {
		listDataJSON.PayForList = make([]models.PayforInfo, 0)
		c.GenerateJSON(listDataJSON)
		return
	}

	listDataJSON.StartIndex = c.Offset
	listDataJSON.PayForList = models.GetPayForByMap(params, c.DisplayCount, c.Offset)
	for index, p := range listDataJSON.PayForList {
		if p.MerchantName == "" {
			listDataJSON.PayForList[index].MerchantName = "任意下发"
		}
		if p.MerchantOrderId == "" {
			listDataJSON.PayForList[index].MerchantOrderId = "任意发下"
		}
		if p.RoadName == "" {
			listDataJSON.PayForList[index].RoadName = "无"
		}
	}
	c.GenerateJSON(listDataJSON)
}

func (c *GetController) GetOnePayFor() {
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))

	payForInfo := models.GetPayForByBankOrderId(bankOrderId)

	listDataJSON := new(PayForDataJSON)
	listDataJSON.Code = 200
	listDataJSON.PayForList = append(listDataJSON.PayForList, payForInfo)

	c.GenerateJSON(listDataJSON)
}

func (c *GetController) GetBalance() {
	roadName := strings.TrimSpace(c.GetString("roadName"))
	roadUid := strings.TrimSpace(c.GetString("roadUid"))

	var roadInfo models.RoadInfo
	if roadUid != "" {
		roadInfo = models.GetRoadInfoByRoadUid(roadUid)
	} else {
		roadInfo = models.GetRoadInfoByName(roadName)
	}

	balanceDataJSON := new(BalanceDataJSON)
	balanceDataJSON.Code = 200

	supplier := controller.GetPaySupplierByCode(roadInfo.ProductUid)
	if supplier == nil {
		balanceDataJSON.Code = -1
		balanceDataJSON.Balance = -1.00
	} else {
		balance := supplier.BalanceQuery(roadInfo)
		balanceDataJSON.Balance = balance
	}

	c.GenerateJSON(balanceDataJSON)
}

func (c *GetController) GetNotifyBankOrderIdList() {
	startTime := strings.TrimSpace(c.GetString("startTime"))
	endTime := strings.TrimSpace(c.GetString("endTime"))
	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))
	notifyType := strings.TrimSpace(c.GetString("notifyType"))

	params := make(map[string]string)
	params["create_time__gte"] = startTime
	params["create_time_lte"] = endTime
	params["merchant_uid"] = merchantUid
	params["type"] = notifyType

	bankOrderIdListJSON := new(NotifyBankOrderIdListJSON)
	bankOrderIdListJSON.Code = 200
	bankOrderIdListJSON.NotifyIdList = models.GetNotifyBankOrderIdListByParams(params)
	c.GenerateJSON(bankOrderIdListJSON)
}

/*
* 获取利润表
 */
func (c *GetController) GetProfit() {
	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))
	agentUid := strings.TrimSpace(c.GetString("agentUid"))
	supplierUid := strings.TrimSpace(c.GetString("supplierUid"))
	payType := strings.TrimSpace(c.GetString("payType"))
	startTime := strings.TrimSpace(c.GetString("startTime"))
	endTime := strings.TrimSpace(c.GetString("endTime"))

	params := make(map[string]string)
	params["merchant_uid"] = merchantUid
	params["agent_uid"] = agentUid
	params["pay_product_code"] = supplierUid
	params["pay_type_code"] = payType
	params["create_time__gte"] = startTime
	params["create_time__lte"] = endTime

	profitListJSON := new(ProfitListJSON)
	profitListJSON.Code = 200
	profitListJSON.ProfitList = models.GetPlatformProfitByMap(params)

	profitListJSON.TotalAmount = 0.00
	profitListJSON.PlatformTotalProfit = 0.00
	profitListJSON.AgentTotalProfit = 0.00

	for i, p := range profitListJSON.ProfitList {
		profitListJSON.TotalAmount += p.OrderAmount
		profitListJSON.PlatformTotalProfit += p.PlatformProfit
		profitListJSON.AgentTotalProfit += p.AgentProfit
		if p.AgentName == "" {
			p.AgentName = "无代理商"
		}
		profitListJSON.ProfitList[i] = p
	}

	c.GenerateJSON(profitListJSON)
}
