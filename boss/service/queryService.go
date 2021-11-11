package service

import (
	"boss/datas"
	"boss/models/order"
	"boss/models/payfor"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type QueryService struct {
}

func OrderQuery(bankOrderId string) string {

	orderInfo := order.GetOrderByBankOrderId(bankOrderId)

	if orderInfo.BankOrderId == "" || len(orderInfo.BankOrderId) == 0 {
		logs.Error("不存在这样的订单，订单查询结束")
		return "不存在这样的订单"
	}

	if orderInfo.Status != "" && orderInfo.Status != "wait" {
		logs.Error(fmt.Sprintf("该订单=%s，已经处理完毕，", bankOrderId))
		return "该订单已经处理完毕"
	}

	// 向gateway发送请求，请求上游的支付结果
	gUrl, _ := web.AppConfig.String("gateway::host")
	gUrl = gUrl + "supplier/order/query" + "?" + "bankOrderId=" + bankOrderId
	res, err := httplib.Get(gUrl).String()
	if err != nil {
		logs.Error("获取gateway上游订单查询结果失败：" + err.Error())
	}

	if res == "success" {
		return res
	}

	return "fail"

}

func (c *QueryService) SupplierOrderQuery(bankOrderId string) *datas.KeyDataJSON {

	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = 200
	exist := order.BankOrderIdIsEixst(bankOrderId)
	if !exist {
		keyDataJSON.Msg = "该订单不存在"
		keyDataJSON.Code = -1
	}

	msg := OrderQuery(bankOrderId)

	keyDataJSON.Msg = msg
	return keyDataJSON
}

func (c *QueryService) SupplierPayForQuery(bankOrderId string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = 200

	if bankOrderId == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "不存在这样的代付订单"
	} else {
		payFor := payfor.GetPayForByBankOrderId(bankOrderId)
		if payFor.RoadUid == "" {
			keyDataJSON.Msg = "该代付订单没有对应的通道uid"
		} else {
			result := querySupplierPayForResult(bankOrderId)
			if result {
				keyDataJSON.Msg = "处理成功！"
			} else {
				keyDataJSON.Msg = "处理失败！"
			}
		}
	}
	return keyDataJSON
}

func querySupplierPayForResult(bankOrderId string) bool {
	payforUrl, _ := web.AppConfig.String("gateway::host")
	u := payforUrl + "gateway/supplier/payfor/query" + "?bankOrderId=" + bankOrderId
	s, err := httplib.Get(u).String()
	if err != nil {
		logs.Error("处理代付查询请求gateway失败：", err)
		return false
	}
	if s == "fail" {
		return false
	} else {
		return true
	}
}
