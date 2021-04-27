/***************************************************
 ** @Desc : This file for 系统常量
 ** @Time : 19.11.30 11:28
 ** @Author : Joker
 ** @File : constant
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.11.30 11:28
 ** @Software: GoLand
****************************************************/
package enum

const (
	UserSession = "business_user"
	UserCookie  = "user_cookie_md5"

	DoMainUrl = "/index/ui/"

	PublicAccount       = "1"               // 对公帐户
	PrivateDebitAccount = "0"               // 对私借记卡
	SettleSingle        = "SELFHELP_SETTLE" // 单笔代付

	ExcelModelName    = "batch_daifa_template.xlsx"
	ExcelModelPath    = "static/excel/batch_daifa_template.xlsx"
	ExcelPath         = "static/excel/temp/"
	ExcelDownloadPath = "static/excel/download/"
)
