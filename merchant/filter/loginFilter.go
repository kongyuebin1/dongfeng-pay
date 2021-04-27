package filter

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/server/web/context"
	"strings"
)

/**
** 对所有的用户请求进行登录判断，如果未登录跳转到登录页面
** 但是对登录请求不进行过滤
 */
var LoginFilter = func(ctx *context.Context) {

	_, ok := ctx.Input.Session("userName").(string)
	if !ok {
		if !strings.Contains(ctx.Request.RequestURI, "/login") {
			ctx.Redirect(302, "/login.html")
		} else {
			logs.Info("该用户没有登录.......")
		}
	}

}
