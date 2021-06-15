package routers

import (
	"FirstHttpServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "GET:Index")
	beego.Router("/login", &controllers.UserController{}, "*:Login")
	beego.Router("/register", &controllers.UserController{}, "*:Register")
}
