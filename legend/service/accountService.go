package service

import "legend/models/fast"

type AccountService struct {
	BaseService
}

func (c *AccountService) GetAccountInfo(userName string) *fast.RpAccount {
	userInfo := fast.GetUserInfoByUserName(userName)

	accountInfo := fast.GetAccontInfo(userInfo.UserNo)

	return accountInfo
}
