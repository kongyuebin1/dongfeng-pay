package gateway

import (
	"gateway/conf"
	"gateway/models/order"
	"gateway/query"
	"gateway/service"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/server/web"
)

type OrderController struct {
	web.Controller
}

func (c *OrderController) OrderQuery() {
	bankOrderId := c.GetString("bankOrderId")
	logs.Debug("获取到boss后台的银行id = " + bankOrderId)

	qy := query.SupplierOrderQueryResult(bankOrderId)

	if qy {
		c.Ctx.WriteString("success")
	} else {
		c.Ctx.WriteString("fail")
	}
	c.StopRun()
}

func (c *OrderController) OrderUpdate() {
	bankOrderId := c.GetString("bankOrderId")
	solveType := c.GetString("solveType")
	orderInfo := order.GetOrderByBankOrderId(bankOrderId)
	flag := false
	if orderInfo.BankOrderId == "" {
		logs.Error("该订单不存在,bankOrderId=", bankOrderId)
	} else if orderInfo.Status != conf.SUCCESS {
		logs.Notice("该订单没有完成支付，不能进行此操作，bankOrderId = ", bankOrderId)
	} else {
		switch solveType {
		case conf.SUCCESS:
			flag = service.SolvePaySuccess(bankOrderId, orderInfo.FactAmount, orderInfo.BankTransId)
		case conf.FAIL:
			flag = service.SolvePayFail(bankOrderId, orderInfo.BankTransId)
		case conf.FREEZE_AMOUNT:
			//将这笔订单进行冻结
			flag = service.SolveOrderFreeze(bankOrderId)
		case conf.UNFREEZE_AMOUNT:
			//将这笔订单金额解冻
			flag = service.SolveOrderUnfreeze(bankOrderId)
		case conf.REFUND:
			if orderInfo.Status == conf.SUCCESS {
				flag = service.SolveRefund(bankOrderId)
			}
		case conf.ORDERROLL:
			if orderInfo.Status == conf.SUCCESS {
				flag = service.SolveOrderRoll(bankOrderId)
			}
		default:
			logs.Error("不存在这样的处理类型")
		}
		if flag {
			c.Ctx.WriteString(conf.SUCCESS)
		} else {
			c.Ctx.WriteString(conf.FAIL)
		}
	}

	c.StopRun()
}
