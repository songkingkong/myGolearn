package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}
type User struct {
	Id       int    `json:"id"`
	Username string `json:"name"`
	Phone    string
}

func (c *UserController) Get() {
	user := User{1, "tizi365", "13089818901"}
	c.Ctx.Output.Header("Content-Type", "message/http")
	c.Ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Data["json"] = &user
	c.ServeJSON()
}