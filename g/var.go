package g

type CommonResult struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type ImUser struct {
	Userid string `json:"userid"`
	Name   string `json:"name"`
}

type SendMsgResult struct {
	CommonResult
	InvalidUser string `json:"invaliduser"`
	InvalidTag  string `json:"invalidtag"`
}
