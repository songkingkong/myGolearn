package routers

import (
	"FirstHttpServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/data", &controllers.UserController{})
}
