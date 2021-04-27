/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/19 14:17
 ** @Author : yuebin
 ** @File : transaction
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/19 14:17
 ** @Software: GoLand
****************************************************/
package models

import (
	"boss/common"
	"boss/utils"
	"context"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func OperatorAccount(accountUid, operatorType string, amount float64) (string, bool) {
	o := orm.NewOrm()

	msg := ""
	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		//处理事务
		accountInfo := new(AccountInfo)
		if err := txOrm.Raw("select * from account_info where account_uid = ? for update", accountUid).QueryRow(accountInfo); err != nil || accountInfo.AccountUid == "" {
			logs.Error("operator account get account info for update fail: ", err)
			return err
		}

		accountInfo.UpdateTime = utils.GetBasicDateTime()
		flag := true

		switch operatorType {
		case common.PLUS_AMOUNT: //处理加款操作
			accountInfo.Balance = accountInfo.Balance + amount
			accountInfo.SettleAmount = accountInfo.SettleAmount + amount
		case common.SUB_AMOUNT: //处理减款
			if accountInfo.Balance >= amount && accountInfo.SettleAmount >= amount {
				accountInfo.Balance = accountInfo.Balance - amount
				accountInfo.SettleAmount = accountInfo.SettleAmount - amount
			} else {
				msg = "账户余额不够减"
				flag = false
			}
		case common.FREEZE_AMOUNT: //处理冻结款
			accountInfo.FreezeAmount = accountInfo.FreezeAmount + amount
		case common.UNFREEZE_AMOUNT: //处理解冻款
			if accountInfo.FreezeAmount >= amount {
				accountInfo.FreezeAmount = accountInfo.FreezeAmount - amount
			} else {
				msg = "账户冻结金额不够解冻款"
				flag = false
			}
		}
		if !flag {
			return errors.New("处理失败")
		}

		if _, err := txOrm.Update(accountInfo); err != nil {
			logs.Error("operator account update account fail: ", err)
			return err
		}
		//往account_history表中插入一条动账记录
		accountHistory := AccountHistoryInfo{AccountUid: accountUid, AccountName: accountInfo.AccountName, Type: operatorType,
			Amount: amount, Balance: accountInfo.Balance, CreateTime: utils.GetBasicDateTime(), UpdateTime: utils.GetBasicDateTime()}

		if _, err := txOrm.Insert(&accountHistory); err != nil {
			logs.Error("operator account insert account history fail: ", err)
			return err
		}

		return nil
	}); err != nil {
		return msg, false
	}

	return msg, true
}
