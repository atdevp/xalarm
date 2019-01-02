package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"xalarm/models"
)

type MsgController struct {
	beego.Controller
}

func (o *MsgController) Get() {

	touser := o.GetString("touser")
	totag := o.GetString("totag")
	msg := o.GetString("msg")
	ok := models.MsgSet.Send(touser, totag, msg)

	if ok != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  ok.Error(),
		}
		o.ServeJSON()
	}

	o.Data["json"] = map[string]interface{}{
		"errcode": 0,
		"errmsg":  "",
	}
	o.ServeJSON()
}

func (o *MsgController) Post() {

	var data models.ImMsg
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &data)

	if err != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  "参数不合法",
		}
		o.ServeJSON()
	}

	ok := models.MsgSet.Send(data.ToUser, data.ToTag, data.Msg)
	if ok != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  ok.Error(),
		}
		o.ServeJSON()
	}

	o.Data["json"] = map[string]interface{}{
		"errcode": 0,
		"errmsg":  "",
	}
	o.ServeJSON()
}
