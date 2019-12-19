
/**************************************时间格式化处理************************************/
function dateFtt(fmt,date)   
{ //author: meizz   
  var o = {   
    "M+" : date.getMonth()+1,                 //月份   
    "d+" : date.getDate(),                    //日   
    "h+" : date.getHours(),                   //小时   
    "m+" : date.getMinutes(),                 //分   
    "s+" : date.getSeconds(),                 //秒   
    "q+" : Math.floor((date.getMonth()+3)/3), //季度   
    "S"  : date.getMilliseconds()             //毫秒   
  };   
  if(/(y+)/.test(fmt))   
    fmt=fmt.replace(RegExp.$1, (date.getFullYear()+"").substr(4 - RegExp.$1.length));   
  for(var k in o)   
    if(new RegExp("("+ k +")").test(fmt))   
  fmt = fmt.replace(RegExp.$1, (RegExp.$1.length==1) ? (o[k]) : (("00"+ o[k]).substr((""+ o[k]).length)));   
  return fmt;   
}

$(".start-time").on('changeDate', function() {
	let startTime = $(".start-time").val();
	if (startTime) {
		$(".end-time").datetimepicker("setStartDate", startTime);
	} else {
		$(".end-time").datetimepicker("setStartDate", new Date(-8639968443048000));
	}
});
$("#end-time").on('changeDate', function() {
	let endTime = $(".end-time").val();
	if (endTime) {
		$(".start-time").datetimepicker('setEndDate', endTime);
	} else {
		$(".start-time").datetimepicker();
	}
});
$(".start-time, .end-time").datetimepicker({
	language: 'zh-CN',
    format: "yyyy-mm-dd hh:ii:00",
    clearBtn: true,
    todayBtn: true,
    autoclose: true,
    startView:2,
    minView: 0,//最低视图 小时视图
    maxView: 4, //最高视图 十年视图
    showSecond : true,
    showHours : true,
    minuteStep:1
});

//将上游通道供应商写入
function setSupplier() {
    $.ajax({
        url: "/get/product",
        success: function (res) {
            if (res.Code == 404) {window.parent.location = "/login.html";}
            else if (res.Code == -1) {alert("没有获取到上游供应商数据");}
            else {
                let  str = '<option value="' + "" + '">' + "请选择" + '</option>';
                for (let key in res.ProductMap) {
                    let  v = res.ProductMap[key];
                    str = str + '<option value="' + key + '">' + v + '</option>'
                }
                $("#search-order-supplier-name").html(str);
            }
        },
        error: function () {
            alert("系统异常，请稍后再试");
        }
    });
}

//动态获取商户名
function setMerchant() {
    $.ajax({
        url:"/get/all/merchant",
        data:{},
        success: function (res) {
            let str = '<option value="' + "" + '">' + "请选择" + '</option>';
            for (let i = 0; i < res.MerchantList.length; i ++) {
                let merchant = res.MerchantList[i];
                str = str + '<option value="' + merchant.MerchantUid + '">' + merchant.MerchantName + '</option>';
            }

            $("#select-merchant-name").html(str);
        },
        error: function () {
            alert("系统异常，请稍后再试");
        }
    });
}

//动态获取代理名称
function setAgent() {
    $.ajax({
        url: "/get/all/agent",
        data:{},
        success:function (res) {
            if (res.Code == 404) {
                window.parent.location = "/login.html";
            } else {
                let str = '<option value="' + "" + '">' + "请选择" + '</option>';
                for (let i = 0; i < res.AgentList.length; i ++) {
                    let agent = res.AgentList[i];
                    str = str + '<option value="' + agent.AgentUid + '">' + agent.AgentName + '</option>';
                }
                $("#select-agent-name").html(str);
            }
        },
        error: function () {
            alert("系统异常，请稍后再试")
        }
    });
}

