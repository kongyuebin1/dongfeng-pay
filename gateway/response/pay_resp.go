package response

import (
	"gateway/models/merchant"
	"gateway/models/road"
)

type PayBaseResp struct {
	Params       map[string]string     //请求的基本参数
	ClientIp     string                //商户ip
	MerchantInfo merchant.MerchantInfo //商户信息
	Msg          string                //信息
	Code         int                   //状态码 200正常
	RoadInfo     road.RoadInfo
	RoadPoolInfo road.RoadPoolInfo
	OrderAmount  float64
	PayWayCode   string
	PlatformRate float64
	AgentRate    float64
}

type ScanSuccessData struct {
	OrderNo    string `json:"orderNo"`
	Sign       string `json:"sign"`
	OrderPrice string `json:"orderPrice"`
	PayKey     string `json:"payKey"`
	PayUrl     string `json:"payURL"`
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}

type ScanFailData struct {
	PayKey     string `json:"payKey"`
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}
