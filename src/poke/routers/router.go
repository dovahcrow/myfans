package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"poke/controllers"
)

func init() {
	beego.Get("/", func(c *context.Context) { c.Redirect(302, beego.UrlFor("BasicController.Index")) })

	beego.Router("/discuss", &controllers.BasicController{}, "get:Index")
	beego.Router("/discuss/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/discuss/reply/rest", &controllers.AjaxController{}, `get:UwantsReply`)
	beego.Router("/discuss/topic/rest", &controllers.AjaxController{}, `get:UwantsTopic`)
	beego.Router("/discuss/reply", &controllers.Uwants{}, "get:GetReply;post:PostReply")
	beego.Router("/discuss/topic", &controllers.Uwants{}, "get:GetTopic;post:PostTopic")

	beego.Router("/discuss/records", &controllers.TopicRecords{}, "get:Records")
	beego.Router(`/discuss/records/search`, &controllers.TopicRecords{}, "get:Search")

	beego.Router("/discuss/replyrecords", &controllers.ThreadRecords{}, "get:Records")
	beego.Router(`/discuss/replyrecords/search`, &controllers.ThreadRecords{}, "get:Search")

	beego.Router("/discuss/users/single", &controllers.UserController{}, "get:Single")
	beego.Router("/discuss/users/bunch", &controllers.UserController{}, "get:Bunch;post:BunchAdd")

	beego.Router(`/discuss/users/update/:id(\d+)`, &controllers.UserController{}, "post:UpdateUser")
	beego.Router("/discuss/users/create", &controllers.UserController{}, "post:CreateUser")
	beego.Router(`/discuss/users/delete/:id(\d+)`, &controllers.UserController{}, "get:DeleteUser")

	beego.Router("/discuss/threads/single", &controllers.ThreadController{}, "get:Single")
	beego.Router("/discuss/threads/bunch", &controllers.ThreadController{}, "get:Bunch;post:BunchAdd")

	beego.Router(`/discuss/threads/update/:id(\d+)`, &controllers.ThreadController{}, "post:UpdateThread")
	beego.Router("/discuss/threads/create", &controllers.ThreadController{}, "post:CreateThread")
	beego.Router(`/discuss/threads/delete/:id(\d+)`, &controllers.ThreadController{}, "get:DeleteThread")

}
