package cron

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"time"
	"xalarm/g"
)

type TokenResponse struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
}

func SyncToken() {

	interval := beego.AppConfig.DefaultInt64("access_token_interval", 7000)

	duration := time.Duration(interval) * time.Second
	for {
		syncToken()
		time.Sleep(duration)
	}
}

func syncToken() {
	token, ok := getToken()
	if ok != nil {
		fmt.Println(ok.Error())
	}
	g.TokenSet.Reinit(token)
}

func getToken() (string, error) {

	corpid := beego.AppConfig.String("corpid")
	corpsecret := beego.AppConfig.String("corpsecret")

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpid, corpsecret)

	req := httplib.Get(url)
	result, ok := req.Response()
	if ok != nil {
		return "", ok
	}

	code := result.StatusCode
	if code != 200 {
		return "", errors.New("get wechat token faild")
	}

	data, ok := ioutil.ReadAll(result.Body)
	if ok != nil {
		return "", ok
	}

	var msg TokenResponse
	json.Unmarshal(data, &msg)

	errCode := msg.Errcode
	if errCode != 0 {
		return "", errors.New(msg.Errmsg)
	}

	return msg.AccessToken, nil
}
