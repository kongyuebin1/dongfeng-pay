/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/26 11:08
 ** @Author : yuebin
 ** @File : sign_verify
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/26 11:08
 ** @Software: GoLand
****************************************************/
package utils

func GetMD5Sign(params map[string]string, keys []string, paySecret string) string {
	str := ""
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		if len(params[k]) == 0 {
			continue
		}
		str += k + "=" + params[k] + "&"
	}
	str += "paySecret=" + paySecret
	sign := GetMD5Upper(str)
	return sign
}

/*
* 验签
 */
func Md5Verify(params map[string]string, paySecret string) bool {
	sign := params["sign"]
	if sign == "" {
		return false
	}

	delete(params, "sign")
	keys := SortMap(params)
	tmpSign := GetMD5Sign(params, keys, paySecret)
	if tmpSign != sign {
		return false
	} else {
		return true
	}
}
