package service

import (
	"legend/models/fast"
	"legend/utils"
)

type AccountService struct {
	BaseService
}

func (c *AccountService) GetAccountInfo(userName string) *fast.AccountInfo {

	merchantInfo := fast.GetMerchantInfoByUserName(userName)

	accountInfo := fast.GetAccountInfo(merchantInfo.MerchantUid)

	return accountInfo
}

/**
** 获取当天的充值金额
 */
func (c *AccountService) GetTodayIncome() float64 {
	startTime := utils.GetNowDate() + " 00:00:00"
	endTime := utils.GetNowDate() + " 23:59:59"

	todayIncome := fast.GetRangeDateIncome(startTime, endTime)

	return todayIncome
}
