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
	"fmt"
	"github.com/rs/xid"
	"merchant/conf"
	"merchant/models"
	"merchant/sys/enum"
	"merchant/utils"
	"regexp"
	"strconv"
	"strings"
)

type Withdraw struct {
	KeepSession
}

func (c *Withdraw) ShowWithdrawUI() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	c.Data["bankInfo"] = enum.GetBankInfo()
	c.Data["userName"] = u.MerchantName
	c.TplName = "withdraw/withdraw.html"
}

// 获取账户提现余额
func (c *Withdraw) UserBalance() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

	ac := models.GetAccountByUid(u.MerchantUid)
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

// 发送提现短信
func (c *Withdraw) SendMsgForWithdraw() {
	bankCode := strings.TrimSpace(c.GetString("bankCode"))
	accountName := strings.TrimSpace(c.GetString("accountName"))
	cardNo := strings.TrimSpace(c.GetString("cardNo"))
	bankAccountType := strings.TrimSpace(c.GetString("bankAccountType"))
	province := strings.TrimSpace(c.GetString("province"))
	city := strings.TrimSpace(c.GetString("city"))
	bankAccountAddress := strings.TrimSpace(c.GetString("bankAccountAddress"))
	moblieNo := strings.TrimSpace(c.GetString("moblieNo"))
	money := strings.TrimSpace(c.GetString("amount"))

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
		matched    bool

		smsSession interface{}
	)

	ac := models.GetAccountByUid(u.MerchantUid)
	matched, msg = verifyAccountAndMoney(bankCode, accountName, cardNo, bankAccountType, province, city,
		bankAccountAddress, moblieNo, money, ac)
	if !matched {
		goto stopRun
	}

	isContains = strings.Contains(u.WhiteIps, ip)
	if isContains {
		msg = "已有IP白名单，无需发送验证码，输入任意字符即可!"
		goto stopRun
	}

	// 以免cookie失效了，但是session还没失效
	smsCookie = c.Ctx.GetCookie("do_pay_sms_cookie")
	if smsCookie == "" {
		c.SetSession("do_pay_sms_cookie", "")
	}
	smsSession = c.GetSession("do_pay_sms_cookie")
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

	c.SetSession("do_pay_code", code)
	code = pubMethod.RandomString(6)
	c.Ctx.SetCookie("do_pay_sms_cookie", code, enum.SmsCookieExpireTime)
	c.SetSession("do_pay_sms_cookie", code)

stopRun:
	c.Data["json"] = pubMethod.JsonFormat(flag, "", msg, "")
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
	u := us.(models.MerchantInfo)
	//获取客户端ip
	ip := strings.TrimSpace(c.Ctx.Input.IP())

	var (
		msg  = enum.FailedString
		flag = enum.FailedFlag

		matched    bool
		isContains bool

		amount float64

		ac   models.AccountInfo
		sett models.PayforInfo
		url  string
	)

	isContains = strings.Contains(u.WhiteIps, ip)
	if !isContains {
		code := c.GetSession("do_pay_code")
		smsCookie := c.Ctx.GetCookie("do_pay_sms_cookie")
		if smsCookie == "" || code == nil {
			msg = "请发送提现验证码！"
			goto stopRun
		}
		if strings.Compare(code.(string), mobileCode) != 0 {
			msg = "短信验证码输入错误！"
			goto stopRun
		}
	}

	ac = models.GetAccountByUid(u.MerchantUid)
	matched, msg = verifyAccountAndMoney(bankCode, accountName, cardNo, bankAccountType, province, city,
		bankAccountAddress, moblieNo, money, ac)
	if !matched {
		goto stopRun
	}

	u = models.GetMerchantByPhone(u.LoginAccount)
	if strings.Compare(enum.ACTIVE, u.Status) != 0 {
		msg = "账户状态异常，请联系管理人员!"
		goto stopRun
	}

	amount, _ = strconv.ParseFloat(money, 10)
	sett = models.PayforInfo{
		PayforUid:          "pppp" + xid.New().String(),
		MerchantUid:        u.MerchantUid,
		MerchantName:       u.MerchantName,
		PhoneNo:            u.LoginAccount,
		MerchantOrderId:    xid.New().String(),
		BankOrderId:        "4444" + xid.New().String(),
		PayforFee:          conf.PAYFOR_FEE,
		Type:               conf.SELF_MERCHANT,
		PayforAmount:       amount,
		PayforTotalAmount:  amount + conf.PAYFOR_FEE,
		BankCode:           bankCode,
		BankName:           enum.GetBankInfo()[bankCode],
		IsSend:             conf.NO,
		BankAccountName:    accountName,
		BankAccountNo:      cardNo,
		BankAccountType:    bankAccountType,
		BankAccountAddress: province + city + bankAccountAddress,
		Status:             conf.PAYFOR_COMFRIM,
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
	u := us.(models.MerchantInfo)

	ranMd5 := encrypt.EncodeMd5([]byte(pubMethod.RandomString(46)))
	c.Ctx.SetCookie(enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.Ctx.SetSecureCookie(ranMd5, enum.UserCookie, ranMd5, enum.CookieExpireTime)
	c.SetSession(enum.UserCookie, ranMd5)

	c.Data["payType"] = enum.GetSettlementStatus()
	c.Data["userName"] = u.MerchantName
	c.TplName = "withdraw/withdraw_record.html"
}

func (c *Withdraw) WithdrawQueryAndListPage() {
	us := c.GetSession(enum.UserSession)
	u := us.(models.MerchantInfo)

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
	in["merchant_uid"] = u.MerchantUid

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
