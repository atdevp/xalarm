package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type ImController struct {
	beego.Controller
}

type ImData struct {
	Userid      string `json:"userid"`
	Content     string `json:"content"`
	Departament string `json:"department"`
}

func (this *ImController) Get() {
	userid := this.GetString("userid")

	this.Data["json"] = map[string]interface{}{
		"success": 0,
		"message": userid,
	}
	this.ServeJSON()
}

func (t *ImController) Post() {

	var data ImData

	err := json.Unmarshal(t.Ctx.Input.RequestBody, &data)

	if err != nil {
		t.Data["json"] = map[string]interface{}{
			"success": 400,
			"message": "参数不合法",
		}
		t.ServeJSON()
	}

	t.Data["json"] = map[string]interface{}{
		"success": 0,
		"message": data,
	}
	t.ServeJSON()
}
