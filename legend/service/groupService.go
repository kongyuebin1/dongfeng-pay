package service

import (
	"github.com/rs/xid"
	"legend/models/legend"
	"legend/response"
	"legend/utils"
)

type GroupService struct {
	BaseService
}

func (c *GroupService) GroupAdd(groupName string) *response.BaseResp {

	resp := new(response.BaseResp)
	resp.Code = -1

	group := new(legend.Group)

	uid := xid.New().String()
	group.GroupName = groupName
	group.Uid = uid
	group.CreateTime = utils.GetNowTime()
	group.UpdateTime = utils.GetNowTime()

	if legend.InsertGroup(group) {
		resp.Code = 0
		resp.Msg = "添加分组成功"
	} else {
		resp.Msg = "添加分组失败"
	}

	return resp
}

func (c *GroupService) GroupList(page, limit int) *response.GroupListResp {

	offset := utils.CountOffset(page, limit)
	count := legend.GetGroupAllCont()
	groups := legend.GetGroupList(offset, limit)

	for i, _ := range groups {
		groups[i].Id = offset + i + 1
	}

	groupListResp := new(response.GroupListResp)
	groupListResp.Count = count
	groupListResp.Code = 0
	groupListResp.Data = groups

	return groupListResp
}

func (c *GroupService) GroupDelete(uid string) *response.BaseResp {
	resp := new(response.BaseResp)
	resp.Code = 0
	if legend.DeleteGroupByUid(uid) {
		resp.Msg = "删除成功"
	} else {
		resp.Msg = "删除分组信息失败"
		resp.Code = -1
	}

	return resp
}

func (c *GroupService) GroupEdit(uid, groupName string) *response.BaseResp {
	resp := new(response.BaseResp)
	resp.Code = -1

	group := legend.GetGroupByUid(uid)
	if group == nil || group.Uid == "" {
		resp.Msg = "不存在这样的分组信息"
	} else {
		group.UpdateTime = utils.GetNowTime()
		group.GroupName = groupName
		if legend.UpdateGroup(group) {
			resp.Code = 0
			resp.Msg = "更新成功"
		} else {
			resp.Msg = "更新失败"
		}
	}

	return resp
}
