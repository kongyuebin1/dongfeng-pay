package expection

import (
	"github.com/beego/beego/v2/server/web"
	"html/template"
	"net/http"
)

func PageNotFind(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(web.BConfig.WebConfig.ViewsPath + "/error/404.html")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}

func SystemExpection(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("503.html").ParseFiles(web.BConfig.WebConfig.ViewsPath + "/error/503.html")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
