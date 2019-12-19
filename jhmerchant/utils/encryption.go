/***************************************************
 ** @Desc : This file for 加密、解密方法
 ** @Time : 2018.12.28 14:10
 ** @Author : Joker
 ** @File : encryption
 ** @Last Modified by : Joker
 ** @Last Modified time: 2019-11-30 10:19:33
 ** @Software: GoLand
****************************************************/
package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

type Encrypt struct{}

//将字符串加密成 md5
func (*Encrypt) EncodeMd5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return hex.EncodeToString(hash.Sum(nil))
}

//base64编码
func (*Encrypt) Base64Encode(raw []byte) string {
	t := base64.StdEncoding.EncodeToString(raw)
	t = strings.TrimSpace(t)
	t = strings.Replace(t, "\r", "", -1)
	t = strings.Replace(t, "\n", "", -1)
	t = strings.Replace(t, "\n\r", "", -1)
	t = strings.Replace(t, "\r\n", "", -1)
	return t
}

//base64解码
func (*Encrypt) Base64Decode(raw string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(raw)
}
