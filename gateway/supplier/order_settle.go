/***************************************************
 ** @Desc : 订单结算，将订单上面的钱加入到账户余额中
 ** @Time : 2019/11/22 11:34
 ** @Author : yuebin
 ** @File : order_settle
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/22 11:34
 ** @Software: GoLand
****************************************************/
package controller

import (
	"context"
	"errors"
	"fmt"
	"gateway/common"
	"gateway/models"
	"gateway/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

//订单结算，将那些支付成功的订单金额加入到商户账户的结算金额中
func OrderSettle() {

	params := make(map[string]string)
	params["is_allow_settle"] = common.YES
	params["is_complete_settle"] = common.NO
	orderSettleList := models.GetOrderSettleListByParams(params)
	for _, orderSettle := range orderSettleList {
		orderProfitInfo := models.GetOrderProfitByBankOrderId(orderSettle.BankOrderId)
		if !settle(orderSettle, orderProfitInfo) {
			logs.Error(fmt.Sprintf("结算订单bankOrderId=%s， 执行失败", orderSettle.BankOrderId))
		} else {
			logs.Info(fmt.Sprintf("结算订单bankOrderId=%s，执行成功", orderSettle.BankOrderId))
		}
	}
}

func settle(orderSettle models.OrderSettleInfo, orderProfit models.OrderProfitInfo) bool {
	o := orm.NewOrm()

	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		tmpSettle := new(models.OrderSettleInfo)
		if err := txOrm.Raw("select * from order_settle_info where bank_order_id=? for update", orderSettle.BankOrderId).QueryRow(tmpSettle); err != nil || tmpSettle == nil {
			logs.Error("获取tmpSettle失败，bankOrderId=%s", orderSettle.BankOrderId)
			return err
		}
		tmpSettle.UpdateTime = utils.GetBasicDateTime()
		tmpSettle.IsCompleteSettle = common.YES
		if _, err := txOrm.Update(tmpSettle); err != nil {
			logs.Error("更新tmpSettle失败，错误：", err)
			return err
		}

		accountInfo := new(models.AccountInfo)
		if err := txOrm.Raw("select * from account_info where account_uid=? for update", orderSettle.MerchantUid).QueryRow(accountInfo); err != nil || accountInfo == nil {
			logs.Error("结算select account info失败，错误信息：", err)
			return err
		}
		accountInfo.UpdateTime = utils.GetBasicDateTime()
		accountInfo.SettleAmount += orderProfit.FactAmount
		if _, err := txOrm.Update(accountInfo); err != nil {
			logs.Error("结算update account 失败，错误信息：", err)
			return err
		}

		merchantDeployInfo := models.GetMerchantDeployByUidAndPayType(accountInfo.AccountUid, orderSettle.PayTypeCode)
		if merchantDeployInfo.IsLoan == common.YES {
			loadAmount := merchantDeployInfo.LoanRate * 0.01 * orderProfit.FactAmount
			date := utils.GetDate()
			params := make(map[string]string)
			params["merchant_uid"] = tmpSettle.MerchantUid
			params["road_uid"] = tmpSettle.RoadUid
			params["load_date"] = date
			if !models.IsExistMerchantLoadByParams(params) {
				tmp := models.MerchantLoadInfo{Status: common.NO, MerchantUid: orderSettle.MerchantUid, RoadUid: orderSettle.RoadUid,
					LoadDate: date, LoadAmount: loadAmount, UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}
				if _, err := txOrm.Insert(tmp); err != nil {
					logs.Error("結算插入merchantLoad失敗，失败信息：", err)
					return err
				} else {
					logs.Info("结算插入新的merchantLoad信息成功")
				}
			} else {
				merchantLoad := new(models.MerchantLoadInfo)
				if err := txOrm.Raw("select * from merchant_load_info where merchant_uid=? and road_uid=? and load_date=? for update").
					QueryRow(merchantLoad); err != nil || merchantLoad == nil {
					logs.Error(fmt.Sprintf("结算过程，select merchant load info失败，错误信息：%s", err))
					return err
				} else {
					merchantLoad.UpdateTime = utils.GetBasicDateTime()
					merchantLoad.LoadAmount += loadAmount
					if _, err := txOrm.Update(merchantLoad); err != nil {
						logs.Error(fmt.Sprintf("结算过程，update merchant load info失败，失败信息：%s", err))
						return err
					}
				}
			}
		} else {
			logs.Info(fmt.Sprintf("结算过程中，该商户不需要押款，全款结算"))
		}
		return nil
	}); err != nil {
		return false
	}
	return true
}

