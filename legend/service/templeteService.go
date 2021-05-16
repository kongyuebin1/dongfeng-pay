package service

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"legend/models/legend"
	"legend/request"
	"legend/response"
	"legend/utils"
)

type TemplateService struct {
	BaseService
}

/**
** 添加比例模板
 */
func (c *TemplateService) AddTemplate(req *request.AddTemplateReq, merchantUid string) *response.AddTemplateResp {

	template := new(legend.ScaleTemplate)
	template.MerchantUid = merchantUid
	template.TemplateName = req.ScaleTemplateName
	template.UserUid = req.ScaleUserName
	template.UserWarn = req.ScaleUserNamePoint
	template.MoneyType = req.MoneyType
	template.PresentType = req.PresentType
	template.UpdateTime = utils.GetNowTime()
	template.CreateTime = utils.GetNowTime()

	addTemplateResp := new(response.AddTemplateResp)
	addTemplateResp.Code = -1

	if template.TemplateName == "" {
		addTemplateResp.Msg = "模板名称不能为空"
		return addTemplateResp
	}

	if legend.IsExistsScaleTemplateByName(template.TemplateName) {
		addTemplateResp.Msg = "模板名称重复，请换一个名称！"
		return addTemplateResp
	}

	if legend.InsertScaleTemplate(template) {
		addTemplateResp.Code = 0
		addTemplateResp.Msg = "添加比例模板成功"
		if !c.AddRandMoney(req) {
			addTemplateResp.Code = -1
			addTemplateResp.Msg = "随机金额添加失败，请检查参数是否合理"
		}
		if !c.AddFixMoney(req) {
			addTemplateResp.Code = -1
			addTemplateResp.Msg = "添加固定金额失败,请检查参数是否合理"
		}
		if !c.AddPresentFixMoney(req) {
			addTemplateResp.Code = -1
			addTemplateResp.Msg = "添加固定金额赠送失败,请检查参数是否合理"
		}
		if !c.AddPresentScaleMoney(req) {
			addTemplateResp.Code = -1
			addTemplateResp.Msg = "添加金额比例赠送失败,请检查参数是否合理"
		}

		// 只要有一个添加失败，全部删除
		if addTemplateResp.Code == -1 {
			legend.DeleteScaleTemplate(template.TemplateName)
			legend.DeleteAnyMoney(template.TemplateName)
			legend.DeleteFixMoney(template.TemplateName)
			legend.DeleteFixPresent(template.TemplateName)
			legend.DeleteScalePresent(template.TemplateName)
		}
	} else {
		addTemplateResp.Msg = "添加比例模板失败"
	}

	return addTemplateResp
}

/**
** 编辑比例模板逻辑
 */
func (c *TemplateService) UpdateTemplate(req *request.AddTemplateReq, merchantUid string) *response.BaseResp {
	resp := new(response.BaseResp)
	resp.Code = -1
	resp.Msg = "更新失败"

	template := legend.GetScaleTemplateByNameAndMerchantUid(req.ScaleTemplateName, merchantUid)
	template.UserUid = req.ScaleUserName
	template.UserWarn = req.ScaleUserNamePoint
	template.MoneyType = req.MoneyType
	template.PresentType = req.PresentType
	template.UpdateTime = utils.GetNowTime()

	if !legend.UpdateScaleTemplate(template) {
		logs.Error("更新比例模板基础数据失败")
		return resp
	}

	if !c.updateAnyMoney(req) {
		logs.Error("更新任意金额数据失败")
		return resp
	}

	if !c.updateFixMoney(req) {
		logs.Error("更新固定金额数据失败")
		return resp
	}

	if !c.updateFixPresent(req) {
		logs.Error("更新赠送固定金额失败")
		return resp
	}

	if !c.updateScalePresent(req) {
		logs.Error("更新赠送比例数据失败")
		return resp
	}

	resp.Code = 0
	resp.Msg = "更新成功"

	return resp

}

/**
** 更新任意金额
 */
