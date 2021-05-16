package service

import (
	"github.com/beego/beego/v2/core/logs"
	"legend/models/fast"
)

type MerchantService struct {
	BaseService
}

func (c *MerchantService) GetMerchantBankInfo(mobile string) (*fast.MerchantInfo, *fast.BankCardInfo) {

	merchantInfo := fast.GetMerchantInfoByUserName(mobile)
	bankInfo := fast.GetBankCardInfoByUserNo(merchantInfo.MerchantUid)

	return merchantInfo, bankInfo
}

/**
** 获取商户的密钥等信息
 */
/*func (c *MerchantService) UserPayConfig(userName string) map[string]string {

	merchantMapData := make(map[string]string)

	userInfo := fast.GetMerchantInfoByUserName(userName)

	if userInfo == nil || userInfo.LoginAccount == "" {
		return merchantMapData
	}

	userNo := userInfo.LoginAccount

	userPayConfig := fast.GetUserPayConfigByUserNo(userNo)
	if nil == userPayConfig || userPayConfig.UserNo == "" {
		return merchantMapData
	}

	return merchantMapData
}*/

/**
** 获取商户信息
 */
func (c *MerchantService) MerchantInfo(mobile string) *fast.MerchantInfo {
	userInfo := fast.GetMerchantInfoByUserName(mobile)
	if nil == userInfo || userInfo.LoginAccount == "" {
		logs.Error("获取用户信息失败")
	}

	//logs.Debug("用户信息：", userInfo)
	return userInfo
}
