// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
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
	
	beego.Router("/api/users/list", &controllers.ImUserController{},"*:ListAllUser")
	beego.Router("/api/users/create", &controllers.ImUserController{},"*:CreateUser")
	beego.Router("/api/users/delete", &controllers.ImUserController{},"*:DeleteUser")

	beego.Router("/api/tag/users/list", &controllers.ImtagController{}, "*:ListTagMember")
	beego.Router("/api/tag/users/create", &controllers.ImtagController{}, "post:CreateTagMember")
	beego.Router("/api/tag/users/delete", &controllers.ImtagController{}, "post:DeleteTagMember") 

	beego.Router("/send", &controllers.MsgController{})
}

