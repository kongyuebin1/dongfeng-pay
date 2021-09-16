package test

import (
	_ "gateway/message"
	_ "gateway/models"
	"gateway/service"
	"gateway/utils"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/rs/xid"
	"net/url"
	"testing"
)
import _ "gateway/routers"

/*
** 充值测试
 */
func TestPay(t *testing.T) {
	params := make(map[string]string)
	params["orderNo"] = xid.New().String()
	params["productName"] = "kongyuhebin"
	params["orderPeriod"] = "1"
	params["orderPrice"] = "100.00"
	params["payWayCode"] = "WEIXIN_SCAN"
	params["osType"] = "1"
	params["notifyUrl"] = "http://localhost:12309/shop/notify"
	params["payKey"] = "kkkkc254gk8isf001cqrj6p0"
	keys := utils.SortMap(params)
	params["sign"] = utils.GetMD5Sign(params, keys, "ssssc254gk8isf001cqrj6pg")

	u := url.Values{}
	for k, v := range params {
		u.Add(k, v)
	}

	l := "http://localhost:12309/gateway/scan?" + u.Encode()
	logs.Info("请求url：" + l)

	resp := httplib.Get(l)
	s, err := resp.String()

	if err != nil {
		logs.Error("请求错误：" + err.Error())

	}

	logs.Info("微信扫码返回结果：" + s)
}

/**
** 充值失败回调
 */
func TestPayFail(t *testing.T) {
	service.SolvePayFail("6666c50bd567matj5v6g30dg", "")
}

/**
** 充值成功
 */
func TestPaySuccess(t *testing.T) {
	service.SolvePaySuccess("6666c50mhcu7matjtv0a4330", 0, "")
}
