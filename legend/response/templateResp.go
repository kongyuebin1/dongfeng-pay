package response

import "legend/models/legend"

type AddTemplateResp struct {
	Code int
	Msg  string
}

type TemplateListResp struct {
	Code  int                    `json:"code"`
	Msg   string                 `json:"msg"`
	Count int                    `json:"count"`
	Data  []legend.ScaleTemplate `json:"data"`
}

type TemplateAllInfoResp struct {
	AddTemplateResp
	TemplateInfo           *legend.ScaleTemplate
	AnyMoneyInfo           *legend.AnyMoney
	FixMoneyInfos          []legend.FixMoney
	PresentFixMoneyInfos   []legend.FixPresent
	PresentScaleMoneyInfos []legend.ScalePresent
}
