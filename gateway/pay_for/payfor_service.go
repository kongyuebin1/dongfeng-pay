/***************************************************
 ** @Desc : 代付处理
 ** @Time : 2019/11/28 18:52
 ** @Author : yuebin
 ** @File : payfor_service
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/28 18:52
 ** @Software: GoLand
****************************************************/
package pay_for

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/conf"
	"gateway/message"
	"gateway/models/accounts"
	"gateway/models/merchant"
	"gateway/models/payfor"
	"gateway/models/road"
	"gateway/response"
	"gateway/supplier/third_party"
	"gateway/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/rs/xid"
	"strconv"
	"strings"
)

/**
** 程序自动代付
 */
func AutoPayFor(params map[string]string, giveType string) *response.PayForResponse {

	payForResponse := new(response.PayForResponse)

	merchantInfo := merchant.GetMerchantByPaykey(params["merchantKey"])
	if !utils.Md5Verify(params, merchantInfo.MerchantSecret) {
		logs.Error(fmt.Sprintf("下游商户代付请求，签名失败，商户信息: %+v", merchantInfo))
		payForResponse.ResultCode = "01"
		payForResponse.ResultMsg = "下游商户代付请求，签名失败。"
		return payForResponse
	} else {
		res, msg := checkSettAmount(params["amount"])
		if !res {
			payForResponse.ResultCode = "01"
			payForResponse.ResultMsg = msg

			return payForResponse
		}

		exist := payfor.IsExistPayForByMerchantOrderId(params["merchantOrderId"])
		if exist {
			logs.Error(fmt.Sprintf("代付订单号重复：merchantOrderId = %s", params["merchantOrderId"]))
			payForResponse.ResultMsg = "商户订单号重复"
			payForResponse.ResultCode = "01"

			return payForResponse
		}

		settAmount, err := strconv.ParseFloat(params["amount"], 64)
		if err != nil {
			logs.Error("代付的金额错误：", err)
			payForResponse.ResultMsg = "代付金额错误"
			payForResponse.ResultCode = "01"
			return payForResponse
		}

		p := payfor.PayforInfo{
			PayforUid:       "pppp" + xid.New().String(),
			MerchantUid:     merchantInfo.MerchantUid,
			MerchantName:    merchantInfo.MerchantName,
			MerchantOrderId: params["merchantOrderId"],
			BankOrderId:     "4444" + xid.New().String(),
			PayforAmount:    settAmount,
			Status:          conf.PAYFOR_COMFRIM,
			BankAccountName: params["realname"],
			BankAccountNo:   params["cardNo"],
			BankAccountType: params["accType"],
			City:            params["city"],
			Ares:            params["province"] + params["city"],
			PhoneNo:         params["mobileNo"],
			GiveType:        giveType,
			UpdateTime:      utils.GetBasicDateTime(),
			CreateTime:      utils.GetBasicDateTime(),
			RequestTime:     utils.GetBasicDateTime(),
		}

		// 获取银行编码和银行名称
		p.BankCode = utils.GetBankCodeByBankCardNo(p.BankAccountNo)
		p.BankName = utils.GetBankNameByCode(p.BankCode)

		if !payfor.InsertPayfor(p) {
			payForResponse.ResultCode = "01"
			payForResponse.ResultMsg = "代付记录插入失败"
		} else {
			payForResponse.ResultMsg = "代付订单已生成"
			payForResponse.ResultCode = "00"
			payForResponse.SettAmount = params["amount"]
			payForResponse.MerchantOrderId = params["MerchantOrderId"]

			p = payfor.GetPayForByBankOrderId(p.BankOrderId)

			if findPayForRoad(p) {
				payForResponse.ResultCode = "00"
				payForResponse.ResultMsg = "银行处理中"
			} else {
				payForResponse.ResultCode = "01"
				payForResponse.ResultMsg = "系统处理失败"
			}

		}

		return payForResponse
	}

}

/**
* 返回1表示需要手动打款，返回0表示银行已经受理，-1表示系统处理失败
 */