func (c *TemplateService) updateAnyMoney(req *request.AddTemplateReq) bool {
	if req.GameMoneyScale <= 0 && req.LimitLowMoney <= 0 {
		logs.Debug("任意金额的2个关键参数均小于等于0")
		anyMoney := legend.GetAnyMoneyByName(req.ScaleTemplateName)
		if anyMoney != nil && anyMoney.TemplateName != "" {
			return legend.DeleteAnyMoney(req.ScaleTemplateName)
		} else {
			logs.Error("不存在这样的任意金额")
			return false
		}
	} else {
		anyMoney := legend.GetAnyMoneyByName(req.ScaleTemplateName)
		anyMoney.GameMoneyScale = req.GameMoneyScale
		anyMoney.GameMoneyName = req.GameMoneyName
		anyMoney.LimitLow = req.LimitLowMoney
		anyMoney.UpdateTime = utils.GetNowTime()
		anyMoney.CreateTime = utils.GetNowTime()

		return legend.UpdateAnyMoney(anyMoney)
	}

}

/**
** 更新固定金额
 */
func (c *TemplateService) updateFixMoney(req *request.AddTemplateReq) bool {

	for _, fixMoney := range legend.GetFixMoneyByName(req.ScaleTemplateName) {
		if !c.isExist(fixMoney.Uid, req.FixUids) {
			// 假如不存在了，那么需要删除这条记录
			legend.DeleteFixMoneyByUid(fixMoney.Uid)
		}
	}

	for i, _ := range req.FixUids {
		if req.FixUids[i] == "" {
			continue
		}

		fixMoney := legend.GetFixMoneyByUid(req.FixUids[i])

		fixMoney.UpdateTime = utils.GetNowTime()
		fixMoney.Uid = req.FixUids[i]
		fixMoney.Price = req.FixPrices[i]
		fixMoney.GoodsName = req.GoodsNames[i]
		fixMoney.GoodsNo = req.GoodsNos[i]
		fixMoney.BuyTimes = req.Limits[i]

		if fixMoney.TemplateName == "" {
			fixMoney.CreateTime = utils.GetNowTime()
			fixMoney.TemplateName = req.ScaleTemplateName
			legend.InsertFixMoney(fixMoney)
		} else {
			legend.UpdateFixMoney(fixMoney)
		}

	}

	return true
}

/**
** 更新固定金额赠送参数
 */
func (c *TemplateService) updateFixPresent(req *request.AddTemplateReq) bool {

	for _, fixPresentMoney := range legend.GetFixPresentsByName(req.ScaleTemplateName) {
		if !c.isExist(fixPresentMoney.Uid, req.PresentFixUids) {
			legend.DeleteFixPresentByUid(fixPresentMoney.Uid)
		}
	}

	for i, _ := range req.PresentFixUids {
		if req.PresentFixUids[i] == "" {
			continue
		}

		fixPresent := legend.GetFixPresentByUid(req.PresentFixUids[i])

		fixPresent.UpdateTime = utils.GetNowTime()
		fixPresent.Uid = req.PresentFixUids[i]
		fixPresent.Money = req.PresentFixMoneys[i]
		fixPresent.PresentMoney = req.PresentFixPresentMoneys[i]

		if fixPresent.TemplateName == "" {
			fixPresent.CreateTime = utils.GetNowTime()
			fixPresent.TemplateName = req.ScaleTemplateName
			legend.InsertFixPresent(fixPresent)
		} else {
			legend.UpdatePresentFixMoney(fixPresent)
		}
	}

	return true
}

/**
** 更新比例赠送参数
 */
func (c *TemplateService) updateScalePresent(req *request.AddTemplateReq) bool {

	for _, scalePresent := range legend.GetScalePresentsByName(req.ScaleTemplateName) {
		if !c.isExist(scalePresent.Uid, req.PresentScaleUids) {
			legend.DeleteScalePresentByUid(scalePresent.Uid)
		}
	}

	for i, _ := range req.PresentScaleUids {
		if req.PresentScaleUids[i] == "" {
			continue
		}

		scalePresent := legend.GetScalePresentByUid(req.PresentScaleUids[i])

		scalePresent.UpdateTime = utils.GetNowTime()
		scalePresent.Uid = req.PresentScaleUids[i]
		scalePresent.Money = req.PresentScaleMoneys[i]
		scalePresent.PresentScale = req.PresentScales[i]

		if scalePresent.TemplateName == "" {
			scalePresent.TemplateName = req.ScaleTemplateName
			scalePresent.CreateTime = utils.GetNowTime()
			legend.InsertScalePresent(scalePresent)
		} else {
			legend.UpdateScalePresent(scalePresent)
		}
	}
	return true
}

