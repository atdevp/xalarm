package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
	"xalarm/g"
)

type ImMsg struct {
	MsgType string `json:"msgtype"`
	AgentID int64  `json:"agentid"`
	ToUser  string `json:"touser"`
	ToTag   string `json:"totag"`
	Msg     string `json:"msg"`
}

var (
	MsgSet = &ImMsg{}
)

func (t *ImMsg) Send(touser string, totag string, msg string) error {

	token := g.GlobalTokenSet.Get()
	msgtype := "text"
	agentid := beego.AppConfig.String("agent_id")

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	req := httplib.Post(url)
	text := make(map[string]string)
	text["content"] = msg

	req.JSONBody(map[string]interface{}{"msgtype": msgtype, "agentid": agentid, "touser": touser, "totag": totag, "text": text})
	req.SetTimeout(1*time.Second, 3*time.Second)

	var ret g.SendMsgResult
	ok := req.ToJSON(&ret)

	if ok != nil {
		return ok
	}

	errCode := ret.Errcode
	if errCode != 0 {
		return errors.New(ret.Errmsg)
	}

	return nil
}
