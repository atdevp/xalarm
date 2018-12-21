package main

import (
	_ "xalarm/routers"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"os"
	"path/filepath"
	"xalarm/cron"
	"xalarm/utils"
	"xalarm/g"
	"fmt"
	"time"
)

func init() {
	//log := logs.NewLogger(10000)
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	log_path := path + "/logs"
	err := utils.CreateDirIfNotExist(log_path)
	if err != nil {
		panic(err)
	}
	//logs.SetLogger(file,`{"filename":"xalarm.log"}`)

}

func startServer() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func main() {
	go cron.SyncToken()

	for {
		token := g.TokenSet.Get()
		fmt.Println(token)
		time.Sleep(1*time.Second)
	}

	startServer()
}
