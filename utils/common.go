package utils

import (
	"github.com/astaxie/beego/logs"
)

var FileLogs *logs.BeeLogger

func init() {
	FileLogs = logs.NewLogger(1000)
	FileLogs.SetLogger("file", `{"filename":"logs/xalarm.log"}`)
}

func GetMaxNumber(arr []int64) int64 {
	var biggest,n int64 
	for _,v := range arr {
		if v>n {
		  n = v
		  biggest = n
		} 
	  }
	  return biggest
}