func findPayForRoad(p payfor.PayforInfo) bool {

	m := merchant.GetMerchantByUid(p.MerchantUid)
	// 检查商户是否设置了自动代付
	if m.AutoPayFor == conf.NO || m.AutoPayFor == "" {
		logs.Notice(fmt.Sprintf("该商户uid=%s， 没有开通自动代付功能", p.MerchantUid))
		p.Type = conf.PAYFOR_HAND
		payfor.UpdatePayFor(p)
	} else {

		if m.SinglePayForRoadUid != "" {
			p.RoadUid = m.SinglePayForRoadUid
			p.RoadName = m.SinglePayForRoadName
		} else {
			roadPoolInfo := road.GetRoadPoolByRoadPoolCode(m.RollPayForRoadCode)
			roadUids := strings.Split(roadPoolInfo.RoadUidPool, "||")
			roadInfoList := road.GetRoadInfosByRoadUids(roadUids)
			if len(roadUids) == 0 || len(roadInfoList) == 0 {
				logs.Error(fmt.Sprintf("通道轮询池=%s, 没有配置通道", m.RollPayForRoadCode))
			} else {
				p.RoadUid = roadInfoList[0].RoadUid
				p.RoadName = roadInfoList[0].RoadName
			}
		}

		if !payfor.UpdatePayFor(p) {
			return false
		}

		if len(p.RoadUid) > 0 {
			roadInfo := road.GetRoadInfoByRoadUid(p.RoadUid)
			p.PayforFee = roadInfo.SettleFee
			p.PayforTotalAmount = p.PayforFee + p.PayforAmount

			if m.PayforFee > conf.ZERO {
				logs.Info(fmt.Sprintf("商户uid=%s，有单独的代付手续费。", m.MerchantUid))
				p.PayforFee = m.PayforFee
				p.PayforTotalAmount = p.PayforFee + p.PayforAmount
			}

			if !payfor.UpdatePayFor(p) {
				return false
			}

			if p.GiveType == conf.SELF_HELP {
				if !MerchantSelf(p) {
					return false
				}
			} else {
				if !SendPayFor(p) {
					return false
				}
			}
		} else {
			p.Status = conf.PAYFOR_FAIL
			if !payfor.UpdatePayFor(p) {
				return false
			}
			p.ResponseContent = "没有设置代付通道"
		}
	}

	return true
}

/**
** 商户自己体现
 */
func MerchantSelf(p payfor.PayforInfo) bool {
	o := orm.NewOrm()

	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		p.UpdateTime = utils.GetBasicDateTime()
		p.Status = conf.PAYFOR_BANKING
		p.RequestTime = utils.GetBasicDateTime()
		p.IsSend = conf.YES
		if _, err := txOrm.Update(&p); err != nil {
			return err
		}

		RequestPayFor(p)

		return nil

	}); err != nil {
		return false
	}
	return true
}

func SendPayFor(p payfor.PayforInfo) bool {
	o := orm.NewOrm()

	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		var account accounts.AccountInfo
		if err := txOrm.Raw("select * from account_info where account_uid = ? for update", p.MerchantUid).QueryRow(&account); err != nil || account.AccountUid == "" {
			logs.Error("send payfor select account fail：", err)
			return err
		}

		//支付金额不足，将直接判定为失败，不往下面邹逻辑了
		if account.SettleAmount-account.PayforAmount < p.PayforAmount+p.PayforFee {
			p.Status = conf.PAYFOR_FAIL
			p.UpdateTime = utils.GetBasicDateTime()

			if _, err := txOrm.Update(&p); err != nil {
				return err
			} else {
				return nil
			}
		}

		account.UpdateTime = utils.GetBasicDateTime()
		account.PayforAmount = account.PayforAmount + p.PayforAmount + p.PayforFee

		if _, err := txOrm.Update(&account); err != nil {
			logs.Error(fmt.Sprintf("商户uid=%s，在发送代付给上游的处理中，更新账户表出错, err: %s", p.MerchantUid, err))
			return err
		}

		p.IsSend = conf.YES
		p.Status = conf.PAYFOR_BANKING //变为银行处理中
		p.GiveType = conf.PAYFOR_ROAD
		p.RequestTime = utils.GetBasicDateTime()
		p.UpdateTime = utils.GetBasicDateTime()

		if _, err := txOrm.Update(&p); err != nil {
			logs.Error(fmt.Sprintf("商户uid=%s，在发送代付给上游的处理中，更代付列表出错， err：%s", p.MerchantUid, err))
			return err
		}

		RequestPayFor(p)

		return nil
	}); err != nil {
		return false
	}
	return true
}

func RequestPayFor(p payfor.PayforInfo) {
	if p.RoadUid == "" {
		return
	}
	p.Type = conf.PAYFOR_ROAD
	roadInfo := road.GetRoadInfoByRoadUid(p.RoadUid)
	supplierCode := roadInfo.ProductUid
	supplier := third_party.GetPaySupplierByCode(supplierCode)
	res := supplier.PayFor(p)
	logs.Info(fmt.Sprintf("代付uid=%s，上游处理结果为：%s", p.PayforUid, res))
	//将代付订单号发送到消息队列
	message.SendMessage(conf.MQ_PAYFOR_QUERY, p.BankOrderId)
}

