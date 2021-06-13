package controllers

import (
	"FirstHttpServer/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) GetUser() {
	c.TplName = "login.tpl"
}
func (c *LoginController) SetUser() {
	user := new(models.User)
	user.UserName = c.GetString("username")
	user.PassWd = c.GetString("password")
	c.Data["json"], _ = json.Marshal(user)
	c.ServeJSON()
}