/*
* 商户的押款释放处理，根据商户的押款时间进行处理
 */
func MerchantLoadSolve() {
	hour := time.Now().Hour()
	merchantDeployList := models.GetMerchantDeployByHour(hour)
	for _, merchantDeploy := range merchantDeployList {
		logs.Info(fmt.Sprintf("开始执行商户uid=%s，进行解款操作", merchantDeploy.MerchantUid))

		loadDate := utils.GetDateBeforeDays(merchantDeploy.LoanDays)
		params := make(map[string]string)
		params["status"] = common.NO
		params["merchant_uid"] = merchantDeploy.MerchantUid
		params["load_date"] = loadDate

		merchantLoadList := models.GetMerchantLoadInfoByMap(params)
		for _, merchantLoad := range merchantLoadList {
			if MerchantAbleAmount(merchantLoad) {
				logs.Info(fmt.Sprintf("商户uid=%s，押款金额=%f，押款通道=%s, 解款成功", merchantLoad.MerchantUid, merchantLoad.LoadAmount, merchantLoad.RoadUid))
			} else {
				logs.Error(fmt.Sprintf("商户uid=%s，押款金额=%f，押款通道=%s, 解款失败", merchantLoad.MerchantUid, merchantLoad.LoadAmount, merchantLoad.RoadUid))
			}
		}
	}
}

/*
* 对应的商户的账户可用金额进行调整操作
 */
func MerchantAbleAmount(merchantLoad models.MerchantLoadInfo) bool {
	o := orm.NewOrm()

	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		tmpLoad := new(models.MerchantLoadInfo)
		if err := txOrm.Raw("select * from merchant_load_info where merchant_uid=? and road_uid=? and load_date=? for update",
			merchantLoad.MerchantUid, merchantLoad.RoadUid, merchantLoad.LoadDate).QueryRow(tmpLoad); err != nil || tmpLoad == nil {
			logs.Error(fmt.Sprintf("解款操作获取商户押款信息失败，fail： %s", err))
			return err

		}
		if tmpLoad.Status != common.NO {
			logs.Error(fmt.Sprintf("押款信息merchantuid=%s，通道uid=%s， 押款日期=%s,已经解款过，不需要再进行处理了", tmpLoad.MerchantUid, tmpLoad.RoadUid, tmpLoad.LoadDate))
			return errors.New("已经解款过，不需要再进行处理了")
		}

		tmpLoad.UpdateTime = utils.GetBasicDateTime()
		tmpLoad.Status = common.YES
		if _, err := txOrm.Update(tmpLoad); err != nil {
			logs.Error(fmt.Sprintf("解款操作更新merchant load info 失败：%s", err))
			return err
		}

		accountInfo := new(models.AccountInfo)
		accountInfo.UpdateTime = utils.GetBasicDateTime()
		if accountInfo.LoanAmount >= tmpLoad.LoadAmount {
			accountInfo.LoanAmount -= tmpLoad.LoadAmount
		} else {
			accountInfo.LoanAmount = common.ZERO
		}

		if _, err := o.Update(accountInfo); err != nil {
			logs.Error(fmt.Sprintf("解款操作更新account info 失败：%s，账户uid=%s", err, accountInfo.AccountUid))
			return err
		}

		return nil

	}); err != nil {
		return false
	}
	return true
}