/**
* 代付结果查询
 */
func PayForResultQuery(params map[string]string) string {

	query := make(map[string]string)
	query["merchantOrderId"] = params["merchantOrderId"]
	merchantInfo := merchant.GetMerchantByPaykey(params["merchantKey"])
	if !utils.Md5Verify(params, merchantInfo.MerchantSecret) {
		query["resultMsg"] = "签名错误"
		query["settStatus"] = "03"
		query["sign"] = utils.GetMD5Sign(params, utils.SortMap(params), merchantInfo.MerchantSecret)
	} else {
		payForInfo := payfor.GetPayForByMerchantOrderId(params["merchantOrderId"])
		if payForInfo.BankOrderId == "" {
			query["resultMsg"] = "不存在这样的代付订单"
			query["settStatus"] = "03"
			query["sign"] = utils.GetMD5Sign(params, utils.SortMap(params), merchantInfo.MerchantSecret)
		} else {
			switch payForInfo.Status {
			case conf.PAYFOR_BANKING:
				query["resultMsg"] = "打款中"
				query["settStatus"] = "02"
			case conf.PAYFOR_SOLVING:
				query["resultMsg"] = "打款中"
				query["settStatus"] = "02"
			case conf.PAYFOR_COMFRIM:
				query["resultMsg"] = "打款中"
				query["settStatus"] = "02"
			case conf.PAYFOR_SUCCESS:
				query["resultMsg"] = "打款成功"
				query["settStatus"] = "00"
				query["settAmount"] = strconv.FormatFloat(payForInfo.PayforAmount, 'f', 2, 64)
				query["settFee"] = strconv.FormatFloat(payForInfo.PayforFee, 'f', 2, 64)
			case conf.PAYFOR_FAIL:
				query["resultMsg"] = "打款失败"
				query["settStatus"] = "01"
			}
			query["sign"] = utils.GetMD5Sign(query, utils.SortMap(query), merchantInfo.MerchantSecret)
		}
	}

	mJson, err := json.Marshal(query)
	if err != nil {
		logs.Error("PayForQuery json marshal fail：", err)
		return fmt.Sprintf("PayForQuery json marshal fail：%s", err.Error())
	} else {
		return string(mJson)
	}
}

/**
* 商户查询余额
 */
func BalanceQuery(params map[string]string) string {

	balanceResponse := new(response.BalanceResponse)
	str := ""
	merchantInfo := merchant.GetMerchantByPaykey(params["merchantKey"])
	if !utils.Md5Verify(params, merchantInfo.MerchantSecret) {
		balanceResponse.ResultCode = "-1"
		balanceResponse.ResultMsg = "签名错误"
		mJson, _ := json.Marshal(balanceResponse)
		str = string(mJson)
	} else {
		accountInfo := accounts.GetAccountByUid(merchantInfo.MerchantUid)
		tmp := make(map[string]string)
		tmp["resultCode"] = "00"
		tmp["balance"] = strconv.FormatFloat(accountInfo.Balance, 'f', 2, 64)
		tmp["availableAmount"] = strconv.FormatFloat(accountInfo.SettleAmount, 'f', 2, 64)
		tmp["freezeAmount"] = strconv.FormatFloat(accountInfo.FreezeAmount, 'f', 2, 64)
		tmp["waitAmount"] = strconv.FormatFloat(accountInfo.WaitAmount, 'f', 2, 64)
		tmp["loanAmount"] = strconv.FormatFloat(accountInfo.LoanAmount, 'f', 2, 64)
		tmp["payforAmount"] = strconv.FormatFloat(accountInfo.PayforAmount, 'f', 2, 64)
		tmp["resultMsg"] = "查询成功"
		tmp["sign"] = utils.GetMD5Sign(tmp, utils.SortMap(tmp), merchantInfo.MerchantSecret)
		mJson, _ := json.Marshal(tmp)
		str = string(mJson)
	}

	return str
}

func checkSettAmount(settAmount string) (bool, string) {
	_, err := strconv.ParseFloat(settAmount, 64)
	if err != nil {
		logs.Error(fmt.Sprintf("代付金额有误，settAmount = %s", settAmount))
		return false, "代付金额有误"
	}
	return true, ""
}