func (c *TemplateService) isExist(j string, ss []string) bool {
	for _, s := range ss {
		if s == j {
			return true
		}
	}

	return false
}

/**
** 添加随机金额
 */
func (c *TemplateService) AddRandMoney(req *request.AddTemplateReq) bool {

	if req.LimitLowMoney < 0 {
		logs.Error("随机金额中的最低充值金额不能小于0")
		return false
	}
	if req.GameMoneyName == "" && req.GameMoneyScale <= 0 {
		logs.Info("不需要添加随机金额选项")
		return true
	}

	anyMoney := new(legend.AnyMoney)
	anyMoney.TemplateName = req.ScaleTemplateName
	anyMoney.GameMoneyName = req.GameMoneyName
	anyMoney.GameMoneyScale = req.GameMoneyScale
	anyMoney.LimitLow = req.LimitLowMoney
	anyMoney.UpdateTime = utils.GetNowTime()
	anyMoney.CreateTime = utils.GetNowTime()

	if legend.InsertAnyMoney(anyMoney) {
		logs.Info("添加随机金额成功！")
	} else {
		return false
	}
	return true
}

/**
**添加固定金额
 */
func (c *TemplateService) AddFixMoney(req *request.AddTemplateReq) bool {

	l := len(req.FixUids)

	if l == 0 || (l == 1 && req.FixUids[0] == "") {
		logs.Error("该模板没有添加固定金额选项")
		return true
	}

	if l != len(req.GoodsNames) || l != len(req.FixPrices) || l != len(req.GoodsNos) || l != len(req.Limits) {
		logs.Error("固定金额参数有误，长度不一致")
		return false
	}

	for i := 0; i < l; i++ {
		fixUid := req.FixUids[i]
		fixPrice := req.FixPrices[i]
		goodName := req.GoodsNames[i]
		goodNo := req.GoodsNos[i]
		limit := req.Limits[i]

		if fixUid == "0" && fixPrice <= 0 && goodName == "0" && goodNo == "0" && limit <= 0 {
			logs.Error("固定金额4个参数都为空！")
			continue
		}

		if fixUid == "0" || fixPrice <= 0 || goodName == "0" || goodNo == "0" || limit <= 0 {
			logs.Error("固定金额参数中有一个缺失: ", fmt.Sprintf("fixUid = %s, fixPrice = %f, goodName = %s, goodNo = %s, limit = %d",
				fixUid, fixPrice, goodName, goodNo, limit))
			return false
		}

		fixMoney := new(legend.FixMoney)
		fixMoney.Uid = fixUid
		fixMoney.TemplateName = req.ScaleTemplateName
		fixMoney.Price = fixPrice
		fixMoney.GoodsName = goodName
		fixMoney.GoodsNo = goodNo
		fixMoney.BuyTimes = limit
		fixMoney.UpdateTime = utils.GetNowTime()
		fixMoney.CreateTime = utils.GetNowTime()

		if !legend.InsertFixMoney(fixMoney) {
			logs.Error("该次固定金额插入数据库失败")
		}

	}
	return true
}

/**
** 添加赠送固定金额赠送
 */
func (c *TemplateService) AddPresentFixMoney(req *request.AddTemplateReq) bool {

	l := len(req.PresentFixUids)
	if l == 0 || (l == 1 && req.PresentScaleUids[0] == "") {
		logs.Error("该模板没有添加固定金额赠送选项")
		return true
	}

	if l != len(req.PresentFixMoneys) || l != len(req.PresentFixPresentMoneys) {
		logs.Error("固定金额赠送选项参数个数不一致")
		return false
	}

	for i := 0; i < l; i++ {
		uid := req.PresentFixUids[i]
		fixMoney := req.PresentFixMoneys[i]
		presentMoney := req.PresentFixPresentMoneys[i]

		if uid == "0" && fixMoney <= 0 && presentMoney <= 0 {
			continue

		}

		if uid == "0" || fixMoney <= 0 || presentMoney <= 0 {
			logs.Error("固定金额参数中有一个缺失: ", fmt.Sprintf("fixUid = %s, fixPrice = %f, presentMoney = %f",
				uid, fixMoney, presentMoney))
			return false

		}

		fixPresent := new(legend.FixPresent)
		fixPresent.Uid = uid
		fixPresent.TemplateName = req.ScaleTemplateName
		fixPresent.Money = fixMoney
		fixPresent.PresentMoney = presentMoney
		fixPresent.UpdateTime = utils.GetNowTime()
		fixPresent.CreateTime = utils.GetNowTime()

		if !legend.InsertFixPresent(fixPresent) {
			logs.Error("该次固定金额赠送插入数据库失败")
		}
	}
	return true
}

