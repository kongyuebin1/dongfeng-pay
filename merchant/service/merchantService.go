package service

import (
	"github.com/astaxie/beego/logs"
	"merchant/models/fast"
)

type MerchantService struct {
	BaseService
}

func (c *MerchantService) GetMerchantBankInfo(mobile string) (*fast.RpUserInfo, *fast.RpUserBankAccount, *fast.RpUserPayConfig) {

	userInfo := fast.GetUserInfoByUserName(mobile)
	bankInfo := fast.GetBankInfoByUserNo(userInfo.UserNo)
	userPayConfig := fast.GetUserPayConfigByUserNo(userInfo.UserNo)

	return userInfo, bankInfo, userPayConfig
}

/**
** 获取商户的密钥等信息
 */
func (c *MerchantService) UserPayConfig(userName string) map[string]string {

	merchantMapData := make(map[string]string)

	userInfo := fast.GetUserInfoByUserName(userName)

	if userInfo == nil || userInfo.Mobile == "" {
		return merchantMapData
	}

	userNo := userInfo.UserNo

	userPayConfig := fast.GetUserPayConfigByUserNo(userNo)
	if nil == userPayConfig || userPayConfig.UserNo == "" {
		return merchantMapData
	}

	return merchantMapData
}

/**
** 获取商户信息
 */
func (c *MerchantService) MerchantInfo(mobile string) *fast.RpUserInfo {
	userInfo := fast.GetUserInfoByUserName(mobile)
	if nil == userInfo || userInfo.UserNo == "" {
		logs.Error("获取用户信息失败")
	}

	//logs.Debug("用户信息：", userInfo)
	return userInfo
}
