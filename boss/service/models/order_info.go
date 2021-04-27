/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/28 10:15
 ** @Author : yuebin
 ** @File : order_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/28 10:15
 ** @Software: GoLand
****************************************************/
package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type OrderInfo struct {
	Id              int
	ShopName        string  //商品名称
	OrderPeriod     string  //订单有效时间
	MerchantOrderId string  //商户订单id
	BankOrderId     string  //本系统订单id
	BankTransId     string  //上游流水id
	OrderAmount     float64 //订单提交的金额
	ShowAmount      float64 //待支付的金额
	FactAmount      float64 //用户实际支付金额
	RollPoolCode    string  //轮询池编码
	RollPoolName    string  //轮询池名臣
	RoadUid         string  //通道标识
	RoadName        string  //通道名称
	PayProductName  string  //上游支付公司的名称
	PayProductCode  string  //上游支付公司的编码代号
	PayTypeCode     string  //支付产品编码
	PayTypeName     string  //支付产品名称
	OsType          string  //操作系统类型
	Status          string  //订单支付状态
	Refund          string  //退款状态
	RefundTime      string  //退款操作时间
	Freeze          string  //冻结状态
	FreezeTime      string  //冻结时间
	Unfreeze        string  //是否已经解冻
	UnfreezeTime    string  //解冻时间
	ReturnUrl       string  //支付完跳转地址
	NotifyUrl       string  //下游回调地址
	MerchantUid     string  //商户id
	MerchantName    string  //商户名称
	AgentUid        string  //该商户所属代理
	AgentName       string  //该商户所属代理名称
	UpdateTime      string
	CreateTime      string
}

const ORDER_INFO = "order_info"

func InsertOrder(orderInfo OrderInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&orderInfo)
	if err != nil {
		logs.Error("insert order info fail: ", err)
		return false
	}
	return true
}

func OrderNoIsEixst(orderId string) bool {
	o := orm.NewOrm()
	exits := o.QueryTable(ORDER_INFO).Filter("merchant_order_id", orderId).Exist()
	return exits
}

func BankOrderIdIsEixst(bankOrderId string) bool {
	o := orm.NewOrm()
	exists := o.QueryTable(ORDER_INFO).Filter("bank_order_id", bankOrderId).Exist()
	return exists
}

func GetOrderLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(ORDER_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, _ := qs.Limit(-1).Count()
	return int(cnt)
}

func GetOrderByMap(params map[string]string, display, offset int) []OrderInfo {
	o := orm.NewOrm()
	var orderInfoList []OrderInfo
	qs := o.QueryTable(ORDER_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(display, offset).OrderBy("-update_time").All(&orderInfoList)
	if err != nil {
		logs.Error("get order by map fail: ", err)
	}
	return orderInfoList
}

func GetSuccessRateByMap(params map[string]string) string {
	o := orm.NewOrm()
	qs := o.QueryTable(ORDER_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	successRate := "0%"
	allCount, _ := qs.Limit(-1).Count()
	successCount, _ := qs.Filter("status", "success").Limit(-1).Count()
	if allCount == 0 {
		return successRate
	}
	tmp := float64(successCount) / float64(allCount) * 100
	successRate = fmt.Sprintf("%.1f", tmp)
	return successRate + "%"
}

func GetAllAmountByMap(params map[string]string) float64 {
	o := orm.NewOrm()
	condition := "select sum(order_amount) as allAmount from order_info "
	for _, v := range params {
		if len(v) > 0 {
			condition = condition + "where "
			break
		}
	}
	flag := false
	if params["create_time__gte"] != "" {
		flag = true
		condition = condition + " create_time >= '" + params["create_time__gte"] + "'"
	}
	if params["create_time__lte"] != "" {
		if flag {
			condition = condition + " and "
		}
		condition = condition + " create_time <= '" + params["create_time__lte"] + "'"
	}
	if params["merchant_name__icontains"] != "" {
		if flag {
			condition = condition + " and "
		}
		condition = condition + "merchant_name like %'" + params["merchant_name__icontains"] + "'% "
	}
	if params["merchant_order_id"] != "" {
		if flag {
			condition = condition + " and "
		}
		condition = condition + " merchant_order_id = '" + params["merchant_order_id"] + "'"
	}
	if params["bank_order_id"] != "" {
		if flag {
			condition = condition + " and "
		}
		condition = condition + " bank_order_id = '" + params["bank_order_id"] + "'"
	}
	if params["status"] != "" {
		if flag {
			condition = condition + " and "
		}
		condition = condition + "status = '" + params["status"] + "'"
	}
	if params["pay_product_code"] != "" {
		if flag {
			condition = condition + " and "
		}
		condition = condition + "pay_product_code = " + params["pay_product_code"] + "'"
	}
	if params["pay_type_code"] != "" {
		if flag {
			condition = condition + " and "
		}
		condition = condition + "pay_type_code = " + params["pay_type_code"]
	}
	logs.Info("get order amount str = ", condition)
	var maps []orm.Params
	allAmount := 0.00
	num, err := o.Raw(condition).Values(&maps)
	if err == nil && num > 0 {
		allAmount, _ = strconv.ParseFloat(maps[0]["allAmount"].(string), 64)
	}
	return allAmount
}

func GetOrderByBankOrderId(bankOrderId string) OrderInfo {
	o := orm.NewOrm()
	var orderInfo OrderInfo
	_, err := o.QueryTable(ORDER_INFO).Filter("bank_order_id", bankOrderId).Limit(1).All(&orderInfo)
	if err != nil {
		logs.Error("get order info by bankOrderId fail: ", err)
	}
	return orderInfo
}

func GetOrderByMerchantOrderId(merchantOrderId string) OrderInfo {
	o := orm.NewOrm()
	var orderInfo OrderInfo
	_, err := o.QueryTable(ORDER_INFO).Filter("merchant_order_id", merchantOrderId).Limit(1).All(&orderInfo)
	if err != nil {
		logs.Error("get order by merchant_order_id: ", err.Error())
	}
	return orderInfo
}

func GetOneOrder(bankOrderId string) OrderInfo {
	o := orm.NewOrm()
	var orderInfo OrderInfo
	_, err := o.QueryTable(ORDER_INFO).Filter("bank_order_id", bankOrderId).Limit(1).All(&orderInfo)
	if err != nil {
		logs.Error("get one order fail: ", err)
	}

	return orderInfo
}
