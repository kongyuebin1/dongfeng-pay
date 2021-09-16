package utils

import (
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/widuu/gojson"
)

const bankUrl = "https://ccdcapi.alipay.com/validateAndCacheCardInfo.json?_input_charset=utf-8&cardBinCheck=true&cardNo="
const CodeToName = "{\n  \"SRCB\": \"深圳农村商业银行\", \n  \"BGB\": \"广西北部湾银行\", \n  \"SHRCB\": \"上海农村商业银行\", \n  \"BJBANK\": \"北京银行\", \n  \"WHCCB\": \"威海市商业银行\", \n  \"BOZK\": \"周口银行\", \n  \"KORLABANK\": \"库尔勒市商业银行\", \n  \"SPABANK\": \"平安银行\", \n  \"SDEB\": \"顺德农商银行\", \n  \"HURCB\": \"湖北省农村信用社\", \n  \"WRCB\": \"无锡农村商业银行\", \n  \"BOCY\": \"朝阳银行\", \n  \"CZBANK\": \"浙商银行\", \n  \"HDBANK\": \"邯郸银行\", \n  \"BOC\": \"中国银行\", \n  \"BOD\": \"东莞银行\", \n  \"CCB\": \"中国建设银行\", \n  \"ZYCBANK\": \"遵义市商业银行\", \n  \"SXCB\": \"绍兴银行\", \n  \"GZRCU\": \"贵州省农村信用社\", \n  \"ZJKCCB\": \"张家口市商业银行\", \n  \"BOJZ\": \"锦州银行\", \n  \"BOP\": \"平顶山银行\", \n  \"HKB\": \"汉口银行\", \n  \"SPDB\": \"上海浦东发展银行\", \n  \"NXRCU\": \"宁夏黄河农村商业银行\", \n  \"NYNB\": \"广东南粤银行\", \n  \"GRCB\": \"广州农商银行\", \n  \"BOSZ\": \"苏州银行\", \n  \"HZCB\": \"杭州银行\", \n  \"HSBK\": \"衡水银行\", \n  \"HBC\": \"湖北银行\", \n  \"JXBANK\": \"嘉兴银行\", \n  \"HRXJB\": \"华融湘江银行\", \n  \"BODD\": \"丹东银行\", \n  \"AYCB\": \"安阳银行\", \n  \"EGBANK\": \"恒丰银行\", \n  \"CDB\": \"国家开发银行\", \n  \"TCRCB\": \"江苏太仓农村商业银行\", \n  \"NJCB\": \"南京银行\", \n  \"ZZBANK\": \"郑州银行\", \n  \"DYCB\": \"德阳商业银行\", \n  \"YBCCB\": \"宜宾市商业银行\", \n  \"SCRCU\": \"四川省农村信用\", \n  \"KLB\": \"昆仑银行\", \n  \"LSBANK\": \"莱商银行\", \n  \"YDRCB\": \"尧都农商行\", \n  \"CCQTGB\": \"重庆三峡银行\", \n  \"FDB\": \"富滇银行\", \n  \"JSRCU\": \"江苏省农村信用联合社\", \n  \"JNBANK\": \"济宁银行\", \n  \"CMB\": \"招商银行\", \n  \"JINCHB\": \"晋城银行JCBANK\", \n  \"FXCB\": \"阜新银行\", \n  \"WHRCB\": \"武汉农村商业银行\", \n  \"HBYCBANK\": \"湖北银行宜昌分行\", \n  \"TZCB\": \"台州银行\", \n  \"TACCB\": \"泰安市商业银行\", \n  \"XCYH\": \"许昌银行\", \n  \"CEB\": \"中国光大银行\", \n  \"NXBANK\": \"宁夏银行\", \n  \"HSBANK\": \"徽商银行\", \n  \"JJBANK\": \"九江银行\", \n  \"NHQS\": \"农信银清算中心\", \n  \"MTBANK\": \"浙江民泰商业银行\", \n  \"LANGFB\": \"廊坊银行\", \n  \"ASCB\": \"鞍山银行\", \n  \"KSRB\": \"昆山农村商业银行\", \n  \"YXCCB\": \"玉溪市商业银行\", \n  \"DLB\": \"大连银行\", \n  \"DRCBCL\": \"东莞农村商业银行\", \n  \"GCB\": \"广州银行\", \n  \"NBBANK\": \"宁波银行\", \n  \"BOYK\": \"营口银行\", \n  \"SXRCCU\": \"陕西信合\", \n  \"GLBANK\": \"桂林银行\", \n  \"BOQH\": \"青海银行\", \n  \"CDRCB\": \"成都农商银行\", \n  \"QDCCB\": \"青岛银行\", \n  \"HKBEA\": \"东亚银行\", \n  \"HBHSBANK\": \"湖北银行黄石分行\", \n  \"WZCB\": \"温州银行\", \n  \"TRCB\": \"天津农商银行\", \n  \"QLBANK\": \"齐鲁银行\", \n  \"GDRCC\": \"广东省农村信用社联合社\", \n  \"ZJTLCB\": \"浙江泰隆商业银行\", \n  \"GZB\": \"赣州银行\", \n  \"GYCB\": \"贵阳市商业银行\", \n  \"CQBANK\": \"重庆银行\", \n  \"DAQINGB\": \"龙江银行\", \n  \"CGNB\": \"南充市商业银行\", \n  \"SCCB\": \"三门峡银行\", \n  \"CSRCB\": \"常熟农村商业银行\", \n  \"SHBANK\": \"上海银行\", \n  \"JLBANK\": \"吉林银行\", \n  \"CZRCB\": \"常州农村信用联社\", \n  \"BANKWF\": \"潍坊银行\", \n  \"ZRCBANK\": \"张家港农村商业银行\", \n  \"FJHXBC\": \"福建海峡银行\", \n  \"ZJNX\": \"浙江省农村信用社联合社\", \n  \"LZYH\": \"兰州银行\", \n  \"JSB\": \"晋商银行\", \n  \"BOHAIB\": \"渤海银行\", \n  \"CZCB\": \"浙江稠州商业银行\", \n  \"YQCCB\": \"阳泉银行\", \n  \"SJBANK\": \"盛京银行\", \n  \"XABANK\": \"西安银行\", \n  \"BSB\": \"包商银行\", \n  \"JSBANK\": \"江苏银行\", \n  \"FSCB\": \"抚顺银行\", \n  \"HNRCU\": \"河南省农村信用\", \n  \"COMM\": \"交通银行\", \n  \"XTB\": \"邢台银行\", \n  \"CITIC\": \"中信银行\", \n  \"HXBANK\": \"华夏银行\", \n  \"HNRCC\": \"湖南省农村信用社\", \n  \"DYCCB\": \"东营市商业银行\", \n  \"ORBANK\": \"鄂尔多斯银行\", \n  \"BJRCB\": \"北京农村商业银行\", \n  \"XYBANK\": \"信阳银行\", \n  \"ZGCCB\": \"自贡市商业银行\", \n  \"CDCB\": \"成都银行\", \n  \"HANABANK\": \"韩亚银行\", \n  \"CMBC\": \"中国民生银行\", \n  \"LYBANK\": \"洛阳银行\", \n  \"GDB\": \"广东发展银行\", \n  \"ZBCB\": \"齐商银行\", \n  \"CBKF\": \"开封市商业银行\", \n  \"H3CB\": \"内蒙古银行\", \n  \"CIB\": \"兴业银行\", \n  \"CRCBANK\": \"重庆农村商业银行\", \n  \"SZSBK\": \"石嘴山银行\", \n  \"DZBANK\": \"德州银行\", \n  \"SRBANK\": \"上饶银行\", \n  \"LSCCB\": \"乐山市商业银行\", \n  \"JXRCU\": \"江西省农村信用\", \n  \"ICBC\": \"中国工商银行\", \n  \"JZBANK\": \"晋中市商业银行\", \n  \"HZCCB\": \"湖州市商业银行\", \n  \"NHB\": \"南海农村信用联社\", \n  \"XXBANK\": \"新乡银行\", \n  \"JRCB\": \"江苏江阴农村商业银行\", \n  \"YNRCC\": \"云南省农村信用社\", \n  \"ABC\": \"中国农业银行\", \n  \"GXRCU\": \"广西省农村信用\", \n  \"PSBC\": \"中国邮政储蓄银行\", \n  \"BZMD\": \"驻马店银行\", \n  \"ARCU\": \"安徽省农村信用社\", \n  \"GSRCU\": \"甘肃省农村信用\", \n  \"LYCB\": \"辽阳市商业银行\", \n  \"JLRCU\": \"吉林农信\", \n  \"URMQCCB\": \"乌鲁木齐市商业银行\", \n  \"XLBANK\": \"中山小榄村镇银行\", \n  \"CSCB\": \"长沙银行\", \n  \"JHBANK\": \"金华银行\", \n  \"BHB\": \"河北银行\", \n  \"NBYZ\": \"鄞州银行\", \n  \"LSBC\": \"临商银行\", \n  \"BOCD\": \"承德银行\", \n  \"SDRCU\": \"山东农信\", \n  \"NCB\": \"南昌银行\", \n  \"TCCB\": \"天津银行\", \n  \"WJRCB\": \"吴江农商银行\", \n  \"CBBQS\": \"城市商业银行资金清算中心\", \n  \"HBRCU\": \"河北省农村信用社\"\n}"

/**
** 根据银行卡号获取银行信息
 */
func GetBankCodeByBankCardNo(bankCardNo string) string {
	bUrl := bankUrl + bankCardNo
	request := httplib.Get(bUrl)
	s, err := request.String()
	if err != nil {
		logs.Error("获取银行信息失败：" + err.Error())
		return ""
	}

	logs.Info("获取到银行信息为：" + s)
	code := gojson.Json(s).Get("bank").Tostring()
	return code

}

/**
** 根据银行卡号获取银行名称
 */
func GetBankNameByCardNo(cardNo string) string {
	code := GetBankCodeByBankCardNo(cardNo)
	if code == "" {
		return code
	}

	return GetBankNameByCode(code)
}

/**
** 根据银行编码获取银行名称
 */
func GetBankNameByCode(code string) string {
	return gojson.Json(CodeToName).Get(code).Tostring()
}
