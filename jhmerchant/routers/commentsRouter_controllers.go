package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:DealExcel"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:DealExcel"],
		beego.ControllerComments{
			Method: "DownloadExcelModel",
			Router: `/excel/download`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:DealExcel"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:DealExcel"],
		beego.ControllerComments{
			Method: "DownloadRecordExcel",
			Router: `/excel/download_excel/?:params`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:DealExcel"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:DealExcel"],
		beego.ControllerComments{
			Method: "MakeComplaintExcel",
			Router: `/excel/make_complaint_record_excel/?:params`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:DealExcel"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:DealExcel"],
		beego.ControllerComments{
			Method: "MakeOrderExcel",
			Router: `/excel/make_order_excel/?:params`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:DealExcel"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:DealExcel"],
		beego.ControllerComments{
			Method: "MakeWithdrawExcel",
			Router: `/excel/make_withdraw_record_excel/?:params`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:History"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:History"],
		beego.ControllerComments{
			Method: "HistoryQueryAndListPage",
			Router: `/history/list_history_record/?:params`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:History"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:History"],
		beego.ControllerComments{
			Method: "ShowHistoryListUI",
			Router: `/history/show_history_list_ui`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"],
		beego.ControllerComments{
			Method: "LoadUserAccountInfo",
			Router: `/index/loadInfo/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"],
		beego.ControllerComments{
			Method: "LoadOrderCount",
			Router: `/index/loadOrders`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"],
		beego.ControllerComments{
			Method: "LoadCountOrder",
			Router: `/index/load_count_order`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"],
		beego.ControllerComments{
			Method: "LoadUserPayWay",
			Router: `/index/pay_way`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"],
		beego.ControllerComments{
			Method: "LoadUserPayWayUI",
			Router: `/index/show_pay_way_ui`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Index"],
		beego.ControllerComments{
			Method: "ShowUI",
			Router: `/index/ui/`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"],
		beego.ControllerComments{
			Method: "FlushCaptcha",
			Router: `/flushCaptcha.py/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"],
		beego.ControllerComments{
			Method: "UserLogin",
			Router: `/login.py/?:params`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"],
		beego.ControllerComments{
			Method: "LoginOut",
			Router: `/loginOut.py`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"],
		beego.ControllerComments{
			Method: "PayDoc",
			Router: `/pay_doc.py`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Login"],
		beego.ControllerComments{
			Method: "VerifyCaptcha",
			Router: `/verifyCaptcha.py/:value/:chaId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:MultiWithdraw"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:MultiWithdraw"],
		beego.ControllerComments{
			Method: "ShowMultiWithdrawUI",
			Router: `/multi_withdraw/show_multi_ui`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:TradeRecord"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:TradeRecord"],
		beego.ControllerComments{
			Method: "ComplaintQueryAndListPage",
			Router: `/trade/complaint/?:params`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:TradeRecord"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:TradeRecord"],
		beego.ControllerComments{
			Method: "TradeQueryAndListPage",
			Router: `/trade/list/?:params`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:TradeRecord"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:TradeRecord"],
		beego.ControllerComments{
			Method: "ShowComplaintUI",
			Router: `/trade/show_complaint_ui`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:TradeRecord"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:TradeRecord"],
		beego.ControllerComments{
			Method: "ShowUI",
			Router: `/trade/show_ui`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:UserInfo"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:UserInfo"],
		beego.ControllerComments{
			Method: "ConfirmOriginPwd",
			Router: `/user_info/confirm_pwd/?:params`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:UserInfo"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:UserInfo"],
		beego.ControllerComments{
			Method: "ModifyUserInfo",
			Router: `/user_info/modify_userInfo/?:params`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:UserInfo"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:UserInfo"],
		beego.ControllerComments{
			Method: "ShowModifyUserInfoUI",
			Router: `/user_info/show_modify_ui`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:UserInfo"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:UserInfo"],
		beego.ControllerComments{
			Method: "ShowUserInfoUI",
			Router: `/user_info/show_ui`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"],
		beego.ControllerComments{
			Method: "LaunchMultiWithdraw",
			Router: `/multi_withdraw/launch_multi_withdraw/?:params`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"],
		beego.ControllerComments{
			Method: "SendMsgForMultiWithdraw",
			Router: `/multi_withdraw/send_msg_for_multi/?:params`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"],
		beego.ControllerComments{
			Method: "UserBalance",
			Router: `/withdraw/balance`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"],
		beego.ControllerComments{
			Method: "LaunchSingleWithdraw",
			Router: `/withdraw/launch_single_withdraw/?:params`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"],
		beego.ControllerComments{
			Method: "WithdrawQueryAndListPage",
			Router: `/withdraw/list_record/?:params`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"],
		beego.ControllerComments{
			Method: "SendMsgForWithdraw",
			Router: `/withdraw/send_msg/?:params`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"],
		beego.ControllerComments{
			Method: "ShowListUI",
			Router: `/withdraw/show_list_ui`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"] = append(beego.GlobalControllerRouter["dongfeng-pay/jhmerchant/controllers:Withdraw"],
		beego.ControllerComments{
			Method: "ShowWithdrawUI",
			Router: `/withdraw/show_ui`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
