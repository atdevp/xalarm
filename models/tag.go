package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
	"xalarm/g"
	"xalarm/utils"
)

var (
	TagSet = &ImTag{}
)

type ImTag struct {
	TagID   int64  `json:"tagid"`
	TagName string `json:"tagname"`
}

type ImTagUser struct {
	ImTag
	UserList []string `json:"userlist"`
}

type MemberForTagResult struct {
	g.CommonResult
	Tagname  string              `json:"tagname"`
	UserList []map[string]string `json:"userlist"`
}

type TagListResult struct {
	g.CommonResult
	Taglist []ImTag
}

func (t *ImTag) Getall() ([]ImTag, error) {

	token := g.GlobalTokenSet.Get()

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=%s", token)
	req := httplib.Get(url)
	req.SetTimeout(1*time.Second, 3*time.Second)

	var ret TagListResult
	ok := req.ToJSON(&ret)

	if ok != nil {
		return []ImTag{}, ok
	}

	errCode := ret.Errcode
	if errCode != 0 {
		return []ImTag{}, errors.New(ret.Errmsg)
	}

	return ret.Taglist, nil
}

func (t *ImTag) Create(name string) error {

	token := g.GlobalTokenSet.Get()

	var ret TagListResult
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=%s", token)
	req := httplib.Get(url)
	req.SetTimeout(1*time.Second, 3*time.Second)
	ok := req.ToJSON(&ret)

	var tagid int64
	if ok != nil {
		tagid = 100
	}

	tagList := ret.Taglist
	if len(tagList) == 0 {
		tagid = 1
	}
	tmp := []int64{}
	for _, v := range tagList {
		tmp = append(tmp, v.TagID)
	}

	biggest := utils.GetMaxNumber(tmp)
	tagid = biggest + 1

	url = fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=%s", token)
	req = httplib.Post(url)
	req.JSONBody(map[string]interface{}{"tagid": tagid, "tagname": name})
	req.SetTimeout(1*time.Second, 3*time.Second)

	var retc g.CommonResult
	ok = req.ToJSON(&ret)

	if ok != nil {
		return ok
	}

	errCode := retc.Errcode
	if errCode != 0 {
		return errors.New(retc.Errmsg)
	}

	return nil
}

func (t *ImTag) UpdateName(id int64, name string) error {

	token := g.GlobalTokenSet.Get()

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=%s", token)
	req := httplib.Post(url)
	req.JSONBody(map[string]interface{}{"tagid": id, "tagname": name})
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

func (t *ImTag) Delete(id int64) error {

	token := g.GlobalTokenSet.Get()

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=%s&tagid=%d", token, id)
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

func (t *ImTag) GetMember(id int64) ([]map[string]string, error) {
	token := g.LocalTokenSet.Get()

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=%s&tagid=%d", token, id)
	req := httplib.Get(url)
	req.SetTimeout(1*time.Second, 3*time.Second)

	var ret MemberForTagResult
	ok := req.ToJSON(&ret)
	fmt.Println(ret.UserList)

	if ok != nil {
		return []map[string]string{}, ok
	}

	errCode := ret.Errcode
	if errCode != 0 {
		return []map[string]string{}, errors.New(ret.Errmsg)
	}

	return ret.UserList, nil
}

func (t *ImTag) AddMember(id int64, userlist []string) error {
	token := g.GlobalTokenSet.Get()

	fmt.Println(id, userlist)
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=%s", token)

	req := httplib.Post(url)
	req.JSONBody(map[string]interface{}{"tagid": id, "userlist": userlist})
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

func (t *ImTag) DeleteMember(id int64, userlist []string) error {
	token := g.GlobalTokenSet.Get()

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=%s", token)
	req := httplib.Post(url)
	req.JSONBody(map[string]interface{}{"tagid": id, "userlist": userlist})
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
