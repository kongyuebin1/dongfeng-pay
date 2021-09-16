/***************************************************
 ** @Desc : 处理网关模块的一些需要操作数据库的功能
 ** @Time : 2019/12/7 16:40
 ** @Author : yuebin
 ** @File : gateway_solve
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/12/7 16:40
 ** @Software: GoLand
****************************************************/
package service

import (
	"fmt"
	"gateway/conf"
	"gateway/models/merchant"
	"gateway/models/order"
	"gateway/models/road"
	"gateway/response"
	"gateway/supplier"
	"gateway/utils"
	"github.com/beego/beego/v2/core/logs"
	"github.com/rs/xid"
	"strings"
	"time"
)

//选择通道
func ChooseRoad(c *response.PayBaseResp) *response.PayBaseResp {
	payWayCode := c.Params["payWayCode"]
	merchantUid := c.MerchantInfo.MerchantUid
	//通道配置信息
	deployInfo := merchant.GetMerchantDeployByUidAndPayType(merchantUid, payWayCode)
	if deployInfo.MerchantUid == "" {
		c.Code = -1
		c.Msg = "该商户没有配置通道信息"
		return c
	}

	singleRoad := road.GetRoadInfoByRoadUid(deployInfo.SingleRoadUid)
	c.RoadPoolInfo = road.GetRoadPoolByRoadPoolCode(deployInfo.RollRoadCode)
	if RoadIsValid(singleRoad, c) {
		c.RoadInfo = singleRoad
		c.PlatformRate = deployInfo.SingleRoadPlatformRate
		c.AgentRate = deployInfo.SingleRoadAgentRate
		return c
	}
	//如果单通道没有有效的，那么寻找通道池里面的通道
	if c.RoadPoolInfo.RoadPoolCode == "" {
		c.Code = -1
		c.Msg = "该商户没有配置通道"
		return c
	}
	roadUids := strings.Split(c.RoadPoolInfo.RoadUidPool, "||")
	roadInfos := road.GetRoadInfosByRoadUids(roadUids)
	for _, roadInfo := range roadInfos {
		if RoadIsValid(roadInfo, c) {
			c.RoadInfo = roadInfo
			c.PlatformRate = deployInfo.RollRoadPlatformRate
			c.AgentRate = deployInfo.RollRoadAgentRate
			return c
		}
	}
	if c.RoadInfo.RoadUid == "" {
		c.Code = -1
		c.Msg = "该商户没有配置通道或者通道不可用"
	}

	return c
}

//判断通道是否是合法的
func RoadIsValid(roadInfo road.RoadInfo, c *response.PayBaseResp) bool {
	if roadInfo.RoadUid == "" || len(roadInfo.RoadUid) == 0 {
		return false
	}
	FORMAT := fmt.Sprintf("该通道：%s;", roadInfo.RoadName)
	if roadInfo.Status != "active" {
		logs.Notice(FORMAT + "不是激活状态")
		return false
	}
	hour := time.Now().Hour()
	s := roadInfo.StarHour
	e := roadInfo.EndHour
	if hour < s || hour > e {
		logs.Notice(FORMAT)
		return false
	}
	minAmount := roadInfo.SingleMinLimit
	maxAmount := roadInfo.SingleMaxLimit
	if minAmount > c.OrderAmount || maxAmount < c.OrderAmount {
		logs.Error(FORMAT + "订单金额超限制")
		return false
	}
	todayLimit := roadInfo.TodayLimit
	totalLimit := roadInfo.TotalLimit
	todayIncome := roadInfo.TodayIncome
	totalIncome := roadInfo.TotalIncome
	if (todayIncome + c.OrderAmount) > todayLimit {
		logs.Error(FORMAT + "达到了每天金额上限")
		return false
	}
	if (totalIncome + c.OrderAmount) > totalLimit {
		logs.Error(FORMAT + "达到了总量限制")
		return false
	}
	//如果通道被选中，那么总请求数+1
	roadInfo.RequestAll = roadInfo.RequestAll + 1
	roadInfo.UpdateTime = utils.GetBasicDateTime()
	road.UpdateRoadInfo(roadInfo)
	return true
}

//获取基本订单记录
func GenerateOrderInfo(c *response.PayBaseResp) order.OrderInfo {
	//6666是自己系统订单号
	bankOrderNo := "6666" + xid.New().String()
	//获取支付类型的名称，例如支付宝扫码等
	payTypeName := conf.GetNameByPayWayCode(c.Params["payWayCode"])
	orderInfo := order.OrderInfo{
		MerchantUid:     c.MerchantInfo.MerchantUid,
		MerchantName:    c.MerchantInfo.MerchantName,
		MerchantOrderId: c.Params["orderNo"],
		BankOrderId:     bankOrderNo,
		OrderAmount:     c.OrderAmount,
		FactAmount:      c.OrderAmount,
		ShowAmount:      c.OrderAmount,
		RollPoolCode:    c.RoadPoolInfo.RoadPoolCode,
		RollPoolName:    c.RoadPoolInfo.RoadPoolName,
		RoadUid:         c.RoadInfo.RoadUid,
		RoadName:        c.RoadInfo.RoadName,
		PayProductName:  c.RoadInfo.ProductName,
		ShopName:        c.Params["productName"],
		Freeze:          conf.NO,
		Refund:          conf.NO,
		Unfreeze:        conf.NO,
		PayProductCode:  c.RoadInfo.ProductUid,
		PayTypeCode:     c.PayWayCode,
		PayTypeName:     payTypeName,
		OsType:          c.Params["osType"],
		Status:          conf.WAIT,
		NotifyUrl:       c.Params["notifyUrl"],
		ReturnUrl:       c.Params["returnUrl"],
		OrderPeriod:     c.Params["orderPeriod"],
		UpdateTime:      utils.GetBasicDateTime(),
		CreateTime:      utils.GetBasicDateTime(),
	}

	if c.MerchantInfo.BelongAgentUid != "" || c.AgentRate > conf.ZERO {
		orderInfo.AgentUid = c.MerchantInfo.BelongAgentUid
		orderInfo.AgentName = c.MerchantInfo.BelongAgentName
	}
	return orderInfo
}

