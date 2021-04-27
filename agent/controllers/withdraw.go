/***************************************************
 ** @Desc : This file for 代付申请和记录
 ** @Time : 19.12.4 10:50
 ** @Author : Joker
 ** @File : withdraw
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.4 10:50
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"agent/common"
	"agent/models"
	"agent/sys/enum"
	"agent/utils"
	"fmt"
	"github.com/rs/xid"
	"regexp"
	"strconv"
	"strings"
)

type Withdraw struct {
	KeepSession
}

func (c *Withdraw) ShowWithdrawUI() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	c.Data["bankInfo"] = enum.GetBankInfo()
	c.Data["userName"] = u.AgentName
	c.TplName = "withdraw/withdraw.html"
}

// 获取账户提现余额
func (c *Withdraw) UserBalance() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	ac := models.GetAccountByUid(u.AgentUid)
	balance := ac.Balance
	if balance < 0 {
		balance = 0
	}

	out := make(map[string]interface{})
	out["fee"] = enum.SettlementFee
	out["balance"] = fmt.Sprintf("%0.2f", balance)

	c.Data["json"] = out
	c.ServeJSON()
	c.StopRun()
}

// 验证卡号和金额是否正确
func verifyAccountAndMoney(bankCode, accountName, cardNo, bankAccountType, province, city, bankAccountAddress, moblieNo, amount string, ac models.AccountInfo) (b bool, msg string) {
	if bankCode == "" || accountName == "" || cardNo == "" {
		msg = "银行名、账户名或卡号不能为空!"
		return false, msg
	}
	if moblieNo == "" || amount == "" {
		msg = "手机号或金额不能为空!"
		return false, msg
	}
	matched, _ := regexp.MatchString(enum.MobileReg, moblieNo)
	if !matched {
		msg = "请输入正确的手机号!"
		return false, msg
	}
	matched, _ = regexp.MatchString(enum.MoneyReg, amount)
	if !matched {
		msg = "请输入正确的金额!"
		return false, msg
	}
	f, err := strconv.ParseFloat(amount, 10)
	if err != nil {
		msg = "请输入正确的金额!"
		return false, msg
	}

	if strings.Compare(enum.PublicAccount, bankAccountType) == 0 {
		if province == "" || city == "" || bankAccountAddress == "" {
			msg = "开户行全称、所在省份或所在城市不能为空!"
			return false, msg
		}
	}

	if f > enum.WithdrawalMaxAmount || f < enum.WithdrawalMinAmount || f+enum.SettlementFee > ac.WaitAmount {
		utils.LogInfo(fmt.Sprintf("提现金额超出限制，提现金额：%f，账户可结算余额：%f，提现最小额：%d，最大额：%d，手续费：%d",
			f, ac.WaitAmount, enum.WithdrawalMinAmount, enum.WithdrawalMaxAmount, enum.SettlementFee))
		msg = "提现金额超出限制!"
		return false, msg
	}

	if f+enum.SettlementFee > ac.Balance || ac.Balance <= 0 {
		utils.LogInfo(fmt.Sprintf("账户金额不足，提现金额：%f，账户余额：%f，手续费：%d",
			f, ac.Balance, enum.SettlementFee))
		msg = "账户可用金额不足!"
		return false, msg
	}
	// 由于没有发生错误，必须把msg重置为初始值，而不是空值
	return true, enum.FailedToAdmin
}

// 单笔提现申请
func (c *Withdraw) LaunchSingleWithdraw() {
	bankCode := strings.TrimSpace(c.GetString("bankCode"))
	accountName := strings.TrimSpace(c.GetString("accountName"))
	cardNo := strings.TrimSpace(c.GetString("cardNo"))
	bankAccountType := strings.TrimSpace(c.GetString("bankAccountType"))
	province := strings.TrimSpace(c.GetString("province"))
	city := strings.TrimSpace(c.GetString("city"))
	bankAccountAddress := strings.TrimSpace(c.GetString("bankAccountAddress"))
	moblieNo := strings.TrimSpace(c.GetString("moblieNo"))
	money := strings.TrimSpace(c.GetString("amount"))
	mobileCode := strings.TrimSpace(c.GetString("smsVerifyCode"))

	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	var (
		msg  = enum.FailedString
		flag = enum.FailedFlag

		matched bool

		amount float64

		ac   models.AccountInfo
		sett models.PayforInfo
		url  string
	)

	if u.PayPassword == "" {
		msg = "请设置支付密码！"
		goto stopRun
	}
	if strings.Compare(strings.ToUpper(encrypt.EncodeMd5([]byte(mobileCode))), u.PayPassword) != 0 {
		msg = "支付密码输入错误！"
		goto stopRun
	}

	ac = models.GetAccountByUid(u.AgentUid)
	matched, msg = verifyAccountAndMoney(bankCode, accountName, cardNo, bankAccountType, province, city,
		bankAccountAddress, moblieNo, money, ac)
	if !matched {
		goto stopRun
	}

	u = models.GetAgentInfoByAgentUid(u.AgentUid)
	if strings.Compare(enum.ACTIVE, u.Status) != 0 {
		msg = "账户状态异常，请联系管理人员!"
		goto stopRun
	}

	amount, _ = strconv.ParseFloat(money, 10)
	sett = models.PayforInfo{
		PayforUid:          "pppp" + xid.New().String(),
		MerchantUid:        u.AgentUid,
		MerchantName:       u.AgentName,
		PhoneNo:            u.AgentPhone,
		MerchantOrderId:    xid.New().String(),
		BankOrderId:        "4444" + xid.New().String(),
		PayforFee:          common.PAYFOR_FEE,
		Type:               common.SELF_MERCHANT,
		PayforAmount:       amount,
		PayforTotalAmount:  amount + common.PAYFOR_FEE,
		BankCode:           bankCode,
		BankName:           enum.GetBankInfo()[bankCode],
		IsSend:             common.NO,
		BankAccountName:    accountName,
		BankAccountNo:      cardNo,
		BankAccountType:    bankAccountType,
		BankAccountAddress: province + city + bankAccountAddress,
		Status:             common.PAYFOR_COMFRIM,
		CreateTime:         pubMethod.GetNowTime(),
		UpdateTime:         pubMethod.GetNowTime(),
	}

	matched = models.InsertPayfor(sett)
	if matched {
		flag = enum.SuccessFlag
		msg = "提交成功，等待审核中，请在结算信息中查询状态！"
		url = "/withdraw/show_list_ui"
	}

stopRun:
	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, url)
	c.ServeJSON()
	c.StopRun()
}

// 提现列表
func (c *Withdraw) ShowListUI() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	c.Data["payType"] = enum.GetSettlementStatus()
	c.Data["userName"] = u.AgentName
	c.TplName = "withdraw/withdraw_record.html"
}

func (c *Withdraw) WithdrawQueryAndListPage() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.AgentInfo)

	// 分页参数
	page, _ := strconv.Atoi(c.GetString("page"))
	limit, _ := strconv.Atoi(c.GetString("limit"))
	if limit == 0 {
		limit = 15
	}

	// 查询参数
	in := make(map[string]string)
	merchantNo := strings.TrimSpace(c.GetString("MerchantNo"))
	bankNo := strings.TrimSpace(c.GetString("BankNo"))
	start := strings.TrimSpace(c.GetString("start"))
	end := strings.TrimSpace(c.GetString("end"))
	status := strings.TrimSpace(c.GetString("status"))

	in["bank_order_id"] = bankNo
	in["merchant_order_id"] = merchantNo
	in["status"] = status
	in["merchant_uid"] = u.AgentUid

	if start != "" {
		in["create_time__gte"] = start
	}
	if end != "" {
		in["create_time__lte"] = end
	}

	// 计算分页数
	count := models.GetPayForLenByMap(in)
	totalPage := count / limit // 计算总页数
	if count%limit != 0 {      // 不满一页的数据按一页计算
		totalPage++
	}

	// 数据获取
	var list []models.PayforInfo
	if page <= totalPage {
		list = models.GetPayForByMap(in, limit, (page-1)*limit)
	}

	// 数据回显
	out := make(map[string]interface{})
	out["limit"] = limit // 分页数据
	out["page"] = page
	out["totalPage"] = totalPage
	out["root"] = list // 显示数据

	c.Data["json"] = out
	c.ServeJSON()
	c.StopRun()
}
