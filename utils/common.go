package utils

import (
	"github.com/astaxie/beego/logs"
)

var FileLogs *logs.BeeLogger

func init() {
	FileLogs = logs.NewLogger(1000)
	FileLogs.SetLogger("file", `{"filename":"logs/xalarm.log"}`)
}
