/***************************************************
 ** @Desc : This file for 导出Excel文件
 ** @Time : 19.12.7 16:48
 ** @Author : Joker
 ** @File : deal_excel
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.7 16:48
 ** @Software: GoLand
 ****************************************************/

let excel = {
    download_trade_order_excel: function () {
        let merchantName = $("#merchant_Name").val();
        let startTime = $("#startTime").val();
        let endTime = $("#endTime").val();
        let payType = $("#payType").val();
        let uStatus = $("#uStatus").val();
        $.ajax({
            type: "GET",
            url: "/excel/make_order_excel/",
            data: {start: startTime, end: endTime, pay_type: payType, status: uStatus, merchantName: merchantName,},
            cache: true,
            success: function (res) {
                if (res.code === 9) {
                    let $form = $("<form method='get'></form>");
                    $form.attr("action", "/excel/download_excel/" + res.msg);
                    $(document.body).append($form);
                    $form.submit()
                } else {
                    toastr.error(res.msg)
                }
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }, download_complaint_record_excel: function () {
        let startTime = $("#startTime").val();
        let merchantName = $("#merchant_Name").val();
        let endTime = $("#endTime").val();
        let payType = $("#payType").val();
        let uStatus = $("#uStatus").val();
        $.ajax({
            type: "GET",
            url: "/excel/make_complaint_record_excel/",
            data: {start: startTime, end: endTime, pay_type: payType, status: uStatus, merchantName: merchantName,},
            cache: true,
            success: function (res) {
                if (res.code === 9) {
                    let $form = $("<form method='get'></form>");
                    $form.attr("action", "/excel/download_excel/" + res.msg);
                    $(document.body).append($form);
                    $form.submit()
                } else {
                    toastr.error(res.msg)
                }
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }, download_withdraw_record_excel: function () {
        let startTime = $("#startTime").val();
        let endTime = $("#endTime").val();
        let uStatus = $("#uStatus").val();
        $.ajax({
            type: "GET",
            url: "/excel/make_withdraw_record_excel/",
            data: {start: startTime, end: endTime, status: uStatus,},
            cache: true,
            success: function (res) {
                if (res.code === 9) {
                    let $form = $("<form method='get'></form>");
                    $form.attr("action", "/excel/download_excel/" + res.msg);
                    $(document.body).append($form);
                    $form.submit()
                } else {
                    toastr.error(res.msg)
                }
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }, download_recharge_record_excel: function () {
        let startTime = $("#startTime").val();
        let endTime = $("#endTime").val();
        let uStatus = $("#uStatus").val();
        $.ajax({
            type: "GET",
            url: "/excel/make_recharge_record_excel/",
            data: {start: startTime, end: endTime, status: uStatus,},
            cache: true,
            success: function (res) {
                if (res.code === 9) {
                    let $form = $("<form method='get'></form>");
                    $form.attr("action", "/excel/download_excel/" + res.msg);
                    $(document.body).append($form);
                    $form.submit()
                } else {
                    toastr.error(res.msg)
                }
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }
};