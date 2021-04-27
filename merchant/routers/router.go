package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dchest/captcha"
	"merchant/controllers"
)

func init() {
	beego.Router("/", &controllers.Login{}, "*:Index")
	beego.Router("/index/ui/", &controllers.Index{}, "*:ShowUI")
	//生产登录验证码
	beego.Handler("/img.do/*.png", captcha.Server(130, 40))
	beego.Router("/history/show_history_list_ui", &controllers.History{}, "*:ShowHistoryListUI")
	beego.Router("/history/list_history_record/?:params", &controllers.History{}, "*:HistoryQueryAndListPage")

	beego.Router("/excel/download", &controllers.DealExcel{}, "*:DownloadExcelModel")
	beego.Router("/excel/make_order_excel/?:params",&controllers.DealExcel{}, "*:MakeOrderExcel" )
	beego.Router("/excel/download_excel/?:params", &controllers.DealExcel{}, "*:DownloadRecordExcel")
	beego.Router("/excel/make_complaint_record_excel/?:params", &controllers.DealExcel{}, "*:MakeComplaintExcel")
	beego.Router("/excel/make_withdraw_record_excel/?:params", &controllers.DealExcel{}, "*:MakeWithdrawExcel")

	beego.Router("/index/ui/", &controllers.Index{}, "*:ShowUI")
	beego.Router("/index/loadInfo/", &controllers.Index{}, "*:LoadUserAccountInfo")
	beego.Router("/index/load_count_order", &controllers.Index{}, "*:LoadCountOrder")
	beego.Router("index/loadOrders", &controllers.Index{}, "*:LoadOrderCount")
	beego.Router("/index/show_pay_way_ui", &controllers.Index{}, "*:LoadUserPayWayUI")
	beego.Router("/index/pay_way", &controllers.Index{}, "*:LoadUserPayWay")

	beego.Router("/login.py/?:params", &controllers.Login{}, "*:UserLogin")
	beego.Router("/verifyCaptcha.py/:value/:chaId", &controllers.Login{}, "*:VerifyCaptcha")
	beego.Router("/flushCaptcha.py/", &controllers.Login{}, "*:FlushCaptcha")
	beego.Router("/loginOut.py", &controllers.Login{}, "*:LoginOut")
	beego.Router("/pay_doc.py", &controllers.Login{}, "*:PayDoc")

	beego.Router("/multi_withdraw/show_multi_ui", &controllers.MultiWithdraw{}, "*:ShowMultiWithdrawUI")
	beego.Router("/multi_withdraw/send_msg_for_multi/?:params", &controllers.Withdraw{}, "*:SendMsgForMultiWithdraw")
	beego.Router("/multi_withdraw/launch_multi_withdraw/?:params", &controllers.Withdraw{}, "*:LaunchMultiWithdraw")

	beego.Router("/trade/show_ui", &controllers.TradeRecord{}, "*:ShowUI")
	beego.Router("/trade/list/?:params", &controllers.TradeRecord{}, "*:TradeQueryAndListPage")
	beego.Router("/trade/show_complaint_ui", &controllers.TradeRecord{}, "*:ShowComplaintUI")
	beego.Router("/trade/complaint/?:params", &controllers.TradeRecord{}, "*:ComplaintQueryAndListPage")

	beego.Router("/user_info/show_modify_ui", &controllers.UserInfo{}, "*:ShowModifyUserInfoUI")
	beego.Router("/user_info/modify_userInfo/?:params", &controllers.UserInfo{}, "*:ModifyUserInfo")
	beego.Router("/user_info/confirm_pwd/?:params", &controllers.UserInfo{}, "*:ConfirmOriginPwd")
	beego.Router("/user_info/show_ui", &controllers.UserInfo{}, "*:ShowUserInfoUI")

	beego.Router("/withdraw/show_ui", &controllers.Withdraw{}, "*:ShowWithdrawUI")
	beego.Router("/withdraw/balance", &controllers.Withdraw{}, "*:UserBalance")
	beego.Router("/withdraw/send_msg/?:params", &controllers.Withdraw{}, "*:SendMsgForWithdraw")
	beego.Router("/withdraw/launch_single_withdraw/?:params", &controllers.Withdraw{}, "*:LaunchSingleWithdraw")
	beego.Router("/withdraw/show_list_ui", &controllers.Withdraw{}, "*:ShowListUI")
	beego.Router("/withdraw/list_record/?:params", &controllers.Withdraw{}, "*:WithdrawQueryAndListPage")
}
