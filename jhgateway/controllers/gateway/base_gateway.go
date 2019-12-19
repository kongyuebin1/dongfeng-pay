/***************************************************
 ** @Desc : 处理下游请求的一些公用的逻辑
 ** @Time : 2019/10/28 18:09
 ** @Author : yuebin
 ** @File : base_gateway
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/28 18:09
 ** @Software: GoLand
****************************************************/
package gateway

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/rs/xid"
	"dongfeng-pay/service/common"
	"dongfeng-pay/service/controller"
	"dongfeng-pay/service/models"
	"dongfeng-pay/service/utils"
	"strconv"
	"strings"
	"time"
)

type BaseGateway struct {
	beego.Controller
	Params       map[string]string   //请求的基本参数
	ClientIp     string              //商户ip
	MerchantInfo models.MerchantInfo //商户信息
	Msg          string              //信息
	Code         int                 //状态码 200正常
	RoadInfo     models.RoadInfo
	RoadPoolInfo models.RoadPoolInfo
	OrderAmount  float64
	PayWayCode   string
	PlatformRate float64
	AgentRate    float64
}

//获取商户请求过来的基本参数参数
func (c *BaseGateway) PayPrepare() {
	c.Params = make(map[string]string)
	//获取客户端的ip
	c.ClientIp = c.Ctx.Input.IP()
	c.Params["orderNo"] = strings.TrimSpace(c.GetString("orderNo"))
	c.Params["productName"] = strings.TrimSpace(c.GetString("productName"))
	c.Params["orderPeriod"] = strings.TrimSpace(c.GetString("orderPeriod"))
	c.Params["orderPrice"] = strings.TrimSpace(c.GetString("orderPrice"))
	c.Params["payWayCode"] = strings.TrimSpace(c.GetString("payWayCode"))
	c.Params["osType"] = strings.TrimSpace(c.GetString("osType"))
	c.Params["notifyUrl"] = strings.TrimSpace(c.GetString("notifyUrl"))
	//c.Params["returnUrl"] = strings.TrimSpace(c.GetString("returnUrl"))
	c.Params["payKey"] = strings.TrimSpace(c.GetString("payKey"))
	c.Params["sign"] = strings.TrimSpace(c.GetString("sign"))

	c.GetMerchantInfo()
	c.JudgeParams()

	if c.Code != -1 {
		c.Code = 200
	}
}

//判断参数的
func (c *BaseGateway) JudgeParams() {
	//c.ReturnUrlIsValid()
	c.OrderIsValid()
	c.NotifyUrlIsValid()
	c.OsTypeIsValid()
	c.PayWayCodeIsValid()
	c.ProductIsValid()
	c.OrderPeriodIsValid()
	c.IpIsWhite()
	c.OrderPriceIsValid()
}

func (c *BaseGateway) ReturnUrlIsValid() {
	if c.Params["returnUrl"] == "" || len(c.Params["returnUrl"]) == 0 {
		c.Code = -1
		c.Msg = "支付成功后跳转地址不能为空"
	}
}

func (c *BaseGateway) NotifyUrlIsValid() {
	if c.Params["notifyUrl"] == "" || len(c.Params["notifyUrl"]) == 0 {
		c.Code = -1
		c.Msg = "支付成功订单回调地址不能空位"
	}
}

func (c *BaseGateway) OsTypeIsValid() {
	if c.Params["osType"] == "" || len(c.Params["osType"]) == 0 {
		c.Code = -1
		c.Msg = "支付设备系统类型不能为空，默认填写\"1\"即可"
	}
}

func (c *BaseGateway) PayWayCodeIsValid() {
	if c.Params["payWayCode"] == "" || len(c.Params["payWayCode"]) == 0 {
		c.Code = -1
		c.Msg = "支付类型字段不能为空"
		return
	}

	if !strings.Contains(c.Params["payWayCode"], "SCAN") {
		c.Code = -1
		c.Msg = "扫码支付不支持这种支付类型"
	} else {
		scanPayWayCodes := common.GetScanPayWayCodes()
		for _, v := range scanPayWayCodes {
			if c.Params["payWayCode"] == v {
				c.PayWayCode = strings.Replace(c.Params["payWayCode"], "-", "_", -1)
				return
			}
		}
		c.Code = -1
		c.Msg = "不存在这种支付类型，请仔细阅读对接文档"
	}
}

func (c *BaseGateway) ProductIsValid() {
	if c.Params["productName"] == "" || len(c.Params["productName"]) == 0 {
		c.Code = -1
		c.Msg = "商品描述信息字段不能为空"
	}
}

