/***************************************************
 ** @Desc : 上有支付公司的编号
 ** @Time : 2019/10/28 10:47
 ** @Author : yuebin
 ** @File : supplier
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/28 10:47
 ** @Software: GoLand
****************************************************/
package common

//添加新的上游通道时，需要添加这里
var supplierCode2Name = map[string]string{
	"KF":     "快付支付",
	"WEIXIN": "官方微信",
	"ALIPAY": "官方支付宝",
	"DAILI":  "代丽支付",
}

func GetSupplierMap() map[string]string {
	return supplierCode2Name
}
