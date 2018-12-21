package im

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
	"xalarm/g"
)

type Tag struct {
	TagID   int64  `json:"tagid"`
	TagName string `json:tagname`
}

func (t *Tag) Create(id, name) error {

	token := g.TokenSet.Get()
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=%s", token)

	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(1*time.Second, 3*time.Second)
}