func (c *BaseGateway) OrderPeriodIsValid() {
	if c.Params["orderPeriod"] == "" || len(c.Params["orderPeriod"]) == 0 {
		c.Code = -1
		c.Msg = "订单过期时间不能为空，默认填写\"1\"即可"
	}
}

//获取商户信息
func (c *BaseGateway) GetMerchantInfo() {
	merchantInfo := models.GetMerchantByPaykey(c.Params["payKey"])
	if merchantInfo.MerchantUid == "" || len(merchantInfo.MerchantUid) == 0 {
		c.Code = -1
		c.Msg = "商户不存在，或者paykey有误，请联系管理员"
	} else if merchantInfo.Status != common.ACTIVE {
		c.Code = -1
		c.Msg = "商户状态已经被冻结或者被删除，请联系管理员！"
	} else {
		c.MerchantInfo = merchantInfo
	}
}

//判断订单金额
func (c *BaseGateway) OrderPriceIsValid() {
	if c.Params["orderPrice"] == "" || len(c.Params["orderPrice"]) == 0 {
		c.Code = -1
		c.Msg = "订单金额不能为空"
		return
	}

	a, err := strconv.ParseFloat(c.Params["orderPrice"], 64)
	if err != nil {
		logs.Error("order price is invalid： ", c.Params["orderPrice"])
		c.Code = -1
		c.Msg = "订单金额非法"
	}
	c.OrderAmount = a
}

//判断金额订单号是否为空或者有重复
func (c *BaseGateway) OrderIsValid() {
	if c.Params["orderNo"] == "" || len(c.Params["orderNo"]) == 0 {
		c.Code = -1
		c.Msg = "商户订单号不能为空"
		return
	}
	if models.OrderNoIsEixst(c.Params["orderNo"]) {
		c.Code = -1
		c.Msg = "商户订单号重复"
	}
}

//判断ip是否在白名单中
func (c *BaseGateway) IpIsWhite() bool {
	//TODO
	return true
}

//选择通道
func (c *BaseGateway) ChooseRoad() {
	payWayCode := c.Params["payWayCode"]
	merchantUid := c.MerchantInfo.MerchantUid
	//通道配置信息
	deployInfo := models.GetMerchantDeployByUidAndPayType(merchantUid, payWayCode)
	if deployInfo.MerchantUid == "" {
		c.Code = -1
		c.Msg = "该商户没有配置"
		return
	}

	singleRoad := models.GetRoadInfoByRoadUid(deployInfo.SingleRoadUid)
	c.RoadPoolInfo = models.GetRoadPoolByRoadPoolCode(deployInfo.RollRoadCode)
	if c.RoadIsValid(singleRoad) {
		c.RoadInfo = singleRoad
		c.PlatformRate = deployInfo.SingleRoadPlatformRate
		c.AgentRate = deployInfo.SingleRoadAgentRate
		return
	}
	//如果单通道没有有效的，那么寻找通道池里面的通道
	if c.RoadPoolInfo.RoadPoolCode == "" {
		c.Code = -1
		c.Msg = "该商户没有配置通道"
		return
	}
	roadUids := strings.Split(c.RoadPoolInfo.RoadUidPool, "||")
	roadInfos := models.GetRoadInfosByRoadUids(roadUids)
	for _, roadInfo := range roadInfos {
		if c.RoadIsValid(roadInfo) {
			c.RoadInfo = roadInfo
			c.PlatformRate = deployInfo.RollRoadPlatformRate
			c.AgentRate = deployInfo.RollRoadAgentRate
			return
		}
	}
	if c.RoadInfo.RoadUid == "" {
		c.Code = -1
		c.Msg = "该商户没有配置通道或者通道不可用"
	}
}

//判断通道是否是合法的
func (c *BaseGateway) RoadIsValid(roadInfo models.RoadInfo) bool {
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
	models.UpdateRoadInfo(roadInfo)
	return true
}

