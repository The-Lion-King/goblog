package routers

import (
	"code/2/blog/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.HomeController{})
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/category",&controllers.CategoryController{})
	beego.Router("/topic",&controllers.TopicController{})
	beego.Router("/reply",&controllers.ReplyController{})
	beego.Router("/attachment/:all",&controllers.AttachController{})
	beego.Router("/reply/add",&controllers.ReplyController{},"post:Add")
	beego.Router("/reply/delete",&controllers.ReplyController{},"get:Delete")
	beego.AutoRouter(&controllers.TopicController{})


}
