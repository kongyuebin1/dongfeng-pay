/***************************************************
 ** @Desc : This file for 枚举
 ** @Time : 2018-7-26 10:13
 ** @Author : Joker
 ** @File : enums.go
 ** @Last Modified by : Joker
 ** @Last Modified time: 2018-08-30 16:32:33
 ** @Software: GoLand
****************************************************/
package enums

/*支付方式*/
var paySubType = map[string]string{
	"":            "所有",
	"WEIXIN_SCAN": "微信扫码",
	"UNION_SCAN":  "银联扫码",
	"ALI_SCAN":    "支付宝扫码",

	"WEIXIN_H5": "微信H5",
	"ALI_H5":    "支付宝H5",

	"UNION_FAST": "银联快捷",
}

func GetPaySubType() map[string]string {
	return paySubType
}

/*银行编码*/
var bankCode = map[string]string{
	"01020000": "ICBC",  //工商银行
	"01030000": "ABC",   //农业银行
	"01040000": "BOC",   //中国银行
	"01050000": "CCB",   //建设银行
	"03010000": "BOCOM", //交通银行
	"03020000": "CNCB",  //中信银行
	"03030000": "CEB",   //中信银行
	"03040000": "HXB",   //光大银行
	"03050000": "CMBC",  //民生银行
	"03060000": "GDB",   //广发银行
	"04100000": "PAB",   //平安银行
	"03080000": "CMB",   //招商银行
	"03090000": "CIB",   //兴业银行
	"03170000": "BOHC",  //渤海银行
	"03200000": "BEAI",  //东亚银行
	"04012900": "BOS",   //上海银行
	"04031000": "BCCB",  //北京银行
	"04083320": "NBBC",  //宁波银行
	"04243010": "NJBC",  //南京银行
	"64296510": "CDSBC", //成都银行
}
