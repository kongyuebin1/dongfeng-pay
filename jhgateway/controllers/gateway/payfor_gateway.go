/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/12/5 14:05
 ** @Author : yuebin
 ** @File : payfor_gateway
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/12/5 14:05
 ** @Software: GoLand
****************************************************/
package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/rs/xid"
	"juhe/service/common"
	"juhe/service/models"
	"juhe/service/utils"
	"strconv"
	"strings"
)

type PayForGateway struct {
	beego.Controller
}

type PayForResponse struct {
	ResultCode      string `json:"resultCode,omitempty"`
	ResultMsg       string `json:"resultMsg,omitempty"`
	MerchantOrderId string `json:"merchantOrderId,omitempty"`
	SettAmount      string `json:"settAmount,omitempty"`
	SettFee         string `json:"settFee,omitempty"`
	Sign            string `json:"sign,omitempty"`
}

type PayForQueryResponse struct {
	ResultMsg       string `json:"resultMsg,omitempty"`
	MerchantOrderId string `json:"merchantOrderId,omitempty"`
	SettAmount      string `json:"settAmount,omitempty"`
	SettFee         string `json:"settFee,omitempty"`
	SettStatus      string `json:"settStatus,omitempty"`
	Sign            string `json:"sign,omitempty"`
}

type BalanceResponse struct {
	resultCode      string `json:"resultCode,omitempty"`
	balance         string `json:"balance,omitempty"`
	availableAmount string `json:"availableAmount,omitempty"`
	freezeAmount    string `json:"freezeAmount,omitempty"`
	waitAmount      string `json:"waitAmount,omitempty"`
	loanAmount      string `json:"loanAmount,omitempty"`
	payforAmount    string `json:"payforAmount,omitempty"`
	resultMsg       string `json:"resultMsg,omitempty"`
	sign            string `json:"sign,omitempty"`
}

/*
* 接受下游商户的代付请求
 */
func (c *PayForGateway) PayFor() {
	params := make(map[string]string)
	params["merchantKey"] = strings.TrimSpace(c.GetString("merchantKey"))
	params["realname"] = strings.TrimSpace(c.GetString("realname"))
	params["cardNo"] = strings.TrimSpace(c.GetString("cardNo"))
	params["bankCode"] = strings.TrimSpace(c.GetString("bankCode"))
	params["accType"] = strings.TrimSpace(c.GetString("accType"))
	params["province"] = strings.TrimSpace(c.GetString("province"))
	params["city"] = strings.TrimSpace(c.GetString("city"))
	params["bankAccountAddress"] = strings.TrimSpace(c.GetString("bankAccountAddress"))
	params["amount"] = strings.TrimSpace(c.GetString("amount"))
	params["moblieNo"] = strings.TrimSpace(c.GetString("moblieNo"))
	params["merchantOrderId"] = strings.TrimSpace(c.GetString("merchantOrderId"))
	params["sign"] = strings.TrimSpace(c.GetString("sign"))

	payForResponse := new(PayForResponse)
	res, msg := c.checkParams(params)
	if !res {
		payForResponse.ResultCode = "01"
		payForResponse.ResultMsg = msg
		c.Data["json"] = payForResponse
		c.ServeJSON()
		return
	}

	merchantInfo := models.GetMerchantByPaykey(params["merchantKey"])
	if !utils.Md5Verify(params, merchantInfo.MerchantSecret) {
		logs.Error(fmt.Sprintf("下游商户代付请求，签名失败，商户信息: %+v", merchantInfo))
		payForResponse.ResultCode = "01"
		payForResponse.ResultMsg = "下游商户代付请求，签名失败。"
	} else {
		res, msg = c.checkSettAmount(params["amount"])
		if !res {
			payForResponse.ResultCode = "01"
			payForResponse.ResultMsg = msg
			c.Data["json"] = payForResponse
			c.ServeJSON()
			return
		}

		exist := models.IsExistPayForByMerchantOrderId(params["merchantOrderId"])
		if exist {
			logs.Error(fmt.Sprintf("代付订单号重复：merchantOrderId = %s", params["merchantOrderId"]))
			payForResponse.ResultMsg = "商户订单号重复"
			payForResponse.ResultCode = "01"
			c.Data["json"] = payForResponse
			c.ServeJSON()
			return
		}

		settAmount, _ := strconv.ParseFloat(params["amount"], 64)
		payFor := models.PayforInfo{PayforUid: "pppp" + xid.New().String(), MerchantUid: merchantInfo.MerchantUid, MerchantName: merchantInfo.MerchantName,
			MerchantOrderId: params["merchantOrderId"], BankOrderId: "4444" + xid.New().String(), PayforAmount: settAmount, Status: common.PAYFOR_COMFRIM,
			BankCode: params["bankCode"], BankName: params["bankAccountAddress"], BankAccountName: params["realname"], BankAccountNo: params["cardNo"],
			BankAccountType: params["accType"], City: params["city"], Ares: params["province"] + params["city"], PhoneNo: params["mobileNo"], Type: common.SELF_API,
			UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime(),
		}

		if !models.InsertPayfor(payFor) {
			payForResponse.ResultCode = "01"
			payForResponse.ResultMsg = "代付记录插入失败"
		} else {
			payForResponse.ResultMsg = "代付订单已生成"
			payForResponse.ResultCode = "00"
			payForResponse.SettAmount = params["amount"]
			payForResponse.MerchantOrderId = params["MerchantOrderId"]

			tmp := make(map[string]string)
			tmp["resultCode"] = payForResponse.ResultCode
			tmp["resultMsg"] = payForResponse.ResultMsg
			tmp["merchantOrderId"] = payForResponse.MerchantOrderId
			tmp["settAmount"] = payForResponse.SettAmount
			keys := utils.SortMap(tmp)
			sign := utils.GetMD5Sign(params, keys, merchantInfo.MerchantSecret)
			tmp["sign"] = sign

			c.Data["json"] = payForResponse
			c.ServeJSON()
		}
	}
}

