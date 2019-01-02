package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"xalarm/models"
)

type ImtagController struct {
	beego.Controller
}

func (o *ImtagController) ListTag() {
	taglist, ok := models.TagSet.Getall()
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
		"taglist": taglist,
	}
	o.ServeJSON()
}

//创建tag
func (o *ImtagController) CreateTag() {

	var data models.ImTag
	ok := json.Unmarshal(o.Ctx.Input.RequestBody, &data)

	if ok != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  "参数不合法",
		}
		o.ServeJSON()
	}

	ok = models.TagSet.Create(data.TagID, data.TagName)
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

//修改tag
func (o *ImtagController) UpdateTag() {

	var data models.ImTag
	ok := json.Unmarshal(o.Ctx.Input.RequestBody, &data)

	if ok != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  "参数不合法",
		}
		o.ServeJSON()
	}

	ok = models.TagSet.UpdateName(data.TagID, data.TagName)
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

//删除tag
func (o *ImtagController) DeleteTag() {

	tagid, ok := o.GetInt64("tagid")
	if ok != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  "参数不合法",
		}
		o.ServeJSON()
	}

	ok = models.TagSet.Delete(tagid)
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

//获取指定tag用户
func (o *ImtagController) ListTagMember() {
	tagid, ok := o.GetInt64("tagid")
	if ok != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  "参数不合法",
		}
		o.ServeJSON()
	}
	userlist, ok := models.TagSet.GetMember(tagid)
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

//添加指定用户
func(o *ImtagController) CreateTagMember() {
	var data models.ImTagUser
	ok := json.Unmarshal(o.Ctx.Input.RequestBody, &data)
	if ok != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  "参数不合法",
		}
		o.ServeJSON()
	}

	ok = models.TagSet.AddMember(data.TagID, data.UserList)
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
	}
	o.ServeJSON()
}

//删除指定用户
func(o *ImtagController) DeleteTagMember() {
	var data models.ImTagUser
	ok := json.Unmarshal(o.Ctx.Input.RequestBody, &data)
	if ok != nil {
		o.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg":  "参数不合法",
		}
		o.ServeJSON()
	}

	ok = models.TagSet.DeleteMember(data.TagID, data.UserList)
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
	}
	o.ServeJSON()
}