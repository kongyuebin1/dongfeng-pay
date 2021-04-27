/***************************************************
 ** @Desc : This file for 交易记录js
 ** @Time : 19.12.3 15:01
 ** @Author : Joker
 ** @File : trade_record
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.3 15:01
 ** @Software: GoLand
 ****************************************************/

let trade = {
    get_last_month_date: function () {
        let date = new Date();
        let daysInMonth = [0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
        let strYear = date.getFullYear();
        let strDay = date.getDate();
        let strMonth = date.getMonth() + 1;
        let hh = date.getHours();
        let mm = date.getMinutes();
        let ss = date.getSeconds();
        if (((strYear % 4) === 0) && ((strYear % 100) !== 0) || ((strYear % 400) === 0)) {
            daysInMonth[2] = 29
        }
        if (strMonth - 1 === 0) {
            strYear -= 1;
            strMonth = 12
        } else {
            strMonth -= 1
        }
        strDay = Math.min(strDay, daysInMonth[strMonth]);
        if (strMonth < 10) {
            strMonth = "0" + strMonth
        }
        if (strDay < 10) {
            strDay = "0" + strDay
        }
        if (hh < 10) {
            hh = "0" + hh
        }
        if (mm < 10) {
            mm = "0" + mm
        }
        if (ss < 10) {
            ss = "0" + ss
        }
        return strYear + "-" + strMonth + "-" + strDay + " " + hh + ":" + mm + ":" + ss
    }, get_time: function (d) {
        let date = new Date(d);
        let strYear = date.getFullYear();
        let strDay = date.getDate();
        let strMonth = date.getMonth() + 1;
        let hh = date.getHours();
        let mm = date.getMinutes();
        let ss = date.getSeconds();
        if (strMonth < 10) {
            strMonth = "0" + strMonth
        }
        if (strDay < 10) {
            strDay = "0" + strDay
        }
        if (hh < 10) {
            hh = "0" + hh
        }
        if (mm < 10) {
            mm = "0" + mm
        }
        if (ss < 10) {
            ss = "0" + ss
        }
        return strYear + "-" + strMonth + "-" + strDay + " " + hh + ":" + mm + ":" + ss
    }, trade_do_paging: function () {
        let merchantNo = $("#merchant_No").val();
        let startTime = $("#startTime").val();
        let endTime = $("#endTime").val();
        let payType = $("#payType").val();
        let uStatus = $("#uStatus").val();
        $.ajax({
            type: "GET",
            url: "/trade/list/",
            data: {
                page: '1',
                limit: "15",
                MerchantNo: merchantNo,
                start: startTime,
                end: endTime,
                pay_type: payType,
                status: uStatus,
            },
            success: function (data) {
                trade.show_trade_data(data.root);
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
                            url: "/trade/list/",
                            type: "GET",
                            data: {
                                page: page,
                                MerchantNo: merchantNo,
                                start: startTime,
                                end: endTime,
                                pay_type: payType,
                                status: uStatus,
                            },
                            success: function (data) {
                                trade.show_trade_data(data.root)
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
    }, show_trade_data: function (list) {
        let con = "";
        $.each(list, function (index, item) {
            let bg_red = "", st = "", t = "";
            switch (item.Status) {
                case"failed":
                    st = "交易失败";
                    break;
                case"wait":
                    st = "等待支付";
                    break;
                case"success":
                    bg_red = ` style="color: green"`;
                    st = "交易成功";
                    t = item.UpdateTime;
                    break
            }
            con += `<tr><th scope="row">` + (index + 1) + `</th><td>` + item.BankOrderId + `</td><td>` + item.MerchantOrderId + `</td><td>` + item.PayProductName + `</td><td>` + item.OrderAmount.toFixed(2) + `</td><td>` + item.UserInAmount.toFixed(2) + `</td><td>` + item.AllProfit.toFixed(2) + `</td><td` + bg_red + `>` + st + `</td><td>` + t + `</td></tr>`
        });
        if (con === "") {
            con += `<tr><td colspan="9">没有检索到数据</td></tr>`
        }
        $("#your_show_time").html(con)
    }, complaint_do_paging: function () {
        let merchantNo = $("#merchant_No").val();
        let startTime = $("#startTime").val();
        let endTime = $("#endTime").val();
        let payType = $("#payType").val();
        let uStatus = $("#uStatus").val();
        $.ajax({
            type: "GET",
            url: "/trade/complaint/",
            data: {
                page: '1',
                limit: "15",
                MerchantNo: merchantNo,
                start: startTime,
                end: endTime,
                pay_type: payType,
                status: uStatus,
            },
            success: function (data) {
                trade.show_complaint_data(data.root);
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
                            url: "/trade/complaint/",
                            type: "GET",
                            data: {
                                page: page,
                                MerchantNo: merchantNo,
                                start: startTime,
                                end: endTime,
                                pay_type: payType,
                                status: uStatus,
                            },
                            success: function (data) {
                                trade.show_complaint_data(data.root)
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
    }, show_complaint_data: function (list) {
        let con = "";
        $.each(list, function (index, item) {
            let st = "";
            switch (item.FreezeOrder) {
                case"yes":
                    st = "已冻结";
                    break;
                case"no":
                    st = "已退款";
                    break
            }
            con += `<tr><th scope="row">` + (index + 1) + `</th><td>` + item.BankOrderId + `</td><td>` + item.MerchantOrderId + `</td><td>` + item.PayProductName + `</td><td>` + item.OrderAmount.toFixed(2) + `</td><td>` + st + `</td><td>` + item.UpdateTime + `</td></tr>`
        });
        if (con === "") {
            con += `<tr><td colspan="9">没有检索到数据</td></tr>`
        }
        $("#your_show_time").html(con)
    },
};