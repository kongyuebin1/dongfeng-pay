package service

import (
	"github.com/rs/xid"
	"legend/models/legend"
	"legend/request"
	"legend/response"
	"legend/utils"
)

type AreaService struct {
	BaseService
}

func (c *AreaService) AddArea(req *request.AreaReq) *response.BaseResp {
	resp := new(response.BaseResp)

	area := new(legend.Area)
	area.AreaName = req.AreaName
	area.GroupName = req.GroupName
	area.TemplateName = req.TemplateName
	area.NotifyUrl = req.NotifyUrl
	area.AttachParams = req.AttachParams
	area.UpdateTime = utils.GetNowTime()
	area.CreateTime = utils.GetNowTime()
	area.Uid = xid.New().String()

	if legend.InsertArea(area) {
		resp.Code = 0
	} else {
		resp.Code = -1
		resp.Msg = "添加分区失败"
	}

	return resp
}

func (c *AreaService) AreaList(page, limit int) *response.AreaListResp {
	offset := utils.CountOffset(page, limit)
	count := legend.GetAreaAllCount()
	areas := legend.GetAreaList(offset, limit)

	for i, _ := range areas {
		areas[i].Id = offset + i + 1
	}

	areaResp := new(response.AreaListResp)
	areaResp.Code = 0
	areaResp.Count = count
	areaResp.Data = areas

	return areaResp
}

func (c *AreaService) DeleteArea(uid string) *response.BaseResp {
	resp := new(response.BaseResp)

	if legend.DeleteAreaByUid(uid) {
		resp.Code = 0
	} else {
		resp.Code = -1
		resp.Msg = "删除分区失败"
	}
	return resp
}

func (c *AreaService) GetArea(uid string) *response.AreaInfoResp {
	resp := new(response.AreaInfoResp)
	resp.Code = 0
	resp.Msg = "请求成功"

	area := legend.GetAreaByUid(uid)
	resp.Area = area

	return resp
}

func (c *AreaService) EditArea(req *request.AreaReq, uid string) *response.BaseResp {
	resp := new(response.BaseResp)
	resp.Code = -1

	area := legend.GetAreaByUid(uid)
	if area == nil || area.AreaName == "" {
		resp.Msg = "更新失败"
	} else {
		area.UpdateTime = utils.GetNowTime()
		area.GroupName = req.GroupName
		area.TemplateName = req.TemplateName
		area.NotifyUrl = req.NotifyUrl
		area.AttachParams = req.AttachParams

		if legend.UpdateArea(area) {
			resp.Msg = "更新失败"
		}
	}

	return resp
}
