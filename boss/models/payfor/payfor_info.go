/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/11/25 14:32
 ** @Author : yuebin
 ** @File : payfor_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/25 14:32
 ** @Software: GoLand
****************************************************/
package payfor

import (
	"boss/common"
	"boss/models/accounts"
	"boss/utils"
	"context"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type PayforInfo struct {
	Id                 int
	PayforUid          string
	MerchantUid        string
	MerchantName       string
	MerchantOrderId    string
	BankOrderId        string
	BankTransId        string
	RoadUid            string
	RoadName           string
	RollPoolCode       string
	RollPoolName       string
	PayforFee          float64
	PayforAmount       float64
	PayforTotalAmount  float64
	BankCode           string
	BankName           string
	BankAccountName    string
	BankAccountNo      string
	BankAccountType    string
	Country            string
	City               string
	Ares               string
	BankAccountAddress string
	PhoneNo            string
	GiveType           string
	Type               string
	NotifyUrl          string
	Status             string
	IsSend             string
	RequestTime        string
	ResponseTime       string
	ResponseContent    string
	Remark             string
	CreateTime         string
	UpdateTime         string
}

const PAYFORINFO = "payfor_info"

func InsertPayfor(payFor PayforInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&payFor)
	if err != nil {
		logs.Error("insert payfor fail: ", err)
		return false
	}
	return true
}

func GetPayForLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(PAYFORINFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	cnt, err := qs.Limit(-1).Count()

	if err != nil {
		logs.Error("get pay for len by map fail: ", err)
	}

	return int(cnt)
}

func GetPayForByMap(params map[string]string, displayCount, offset int) []PayforInfo {
	o := orm.NewOrm()
	var payForList []PayforInfo

	qs := o.QueryTable(PAYFORINFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	_, err := qs.Limit(displayCount, offset).OrderBy("-create_time").All(&payForList)

	if err != nil {
		logs.Error("get agentInfo by map fail: ", err)
	}

	return payForList
}

func GetPayForListByParams(params map[string]string) []PayforInfo {
	o := orm.NewOrm()
	var payForList []PayforInfo
	qs := o.QueryTable(PAYFORINFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(-1).All(&payForList)
	if err != nil {
		logs.Error("GetPayForListByParams fail：", err)
	}
	return payForList
}

func GetPayForByBankOrderId(bankOrderId string) PayforInfo {
	o := orm.NewOrm()
	var payFor PayforInfo
	_, err := o.QueryTable(PAYFORINFO).Filter("bank_order_id", bankOrderId).Limit(1).All(&payFor)

	if err != nil {
		logs.Error("get pay for by bank_order_id fail: ", err)
	}

	return payFor
}

func ForUpdatePayFor(payFor PayforInfo) bool {
	o := orm.NewOrm()

	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		var tmp PayforInfo
		if err := txOrm.Raw("select * from payfor_info where bank_order_id = ? for update", payFor.BankOrderId).QueryRow(&tmp); err != nil || tmp.PayforUid == "" {
			logs.Error("for update payfor select fail:", err)
			return err
		}

		if tmp.Status == common.PAYFOR_FAIL || tmp.Status == common.PAYFOR_SUCCESS {
			return errors.New("订单已经处理")
		}

		//如果是手动打款，并且是需要处理商户金额
		if payFor.Status == common.PAYFOR_SOLVING && tmp.Status == common.PAYFOR_COMFRIM &&
			payFor.GiveType == common.PAYFOR_HAND && payFor.Type != common.SELF_HELP {

			var account accounts.AccountInfo
			if err := txOrm.Raw("select * from account_info where account_uid = ? for update", payFor.MerchantUid).QueryRow(&account); err != nil || account.AccountUid == "" {
				logs.Error("for update payfor select account info，fail：", err)
				return err
			}
			//计算该用户的可用金额
			ableAmount := account.SettleAmount - account.FreezeAmount - account.PayforAmount - account.LoanAmount
			if ableAmount >= payFor.PayforAmount+payFor.PayforFee {
				account.PayforAmount += payFor.PayforFee + payFor.PayforAmount
				account.UpdateTime = utils.GetBasicDateTime()
				if _, err := txOrm.Update(&account); err != nil {
					logs.Error("for update payfor update account fail：", err)
					return err
				}
			} else {
				logs.Error(fmt.Sprintf("商户uid=%s，可用金额不够", payFor.MerchantUid))
				payFor.ResponseContent = "商户可用余额不足"
				payFor.Status = common.PAYFOR_FAIL
			}
		}

		if _, err := txOrm.Update(&payFor); err != nil {
			logs.Error("for update payfor fail: ", err)
			return err
		}

		return nil

	}); err != nil {
		return false
	}
	return true
}

func UpdatePayFor(payFor PayforInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&payFor)

	if err != nil {
		logs.Error("update pay for fail：", err)
		return false
	}

	return true
}
