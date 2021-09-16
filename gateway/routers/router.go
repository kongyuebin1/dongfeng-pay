package routers

import (
	"gateway/controllers/gateway"
	"gateway/supplier/third_party"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	//网关处理函数
	web.Router("/gateway/scan", &gateway.ScanController{}, "*:Scan")
	web.Router("/err/params", &gateway.ErrorGatewayController{}, "*:ErrorParams")
	//代付相关的接口
	web.Router("gateway/payfor", &gateway.PayForGateway{}, "*:PayFor")
	web.Router("/gateway/payfor/query", &gateway.PayForGateway{}, "*:PayForQuery")
	web.Router("/gateway/balance", &gateway.PayForGateway{}, "*:Balance")
	// 接收银行回调
	web.Router("/daili/notify", &third_party.DaiLiImpl{}, "*:PayNotify")
}
