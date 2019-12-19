/***************************************************
 ** @Desc : This file for ...基本js
 ** @Time : 2018-08-27 10:52:14
 ** @Author : Joker
 ** @File :
 ** @Last Modified by : Joker
 ** @Last Modified time: 2018-08-27 11:12:19
 ** @Software: HBuilder
 ****************************************************/

/*限制只能选中一种支付方式*/
$(function () {
    $("input[name=SCAN]").click(function () {
        $("input[name=WAP]").attr("checked", false);
        $("input[name=WY]").attr("checked", false);
        $("input[name=KJ]").attr("checked", false);
        $("input[name=H5]").attr("checked", false);
    });
    $("input[name=WAP]").click(function () {
        $("input[name=SCAN]").attr("checked", false);
        $("input[name=WY]").attr("checked", false);
        $("input[name=KJ]").attr("checked", false);
        $("input[name=H5]").attr("checked", false);
    });
    $("input[name=WY]").click(function () {
        $("input[name=WAP]").attr("checked", false);
        $("input[name=SCAN]").attr("checked", false);
        $("input[name=KJ]").attr("checked", false);
        $("input[name=H5]").attr("checked", false);
    });
    $("input[name=KJ]").click(function () {
        $("input[name=WAP]").attr("checked", false);
        $("input[name=WY]").attr("checked", false);
        $("input[name=SCAN]").attr("checked", false);
        $("input[name=H5]").attr("checked", false);
    });
    $("input[name=H5]").click(function () {
        $("input[name=WAP]").attr("checked", false);
        $("input[name=WY]").attr("checked", false);
        $("input[name=SCAN]").attr("checked", false);
        $("input[name=KJ]").attr("checked", false);
    });
});