package utils

import (
	"github.com/beego/beego/v2/core/logs"
	"strconv"
	"strings"
)

func StringToFloats(s string) []float64 {

	fs := make([]float64, 0)
	if s == "" || len(s) == 0 {
		return fs
	}
	str := strings.Split(s, ",")
	for i := 0; i < len(str); i++ {
		s := str[i]
		logs.Debug("string to float：", s)
		if f, err := strconv.ParseFloat(s, 64); err != nil {
			logs.Error("string to float64 err：", err)
			fs = append(fs, 0)
		} else {
			fs = append(fs, f)
		}
	}

	return fs
}

func StringToInt(s string) []int {
	is := make([]int, 0)
	if s == "" || len(s) == 0 {
		return is
	}

	ss := strings.Split(s, ",")
	for i := 0; i < len(ss); i++ {
		if a, err := strconv.Atoi(ss[i]); err != nil {
			logs.Error("string to int err：", err)
			is = append(is, 0)
		} else {
			is = append(is, a)
		}
	}

	return is
}
