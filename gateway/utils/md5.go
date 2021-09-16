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

/**
** 将map数据变成key=value形式的字符串
 */
func MapToString(m map[string]string) string {

	res := ""
	for k, v := range m {
		res = res + k + "=" + v + "&"
	}

	suffix := strings.TrimSuffix(res, "&")

	return suffix
}
