package routers

import (
	"FirstHttpServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "GET:Index")
	beego.Router("/login", &controllers.LoginController{}, "Get:GetUser")
	beego.Router("/login", &controllers.LoginController{}, "Post:SetUser")
}
