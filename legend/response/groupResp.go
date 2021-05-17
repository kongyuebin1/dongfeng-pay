package response

import "legend/models/legend"

type GroupListResp struct {
	Code  int            `json:"code"`
	Msg   string         `json:"msg"`
	Count int            `json:"count"`
	Data  []legend.Group `json:"data"`
}