//获取基本订单记录
func (c *BaseGateway) GetOrderInfo() models.OrderInfo {
	//6666是自己系统订单号
	bankOrderNo := "6666" + xid.New().String()
	//获取支付类型的名称，例如支付宝扫码等
	payTypeName := common.GetNameByPayWayCode(c.Params["payWayCode"])
	orderInfo := models.OrderInfo{
		MerchantUid: c.MerchantInfo.MerchantUid, MerchantName: c.MerchantInfo.MerchantName, MerchantOrderId: c.Params["orderNo"],
		BankOrderId: bankOrderNo, OrderAmount: c.OrderAmount, FactAmount: c.OrderAmount, ShowAmount: c.OrderAmount,
		RollPoolCode: c.RoadPoolInfo.RoadPoolCode, RollPoolName: c.RoadPoolInfo.RoadPoolName, RoadUid: c.RoadInfo.RoadUid,
		RoadName: c.RoadInfo.RoadName, PayProductName: c.RoadInfo.ProductName, ShopName: c.Params["productName"], Freeze: common.NO,
		Refund: common.NO, Unfreeze: common.NO, PayProductCode: c.RoadInfo.ProductUid, PayTypeCode: c.PayWayCode, PayTypeName: payTypeName,
		OsType: c.Params["osType"], Status: common.WAIT, NotifyUrl: c.Params["notifyUrl"], ReturnUrl: c.Params["returnUrl"],
		OrderPeriod: c.Params["orderPeriod"], UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime(),
	}
	if c.MerchantInfo.BelongAgentUid != "" || c.AgentRate > common.ZERO {
		orderInfo.AgentUid = c.MerchantInfo.BelongAgentUid
		orderInfo.AgentName = c.MerchantInfo.BelongAgentName
	}
	return orderInfo
}

//计算收益，平台利润，代理利润
func (c *BaseGateway) GetOrderProfit(orderInfo models.OrderInfo) models.OrderProfitInfo {
	//因为所有的手续费率都是百分率，所以需要除以100
	payTypeName := common.GetNameByPayWayCode(c.PayWayCode)
	supplierProfit := c.OrderAmount / 100 * c.RoadInfo.BasicFee
	platformProfit := c.OrderAmount / 100 * c.PlatformRate
	agentProfit := c.OrderAmount / 100 * c.AgentRate
	//如果用户没有设置代理，那么代理利润为0.000
	if c.MerchantInfo.BelongAgentUid == "" || len(c.MerchantInfo.BelongAgentUid) == 0 {
		agentProfit = common.ZERO
	}
	allProfit := supplierProfit + platformProfit + agentProfit

	if allProfit >= c.OrderAmount {
		logs.Error("手续费已经超过订单金额，bankOrderId = %s", orderInfo.BankOrderId)
		c.Msg = "手续费已经超过了订单金额"
		c.Code = -1
	}
	orderProfit := models.OrderProfitInfo{
		PayProductCode: c.RoadInfo.ProductUid, PayProductName: c.RoadInfo.ProductName, PayTypeCode: c.PayWayCode, PayTypeName: payTypeName,
		Status: common.WAIT, MerchantOrderId: c.Params["orderNo"], BankOrderId: orderInfo.BankOrderId, OrderAmount: c.OrderAmount,
		FactAmount: c.OrderAmount, ShowAmount: c.OrderAmount, AllProfit: allProfit, UserInAmount: c.OrderAmount - allProfit,
		SupplierProfit: supplierProfit, PlatformProfit: platformProfit, AgentProfit: agentProfit, UpdateTime: utils.GetBasicDateTime(),
		CreateTime: utils.GetBasicDateTime(), MerchantUid: c.MerchantInfo.MerchantUid, MerchantName: orderInfo.MerchantName,
		SupplierRate: c.RoadInfo.BasicFee, PlatformRate: c.PlatformRate, AgentRate: c.AgentRate, AgentName: orderInfo.AgentName, AgentUid: orderInfo.AgentUid,
	}

	//如果该条订单设置了代理利率，并且设置了代理
	if c.MerchantInfo.BelongAgentUid != "" || c.AgentRate > common.ZERO {
		orderProfit.AgentUid = c.MerchantInfo.BelongAgentUid
		orderProfit.AgentName = c.MerchantInfo.BelongAgentName
	}
	return orderProfit
}

/*
* 生成订单一系列的记录
 */
func (c *BaseGateway) GenerateRecord() (models.OrderInfo, models.OrderProfitInfo) {
	//生成订单记录，订单利润利润
	orderInfo := c.GetOrderInfo()
	orderProfit := c.GetOrderProfit(orderInfo)
	if c.Code == -1 {
		return orderInfo, orderProfit
	}
	if !controller.InsertOrderAndOrderProfit(orderInfo, orderProfit) {
		c.Code = -1
		return orderInfo, orderProfit
	}
	logs.Info("插入支付订单记录和支付利润记录成功")
	return orderInfo, orderProfit
}

func (c *BaseGateway) GenerateSuccessData(scanData controller.ScanData) *ScanSuccessData {
	params := make(map[string]string)
	params["orderNo"] = scanData.OrderNo
	params["orderPrice"] = scanData.OrderPrice
	params["payKey"] = c.MerchantInfo.MerchantKey
	params["payURL"] = scanData.PayUrl
	params["statusCode"] = "00"

	keys := utils.SortMap(params)
	sign := utils.GetMD5Sign(params, keys, c.MerchantInfo.MerchantSecret)
	scanSuccessData := new(ScanSuccessData)

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
