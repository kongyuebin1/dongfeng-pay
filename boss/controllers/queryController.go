/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/6 14:03
 ** @Author : yuebin
 ** @File : query.go
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/6 14:03
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"boss/common"
	"boss/models"
	controller "boss/supplier"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
)

type SupplierQuery struct {
	beego.Controller
}

func OrderQuery(bankOrderId string) string {

	orderInfo := models.GetOrderByBankOrderId(bankOrderId)

	if orderInfo.BankOrderId == "" || len(orderInfo.BankOrderId) == 0 {
		logs.Error("不存在这样的订单，订单查询结束")
		return "不存在这样的订单"
	}

	if orderInfo.Status != "" && orderInfo.Status != "wait" {
		logs.Error(fmt.Sprintf("该订单=%s，已经处理完毕，", bankOrderId))
		return "该订单已经处理完毕"
	}

	supplierCode := orderInfo.PayProductCode
	supplier := controller.GetPaySupplierByCode(supplierCode)

	flag := supplier.PayQuery(orderInfo)
	if flag {
		return "查询完毕，返回正确结果"
	} else {
		return "订单还在处理中"
	}

}

func (c *SupplierQuery) SupplierOrderQuery() {

	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))
	exist := models.BankOrderIdIsEixst(bankOrderId)

	keyDataJSON := new(KeyDataJSON)
	if !exist {
		keyDataJSON.Msg = "该订单不存在"
	}

	msg := OrderQuery(bankOrderId)

	keyDataJSON.Msg = msg
	c.Data["json"] = keyDataJSON
	c.ServeJSON()
}

/*
* 向上游查询代付结果
 */
func (c *SupplierQuery) SupplierPayForQuery() {
	bankOrderId := strings.TrimSpace(c.GetString("bankOrderId"))

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = 200

	if bankOrderId == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "不存在这样的代付订单"
	} else {
		payFor := models.GetPayForByBankOrderId(bankOrderId)
		if payFor.RoadUid == "" {
			keyDataJSON.Msg = "该代付订单没有对应的通道uid"
		} else {
			roadInfo := models.GetRoadInfoByRoadUid(payFor.RoadUid)
			supplier := controller.GetPaySupplierByCode(roadInfo.ProductUid)
			result, msg := supplier.PayForQuery(payFor)
			keyDataJSON.Msg = msg
			if result == common.PAYFOR_SUCCESS {
				controller.PayForSuccess(payFor)
			} else if result == common.PAYFOR_FAIL {
				controller.PayForFail(payFor)
			} else {
				logs.Info("银行处理中")
			}
		}
	}

	c.Data["json"] = keyDataJSON
	c.ServeJSON()
}
