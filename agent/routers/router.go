package routers

import (
	"agent/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dchest/captcha"
)

func init() {
	//生产登录验证码
	beego.Handler("/img.do/*.png", captcha.Server(130, 40))
	beego.Router("/", &controllers.Login{}, "*:Index")
	beego.Router("/login.py/?:params", &controllers.Login{}, "post:UserLogin")
	beego.Router("/verifyCaptcha.py/:value/:chaId", &controllers.Login{}, "get:VerifyCaptcha")
	beego.Router("/flushCaptcha.py/", &controllers.Login{}, "get:FlushCaptcha")
	beego.Router("/loginOut.py", &controllers.Login{}, "*:LoginOut")
	beego.Router("/pay_doc.py", &controllers.Login{}, "*:PayDoc")

	beego.Router("/history/show_history_list_ui", &controllers.History{}, "*:ShowHistoryListUI")
	beego.Router("/history/list_history_record/?:params", &controllers.History{}, "*:HistoryQueryAndListPage")

	beego.Router("/excel/download", &controllers.DealExcel{}, "*:DownloadExcelModel")
	beego.Router("/excel/make_order_excel/?:params", &controllers.DealExcel{}, "get:MakeOrderExcel")
	beego.Router("/excel/download_excel/?:params", &controllers.DealExcel{}, "get:DownloadRecordExcel")
	beego.Router("/excel/make_complaint_record_excel/?:params", &controllers.DealExcel{}, "get:MakeComplaintExcel")
	beego.Router("/excel/make_withdraw_record_excel/?:params", &controllers.DealExcel{}, "get:MakeWithdrawExcel")

	beego.Router("/index/ui/", &controllers.Index{}, "*:ShowUI")
	beego.Router("/index/loadInfo/", &controllers.Index{}, "*:LoadUserAccountInfo")
	beego.Router("/index/load_count_order", &controllers.Index{}, "*:LoadCountOrder")
	beego.Router("/index/loadOrders", &controllers.Index{}, "*:LoadOrderCount")
	beego.Router("/index/show_pay_way_ui", &controllers.Index{}, "*:LoadUserPayWayUI")
	beego.Router("/index/pay_way", &controllers.Index{}, "*:LoadUserPayWay")

	beego.Router("/multi_withdraw/show_multi_ui", &controllers.MultiWithdraw{}, "*:ShowMultiWithdrawUI")
	beego.Router("/multi_withdraw/launch_multi_withdraw/?:params", &controllers.Withdraw{}, "post:LaunchMultiWithdraw")

	beego.Router("/trade/show_ui", &controllers.TradeRecord{}, "*:ShowUI")
	beego.Router("/trade/list/?:params", &controllers.TradeRecord{}, "*:TradeQueryAndListPage")
	beego.Router("/trade/show_complaint_ui", &controllers.TradeRecord{}, "*:ShowComplaintUI")
	beego.Router("/trade/complaint/?:params", &controllers.TradeRecord{}, "*:ComplaintQueryAndListPage")

	beego.Router("/user_info/show_modify_ui", &controllers.UserInfo{}, "*:ShowModifyUserInfoUI")
	beego.Router("/user_info/modify_userInfo/?:params", &controllers.UserInfo{}, "*:ModifyUserInfo")
	beego.Router("/user_info/confirm_pwd/?:params", &controllers.UserInfo{}, "*:ConfirmOriginPwd")
	beego.Router("/user_info/set_pay_password/?:params", &controllers.UserInfo{}, "*:SetPayPassword")
	beego.Router("/user_info/confirm_pay_pwd/?:params", &controllers.UserInfo{}, "*:ConfirmOriginPayPwd")
	beego.Router("/user_info/show_ui", &controllers.UserInfo{}, "*:ShowUserInfoUI")
	beego.Router("/user_info/show_merchant_ui", &controllers.UserInfo{}, "*:ShowMerchantUI")
	beego.Router("//user_info/merchant_list/?:params", &controllers.UserInfo{}, "*:MerchantQueryAndListPage")

	beego.Router("/withdraw/show_ui", &controllers.Withdraw{}, "*:ShowWithdrawUI")
	beego.Router("/withdraw/balance", &controllers.Withdraw{}, "*:UserBalance")
	beego.Router("/withdraw/launch_single_withdraw/?:params", &controllers.Withdraw{}, "post:LaunchSingleWithdraw")
	beego.Router("/withdraw/show_list_ui", &controllers.Withdraw{}, "*:ShowListUI")
	beego.Router("/withdraw/list_record/?:params", &controllers.Withdraw{}, "*:WithdrawQueryAndListPage")
}
