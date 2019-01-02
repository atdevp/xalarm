package main

import (
	_ "xalarm/routers"
	_ "xalarm/utils"
	"github.com/astaxie/beego"
	"xalarm/cron"
)

func startServer() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/swagger", "swagger")
	beego.Run()
}

func main() {
	go cron.SyncToken()

	startServer()
}
