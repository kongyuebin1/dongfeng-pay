package test

import (
	"gateway/conf"
	_ "gateway/message"
	_ "gateway/models"
	"gateway/models/payfor"
	"gateway/pay_for"
	"gateway/utils"
	"github.com/beego/beego/v2/core/logs"
	"github.com/rs/xid"
	"testing"
)

func TestAutoPayFor(t *testing.T) {
	params := make(map[string]string)

	params["merchantKey"] = "kkkkc254gk8isf001cqrj6p0"
	params["realname"] = "孔跃彬"
	params["cardNo"] = "6214830200383973"
	params["accType"] = "0"
	params["amount"] = "100"
	paySecret := "ssssc254gk8isf001cqrj6pg"
	params["merchantOrderId"] = xid.New().String()
	keys := utils.SortMap(params)
	params["sign"] = utils.GetMD5Sign(params, keys, paySecret)
	payFor := pay_for.AutoPayFor(params, conf.SELF_API)
	logs.Info(payFor)
}

func TestPayForFail(t *testing.T) {
	p := new(payfor.PayforInfo)
	p.BankOrderId = "4444c4vlk3u7mathho2o8md0"
	res := pay_for.PayForFail(*p)
	logs.Info(res)
}

func TestPayForSuccess(t *testing.T) {
	p := new(payfor.PayforInfo)
	p.BankOrderId = "4444c4vlk3u7mathho2o8md0"
	res := pay_for.PayForSuccess(*p)
	logs.Info(res)
}