func (c *PayForGateway) checkSettAmount(settAmount string) (bool, string) {
	_, err := strconv.ParseFloat(settAmount, 64)
	if err != nil {
		logs.Error(fmt.Sprintf("代付金额有误，settAmount = %s", settAmount))
		return false, "代付金额有误"
	}
	return true, ""
}

func (c *PayForGateway) checkParams(params map[string]string) (bool, string) {
	for k, v := range params {
		if v == "" || len(v) == 0 {
			return false, fmt.Sprintf("字段： %s 为必填！", k)
		}
	}
	return true, ""
}

/*
* 代付结果查询，
 */
func (c *PayForGateway) PayForQuery() {
	params := make(map[string]string)
	params["merchantKey"] = strings.TrimSpace(c.GetString("merchantKey"))
	params["timestamp"] = strings.TrimSpace(c.GetString("timestamp"))
	params["merchantOrderId"] = strings.TrimSpace(c.GetString("merchantOrderId"))
	params["sign"] = strings.TrimSpace(c.GetString("sign"))

	query := make(map[string]string)
	query["merchantOrderId"] = params["merchantOrderId"]
	merchantInfo := models.GetMerchantByPaykey(params["merchantKey"])
	if !utils.Md5Verify(params, merchantInfo.MerchantSecret) {
		query["resultMsg"] = "签名错误"
		query["settStatus"] = "03"
		query["sign"] = utils.GetMD5Sign(params, utils.SortMap(params), merchantInfo.MerchantSecret)
	} else {
		payForInfo := models.GetPayForByMerchantOrderId(params["merchantOrderId"])
		if payForInfo.BankOrderId == "" {
			query["resultMsg"] = "不存在这样的代付订单"
			query["settStatus"] = "03"
			query["sign"] = utils.GetMD5Sign(params, utils.SortMap(params), merchantInfo.MerchantSecret)
		} else {
			switch payForInfo.Status {
			case common.PAYFOR_BANKING:
				query["resultMsg"] = "打款中"
				query["settStatus"] = "02"
			case common.PAYFOR_SOLVING:
				query["resultMsg"] = "打款中"
				query["settStatus"] = "02"
			case common.PAYFOR_COMFRIM:
				query["resultMsg"] = "打款中"
				query["settStatus"] = "02"
			case common.PAYFOR_SUCCESS:
				query["resultMsg"] = "打款成功"
				query["settStatus"] = "00"
				query["settAmount"] = strconv.FormatFloat(payForInfo.PayforAmount, 'f', 2, 64)
				query["settFee"] = strconv.FormatFloat(payForInfo.PayforFee, 'f', 2, 64)
			case common.PAYFOR_FAIL:
				query["resultMsg"] = "打款失败"
				query["settStatus"] = "01"
			}
			query["sign"] = utils.GetMD5Sign(query, utils.SortMap(query), merchantInfo.MerchantSecret)
		}
	}

	mJson, err := json.Marshal(query)
	if err != nil {
		logs.Error("PayForQuery json marshal fail：", err)
	}

	c.Data["json"] = string(mJson)
	c.ServeJSON()
}

/*
* 商户查找余额
 */
func (c *PayForGateway) Balance() {
	params := make(map[string]string)
	params["merchantKey"] = strings.TrimSpace(c.GetString("merchantKey"))
	params["timestamp"] = strings.TrimSpace(c.GetString("timestamp"))
	params["sign"] = strings.TrimSpace(c.GetString("sign"))

	balanceResponse := new(BalanceResponse)
	res, msg := c.checkParams(params)
	if !res {
		balanceResponse.resultCode = "-1"
		balanceResponse.resultMsg = msg
		c.Data["json"] = balanceResponse
	} else {
		merchantInfo := models.GetMerchantByPaykey(params["merchantKey"])
		if !utils.Md5Verify(params, merchantInfo.MerchantSecret) {
			balanceResponse.resultCode = "-1"
			balanceResponse.resultMsg = "签名错误"
			c.Data["json"] = balanceResponse
		} else {
			accountInfo := models.GetAccountByUid(merchantInfo.MerchantUid)
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
			c.Data["json"] = string(mJson)
		}
	}
	c.ServeJSON()
}
