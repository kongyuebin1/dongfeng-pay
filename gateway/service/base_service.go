package service

import (
	"context"
	"gateway/conf"
	"gateway/models/merchant"
	"gateway/models/order"
	"gateway/response"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"strconv"
	"strings"
)

//获取商户信息
func GetMerchantInfo(params map[string]string) *response.PayBaseResp {

	c := new(response.PayBaseResp)
	c.Params = make(map[string]string)
	c.Params = params

	merchantInfo := merchant.GetMerchantByPaykey(params["payKey"])

	if merchantInfo.MerchantUid == "" || len(merchantInfo.MerchantUid) == 0 {
		c.Code = -1
		c.Msg = "商户不存在，或者paykey有误，请联系管理员"
	} else if merchantInfo.Status != conf.ACTIVE {
		c.Code = -1
		c.Msg = "商户状态已经被冻结或者被删除，请联系管理员！"
	} else {
		c.MerchantInfo = merchantInfo
	}

	return c
}

func JudgeParams(c *response.PayBaseResp) *response.PayBaseResp {
	//c.ReturnUrlIsValid()
	c = OrderIsValid(c)
	c = NotifyUrlIsValid(c)
	c = OsTypeIsValid(c)
	c = PayWayCodeIsValid(c)
	c = ProductIsValid(c)
	c = OrderPeriodIsValid(c)
	//c = IpIsWhite()
	c = OrderPriceIsValid(c)

	return c
}

/*
* 插入支付订单记录和订单利润记录，保证一致性
 */
func InsertOrderAndOrderProfit(orderInfo order.OrderInfo, orderProfitInfo order.OrderProfitInfo) bool {
	o := orm.NewOrm()
	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		if _, err := txOrm.Insert(&orderInfo); err != nil {
			logs.Error("insert orderInfo fail: ", err)
			return err
		}
		if _, err := txOrm.Insert(&orderProfitInfo); err != nil {
			logs.Error("insert orderProfit fail: ", err)
			return err
		}

		return nil

	}); err != nil {
		return false
	}
	return true
}

/**
** 判断跳转地址是否符合规则
 */
func ReturnUrlIsValid(c *response.PayBaseResp) *response.PayBaseResp {
	if c.Params["returnUrl"] == "" || len(c.Params["returnUrl"]) == 0 {
		c.Code = -1
		c.Msg = "支付成功后跳转地址不能为空"
	}
	return c
}

/**
** 判断回调地址是否符合规则
 */
func NotifyUrlIsValid(c *response.PayBaseResp) *response.PayBaseResp {
	if c.Params["notifyUrl"] == "" || len(c.Params["notifyUrl"]) == 0 {
		c.Code = -1
		c.Msg = "支付成功订单回调地址不能空位"
	}

	return c
}

/**
** 判断设备类型是否符合规则
 */
func OsTypeIsValid(c *response.PayBaseResp) *response.PayBaseResp {
	if c.Params["osType"] == "" || len(c.Params["osType"]) == 0 {
		c.Code = -1
		c.Msg = "支付设备系统类型不能为空，默认填写\"1\"即可"
	}

	return c
}

/**
** 判断支付类型字段是否符合规则
 */
func PayWayCodeIsValid(c *response.PayBaseResp) *response.PayBaseResp {
	if c.Params["payWayCode"] == "" || len(c.Params["payWayCode"]) == 0 {
		c.Code = -1
		c.Msg = "支付类型字段不能为空"
		return c
	}

	if !strings.Contains(c.Params["payWayCode"], "SCAN") {
		c.Code = -1
		c.Msg = "扫码支付不支持这种支付类型"
	} else {
		scanPayWayCodes := conf.GetScanPayWayCodes()
		for _, v := range scanPayWayCodes {
			if c.Params["payWayCode"] == v {
				c.PayWayCode = strings.Replace(c.Params["payWayCode"], "-", "_", -1)
				return c
			}
		}
		c.Code = -1
		c.Msg = "不存在这种支付类型，请仔细阅读对接文档"
	}

	return c
}

func ProductIsValid(c *response.PayBaseResp) *response.PayBaseResp {
	if c.Params["productName"] == "" || len(c.Params["productName"]) == 0 {
		c.Code = -1
		c.Msg = "商品描述信息字段不能为空"
	}

	return c
}

func OrderPeriodIsValid(c *response.PayBaseResp) *response.PayBaseResp {
	if c.Params["orderPeriod"] == "" || len(c.Params["orderPeriod"]) == 0 {
		c.Code = -1
		c.Msg = "订单过期时间不能为空，默认填写\"1\"即可"
	}

	return c
}

//判断订单金额
func OrderPriceIsValid(c *response.PayBaseResp) *response.PayBaseResp {
	if c.Params["orderPrice"] == "" || len(c.Params["orderPrice"]) == 0 {
		c.Code = -1
		c.Msg = "订单金额不能为空"
		return c
	}

	a, err := strconv.ParseFloat(c.Params["orderPrice"], 64)
	if err != nil {
		logs.Error("order price is invalid： ", c.Params["orderPrice"])
		c.Code = -1
		c.Msg = "订单金额非法"
	}
	c.OrderAmount = a

	return c
}

//判断金额订单号是否为空或者有重复
func OrderIsValid(c *response.PayBaseResp) *response.PayBaseResp {
	if c.Params["orderNo"] == "" || len(c.Params["orderNo"]) == 0 {
		c.Code = -1
		c.Msg = "商户订单号不能为空"
		return c
	}
	if order.OrderNoIsEixst(c.Params["orderNo"]) {
		c.Code = -1
		c.Msg = "商户订单号重复"
	}

	return c
}

//判断ip是否在白名单中
func IpIsWhite() bool {
	//TODO
	return true
}
