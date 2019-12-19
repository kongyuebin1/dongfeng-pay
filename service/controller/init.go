/***************************************************
 ** @Desc : 注册上游支付接口
 ** @Time : 2019/10/28 14:48
 ** @Author : yuebin
 ** @File : init
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/28 14:48
 ** @Software: GoLand
****************************************************/
package controller

import (
	"github.com/astaxie/beego/logs"
	"juhe/service/common"
)

var registerSupplier = make(map[string]PayInterface)

//注册各种上游的支付接口

func Init() {
	registerSupplier["KF"] = new(KuaiFuImpl)
	logs.Notice(common.CheckSupplierByCode("KF"))
	registerSupplier["WEIXIN"] = new(WeiXinImpl)
	logs.Notice(common.CheckSupplierByCode("WEIXIN"))
	registerSupplier["ALIPAY"] = new(AlipayImpl)
	logs.Notice(common.CheckSupplierByCode("ALIPAY"))
}

func GetPaySupplierByCode(code string) PayInterface {
	return registerSupplier[code]
}
