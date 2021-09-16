/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/26 15:30
 ** @Author : yuebin
 ** @File : conf_pro
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/26 15:30
 ** @Software: GoLand
****************************************************/
package conf

const (
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
	DB_USER     = "root"
	DB_PASSWORD = "Kyb^15273031604"
	DB_BASE     = "juhe_pay"
)

const (
	ACTIVE          = "active"
	UNACTIVE        = "unactive"
	DELETE          = "delete"
	REFUND          = "refund"
	ORDERROLL       = "order_roll"
	WAIT            = "wait"
	SUCCESS         = "success"
	FAIL            = "fail"
	YES             = "yes"
	NO              = "no"
	ZERO            = 0.0  //0元手续费
	VERIFY_CODE_LEN = 4    //验证码的长度
	PAYFOR_FEE      = 2.00 //代付手续费
	PAYFOR_INTERVAL = 5    //每过5分钟执行一次代付

	PLUS_AMOUNT     = "plus_amount"     //加款操作
	SUB_AMOUNT      = "sub_amount"      //减款操作
	FREEZE_AMOUNT   = "freeze_amount"   //冻结操作
	UNFREEZE_AMOUNT = "unfreeze_amount" //解冻操作

	PAYFOR_COMFRIM = "payfor_confirm" //下发带审核
	PAYFOR_SOLVING = "payfor_solving" //发下处理中
	PAYFOR_HANDING = "payfor_handing" //手动打款中
	PAYFOR_BANKING = "payfor_banking" //银行处理中
	PAYFOR_FAIL    = "payfor_fail"    //代付失败
	PAYFOR_SUCCESS = "payfor_success" //代付成功

	PAYFOR_ROAD   = "payfor_road"   //通道打款
	PAYFOR_HAND   = "payfor_hand"   //手动打款
	PAYFOR_REFUSE = "payfor_refuse" // 拒绝打款

	SELF_API      = "self_api"      //自助api系统下发
	SELF_MERCHANT = "self_merchant" //管理手动处理商户下发
	SELF_HELP     = "self_help"     //管理自己提现

	PUBLIC  = "public"  //对公卡
	PRIVATE = "private" //对私卡
)

const (
	ICBC   = "ICBC"
	ABC    = "ABC"
	BOC    = "BOC"
	CCB    = "CCB"
	BOCOM  = "BOCOM"
	CNCB   = "CNCB"
	CEB    = "CEB"
	HXB    = "HXB"
	CMBC   = "CMBC"
	GDB    = "GDB"
	CMB    = "CMB"
	CIB    = "CIB"
	SPDB   = "SPDB"
	PSBC   = "PSBC"
	PAB    = "PAB"
	NJCB   = "NJCB"
	NBCB   = "NBCB"
	WZCB   = "WZCB"
	CSCB   = "CSCB"
	CZCB   = "CZCB"
	CCQTGB = "CCQTGB"
	SHRCB  = "SHRCB"
	BJRCB  = "BJRCB"
	SDB    = "SDB"
)

var bankInfo = map[string]string{
	ICBC:   "中国工商银行",
	ABC:    "中国农业银行",
	BOC:    "中国银行",
	CCB:    "中国建设银行",
	BOCOM:  "交通银行",
	CNCB:   "中信银行",
	CEB:    "中国光大银行",
	HXB:    "华夏银行",
	CMBC:   "中国民生银行",
	GDB:    "广发银行",
	CMB:    "招商银行",
	CIB:    "兴业银行",
	SPDB:   "浦发银行",
	PSBC:   "中国邮政储蓄银行",
	PAB:    "平安银行",
	NJCB:   "南京银行",
	NBCB:   "宁波银行",
	WZCB:   "温州市商业银行",
	CSCB:   "长沙银行",
	CZCB:   "浙江稠州商业银行",
	CCQTGB: "重庆三峡银行",
	SHRCB:  "上海农村商业银行",
	BJRCB:  "北京农商行",
	SDB:    "深圳发展银行",
}
