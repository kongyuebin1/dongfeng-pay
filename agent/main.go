package main

import (
	_ "agent/models"
	_ "agent/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.BConfig.WebConfig.Session.SessionOn = true
	web.Run()
}

