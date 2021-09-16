/***************************************************
 ** @Desc : 注册上游支付接口
 ** @Time : 2019/10/28 14:48
 ** @Author : yueBin
 ** @File : init
 ** @Last Modified by : yueBin
 ** @Last Modified time: 2019/10/28 14:48
 ** @Software: GoLand
****************************************************/
package third_party

import (
	"gateway/supplier"
	"github.com/beego/beego/v2/core/logs"
)

//添加新的上游通道时，需要添加这里
var supplierCode2Name = map[string]string{
	"KF":     "快付支付",
	"WEIXIN": "官方微信",
	"ALIPAY": "官方支付宝",
	"DAILI":  "代丽支付",
}

var registerSupplier = make(map[string]supplier.PayInterface)

//注册各种上游的支付接口

func init() {
	registerSupplier["KF"] = new(KuaiFuImpl)
	logs.Notice(CheckSupplierByCode("KF"))

	registerSupplier["DAILI"] = new(DaiLiImpl)
	logs.Notice(CheckSupplierByCode("DAILI"))
}

func GetPaySupplierByCode(code string) supplier.PayInterface {
	return registerSupplier[code]
}

func CheckSupplierByCode(code string) string {
	for k, v := range supplierCode2Name {
		if k == code {
			return v + "，注册完毕"
		}
	}
	return "未找到上游名称，注册有问题。"
}
