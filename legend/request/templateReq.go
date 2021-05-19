package request

type AddTemplateReq struct {
	ScaleTemplateName       string    `form:"scaleTemplateName"`
	ScaleUserName           string    `form:"scaleUserName"`
	ScaleUserNamePoint      string    `form:"scaleUserNamePoint"`
	MoneyType               string    `form:"moneyType"`
	GameMoneyName           string    `form:"gameMoneyName"`
	GameMoneyScale          int       `form:"gameMoneyScale"`
	LimitLowMoney           float64   `form:"limitLowMoney"`
	PresentType             string    `form:"presentType"`
	FixPrices               []float64 `form:"fixPrices"`
	GoodsNames              []string  `form:"goodsNames"`
	GoodsNos                []string  `form:"goodsNos"`
	Limits                  []int     `form:"limits"`
	PresentFixMoneys        []float64 `form:"presentFixMoneys"`
	PresentFixPresentMoneys []float64 `form:"presentFixPresentMoneys"`
	PresentScaleMoneys      []float64 `form:"presentScaleMoneys"`
	PresentScales           []float64 `form:"presentScales"`
	FixUids                 []string  `form:"fixUids"`
	PresentScaleUids        []string  `form:"presentScaleUids"`
	PresentFixUids          []string  `form:"presentFixUids"`
}
