package service

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"legend/common"
	"legend/models/fast"
	"legend/utils"
	"strings"
)

type LoginService struct {
	BaseService
}

type LoginJsonData struct {
	Code int
	Msg  string
}

func (c *LoginService) Login(userName, password string) *LoginJsonData {
	loginJsonData := new(LoginJsonData)
	loginJsonData.Code = 200

	userInfo := fast.GetUserInfoByUserName(userName)
	logs.Info("登录账户信息：", fmt.Sprintf("%+v", userInfo))
	if nil == userInfo || userInfo.Mobile == "" {
		logs.Error("用户不存在，账户：", userName)
		loginJsonData.Code = 404
		loginJsonData.Msg = "用户不存在"
	} else {
		if userInfo.Status == common.UNACTIVE {
			logs.Warn("账号异常，请联系管理员，账号：", userName)
			loginJsonData.Code = 503
			loginJsonData.Msg = "账户已经被冻结"
		} else {
			md5Password := utils.EncodeMd5(password)
			logs.Info("账户密码md5后：", md5Password, "；数据库保存的为：", userInfo.Password)
			if strings.ToLower(utils.EncodeMd5(password)) != strings.ToLower(userInfo.Password) {
				logs.Error("密码错误，账户:", userName)
				loginJsonData.Code = -1
				loginJsonData.Msg = "密码错误"
			} else {
				logs.Info("登录成功")
			}
		}
	}

	return loginJsonData
}

/**
** 更新用户的登录密码
 */
func (c *LoginService) PersonPassword(newPassword, oldPassword, repeatPassword, userName string) *LoginJsonData {

	logoutJsonData := new(LoginJsonData)
	logoutJsonData.Code = -1

	userInfo := fast.GetUserInfoByUserName(userName)
	if userInfo.Password != utils.EncodeMd5(oldPassword) {
		logoutJsonData.Msg = "旧密码输入不正确"
	} else if newPassword != repeatPassword {
		logoutJsonData.Msg = "2次密码不一致"
	} else {
		passwordMd5 := utils.EncodeMd5(newPassword)
		userInfo.Password = passwordMd5
		if !fast.UpdateUserInfo(userInfo) {
			logoutJsonData.Msg = "密码更新失败"
		} else {

			logoutJsonData.Code = 200
		}
	}

	return logoutJsonData
}
