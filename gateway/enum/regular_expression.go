/***************************************************
 ** @Desc : This file for 正则表达式
 ** @Time : 19.12.5 10:25
 ** @Author : Joker
 ** @File : regular_expression
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.5 10:25
 ** @Software: GoLand
****************************************************/
package enum

const (
	PasswordReg = `^[a-zA-Z]{1}([a-zA-Z0-9]|[._]){5,19}$`
	MoneyReg    = `^(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*))$`
	MobileReg   = `^[1]([3-9])[0-9]{9}$`
)