//计算收益，平台利润，代理利润
func GenerateOrderProfit(orderInfo order.OrderInfo, c *response.PayBaseResp) order.OrderProfitInfo {
	//因为所有的手续费率都是百分率，所以需要除以100
	payTypeName := conf.GetNameByPayWayCode(c.PayWayCode)
	supplierProfit := c.OrderAmount / 100 * c.RoadInfo.BasicFee
	platformProfit := c.OrderAmount / 100 * c.PlatformRate
	agentProfit := c.OrderAmount / 100 * c.AgentRate
	//如果用户没有设置代理，那么代理利润为0.000
	if c.MerchantInfo.BelongAgentUid == "" || len(c.MerchantInfo.BelongAgentUid) == 0 {
		agentProfit = conf.ZERO
	}
	allProfit := supplierProfit + platformProfit + agentProfit

	if allProfit >= c.OrderAmount {
		logs.Error("手续费已经超过订单金额，bankOrderId = %s", orderInfo.BankOrderId)
		c.Msg = "手续费已经超过了订单金额"
		c.Code = -1
	}

	orderProfit := order.OrderProfitInfo{
		PayProductCode:  c.RoadInfo.ProductUid,
		PayProductName:  c.RoadInfo.ProductName,
		PayTypeCode:     c.PayWayCode,
		PayTypeName:     payTypeName,
		Status:          conf.WAIT,
		MerchantOrderId: c.Params["orderNo"],
		BankOrderId:     orderInfo.BankOrderId,
		OrderAmount:     c.OrderAmount,
		FactAmount:      c.OrderAmount,
		ShowAmount:      c.OrderAmount,
		AllProfit:       allProfit,
		UserInAmount:    c.OrderAmount - allProfit,
		SupplierProfit:  supplierProfit,
		PlatformProfit:  platformProfit,
		AgentProfit:     agentProfit,
		UpdateTime:      utils.GetBasicDateTime(),
		CreateTime:      utils.GetBasicDateTime(),
		MerchantUid:     c.MerchantInfo.MerchantUid,
		MerchantName:    orderInfo.MerchantName,
		SupplierRate:    c.RoadInfo.BasicFee,
		PlatformRate:    c.PlatformRate,
		AgentRate:       c.AgentRate,
		AgentName:       orderInfo.AgentName,
		AgentUid:        orderInfo.AgentUid,
	}

	//如果该条订单设置了代理利率，并且设置了代理
	if c.MerchantInfo.BelongAgentUid != "" || c.AgentRate > conf.ZERO {
		orderProfit.AgentUid = c.MerchantInfo.BelongAgentUid
		orderProfit.AgentName = c.MerchantInfo.BelongAgentName
	}
	return orderProfit
}

/*
* 生成订单一系列的记录
 */
func GenerateRecord(c *response.PayBaseResp) (order.OrderInfo, order.OrderProfitInfo) {
	//生成订单记录，订单利润利润
	orderInfo := GenerateOrderInfo(c)
	orderProfit := GenerateOrderProfit(orderInfo, c)
	if c.Code == -1 {
		return orderInfo, orderProfit
	}
	if !InsertOrderAndOrderProfit(orderInfo, orderProfit) {
		c.Code = -1
		return orderInfo, orderProfit
	}
	logs.Info("插入支付订单记录和支付利润记录成功")
	return orderInfo, orderProfit
}

func GenerateSuccessData(scanData supplier.ScanData, c *response.PayBaseResp) *response.ScanSuccessData {
	params := make(map[string]string)
	params["orderNo"] = scanData.OrderNo
	params["orderPrice"] = scanData.OrderPrice
	params["payKey"] = c.MerchantInfo.MerchantKey
	params["payURL"] = scanData.PayUrl
	params["statusCode"] = "00"

	keys := utils.SortMap(params)
	sign := utils.GetMD5Sign(params, keys, c.MerchantInfo.MerchantSecret)
	scanSuccessData := new(response.ScanSuccessData)

	scanSuccessData.StatusCode = "00"
	scanSuccessData.PayKey = c.MerchantInfo.MerchantKey
	scanSuccessData.OrderNo = scanData.OrderNo
	scanSuccessData.OrderPrice = scanData.OrderPrice
	scanSuccessData.PayUrl = scanData.PayUrl
	scanSuccessData.PayKey = c.MerchantInfo.MerchantKey
	scanSuccessData.Msg = "请求成功"
	scanSuccessData.Sign = sign

	return scanSuccessData
}
