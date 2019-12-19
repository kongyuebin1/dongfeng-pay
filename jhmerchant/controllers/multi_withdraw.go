/***************************************************
 ** @Desc : This file for 批量提现
 ** @Time : 19.12.6 17:07
 ** @Author : Joker
 ** @File : multi_withdraw
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.6 17:07
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"fmt"
	"github.com/rs/xid"
	"github.com/tealeg/xlsx"
	"dongfeng-pay/jhmerchant/sys/enum"
	"dongfeng-pay/jhmerchant/utils"
	"dongfeng-pay/service/common"
	"dongfeng-pay/service/models"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MultiWithdraw struct {
	KeepSession
}

// @router /multi_withdraw/show_multi_ui
func (c *MultiWithdraw) ShowMultiWithdrawUI() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	c.Data["userName"] = u.MerchantName
	c.TplName = "withdraw/multi_withdraw.html"
}

// 发送批量提现短信
// @router /multi_withdraw/send_msg_for_multi/?:params [post]
func (c *Withdraw) SendMsgForMultiWithdraw() {
	file, header, err := c.GetFile("file")

	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)
	//获取客户端ip
	ip := strings.TrimSpace(c.Ctx.Input.IP())

	var (
		msg  = enum.FailedString
		flag = enum.FailedFlag

		smsCookie  string
		codeCookie string
		code       string

		isSend     bool
		isContains bool

		smsSession interface{}
	)

	fileName := header.Filename
	defer file.Close()
	split := strings.Split(fileName, ".")
	if len(split) < 2 {
		msg = "请选择批量文件!"
		goto stopRun
	}
	if strings.Compare(strings.ToLower(split[1]), "xls") != 0 &&
		strings.Compare(strings.ToLower(split[1]), "xlsx") != 0 {
		msg = "仅支持“xls”、“xlsx”格式文件!"
		goto stopRun
	}
	if err != nil {
		msg = "请上传批量文件! " + err.Error()
		goto stopRun
	}

	isContains = strings.Contains(u.WhiteIps, ip)
	if isContains {
		msg = "已有IP白名单，无需发送验证码，输入任意字符即可!"
		goto stopRun
	}

	// 以免cookie失效了，但是session还没失效
	smsCookie = c.Ctx.GetCookie("do_multi_pay_sms_cookie")
	if smsCookie == "" {
		c.SetSession("do_multi_pay_sms_cookie", "")
	}
	smsSession = c.GetSession("do_multi_pay_sms_cookie")
	if smsSession != nil {
		codeCookie = smsSession.(string)
	}

	if smsCookie != "" || strings.Compare(smsCookie, codeCookie) != 0 {
		msg = fmt.Sprintf("请在%d分钟后送验证码！", enum.SmsCookieExpireTime/60)
		goto stopRun
	}

	code = pubMethod.RandomIntOfString(6)
	isSend = utils.SendSmsForPay(u.LoginAccount, code)
	if !isSend {
		msg = fmt.Sprintf("验证码发送失败，请核实预留手机号：%s 是否正确！", u.LoginAccount)
		goto stopRun
	}

	flag = enum.SuccessFlag
	msg = fmt.Sprintf("验证码已发送到您的手机号：%s，且%d分钟内有效！", u.LoginAccount, enum.SmsCookieExpireTime/60)

	c.SetSession("do_multi_pay_code", code)
	code = pubMethod.RandomString(6)
	c.Ctx.SetCookie("do_multi_pay_sms_cookie", code, enum.SmsCookieExpireTime)
	c.SetSession("do_multi_pay_sms_cookie", code)

stopRun:
	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, "")
	c.ServeJSON()
	c.StopRun()
}

// 申请批量提现
// @router /multi_withdraw/launch_multi_withdraw/?:params [post]
func (c *Withdraw) LaunchMultiWithdraw() {
	mobileCode := strings.TrimSpace(c.GetString("mobileCode"))
	file, header, err := c.GetFile("file")

	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)
	//获取客户端ip
	ip := strings.TrimSpace(c.Ctx.Input.IP())

	var (
		msg  = enum.FailedString
		flag = enum.FailedFlag

		isContains bool
		b          bool

		smsCookie string
		url       string
	)

	fileName := header.Filename
	defer file.Close()
	split := strings.Split(fileName, ".")
	if len(split) < 2 {
		msg = "请选择批量文件!"
		goto stopRun
	}
	if strings.Compare(strings.ToLower(split[1]), "xls") != 0 &&
		strings.Compare(strings.ToLower(split[1]), "xlsx") != 0 {
		msg = "仅支持“xls”、“xlsx”格式文件!"
		goto stopRun
	}
	if err != nil {
		msg = "请上传批量文件! " + err.Error()
		goto stopRun
	}

	isContains = strings.Contains(u.WhiteIps, ip)
	if !isContains {
		code := c.GetSession("do_multi_pay_code")
		smsCookie = c.Ctx.GetCookie("do_multi_pay_sms_cookie")
		if smsCookie == "" || code == nil {
			msg = "请发送提现验证码！"
			goto stopRun
		}
		if strings.Compare(code.(string), mobileCode) != 0 {
			msg = "短信验证码输入错误！"
			goto stopRun
		}
	}

	u = models.GetMerchantByUid(u.MerchantUid)
	if strings.Compare(enum.ACTIVE, u.Status) != 0 {
		msg = "商户状态异常，请联系管理人员!"
		goto stopRun
	}

	b, msg = handleFileContent(split[0], u, c)
	if b {
		flag = enum.SuccessFlag
		url = "/withdraw/show_list_ui"
	}

stopRun:
	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, url)
	c.ServeJSON()
	c.StopRun()
}

func handleFileContent(name string, u models.MerchantInfo, c *Withdraw) (bool, string) {
	// 重命名文件
	fileName := name + " - " + pubMethod.GetNowTimeV2() + pubMethod.RandomString(4) + ".xlsx"

	// 保存文件
	_ = c.SaveToFile("file", path.Join(enum.ExcelPath, fileName))

	// 读取文件内容
	xlFile, err := xlsx.OpenFile(enum.ExcelPath + fileName)
	if err != nil {
		msg := "文件内容错误：" + err.Error()
		utils.LogInfo(msg)
		return false, msg
	}

	// 只读取文档中第一个工作表，忽略其他工作表
	sheet := xlFile.Sheets[0]
	line, err := sheet.Row(0).Cells[1].Int()
	if err != nil {
		msg := "请输入正确的总笔数：" + err.Error()
		utils.LogInfo(msg)
		return false, msg
	}
	if line <= 0 {
		msg := "请输入正确的总笔数！"
		return false, msg
	}
	if line > 300 {
		line = 300
	}

	ac := models.GetAccountByUid(u.MerchantUid)
	if strings.Compare(enum.ACTIVE, ac.Status) != 0 {
		msg := "账户状态异常，请联系管理人员!"
		return false, msg
	}
	// 后台处理文件，不让用户等待
	go func() {
		for k, row := range sheet.Rows {
			if k == 0 || k == 1 {
				continue
			}

			// 数据行数不得超过指定行数
			if k == line+2 || k == 301 {
				break
			}

			// 出现空行，则忽略后面记录
			if row.Cells[0].String() == "" || row.Cells[4].String() == "" {
				break
			}

			bankAccountType := row.Cells[3].String()
			ac := models.GetAccountByUid(u.MerchantUid)
			b, msg, code := verifyFileContent(row.Cells[0].String(), row.Cells[1].String(), row.Cells[2].String(), bankAccountType,
				row.Cells[4].String(), ac)
			if !b {
				utils.LogInfo(fmt.Sprintf("用户：%s 批量代付中，第 %d 行记录出现错误：%s", u.MerchantName, k+1, msg))
				// 账户可用余额不足，终止读取记录
				if code == 5009 {
					break
				}
				continue
			}

			if strings.Compare("对公", bankAccountType) == 0 {
				bankAccountType = enum.PublicAccount
			} else {
				bankAccountType = enum.PrivateDebitAccount
			}

			money, _ := strconv.ParseFloat(row.Cells[4].String(), 10)
			payFor := models.PayforInfo{
				PayforUid:          "pppp" + xid.New().String(),
				MerchantUid:        u.MerchantUid,
				MerchantName:       u.MerchantName,
				PhoneNo:            u.LoginAccount,
				MerchantOrderId:    xid.New().String(),
				BankOrderId:        "4444" + xid.New().String(),
				PayforFee:          common.PAYFOR_FEE,
				Type:               common.SELF_MERCHANT,
				PayforAmount:       money,
				PayforTotalAmount:  money + common.PAYFOR_FEE,
				BankCode:           "C",
				BankName:           row.Cells[2].String(),
				IsSend:             common.NO,
				BankAccountName:    row.Cells[0].String(),
				BankAccountNo:      row.Cells[1].String(),
				BankAccountType:    bankAccountType,
				BankAccountAddress: row.Cells[2].String(),
				Status:             common.PAYFOR_COMFRIM,
				CreateTime:         pubMethod.GetNowTime(),
				UpdateTime:         pubMethod.GetNowTime(),
			}

			models.InsertPayfor(payFor)

			time.Sleep(500 * time.Millisecond)
		}
	}()

	return true, "提交成功，等待审核中，请在结算信息中查询状态！"
}

// 验证文件内容是否规范
func verifyFileContent(accountName, cardNo, bankAccountAddress, bankAccountType, amount string, ac models.AccountInfo) (bool, string, int) {
	if accountName == "" || cardNo == "" {
		msg := "账户名或卡号不能为空!"
		return false, msg, 5001
	}
	// 账户类型
	if strings.Compare("对公", bankAccountType) == 0 {
		if bankAccountAddress == "" {
			msg := "收款方开户机构名称不能为空!"
			return false, msg, 5002
		}
	}
	if amount == "" {
		msg := "金额不能为空!"
		return false, msg, 5003
	}
	matched, _ := regexp.MatchString(enum.MoneyReg, amount)
	if !matched {
		msg := "请输入正确的金额!"
		return false, msg, 5004
	}
	f, err := strconv.ParseFloat(amount, 10)
	if err != nil {
		msg := "请输入正确的金额! " + err.Error()
		return false, msg, 5007
	}

	if f > enum.WithdrawalMaxAmount || f < enum.WithdrawalMinAmount || f+enum.SettlementFee > ac.WaitAmount {
		msg := fmt.Sprintf("单笔提现金额超出限制，提现金额：%f，账户可结算余额：%f，提现最小额：%d，最大额：%d，手续费：%d",
			f, ac.WaitAmount, enum.WithdrawalMinAmount, enum.WithdrawalMaxAmount, enum.SettlementFee)
		return false, msg, 5008
	}

	if f+enum.SettlementFee > ac.Balance || ac.Balance <= 0 {
		msg := fmt.Sprintf("账户金额不足，提现金额：%f，账户余额：%f，手续费：%d",
			f, ac.Balance, enum.SettlementFee)
		return false, msg, 5009
	}
	return true, "", 5000
}
