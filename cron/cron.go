package cron

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
	"xalarm/g"
	"xalarm/utils"
)

type TokenResponse struct {
	g.CommonResult
	AccessToken string `json:"access_token"`
}

func SyncToken() {

	interval := beego.AppConfig.DefaultInt64("access_token_interval", 3600)

	duration := time.Duration(interval) * time.Second
	for {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		utils.FileLogs.Info("cron_sync_token_time: %v", currentTime)
		syncGlobalToken()
		syncLocalToken()
		time.Sleep(duration)
	}
}

func syncGlobalToken() {
	secret := beego.AppConfig.String("corpsecret")
	token, ok := getToken(secret)
	if ok != nil {
		fmt.Println(ok.Error())
	}
	g.GlobalTokenSet.Reinit(token)
}

func syncLocalToken() {
	secret := beego.AppConfig.String("addressbooksecret")
	token, ok := getToken(secret)
	if ok != nil {
		fmt.Println(ok.Error())
	}
	g.LocalTokenSet.Reinit(token)
}

func getToken(secret string) (string, error) {

	corpid := beego.AppConfig.String("corpid")

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpid, secret)
	req := httplib.Get(url)
	req.SetTimeout(1*time.Second, 3*time.Second)

	var ret TokenResponse
	ok := req.ToJSON(&ret)
	if ok != nil {
		return "", ok
	}

	errCode := ret.Errcode
	if errCode != 0 {
		return "", errors.New(ret.Errmsg)
	}

	return ret.AccessToken, nil
}
