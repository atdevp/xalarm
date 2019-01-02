package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
	"xalarm/g"
)

var (
	UserSet = &ImUser{}
)

type ImUser struct {
	Userid string `json:"userid"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Status int64  `json:"status"`
}

type UserListResult struct {
	g.CommonResult
	Userlist []ImUser
}

func (t *ImUser) Getall() ([]ImUser, error) {

	token := g.LocalTokenSet.Get()

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=%s&department_id=1&fetch_child=0", token)
	req := httplib.Get(url)
	req.SetTimeout(1*time.Second, 3*time.Second)

	var ret UserListResult
	ok := req.ToJSON(&ret)

	if ok != nil {
		return []ImUser{}, ok
	}

	errCode := ret.Errcode
	if errCode != 0 {
		return []ImUser{}, errors.New(ret.Errmsg)
	}
	return ret.Userlist, nil
}

func (t *ImUser) Create(userid string, username string, mobile string) error {

	token := g.LocalTokenSet.Get()

	department := beego.AppConfig.DefaultInt64("default_department_id", 1)

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=%s", token)
	req := httplib.Post(url)
	req.JSONBody(map[string]interface{}{"userid": userid, "name": username, "department": department, "mobile": mobile})
	req.SetTimeout(1*time.Second, 3*time.Second)

	var ret g.CommonResult
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

func (t *ImUser) Delete(userid string) error {

	token := g.LocalTokenSet.Get()

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=%s&userid=%s", token, userid)
	req := httplib.Get(url)
	req.SetTimeout(1*time.Second, 3*time.Second)

	var ret g.CommonResult
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
