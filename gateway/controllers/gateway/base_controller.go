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
	"gateway/response"
	"gateway/service"
	"github.com/beego/beego/v2/server/web"
	"strings"
)

type BaseGateway struct {
	web.Controller
}

//获取商户请求过来的基本参数参数
func (c *BaseGateway) PayPrepare() *response.PayBaseResp {
	params := make(map[string]string)
	//获取客户端的ip
	clientIp := c.Ctx.Input.IP()
	params["orderNo"] = strings.TrimSpace(c.GetString("orderNo"))
	params["productName"] = strings.TrimSpace(c.GetString("productName"))
	params["orderPeriod"] = strings.TrimSpace(c.GetString("orderPeriod"))
	params["orderPrice"] = strings.TrimSpace(c.GetString("orderPrice"))
	params["payWayCode"] = strings.TrimSpace(c.GetString("payWayCode"))
	params["osType"] = strings.TrimSpace(c.GetString("osType"))
	params["notifyUrl"] = strings.TrimSpace(c.GetString("notifyUrl"))
	//c.Params["returnUrl"] = strings.TrimSpace(c.GetString("returnUrl"))
	params["payKey"] = strings.TrimSpace(c.GetString("payKey"))
	params["sign"] = strings.TrimSpace(c.GetString("sign"))

	p := service.GetMerchantInfo(params)
	p.ClientIp = clientIp
	p = service.JudgeParams(p)

	if p.Code != -1 {
		p.Code = 200
	}

	return p
}
