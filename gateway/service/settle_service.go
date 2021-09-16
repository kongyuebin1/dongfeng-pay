/***************************************************
 ** @Desc : 订单结算，将订单上面的钱加入到账户余额中
 ** @Time : 2019/11/22 11:34
 ** @Author : yuebin
 ** @File : order_settle
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/11/22 11:34
 ** @Software: GoLand
****************************************************/
package service

import (
	"context"
	"errors"
	"fmt"
	"gateway/conf"
	"gateway/models/accounts"
	"gateway/models/merchant"
	"gateway/models/order"
	"gateway/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

const (
	Interval = 2 //隔多少分钟进行结算
	Minutes  = 1 //每隔15分钟，进行扫码，看有没有隔天押款金额
)

/**
* 订单结算，将那些支付成功的订单金额加入到商户账户的结算金额中
 */
func OrderSettle() {

	params := make(map[string]string)
	params["is_allow_settle"] = conf.YES
	params["is_complete_settle"] = conf.NO
	orderSettleList := order.GetOrderSettleListByParams(params)
	for _, orderSettle := range orderSettleList {
		orderProfitInfo := order.GetOrderProfitByBankOrderId(orderSettle.BankOrderId)
		if !settle(orderSettle, orderProfitInfo) {
			logs.Error(fmt.Sprintf("结算订单bankOrderId = #{orderSettle.BankOrderId}， 执行失败"))
		} else {
			logs.Info(fmt.Sprintf("结算订单bankOrderId= #{orderSettle.BankOrderId}，执行成功"))
		}
	}
}

func settle(orderSettle order.OrderSettleInfo, orderProfit order.OrderProfitInfo) bool {
	o := orm.NewOrm()

	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		tmpSettle := new(order.OrderSettleInfo)
		if err := txOrm.Raw("select * from order_settle_info where bank_order_id=? for update", orderSettle.BankOrderId).QueryRow(tmpSettle); err != nil || tmpSettle.BankOrderId == "" {
			logs.Error("获取tmpSettle失败，bankOrderId=%s", orderSettle.BankOrderId)
			return err
		}

		tmpSettle.UpdateTime = utils.GetBasicDateTime()
		tmpSettle.IsCompleteSettle = conf.YES
		if _, err := txOrm.Update(tmpSettle); err != nil {
			logs.Error("更新tmpSettle失败，错误：", err)
			return err
		}

		accountInfo := new(accounts.AccountInfo)
		if err := txOrm.Raw("select * from account_info where account_uid=? for update", orderSettle.MerchantUid).QueryRow(accountInfo); err != nil || accountInfo.UpdateTime == "" {
			logs.Error("结算select account info失败，错误信息：", err)
			return err
		}
		accountInfo.UpdateTime = utils.GetBasicDateTime()

		// 商户有押款操作
		loadAmount := 0.0
		merchantDeployInfo := merchant.GetMerchantDeployByUidAndPayType(accountInfo.AccountUid, orderSettle.PayTypeCode)
		if merchantDeployInfo.IsLoan == conf.YES {
			loadAmount = merchantDeployInfo.LoanRate * 0.01 * orderProfit.FactAmount
			date := utils.GetDate()
			params := make(map[string]string)
			params["merchant_uid"] = tmpSettle.MerchantUid
			params["road_uid"] = tmpSettle.RoadUid
			params["load_date"] = date
			if !merchant.IsExistMerchantLoadByParams(params) {

				tmp := merchant.MerchantLoadInfo{Status: conf.NO, MerchantUid: orderSettle.MerchantUid, RoadUid: orderSettle.RoadUid,
					LoadDate: utils.GetDateAfterDays(merchantDeployInfo.LoanDays), LoadAmount: loadAmount,
					UpdateTime: utils.GetBasicDateTime(), CreateTime: utils.GetBasicDateTime()}

				if _, err := txOrm.Insert(&tmp); err != nil {
					logs.Error("結算插入merchantLoad失敗，失败信息：", err)
					return err
				} else {
					logs.Info("结算插入新的merchantLoad信息成功")
				}
			} else {
				merchantLoad := new(merchant.MerchantLoadInfo)
				if err := txOrm.Raw("select * from merchant_load_info where merchant_uid=? and road_uid=? and load_date=? for update").
					QueryRow(merchantLoad); err != nil || merchantLoad.UpdateTime == "" {
					logs.Error(fmt.Sprintf("结算过程，select merchant load info失败，错误信息：#{err}"))
					return err
				} else {
					merchantLoad.UpdateTime = utils.GetBasicDateTime()
					merchantLoad.LoadAmount += loadAmount
					if _, err := txOrm.Update(merchantLoad); err != nil {
						logs.Error(fmt.Sprintf("结算过程，update merchant load info失败，失败信息：#{err}"))
						return err
					}
				}
			}
		} else {
			logs.Info(fmt.Sprintf("结算过程中，该商户不需要押款，全款结算"))
		}

		if accountInfo.WaitAmount < orderProfit.UserInAmount {
			logs.Error("系统出现严重故障，账户的带结算金额小于订单结算金额")
			return errors.New("系统出现严重故障，账户的带结算金额小于订单结算金额, 账户 = " + accountInfo.AccountName + "订单id = " + orderProfit.BankOrderId)
		}

		needAmount := orderProfit.UserInAmount - loadAmount

		accountInfo.SettleAmount = accountInfo.SettleAmount + needAmount
		accountInfo.WaitAmount = accountInfo.WaitAmount - orderProfit.UserInAmount
		accountInfo.LoanAmount = accountInfo.LoanAmount + loadAmount

		if _, err := txOrm.Update(accountInfo); err != nil {
			logs.Error("结算update account 失败，错误信息：", err)
			return err
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
	merchantDeployList := merchant.GetMerchantDeployByHour(hour)
	for _, merchantDeploy := range merchantDeployList {
		logs.Info(fmt.Sprintf("开始执行商户uid= #{merchantDeploy.MerchantUid}，进行解款操作"))

		loadDate := utils.GetDateBeforeDays(merchantDeploy.LoanDays)
		params := make(map[string]string)
		params["status"] = conf.NO
		params["merchant_uid"] = merchantDeploy.MerchantUid
		params["load_date"] = loadDate

		merchantLoadList := merchant.GetMerchantLoadInfoByMap(params)
		for _, merchantLoad := range merchantLoadList {
			if MerchantAbleAmount(merchantLoad) {
				logs.Info(fmt.Sprintf("商户uid= %s，押款金额=%f，押款通道= %s, 解款成功", merchantLoad.MerchantUid, merchantLoad.LoadAmount, merchantLoad.RoadUid))
			} else {
				logs.Error(fmt.Sprintf("商户uid=%s，押款金额=%f，押款通道=%s, 解款失败", merchantLoad.MerchantUid, merchantLoad.LoadAmount, merchantLoad.RoadUid))
			}
		}
	}
}

/*
* 对应的商户的账户可用金额进行调整操作
 */
func MerchantAbleAmount(merchantLoad merchant.MerchantLoadInfo) bool {
	o := orm.NewOrm()

	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		tmpLoad := new(merchant.MerchantLoadInfo)
		if err := txOrm.Raw("select * from merchant_load_info where merchant_uid=? and road_uid=? and load_date=? for update",
			merchantLoad.MerchantUid, merchantLoad.RoadUid, merchantLoad.LoadDate).QueryRow(tmpLoad); err != nil || tmpLoad.MerchantUid == "" {
			logs.Error(fmt.Sprintf("解款操作获取商户押款信息失败，fail： %s", err))
			return err

		}
		if tmpLoad.Status != conf.NO {
			logs.Error(fmt.Sprintf("押款信息merchantuid=%s，通道uid=%s， 押款日期=%s,已经解款过，不需要再进行处理了", tmpLoad.MerchantUid, tmpLoad.RoadUid, tmpLoad.LoadDate))
			return errors.New("已经解款过，不需要再进行处理了")
		}

		tmpLoad.UpdateTime = utils.GetBasicDateTime()
		tmpLoad.Status = conf.YES
		if _, err := txOrm.Update(tmpLoad); err != nil {
			logs.Error(fmt.Sprintf("解款操作更新merchant load info 失败：%s", err))
			return err
		}

		accountInfo := new(accounts.AccountInfo)
		if err := txOrm.Raw("select * from account_info where account_uid = ? for update", merchantLoad.MerchantUid).QueryRow(accountInfo); err != nil || accountInfo.AccountUid == "" {
			logs.Error("结款操作获取账户信息失败：", err)
			return err
		}
		accountInfo.UpdateTime = utils.GetBasicDateTime()
		if accountInfo.LoanAmount >= tmpLoad.LoadAmount {
			accountInfo.LoanAmount = accountInfo.LoanAmount - tmpLoad.LoadAmount
			accountInfo.SettleAmount = accountInfo.SettleAmount + tmpLoad.LoadAmount
		} else {
			accountInfo.LoanAmount = conf.ZERO
		}

		if _, err := txOrm.Update(accountInfo); err != nil {
			logs.Error(fmt.Sprintf("解款操作更新account info 失败：%s，账户uid=%s", err, accountInfo.AccountUid))
			return err
		}

		return nil

	}); err != nil {
		return false
	}
	return true
}

func OrderSettleInit() {
	//每隔5分钟，巡查有没有可以进行结算的订单
	go func() {
		settleTimer := time.NewTimer(time.Duration(Interval) * time.Minute)
		oneMinuteTimer := time.NewTimer(time.Duration(Minutes) * time.Minute)
		for {
			select {
			case <-settleTimer.C:
				settleTimer = time.NewTimer(time.Duration(Interval) * time.Minute)
				logs.Info("开始对商户进行支付订单结算>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
				OrderSettle()
			case <-oneMinuteTimer.C:
				oneMinuteTimer = time.NewTimer(time.Duration(Minutes) * time.Minute)
				logs.Info("开始执行商户的解款操作>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
				MerchantLoadSolve()
			}
		}
	}()
}
