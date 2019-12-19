/***************************************************
 ** @Desc : This file for 处理Excel文件
 ** @Time : 19.12.6 16:25
 ** @Author : Joker
 ** @File : deal_excel
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.6 16:25
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"juhe/jhmerchant/sys/enum"
	"juhe/jhmerchant/utils"
	"juhe/service/models"
	"os"
	"recharge/sys"
	"strings"
	"time"
)

type DealExcel struct {
	KeepSession
}

// 下载模板
// @router /excel/download
func (c *DealExcel) DownloadExcelModel() {
	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	c.Ctx.Output.Download(enum.ExcelModelPath, enum.ExcelModelName)
}

// 导出订单记录
// @router /excel/make_order_excel/?:params [get]
func (c *DealExcel) MakeOrderExcel() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	// 查询参数
	in := make(map[string]string)
	start := strings.TrimSpace(c.GetString("start"))
	end := strings.TrimSpace(c.GetString("end"))
	payType := strings.TrimSpace(c.GetString("pay_type"))
	status := strings.TrimSpace(c.GetString("status"))

	in["pay_type_code"] = payType
	in["status"] = status
	in["merchant_uid"] = u.MerchantUid

	if start != "" {
		in["update_time__gte"] = start
	}
	if end != "" {
		in["update_time__lte"] = end
	}

	var (
		msg      = enum.FailedString
		flag     = enum.FailedFlag
		fileName = "trade_order-" + pubMethod.GetNowTimeV2() + pubMethod.RandomString(6) + ".xlsx"

		file  *xlsx.File
		sheet *xlsx.Sheet
		row   *xlsx.Row
		cell  *xlsx.Cell
		err   error
	)

	// 数据获取
	list := models.GetOrderProfitByMap(in, -1, 0)
	if len(list) <= 0 {
		msg = "没有检索到数据！"
		goto stopRun
	}

	// 写入记录
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("订单记录")
	if err != nil {
		utils.LogError(fmt.Sprintf("商户：%s 导出订单记录，发生错误：%v", u.MerchantName, err))
		msg = enum.FailedToAdmin
		goto stopRun
	}

	// 第一行
	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "平台订单号"
	cell = row.AddCell()
	cell.Value = "商户订单号"
	cell = row.AddCell()
	cell.Value = "支付方式"
	cell = row.AddCell()
	cell.Value = "订单金额"
	cell = row.AddCell()
	cell.Value = "收入金额"
	cell = row.AddCell()
	cell.Value = "手续费"
	cell = row.AddCell()
	cell.Value = "状 态"
	cell = row.AddCell()
	cell.Value = "成功支付时间"
	for _, v := range list {
		addRow := sheet.AddRow()
		addCell := addRow.AddCell()
		addCell.Value = v.BankOrderId
		addCell = addRow.AddCell()
		addCell.Value = v.MerchantOrderId
		addCell = addRow.AddCell()
		addCell.Value = v.PayProductName
		addCell = addRow.AddCell()
		addCell.Value = fmt.Sprintf("%f", v.OrderAmount)

		var (
			st = ""
			t  string
		)
		switch v.Status {
		case "failed":
			st = "交易失败"
		case "wait":
			st = "等待支付"
		case "success":
			st = "交易成功"
			t = v.UpdateTime
		}

		addCell = addRow.AddCell()
		addCell.Value = fmt.Sprintf("%f", v.UserInAmount)
		addCell = addRow.AddCell()
		addCell.Value = fmt.Sprintf("%f", v.AllProfit)
		addCell = addRow.AddCell()
		addCell.Value = st
		addCell = addRow.AddCell()
		addCell.Value = t
	}

	err = file.Save(enum.ExcelDownloadPath + fileName)
	if err != nil {
		utils.LogError(fmt.Sprintf("商户：%s 导出订单记录，保存文件发生错误：%v", u.MerchantName, err))
		msg = enum.FailedToAdmin
		goto stopRun
	}

	flag = enum.SuccessFlag
	msg = fileName

stopRun:
	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, "")
	c.ServeJSON()
	c.StopRun()
}

// 下载excel
// @router /excel/download_excel/?:params [get]
func (c *DealExcel) DownloadRecordExcel() {
	fileName := c.GetString(":params")
	file := enum.ExcelDownloadPath + fileName

	defer func() {
		if r := recover(); r != nil {
			sys.LogEmergency(file + " 此文件不存在！")
			time.Sleep(3 * time.Second)
		}
	}()
	// 删除临时文件
	go func() {
		tk := time.NewTicker(5 * time.Minute)
		select {
		case <-tk.C:
			_ = os.Remove(file)
			tk.Stop()
		}
	}()

	c.Ctx.Output.Download(file, fileName)
}

// 导出投诉记录
// @router /excel/make_complaint_record_excel/?:params [get]
func (c *DealExcel) MakeComplaintExcel() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	// 查询参数
	in := make(map[string]string)
	start := strings.TrimSpace(c.GetString("start"))
	end := strings.TrimSpace(c.GetString("end"))
	payType := strings.TrimSpace(c.GetString("pay_type"))
	status := strings.TrimSpace(c.GetString("status"))

	in["pay_type_code"] = payType
	if strings.Compare("YES", status) == 0 {
		in["freeze"] = enum.YES
	} else {
		in["refund"] = enum.YES
	}
	in["merchant_uid"] = u.MerchantUid

	if start != "" {
		in["update_time__gte"] = start
	}
	if end != "" {
		in["update_time__lte"] = end
	}

	var (
		msg      = enum.FailedString
		flag     = enum.FailedFlag
		fileName = "complaint_order-" + pubMethod.GetNowTimeV2() + pubMethod.RandomString(6) + ".xlsx"

		file  *xlsx.File
		sheet *xlsx.Sheet
		row   *xlsx.Row
		cell  *xlsx.Cell
		err   error
	)

	// 数据获取
	list := models.GetOrderByMap(in, -1, 0)
	if len(list) <= 0 {
		msg = "没有检索到数据！"
		goto stopRun
	}

	// 写入记录
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("投诉记录")
	if err != nil {
		utils.LogError(fmt.Sprintf("商户：%s 导出投诉记录，发生错误：%v", u.MerchantName, err))
		msg = enum.FailedToAdmin
		goto stopRun
	}

	// 第一行
	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "平台流水号"
	cell = row.AddCell()
	cell.Value = "商户订单号"
	cell = row.AddCell()
	cell.Value = "支付方式"
	cell = row.AddCell()
	cell.Value = "订单金额"
	cell = row.AddCell()
	cell.Value = "状 态"
	cell = row.AddCell()
	cell.Value = "冻结时间"
	for _, v := range list {
		addRow := sheet.AddRow()
		addCell := addRow.AddCell()
		addCell.Value = v.BankOrderId
		addCell = addRow.AddCell()
		addCell.Value = v.MerchantOrderId
		addCell = addRow.AddCell()
		addCell.Value = v.PayProductName
		addCell = addRow.AddCell()
		addCell.Value = fmt.Sprintf("%f", v.OrderAmount)

		var st = ""
		switch v.Freeze {
		case "yes":
			st = "已冻结"
		case "no":
			st = "已退款"
		}

		addCell = addRow.AddCell()
		addCell.Value = st
		addCell = addRow.AddCell()
		addCell.Value = v.UpdateTime
	}

	err = file.Save(enum.ExcelDownloadPath + fileName)
	if err != nil {
		utils.LogError(fmt.Sprintf("商户：%s 导出投诉记录，保存文件发生错误：%v", u.MerchantName, err))
		msg = enum.FailedToAdmin
		goto stopRun
	}

	flag = enum.SuccessFlag
	msg = fileName

stopRun:
	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, "")
	c.ServeJSON()
	c.StopRun()
}

// 导出提现记录
// @router /excel/make_withdraw_record_excel/?:params [get]
func (c *DealExcel) MakeWithdrawExcel() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	// 查询参数
	in := make(map[string]string)
	start := strings.TrimSpace(c.GetString("start"))
	end := strings.TrimSpace(c.GetString("end"))
	status := strings.TrimSpace(c.GetString("status"))

	in["status"] = status
	in["merchant_uid"] = u.MerchantUid

	if start != "" {
		in["update_time__gte"] = start
	}
	if end != "" {
		in["update_time__lte"] = end
	}

	var (
		msg      = enum.FailedString
		flag     = enum.FailedFlag
		fileName = "withdraw_order-" + pubMethod.GetNowTimeV2() + pubMethod.RandomString(6) + ".xlsx"

		file  *xlsx.File
		sheet *xlsx.Sheet
		row   *xlsx.Row
		cell  *xlsx.Cell
		err   error
	)

	// 数据获取
	list := models.GetPayForByMap(in, -1, 0)
	if len(list) <= 0 {
		msg = "没有检索到数据！"
		goto stopRun
	}

	// 写入记录
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("提现记录")
	if err != nil {
		utils.LogError(fmt.Sprintf("商户：%s 导出提现记录，发生错误：%v", u.MerchantName, err))
		msg = enum.FailedToAdmin
		goto stopRun
	}

	// 第一行
	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "平台订单号"
	cell = row.AddCell()
	cell.Value = "商户订单号"
	cell = row.AddCell()
	cell.Value = "结算金额"
	cell = row.AddCell()
	cell.Value = "手续费"
	cell = row.AddCell()
	cell.Value = "银行名称"
	cell = row.AddCell()
	cell.Value = "开户名"
	cell = row.AddCell()
	cell.Value = "开户账户"
	cell = row.AddCell()
	cell.Value = "状 态"
	cell = row.AddCell()
	cell.Value = "创建时间"
	cell = row.AddCell()
	cell.Value = "打款时间"
	cell = row.AddCell()
	cell.Value = "备注"
	for _, v := range list {
		addRow := sheet.AddRow()
		addCell := addRow.AddCell()
		addCell.Value = v.BankOrderId
		addCell = addRow.AddCell()
		addCell.Value = v.MerchantOrderId
		addCell = addRow.AddCell()
		addCell.Value = fmt.Sprintf("%f", v.PayforTotalAmount)
		addCell = addRow.AddCell()
		addCell.Value = fmt.Sprintf("%f", v.PayforFee)
		addCell = addRow.AddCell()
		addCell.Value = v.BankName
		addCell = addRow.AddCell()
		addCell.Value = v.BankAccountName
		addCell = addRow.AddCell()
		addCell.Value = v.BankAccountNo

		var (
			st = ""
			t  string
		)
		switch v.Status {
		case "payfor_confirm":
			st = "等待审核"
		case "payfor_solving":
			st = "系统处理中"
		case "payfor_banking":
			st = "银行处理中"
		case "success":
			st = "代付成功"
			t = v.UpdateTime
		case "failed":
			st = "代付失败"
		}

		addCell = addRow.AddCell()
		addCell.Value = st
		addCell = addRow.AddCell()
		addCell.Value = v.CreateTime
		addCell = addRow.AddCell()
		addCell.Value = t
		addCell = addRow.AddCell()
		addCell.Value = v.Remark
	}

	err = file.Save(enum.ExcelDownloadPath + fileName)
	if err != nil {
		utils.LogError(fmt.Sprintf("商户：%s 导出提现记录，保存文件发生错误：%v", u.MerchantName, err))
		msg = enum.FailedToAdmin
		goto stopRun
	}

	flag = enum.SuccessFlag
	msg = fileName

stopRun:
	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, "")
	c.ServeJSON()
	c.StopRun()
}
