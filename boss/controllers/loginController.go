package controllers

import (
	"boss/common"
	"boss/models"
	"boss/utils"
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {

}

func (c *LoginController) Login() {

	userID := c.GetString("userID")
	passWD := c.GetString("passwd")
	code := c.GetString("Code")

	dataJSON := new(KeyDataJSON)

	valid := validation.Validation{}

	if v := valid.Required(userID, "userID"); !v.Ok {
		dataJSON.Key = v.Error.Key
		dataJSON.Msg = "手机号不能为空！"
	} else if v := valid.Required(passWD, "passWD"); !v.Ok {
		dataJSON.Key = v.Error.Key
		dataJSON.Msg = "登录密码不能为空！"
	} else if v := valid.Length(code, common.VERIFY_CODE_LEN, "code"); !v.Ok {
		dataJSON.Key = v.Error.Key
		dataJSON.Msg = "验证码不正确！"
	}

	userInfo := models.GetUserInfoByUserID(userID)

	if userInfo.UserId == "" {
		dataJSON.Key = "userID"
		dataJSON.Msg = "用户不存在，请求联系管理员！"
	} else {
		codeInterface := c.GetSession("verifyCode")
		if userInfo.Passwd != utils.GetMD5Upper(passWD) {
			dataJSON.Key = "passWD"
			dataJSON.Msg = "密码不正确！"
		} else if codeInterface == nil {
			dataJSON.Key = "code"
			dataJSON.Msg = "验证码失效！"
		} else if code != codeInterface.(string) {
			dataJSON.Key = "code"
			dataJSON.Msg = "验证码不正确！"
		} else if userInfo.Status == "unactive" {
			dataJSON.Key = "unactive"
			dataJSON.Msg = "用户已被冻结！"
		} else if userInfo.Status == "del" {
			dataJSON.Key = "del"
			dataJSON.Msg = "用户已被删除！"
		}
	}

	go func() {
		userInfo.Ip = c.Ctx.Input.IP()
		models.UpdateUserInfoIP(userInfo)
	}()

	if dataJSON.Key == "" {
		c.SetSession("userID", userID)
		c.DelSession("verifyCode")
	}

	c.Data["json"] = dataJSON
	c.ServeJSON()
}

/*
* 退出登录,删除session中的数据，避免数据量过大，内存吃紧
 */

func (c *LoginController) Logout() {
	dataJSON := new(BaseDataJSON)

	c.DelSession("userID")
	dataJSON.Code = 200

	c.Data["json"] = dataJSON
	c.ServeJSON()
}

/*
* 验证码获取，如果获取成功，并将验证码存到session中
 */
func (c *LoginController) GetVerifyImg() {
	Image, verifyCode := utils.GenerateVerifyCodeImg()
	if Image == nil || len(verifyCode) != common.VERIFY_CODE_LEN {
		logs.Error("获取验证码图片失败！")
	} else {
		c.SetSession("verifyCode", verifyCode)
	}
	logs.Info("验证码：", verifyCode)
	Image.WriteTo(c.Ctx.ResponseWriter)
}
