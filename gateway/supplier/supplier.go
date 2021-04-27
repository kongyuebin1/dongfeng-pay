/***************************************************
 ** @Desc : 上有支付公司的编号
 ** @Time : 2019/10/28 10:47
 ** @Author : yuebin
 ** @File : supplier
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/28 10:47
 ** @Software: GoLand
****************************************************/
package controller

//添加新的上游通道时，需要添加这里
var supplierCode2Name = map[string]string{
	"KF":     "快付支付",
	"WEIXIN": "官方微信",
	"ALIPAY": "官方支付宝",
}

func GetSupplierMap() map[string]string {
	return supplierCode2Name
}

func GetSupplierCodes() []string {
	var supplierCodes []string
	for k := range supplierCode2Name {
		supplierCodes = append(supplierCodes, k)
	}

	return supplierCodes
}

func GetSupplierNames() []string {
	var supplierNames []string
	for _, v := range supplierCode2Name {
		supplierNames = append(supplierNames, v)
	}
	return supplierNames
}

func CheckSupplierByCode(code string) string {
	for k, v := range supplierCode2Name {
		if k == code {
			return v + "，注册完毕"
		}
	}
	return "未找到上游名称，注册有问题。"
}
