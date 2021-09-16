package pay_for

import (
	"context"
	"errors"
	"fmt"
	"gateway/conf"
	"gateway/models/accounts"
	"gateway/models/payfor"
	"gateway/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func PayForFail(p payfor.PayforInfo) bool {

	o := orm.NewOrm()
	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		var tmpForPay payfor.PayforInfo
		if err := txOrm.Raw("select * from payfor_info where bank_order_id = ? for update", p.BankOrderId).QueryRow(&tmpForPay); err != nil || tmpForPay.PayforUid == "" {

			logs.Error("solve pay fail select fail：", err)
			return err
		}

		if tmpForPay.Status == conf.PAYFOR_FAIL || tmpForPay.Status == conf.PAYFOR_SUCCESS {
			logs.Error(fmt.Sprintf("该代付订单uid=%s，状态已经是最终结果", tmpForPay.PayforUid))
			return errors.New("状态已经是最终结果")
		}
		//更新payfor记录的状态
		tmpForPay.Status = conf.PAYFOR_FAIL
		tmpForPay.UpdateTime = utils.GetBasicDateTime()
		if _, err := txOrm.Update(&tmpForPay); err != nil {
			logs.Error("PayForFail update payfor_info fail: ", err)
			return err
		}

		var account accounts.AccountInfo
		if err := txOrm.Raw("select * from account_info where account_uid = ? for update", tmpForPay.MerchantUid).QueryRow(&account); err != nil || account.AccountUid == "" {

			logs.Error("payfor select account fail：", err)
			return err
		}

		account.UpdateTime = utils.GetBasicDateTime()
		if account.PayforAmount < tmpForPay.PayforTotalAmount {
			logs.Error(fmt.Sprintf("商户uid=%s，账户中待代付金额小于代付记录的金额", tmpForPay.MerchantUid))
			return errors.New("账户中待代付金额小于代付记录的金额")
		}
		//将正在打款中的金额减去
		account.PayforAmount = account.PayforAmount - tmpForPay.PayforTotalAmount

		if _, err := txOrm.Update(&account); err != nil {
			logs.Error("PayForFail update account fail: ", err)
			return err
		}

		return nil

	}); err != nil {
		return false
	}
	return true
}

func PayForSuccess(p payfor.PayforInfo) bool {
	o := orm.NewOrm()

	if err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		var tmpPayFor payfor.PayforInfo
		if err := txOrm.Raw("select * from payfor_info where bank_order_id = ? for update", p.BankOrderId).QueryRow(&tmpPayFor); err != nil || tmpPayFor.PayforUid == "" {
			logs.Error("payfor success select payfor fail：", err)
			return err
		}
		if tmpPayFor.Status == conf.PAYFOR_FAIL || tmpPayFor.Status == conf.PAYFOR_SUCCESS {
			logs.Error(fmt.Sprintf("该代付订单uid=#{payFor.PayforUid}，已经是最终结果，不需要处理"))
			return errors.New("已经是最终结果，不需要处理")
		}

		tmpPayFor.UpdateTime = utils.GetBasicDateTime()
		tmpPayFor.Status = conf.PAYFOR_SUCCESS
		_, err := txOrm.Update(&tmpPayFor)
		if err != nil {
			logs.Error("PayForSuccess update payfor fail: ", err)
			return err
		}

		var account accounts.AccountInfo
		if err := txOrm.Raw("select * from account_info where account_uid = ? for update", tmpPayFor.MerchantUid).QueryRow(&account); err != nil || account.AccountUid == "" {
			logs.Error("payfor success select account fail：", err)
			return err
		}

		account.UpdateTime = utils.GetBasicDateTime()
		if account.PayforAmount < tmpPayFor.PayforTotalAmount {
			logs.Error(fmt.Sprintf("商户uid=#{payFor.MerchantUid}，账户中待代付金额小于代付记录的金额"))
			return errors.New("账户中待代付金额小于代付记录的金额")
		}

		//代付打款中的金额减去
		account.PayforAmount = account.PayforAmount - tmpPayFor.PayforTotalAmount
		//减去余额，减去可用金额
		account.Balance = account.Balance - tmpPayFor.PayforTotalAmount
		//已结算金额减去
		account.SettleAmount = account.SettleAmount - tmpPayFor.PayforTotalAmount

		if _, err := txOrm.Update(&account); err != nil {
			logs.Error("PayForSuccess update account fail：", err)
			return err
		}

		//添加一条动账记录
		accountHistory := accounts.AccountHistoryInfo{
			AccountUid:  tmpPayFor.MerchantUid,
			AccountName: tmpPayFor.MerchantName,
			Type:        conf.SUB_AMOUNT,
			Amount:      tmpPayFor.PayforTotalAmount,
			Balance:     account.Balance,
			UpdateTime:  utils.GetBasicDateTime(),
			CreateTime:  utils.GetBasicDateTime(),
		}

		if _, err := txOrm.Insert(&accountHistory); err != nil {
			logs.Error("PayForSuccess insert account history fail: ", err)
			return err
		}

		return nil
	}); err != nil {
		return false
	}

	return true
}
