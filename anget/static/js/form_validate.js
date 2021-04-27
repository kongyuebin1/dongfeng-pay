/***************************************************
 ** @Desc : This file for 表单验证js
 ** @Time : 19.12.3 11:17
 ** @Author : Joker
 ** @File : form_validate
 ** @Last Modified by : Joker
 ** @Last Modified time: 19.12.3 11:17
 ** @Software: GoLand
 ****************************************************/

let form_v = {
    modify_userInfo: function () {
        let or_pwd = $("#or_pwd").val();
        let new_pwd = $("#new_pwd").val();
        let confirm_pwd = $("#confirm_pwd").val();
        let patrn = /^[a-zA-Z]{1}([a-zA-Z0-9]|[._]){5,19}$/;
        if (or_pwd === "" || new_pwd === "" || confirm_pwd === "") {
            toastr.error("密码不能为空!");
            return
        }
        if (!patrn.exec(new_pwd) || !patrn.exec(confirm_pwd)) {
            toastr.error("密码只能输入6-20个以字母开头、可带数字、“_”、“.”的字串!");
            return
        }
        if (new_pwd !== confirm_pwd) {
            toastr.error("两次密码不匹配!");
            return
        }
        $.ajax({
            type: "POST",
            url: "/user_info/confirm_pwd/",
            data: {c: or_pwd,},
            cache: false,
            success: function (res) {
                if (res.code === -9) {
                    toastr.error(res.msg)
                } else {
                    $.ajax({
                        type: "POST",
                        url: "/user_info/modify_userInfo/",
                        data: $("#modify_userInfo").serialize(),
                        cache: false,
                        success: function (res) {
                            if (res.code === 9) {
                                toastr.success(res.msg);
                                setTimeout(function () {
                                    window.location.reload()
                                }, 3000)
                            } else {
                                toastr.error(res.msg)
                            }
                        },
                        error: function (XMLHttpRequest) {
                            toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
                        }
                    })
                }
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }, modify_pay_password: function () {
        let or_pwd = $("#or_pwd").val();
        let new_pwd = $("#new_pwd").val();
        let confirm_pwd = $("#confirm_pwd").val();
        let patrn = /^[a-zA-Z]{1}([a-zA-Z0-9]|[._]){5,19}$/;
        if (new_pwd === "" || confirm_pwd === "") {
            toastr.error("密码不能为空!");
            return
        }
        if (!patrn.exec(new_pwd) || !patrn.exec(confirm_pwd)) {
            toastr.error("密码只能输入6-20个以字母开头、可带数字、“_”、“.”的字串!");
            return
        }
        if (new_pwd !== confirm_pwd) {
            toastr.error("两次密码不匹配!");
            return
        }
        $.ajax({
            type: "POST",
            url: "/user_info/confirm_pay_pwd/",
            data: {c: or_pwd,},
            cache: false,
            success: function (res) {
                if (res.code === -9) {
                    toastr.error(res.msg)
                } else {
                    $.ajax({
                        type: "POST",
                        url: "/user_info/set_pay_password/",
                        data: $("#set_pay_password").serialize(),
                        cache: false,
                        success: function (res) {
                            if (res.code === 9) {
                                toastr.success(res.msg);
                                setTimeout(function () {
                                    window.location.reload()
                                }, 3000)
                            } else {
                                toastr.error(res.msg)
                            }
                        },
                        error: function (XMLHttpRequest) {
                            toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
                        }
                    })
                }
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }, confirm_origin_pwd: function () {
        let or_pwd = $("#or_pwd").val();
        if (or_pwd === "") {
            toastr.error("原始密码不能为空!");
            return
        }
        $.ajax({
            type: "POST",
            url: "/user_info/confirm_pwd/",
            data: {c: or_pwd,},
            cache: false,
            success: function (res) {
                if (res.code === -9) {
                    toastr.error(res.msg)
                }
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }, confirm_new_pwd: function () {
        let new_pwd = $("#new_pwd").val();
        let confirm_pwd = $("#confirm_pwd").val();
        let patrn = /^[a-zA-Z]{1}([a-zA-Z0-9]|[._]){5,19}$/;
        if (new_pwd === "" || confirm_pwd === "") {
            toastr.error("新密码不能为空!");
            return
        }
        if (!patrn.exec(new_pwd) || !patrn.exec(confirm_pwd)) {
            toastr.error("密码只能输入6-20个以字母开头、可带数字、“_”、“.”的字串!");
            return
        }
        if (new_pwd !== confirm_pwd) {
            toastr.error("两次密码不匹配!")
        }
    }, launch_single_withdraw: function () {
        let balance = $("#balance").val();
        let bankCode = $("#bankCode").val();
        let accountName = $("#accountName").val();
        let cardNo = $("#cardNo").val();
        let bankAccountType = $("#bankAccountType").val();
        let province = $("#province").val();
        let city = $("#city").val();
        let bankAccountAddress = $("#bankAccountAddress").val();
        let moblieNo = $("#moblieNo").val();
        let amount = $("#amount").val();
        let smsVerifyCode = $("#smsVerifyCode").val();
        let patrn = /^(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*))$/;
        let patrn2 = /^[1]([3-9])[0-9]{9}$/;
        if (bankCode === "" || accountName === "" || cardNo === "") {
            toastr.error("银行名、账户名或卡号不能为空!");
            return
        }
        if (amount === "" || moblieNo === "") {
            toastr.error("手机号或金额不能为空!");
            return
        }
        if (!patrn2.exec(moblieNo)) {
            toastr.error("请输入正确的手机号!");
            return
        }
        if (!patrn.exec(amount)) {
            toastr.error("请输入正确的金额!");
            return
        }
        if ("PUBLIC_ACCOUNT" === bankAccountType) {
            if (province === "" || city === "" || bankAccountAddress === "") {
                toastr.error("开户行全称、所在省份或所在城市不能为空!");
                return
            }
        }
        if (parseInt(amount) > parseInt(balance) || parseInt(amount) > 50000) {
            toastr.error("提现金额超出限制!");
            return
        }
        if (smsVerifyCode === "") {
            toastr.error("支付密码不能为空!");
            return
        }
        $.ajax({
            type: "POST",
            url: "/withdraw/launch_single_withdraw/",
            data: $("#withdraw").serialize(),
            success: function (resp) {
                if (resp.code === 9) {
                    swal(resp.msg, {icon: "success", type: "success", closeOnClickOutside: false,}).then((con) => {
                        if (con) {
                            if (resp.url !== "") {
                                window.location.href = resp.url
                            }
                        }
                    })
                } else {
                    toastr.error(resp.msg)
                }
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }, launch_multi_withdraw: function () {
        let file = $("#file").val();
        let smsVerifyCode = $("#smsVerifyCode").val();
        let pos = file.lastIndexOf(".");
        let lastname = file.substring(pos, file.length);
        if (lastname.toLowerCase() !== ".xls" && lastname.toLowerCase() !== ".xlsx") {
            toastr.error("仅支持“xls”、“xlsx”格式文件!");
            return
        }
        if (smsVerifyCode === "") {
            toastr.error("支付密码不能为空!");
            return
        }
        let multi_withdraw = document.getElementById('multi_withdraw'), formData = new FormData(multi_withdraw);
        $.ajax({
            type: "POST",
            url: '/multi_withdraw/launch_multi_withdraw/',
            data: formData,
            processData: false,
            contentType: false,
            success: function (res) {
                if (res.code === 9) {
                    swal(res.msg, {icon: "success", type: "success", closeOnClickOutside: false,}).then((con) => {
                        if (con) {
                            if (res.url !== "") {
                                window.location.href = res.url
                            }
                        }
                    })
                } else {
                    swal(res.msg, {icon: "warning", closeOnClickOutside: false,})
                }
            },
            error: function (XMLHttpRequest) {
                toastr.info('something is wrong, code: ' + XMLHttpRequest.status)
            }
        })
    }
};