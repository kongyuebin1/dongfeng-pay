package gateway

import (
	"encoding/json"
	"fmt"
	"gateway/utils"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestPayFor(t *testing.T) {
	u := "http://127.0.0.1:12309/gateway/payfor"
	m := make(map[string]string)
	m["merchantKey"] = "kkkkc254gk8isf001cqrj6p0"
	m["realname"] = "11"
	m["cardNo"] = "123"
	m["accType"] = "private"
	m["amount"] = "0.1"
	merchantOrderId := GenerateOrderNo()
	m["merchantOrderId"] = merchantOrderId
	sec := "ssssc254gk8isf001cqrj6pg"
	keys := utils.SortMap(m)
	sign := utils.GetMD5Sign(m, keys, sec)
	m["sign"] = sign

	m1 := make(map[string]interface{})
	marshal, _ := json.Marshal(m)
	json.Unmarshal(marshal, &m1)
	req := new(utils.Request)
	req.SetParams(m1)
	req.SetURL(u)
	resp, err := req.GET()
	if err != nil {
		t.Fatal("err------>", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)

		t.Fatal("返回的http状态码不是200,body:", string(b))
	}
	b, _ := ioutil.ReadAll(resp.Body)
	log.Println("result--------->", string(b))
	/**
	  result---------> {
	    "resultCode": "00",
	    "resultMsg": "银行处理中",
	    "settAmount": "0.1"
	  }
	*/
}

func GetTimeTick64() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetFormatTime(time time.Time) string {
	return time.Format("20060102")
}

// 基础做法 日期20191025时间戳1571987125435+3位随机数
func GenerateOrderNo() string {
	date := GetFormatTime(time.Now())
	r := rand.Intn(1000)
	code := fmt.Sprintf("%s%d%03d", date, GetTimeTick64(), r)
	return code
}
func TestGenerateCode(t *testing.T) {
	GenerateOrderNo()
}
