/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/29 15:01
 ** @Author : yuebin
 ** @File : pay_way_code
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/29 15:01
 ** @Software: GoLand
****************************************************/
package common

var ScanPayWayCodes = []string{
	"WEIXIN_SCAN",
	"UNION_SCAN",
	"ALI_SCAN",
	"BAIDU_SCAN",
	"JD_SCAN",
	"QQ_SCAN",
}

var H5PayWayCodes = []string{
	"WEIXIN_H5",
	"ALI_H5",
	"QQ_H5",
	"UNION_H5",
	"BAIDU_H5",
	"JD_H5",
}

var SytPayWayCodes = []string{
	"WEIXIN_SYT",
	"ALI_SYT",
	"QQ_SYT",
	"UNION_SYT",
	"BAIDU_SYT",
	"JD_SYT",
}

var FastPayWayCodes = []string{
	"UNION-FAST",
}

var WebPayWayCode = []string{
	"UNION-WAP",
}

func GetScanPayWayCodes() []string {
	return ScanPayWayCodes
}

func GetNameByPayWayCode(code string) string {
	switch code {
	case "WEIXIN_SCAN":
		return "微信扫码"
	case "UNION_SCAN":
		return "银联扫码"
	case "ALI_SCAN":
		return "支付宝扫码"
	case "BAIDU_SCAN":
		return "百度扫码"
	case "JD_SCAN":
		return "京东扫码"
	case "QQ_SCAN":
		return "QQ扫码"

	case "WEIXIN_H5":
		return "微信H5"
	case "UNION_H5":
		return "银联H5"
	case "ALI_H5":
		return "支付宝H5"
	case "BAIDU_H5":
		return "百度H5"
	case "JD_H5":
		return "京东H5"
	case "QQ_H5":
		return "QQ-H5"

	case "WEIXIN_SYT":
		return "微信收银台"
	case "UNION_SYT":
		return "银联收银台"
	case "ALI_SYT":
		return "支付宝收银台"
	case "BAIDU_SYT":
		return "百度收银台"
	case "JD_SYT":
		return "京东收银台"
	case "QQ_SYT":
		return "QQ-收银台"

	case "UNION_FAST":
		return "银联快捷"
	case "UNION_WAP":
		return "银联web"
	default:
		return "未知"
	}
}
