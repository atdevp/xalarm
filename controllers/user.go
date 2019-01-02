package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"xalarm/models"
)

type ImUserController struct {
	beego.Controller
}

func (o *ImUserController) ListAllUser() {
	userlist, ok := models.UserSet.Getall()
	if ok != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  ok.Error(),
		}
		o.ServeJSON()
	}

	o.Data["json"] = map[string]interface{}{
		"errcode":  0,
		"errmsg":   "",
		"userlist": userlist,
	}
	o.ServeJSON()
}

func (o *ImUserController) CreateUser() {

	var data models.ImUser
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &data)

	if err != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  "参数不合法",
		}
		o.ServeJSON()
	}

	ok := models.UserSet.Create(data.Userid, data.Name, data.Mobile)
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

func (o *ImUserController) DeleteUser() {

	userid := o.GetString("userid")

	ok := models.UserSet.Delete(userid)

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
