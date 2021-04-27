/***************************************************
 ** @Desc : This file for 银行编码
 ** @Time : 19.12.4 10:42
 ** @Author : Joker
 ** @File : bank_info
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.4 10:42
 ** @Software: GoLand
****************************************************/
package enum

const (
	ICBC   = "ICBC"
	ABC    = "ABC"
	BOC    = "BOC"
	CCB    = "CCB"
	BOCOM  = "BOCOM"
	CNCB   = "CNCB"
	CEB    = "CEB"
	HXB    = "HXB"
	CMBC   = "CMBC"
	GDB    = "GDB"
	CMB    = "CMB"
	CIB    = "CIB"
	SPDB   = "SPDB"
	PSBC   = "PSBC"
	PAB    = "PAB"
	NJCB   = "NJCB"
	NBCB   = "NBCB"
	WZCB   = "WZCB"
	CSCB   = "CSCB"
	CZCB   = "CZCB"
	CCQTGB = "CCQTGB"
	SHRCB  = "SHRCB"
	BJRCB  = "BJRCB"
	SDB    = "SDB"
)

var bankInfo = map[string]string{
	ICBC:   "中国工商银行",
	ABC:    "中国农业银行",
	BOC:    "中国银行",
	CCB:    "中国建设银行",
	BOCOM:  "交通银行",
	CNCB:   "中信银行",
	CEB:    "中国光大银行",
	HXB:    "华夏银行",
	CMBC:   "中国民生银行",
	GDB:    "广发银行",
	CMB:    "招商银行",
	CIB:    "兴业银行",
	SPDB:   "浦发银行",
	PSBC:   "中国邮政储蓄银行",
	PAB:    "平安银行",
	NJCB:   "南京银行",
	NBCB:   "宁波银行",
	WZCB:   "温州市商业银行",
	CSCB:   "长沙银行",
	CZCB:   "浙江稠州商业银行",
	CCQTGB: "重庆三峡银行",
	SHRCB:  "上海农村商业银行",
	BJRCB:  "北京农商行",
	SDB:    "深圳发展银行",
}

func GetBankInfo() map[string]string {
	return bankInfo
}
