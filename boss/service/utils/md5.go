/***************************************************
 ** @Desc : 获取一个md5的字符串
 ** @Time : 2019/8/9 16:06
 ** @Author : yuebin
 ** @File : md5
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/9 16:06
 ** @Software: GoLand
****************************************************/
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

/*
* 获取小写的MD5
 */
func GetMD5LOWER(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

/*
* 获取大写的MD5
 */
func GetMD5Upper(s string) string {
	return strings.ToUpper(GetMD5LOWER(s))
}
