package response

import "legend/models/legend"

type AreaListResp struct {
	Code  int           `json:"code"`
	Msg   string        `json:"msg"`
	Count int           `json:"count"`
	Data  []legend.Area `json:"data"`
}

type AreaInfoResp struct {
	Code int
	Msg  string
	Area *legend.Area
}
