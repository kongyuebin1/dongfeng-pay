package service

import (
	"legend/models/fast"
)

type AccountService struct {
	BaseService
}

func (c *AccountService) GetAccountInfo(userName string) *fast.AccountInfo {

	merchantInfo := fast.GetMerchantInfoByUserName(userName)

	accountInfo := fast.GetAccountInfo(merchantInfo.MerchantUid)

	return accountInfo
}
