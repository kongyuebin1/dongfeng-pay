/***************************************************
 ** @Desc : This file for 支付方式
 ** @Time : 19.12.3 15:24
 ** @Author : Joker
 ** @File : pay_type
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.3 15:24
 ** @Software: GoLand
****************************************************/
package enum

const (
	WEIXIN_SCAN  = "WEIXIN_SCAN"
	WEIXIN_H5    = "WEIXIN_H5"
	WEIXIN_SYT   = "WEIXIN_SYT"
	ALI_SCAN     = "ALI_SCAN"
	ALI_H5       = "ALI_H5"
	ALI_SYT      = "ALI_SYT"
	QQ_SCAN      = "QQ_SCAN"
	QQ_H5        = "QQ_H5"
	QQ_SYT       = "QQ_SYT"
	UNION_SCAN   = "UNION_SCAN"
	UNION_H5     = "UNION_H5"
	UNION_PC_WAP = "UNION_PC_WAP"
	UNION_SYT    = "UNION_SYT"
	UNION_FAST   = "UNION_FAST"
	BAIDU_SCAN   = "BAIDU_SCAN"
	BAIDU_H5     = "BAIDU_H5"
	BAIDU_SYT    = "BAIDU_SYT"
	JD_SCAN      = "JD_SCAN"
	JD_H5        = "JD_H5"
	JD_SYT       = "JD_SYT"
)

var payType = map[string]string{
	WEIXIN_SCAN:  "微信扫码",
	WEIXIN_H5:    "微信H5",
	WEIXIN_SYT:   "微信收银台",
	ALI_SCAN:     "支付宝扫码",
	ALI_H5:       "支付宝H5",
	ALI_SYT:      "支付宝收银台",
	QQ_SCAN:      "QQ扫码",
	QQ_H5:        "QQ-H5",
	QQ_SYT:       "QQ收银台",
	UNION_SCAN:   "银联扫码",
	UNION_H5:     "银联H5",
	UNION_PC_WAP: "银联pc-web",
	UNION_SYT:    "银联收银台",
	UNION_FAST:   "银联快捷",
	BAIDU_SCAN:   "百度钱包扫码",
	BAIDU_H5:     "百度钱包H5",
	BAIDU_SYT:    "百度钱包收银台",
	JD_SCAN:      "京东扫码",
	JD_H5:        "京东H5",
	JD_SYT:       "京东收银台",
}

func GetPayType() map[string]string {
	return payType
}
