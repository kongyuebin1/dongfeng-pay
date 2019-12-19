/***************************************************
 ** @Desc : This file for ...
 ** @Time : 19.12.6 13:43
 ** @Author : Joker
 ** @File : withdraw_record
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.6 13:43
 ** @Software: GoLand
 ****************************************************/

let pay = {
    withdraw_do_paging: function () {
        let bankNo = $("#bankNo").val();
        let merchantNo = $("#merchant_No").val();
        let startTime = $("#startTime").val();
        let endTime = $("#endTime").val();
        let uStatus = $("#uStatus").val();
        $.ajax({
            type: "GET",
            url: "/withdraw/list_record/",
            data: {
                page: '1',
                limit: "15",
                MerchantNo: merchantNo,
                BankNo: bankNo,
                start: startTime,
                end: endTime,
                status: uStatus,
            },
            success: function (data) {
                pay.show_withdraw_data(data.root);
                let options = {
                    bootstrapMajorVersion: 3,
                    currentPage: data.page,
                    totalPages: data.totalPage,
                    numberOfPages: data.limit,
                    itemTexts: function (type, page) {
                        switch (type) {
                            case"first":
                                return "首页";
                            case"prev":
                                return "上一页";
                            case"next":
                                return "下一页";
                            case"last":
                                return "末页";
                            case"page":
                                return page
                        }
                    },
                    onPageClicked: function (event, originalEvent, type, page) {
                        $.ajax({
                            url: "/withdraw/list_record/",
                            type: "GET",
                            data: {
                                page: page,
                                MerchantNo: merchantNo,
                                BankNo: bankNo,
                                start: startTime,
                                end: endTime,
                                status: uStatus,
                            },
                            success: function (data) {
                                pay.show_withdraw_data(data.root)
                            }
                        })
                    }
                };
                $('#do_paging').bootstrapPaginator(options)
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }, show_withdraw_data: function (list) {
        let con = "";
        $.each(list, function (index, item) {
            let bg_red = "", st, t = "";
            switch (item.Status) {
                case"payfor_confirm":
                    st = "等待审核";
                    break;
                case"payfor_solving":
                    st = "系统处理中";
                    break;
                case"payfor_banking":
                    st = "银行处理中";
                    break;
                case"failed":
                    st = "代付失败";
                    break;
                case"success":
                    bg_red = ` style="color: green"`;
                    st = "打款成功";
                    t = trade.get_time(item.UpdateTime);
                    break;
                default:
                    st = ""
            }
            con += `<tr><th scope="row">` + (index + 1) + `</th><td>` + item.BankOrderId + `</td><td>` + item.MerchantOrderId + `</td><td>` + item.PayforTotalAmount.toFixed(2) + `</td><td>` + item.PayforFee.toFixed(2) + `</td><td>` + item.BankName + `</td><td>` + item.BankAccountName + `</td><td>` + item.BankAccountNo + `</td><td` + bg_red + `>` + st + `</td><td>` + item.CreateTime + `</td><td>` + t + `</td><td>` + item.Remark + `</td></tr>`
        });
        if (con === "") {
            con += `<tr><td colspan="9">没有检索到数据</td></tr>`
        }
        $("#your_show_time").html(con)
    }, recharge_do_paging: function () {
        let startTime = $("#startTime").val();
        let endTime = $("#endTime").val();
        let uStatus = $("#uStatus").val();
        $.ajax({
            type: "GET",
            url: "/recharge/list_recharge_record/",
            data: {page: '1', limit: "15", start: startTime, end: endTime, status: uStatus,},
            success: function (data) {
                pay.show_recharge_data(data.root);
                let options = {
                    bootstrapMajorVersion: 3,
                    currentPage: data.page,
                    totalPages: data.totalPage,
                    numberOfPages: data.limit,
                    itemTexts: function (type, page) {
                        switch (type) {
                            case"first":
                                return "首页";
                            case"prev":
                                return "上一页";
                            case"next":
                                return "下一页";
                            case"last":
                                return "末页";
                            case"page":
                                return page
                        }
                    },
                    onPageClicked: function (event, originalEvent, type, page) {
                        $.ajax({
                            url: "/recharge/list_recharge_record/",
                            type: "GET",
                            data: {page: page, start: startTime, end: endTime, status: uStatus,},
                            success: function (data) {
                                pay.show_recharge_data(data.root)
                            }
                        })
                    }
                };
                $('#do_paging').bootstrapPaginator(options)
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }, show_recharge_data: function (list) {
        let con = "";
        $.each(list, function (index, item) {
            let st;
            switch (item.OperateType) {
                case"plus_amount":
                    st = "加款";
                    break;
                case"sub_amount":
                    st = "减款";
                    break;
                case"freeze_amount":
                    st = "冻结";
                    break;
                case"unfreeze_amount":
                    st = "解冻";
                    break;
                default:
                    st = ""
            }
            con += `<tr><th scope="row">` + (index + 1) + `</th><td>` + item.CreateTime + `</td><td>` + item.Amount.toFixed(2) + `</td><td>` + item.Fee.toFixed(2) + `</td><td>` + item.FreezeAmount.toFixed(2) + `</td><td>` + item.Recharge.toFixed(2) + `</td><td>` + st + `</td></tr>`
        });
        if (con === "") {
            con += `<tr><td colspan="9">没有检索到数据</td></tr>`
        }
        $("#your_show_time").html(con)
    }, history_do_paging: function () {
        let startTime = $("#startTime").val();
        let endTime = $("#endTime").val();
        let uStatus = $("#uStatus").val();
        $.ajax({
            type: "GET",
            url: "/history/list_history_record/",
            data: {page: '1', limit: "15", start: startTime, end: endTime, status: uStatus,},
            success: function (data) {
                pay.show_history_data(data.root);
                let options = {
                    bootstrapMajorVersion: 3,
                    currentPage: data.page,
                    totalPages: data.totalPage,
                    numberOfPages: data.limit,
                    itemTexts: function (type, page) {
                        switch (type) {
                            case"first":
                                return "首页";
                            case"prev":
                                return "上一页";
                            case"next":
                                return "下一页";
                            case"last":
                                return "末页";
                            case"page":
                                return page
                        }
                    },
                    onPageClicked: function (event, originalEvent, type, page) {
                        $.ajax({
                            url: "/history/list_history_record/",
                            type: "GET",
                            data: {page: page, start: startTime, end: endTime, status: uStatus,},
                            success: function (data) {
                                pay.show_history_data(data.root)
                            }
                        })
                    }
                };
                $('#do_paging').bootstrapPaginator(options)
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }, show_history_data: function (list) {
        let con = "";
        $.each(list, function (index, item) {
            let st = "";
            switch (item.Type) {
                case"plus_amount":
                    st = "加款";
                    break;
                case"sub_amount":
                    st = "减款";
                    break;
                case"freeze_amount":
                    st = "冻结";
                    break;
                case"unfreeze_amount":
                    st = "解冻";
                    break
            }
            con += `<tr><th scope="row">` + (index + 1) + `</th>
                    <td>` + item.CreateTime + `</td>
                    <td>` + item.Amount.toFixed(2) + `</td>
                    <td>` + item.Balance.toFixed(2) + `</td>
                    <td>` + st + `</td></tr>`
        });
        if (con === "") {
            con += `<tr><td colspan="9">没有检索到数据</td></tr>`
        }
        $("#your_show_time").html(con)
    },
};