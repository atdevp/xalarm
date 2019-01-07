// @APIVersion 1.0.0
// @Title Xalarm API
// @Description xalarm by wechat
// @Contact root.db.root@gmail.com
package routers

import (
	"xalarm/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/tags/list", &controllers.ImtagController{}, "*:ListTag")
	beego.Router("/api/tags/create", &controllers.ImtagController{}, "post:CreateTag")
	beego.Router("/api/tags/update", &controllers.ImtagController{}, "post:UpdateTag")
	beego.Router("/api/tags/delete", &controllers.ImtagController{}, "get:DeleteTag")

	beego.Router("/api/users/list", &controllers.ImUserController{}, "*:ListAllUser")
	beego.Router("/api/users/create", &controllers.ImUserController{}, "post:CreateUser")
	beego.Router("/api/users/delete", &controllers.ImUserController{}, "get:DeleteUser")

	beego.Router("/api/tag/users/list", &controllers.ImtagController{}, "*:ListTagMember")
	beego.Router("/api/tag/users/create", &controllers.ImtagController{}, "post:CreateTagMember")
	beego.Router("/api/tag/users/delete", &controllers.ImtagController{}, "post:DeleteTagMember")

	beego.Router("/api/alarm/wechat/send", &controllers.MsgController{})
}


