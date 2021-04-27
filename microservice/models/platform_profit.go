/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/12/17 17:50
 ** @Author : yuebin
 ** @File : platform_profit
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/12/17 17:50
 ** @Software: GoLand
****************************************************/
package models

type PlatformProfit struct {
	MerchantName   string
	AgentName      string
	SupplierName   string
	PayTypeName    string
	OrderAmount    float64
	OrderCount     int
	PlatformProfit float64
	AgentProfit    float64
}
