package service

import (
	"boss/common"
	"boss/datas"
	"boss/models/notify"
	"boss/models/order"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"strings"
)

type SendNotifyMerchantService struct {
}

func (c *SendNotifyMerchantService) SendNotifyToMerchant(bankOrderId string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = -1
	orderInfo := order.GetOrderByBankOrderId(bankOrderId)
	if orderInfo.Status == common.WAIT {
		keyDataJSON.Msg = "该订单不是成功状态，不能回调"
	} else {
		notifyInfo := notify.GetNotifyInfoByBankOrderId(bankOrderId)
		notifyUrl := notifyInfo.Url
		logs.Info(fmt.Sprintf("boss管理后台手动触发订单回调，url=%s", notifyUrl))
		req := httplib.Post(notifyUrl)
		response, err := req.String()
		if err != nil {
			logs.Error("回调发送失败，fail：", err)
			keyDataJSON.Msg = fmt.Sprintf("该订单回调发送失败，订单回调，fail：%s", err)
		} else {
			if !strings.Contains(strings.ToLower(response), "success") {
				keyDataJSON.Msg = fmt.Sprintf("该订单回调发送成功，但是未返回success字段， 商户返回内容=%s",
					response)
			} else {
				keyDataJSON.Code = 200
				keyDataJSON.Msg = fmt.Sprintf("该订单回调发送成功")
			}
		}
	}
	return keyDataJSON
}

func (c *SendNotifyMerchantService) SelfSendNotify(bankOrderId string) *datas.KeyDataJSON {
	notifyInfo := notify.GetNotifyInfoByBankOrderId(bankOrderId)

	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = 200

	req := httplib.Post(notifyInfo.Url)

	response, err := req.String()
	if err != nil {
		keyDataJSON.Msg = fmt.Sprintf("订单 bankOrderId=%s，已经发送回调出错：%s", bankOrderId, err)
	} else {
		keyDataJSON.Msg = fmt.Sprintf("订单 bankOrderId=%s，已经发送回调，商户返回内容：%s",
			bankOrderId, response)
	}
	return keyDataJSON
}
