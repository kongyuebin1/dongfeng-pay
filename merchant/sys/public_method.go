/***************************************************
 ** @Desc : This file for 共有方法
 ** @Time : 2019.04.01 11:48
 ** @Author : Joker
 ** @File : public_method
 ** @Last Modified by : Joker
 ** @Last Modified time: 2019-11-29 11:05:28
 ** @Software: GoLand
****************************************************/
package sys

import (
	"fmt"
	"math/rand"
	"merchant/sys/enum"
	"os"
	"strconv"
	"strings"
	"time"
)

type PublicMethod struct{}

// 返回当前时间的字符串:2006-01-02 15:04:05
func (*PublicMethod) GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 返回当前时间的字符串:20060102150405
func (*PublicMethod) GetNowTimeV2() string {
	return time.Now().Format("20060102150405")
}

// 返回格式化的字符串:2006-01-02 15:04:05
func (*PublicMethod) ParseDatetime(t time.Time) string {
	f := t.Format("2006-01-02 15:04:05")
	if strings.Compare("0001-01-01 00:00:00", f) == 0 {
		f = ""
	}
	return f
}

// 比较是否是同一天
func (*PublicMethod) IsSameDay(d string) (bool, string) {
	now := time.Now()
	parse, e := time.Parse("2006-01-02 15:04:05", d)
	if e != nil {
		return false, fmt.Sprintf("比较时间：%s 格式不对, %v", d, e)
	}
	year := now.Year()-parse.Year() == 0
	mouth := strings.Compare(now.Month().String(), parse.Month().String()) == 0
	day := now.Day()-parse.Day() == 0
	if year && mouth && day {
		return true, ""
	}
	return false, fmt.Sprintf("比较时间：%s 与今天不是同一天，当天收益清零", d)
}

// 在数字、大写字母、小写字母范围内生成num位的随机字符串
func (*PublicMethod) RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rand.Intn(26)+65))
		} else {
			result = append(result, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(result, "")
}

// 生成n位随机数字字符串
func (*PublicMethod) RandomIntOfString(length int) string {
	result := make([]string, 0, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		randInt := rand.Intn(10)
		result = append(result, strconv.Itoa(randInt))
	}
	return strings.Join(result, "")
}

// @Description: 返回当前操作数据库的状态
// @Author: Joker
// @Date: 2019.04.01 14:37
// @Param: code: 状态码,msg; 状态信息:url: 跳转地址; data: json内容
// @return: Json串
func (*PublicMethod) JsonFormat(code int, data interface{}, msg string, url string) (json map[string]interface{}) {
	if code == 9 {
		json = map[string]interface{}{
			"code": code,
			"data": data,
			"msg":  msg,
			"url":  url,
		}
	} else {
		json = map[string]interface{}{
			"code": code,
			"msg":  msg,
			"url":  url,
		}
	}
	return json
}

// 返回当前操作数据库的状态:成功/失败
func (*PublicMethod) GetDatabaseStatus(code int) map[string]interface{} {
	msg := enum.FailedString
	if code == enum.SuccessFlag {
		msg = enum.SuccessString
	}
	out := make(map[string]interface{})
	out["code"] = code
	out["msg"] = msg
	return out
}

// 格式化浮点数
func (*PublicMethod) FormatFloat64ToString(f float64) string {
	if f < 0 {
		f = 0
	}
	return fmt.Sprintf("%0.2f", f)
}

// 判断文件或文件夹是否存在
// 使用os.Stat()函数返回的错误值进行判断:
// 如果返回的错误为nil,说明文件或文件夹存在
// 如果返回的错误类型
// 使用os.IsNotExist()判断为true,说明文件或文件夹不存在
// 如果返回的错误为其它类型,则不确定是否在存在
func (*PublicMethod) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
