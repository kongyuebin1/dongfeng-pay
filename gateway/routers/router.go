package routers

import (
	"gateway/controllers/gateway"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//网关处理函数
	beego.Router("/gateway/scan", &gateway.ScanController{}, "*:Scan")
	beego.Router("/err/params", &gateway.ErrorGatewayController{}, "*:ErrorParams")
	//代付相关的接口
	beego.Router("gateway/payfor", &gateway.PayForGateway{}, "*:PayFor")
	beego.Router("/gateway/payfor/query", &gateway.PayForGateway{}, "*:PayForQuery")
	beego.Router("/gateway/balance", &gateway.PayForGateway{}, "*:Balance")
}