/**
** 添加赠送金额比例
 */
func (c *TemplateService) AddPresentScaleMoney(req *request.AddTemplateReq) bool {
	l := len(req.PresentScaleUids)
	if l == 0 || (l == 1 && req.PresentScaleUids[0] == "") {
		logs.Error("该模板没有添加按百分比赠送")
		return true
	}

	if l != len(req.PresentScaleMoneys) || l != len(req.PresentScales) {
		logs.Error("按百分比赠送选项参数个数不一致")
		return false
	}

	for i := 0; i < l; i++ {
		uid := req.PresentScaleUids[i]
		money := req.PresentScaleMoneys[i]
		scale := req.PresentScales[i]

		if money <= 0 {
			logs.Error("金额不能等于0")
			return false
		}

		if uid == "0" && money <= 0 && scale <= 0 {
			continue
		}

		if uid == "0" || money <= 0 || scale <= 0 {
			logs.Error("百分比赠送缺失参数: ", fmt.Sprintf("uid = %s, money = %f, scale = %f",
				uid, money, scale))

			return false

		}

		scalePresent := new(legend.ScalePresent)
		scalePresent.Uid = uid
		scalePresent.TemplateName = req.ScaleTemplateName
		scalePresent.Money = money
		scalePresent.PresentScale = scale
		scalePresent.UpdateTime = utils.GetNowTime()
		scalePresent.CreateTime = utils.GetNowTime()

		if !legend.InsertScalePresent(scalePresent) {
			logs.Error("该次固定金额赠送插入数据库失败")
		}
	}
	return true
}

func (c *TemplateService) GetTemplateList(page, limit int) *response.TemplateListResp {

	offset := utils.CountOffset(page, limit)
	count := legend.GetScaleTemplateAll()
	scaleTemplates := legend.GetScaleTemplateList(offset, limit)

	for i, _ := range scaleTemplates {
		scaleTemplates[i].Id = offset + i + 1
	}

	templateListResp := new(response.TemplateListResp)
	templateListResp.Count = count
	templateListResp.Code = 0
	templateListResp.Data = scaleTemplates

	return templateListResp
}

/**
** 删除比例模板的所有内容
 */
func (c *TemplateService) DeleteTemplate(templateName string) *response.BaseResp {

	baseResp := new(response.BaseResp)
	baseResp.Code = -1

	b := true
	if !legend.DeleteScaleTemplate(templateName) {
		b = false
	}
	if !legend.DeleteAnyMoney(templateName) {
		b = false
	}
	if !legend.DeleteFixMoney(templateName) {
		b = false
	}
	if !legend.DeleteFixPresent(templateName) {
		b = false
	}
	if !legend.DeleteFixPresent(templateName) {
		b = false
	}

	if b {
		baseResp.Msg = "删除成功"
		baseResp.Code = 0
	} else {
		baseResp.Msg = "删除失败"
	}

	return baseResp
}

func (c *TemplateService) AllTemplateInfo(templateName string) *response.TemplateAllInfoResp {

	templateAllInfo := new(response.TemplateAllInfoResp)
	templateAllInfo.TemplateInfo = legend.GetScaleTemplateByName(templateName)
	templateAllInfo.AnyMoneyInfo = legend.GetAnyMoneyByName(templateName)
	templateAllInfo.FixMoneyInfos = legend.GetFixMoneyByName(templateName)
	templateAllInfo.PresentFixMoneyInfos = legend.GetFixPresentsByName(templateName)
	templateAllInfo.PresentScaleMoneyInfos = legend.GetScalePresentsByName(templateName)

	return templateAllInfo
}
