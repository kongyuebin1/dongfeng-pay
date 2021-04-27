/***************************************************
 ** @Desc : This file for 登录js
 ** @Time : 2019.04.01 13:34
 ** @Author : Joker
 ** @File : login.js
 ** @Last Modified by : Joker
 ** @Last Modified time: 2019.04.01 13:34
 ** @Software: GoLand
 ****************************************************/

let login={changeImg:function(){login.setImgSrc($("#rcCaptcha-img"),"reload="+(new Date()).getDate())},setImgSrc:function(obj,reload){var $src=obj[0].src;var $flag=$src.indexOf("?");if($flag>=0){$src=$src.substr(0,$flag)}obj.attr("src",$src+"?"+reload)},flushCaptcha:function(){$.ajax({type:"GET",url:"flushCaptcha.py",success:function(res){$("#rcCaptcha-img").attr("src","/img.do/"+res.data+".png");$("#captchaId").val(res.data)}})},login_action:function(){let userName=$("#userName").val();let Password=$("#Password").val();let captchaCode=$("#captchaCode").val();if(userName===""||Password===""){toastr.error("用户名或者密码不能为空！");return}if(captchaCode===""){toastr.error("验证码不能为空！");return}$.ajax({type:"POST",url:"/login.py/",data:$("#form-validate").serialize(),success:function(data){if(data.code===-9){toastr.error(data.msg,function(){if(data.url==="-9"){login.flushCaptcha()}})}else{window.location.href=data.url}},error:function(XMLHttpRequest){toastr.info('something is wrong, code: '+XMLHttpRequest.status)}})},loginOut:function(){swal({title:"Are you sure?",text:"您确定要退出登录吗？",icon:"warning",closeOnClickOutside:false,buttons:true,dangerMode:true,}).then((willDelete)=>{if(willDelete){$.ajax({type:"get",url:"/loginOut.py",success:function(res){window.location.href=res.url},error:function(XMLHttpRequest){toastr.info('something is wrong, code: '+XMLHttpRequest.status)}})}})}};