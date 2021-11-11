/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/30 11:44
 ** @Author : yuebin
 ** @File : order_profit_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/30 11:44
 ** @Software: GoLand
****************************************************/
package order

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"strings"
)

type OrderProfitInfo struct {
	Id              int
	MerchantName    string
	MerchantUid     string
	AgentName       string
	AgentUid        string
	PayProductCode  string
	PayProductName  string
	PayTypeCode     string
	PayTypeName     string
	Status          string
	MerchantOrderId string
	BankOrderId     string
	BankTransId     string
	OrderAmount     float64
	ShowAmount      float64
	FactAmount      float64
	UserInAmount    float64
	SupplierRate    float64
	PlatformRate    float64
	AgentRate       float64
	AllProfit       float64
	SupplierProfit  float64
	PlatformProfit  float64
	AgentProfit     float64
	UpdateTime      string
	CreateTime      string
}

const ORDER_PROFIT_INFO = "order_profit_info"

func GetOrderProfitByBankOrderId(bankOrderId string) OrderProfitInfo {
	o := orm.NewOrm()
	var orderProfit OrderProfitInfo
	_, err := o.QueryTable(ORDER_PROFIT_INFO).Filter("bank_order_id", bankOrderId).Limit(1).All(&orderProfit)
	if err != nil {
		logs.Error("GetOrderProfitByBankOrderId fail：", err)
	}
	return orderProfit
}

func GetOrderProfitLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(ORDER_PROFIT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, _ := qs.Limit(-1).Count()
	return int(cnt)
}

func GetOrderProfitByMap(params map[string]string, display, offset int) []OrderProfitInfo {
	o := orm.NewOrm()
	var orderProfitInfoList []OrderProfitInfo
	qs := o.QueryTable(ORDER_PROFIT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(display, offset).OrderBy("-update_time").All(&orderProfitInfoList)
	if err != nil {
		logs.Error("get order by map fail: ", err)
	}
	return orderProfitInfoList
}

func GetPlatformProfitByMap(params map[string]string) []PlatformProfit {

	o := orm.NewOrm()

	cond := "select merchant_name, agent_name, pay_product_name as supplier_name, pay_type_name, sum(fact_amount) as order_amount, count(1) as order_count, " +
		"sum(platform_profit) as platform_profit, sum(agent_profit) as agent_profit from " + ORDER_PROFIT_INFO + " where status='success' "
	flag := false
	for k, v := range params {
		if len(v) > 0 {
			if flag {
				cond += " and"
			}
			if strings.Contains(k, "create_time__gte") {
				cond = cond + " create_time>='" + v + "'"
			} else if strings.Contains(k, "create_time__lte") {
				cond = cond + " create_time<='" + v + "'"
			} else {
				cond = cond + " " + k + "='" + v + "'"
			}
			flag = true
		}
	}

	cond += " group by merchant_uid, agent_uid, pay_product_code, pay_type_code"

	var platformProfitList []PlatformProfit
	_, err := o.Raw(cond).QueryRows(&platformProfitList)
	if err != nil {
		logs.Error("get platform profit by map fail:", err)
	}

	return platformProfitList
}
