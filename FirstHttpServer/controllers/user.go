package controllers

import (
	"FirstHttpServer/models"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}
type JsonResult struct {
	Code int         `json:"code"`
	Msg  models.User `json:"msg"`
}

func (c *UserController) Login() {
	method := c.Ctx.Request.Method
	switch {
	case method == "GET":
		c.TplName = "login.tpl"
	case method == "POST":
		var jsonResponse JsonResult
		user := models.User{UserName: c.GetString("username"), PassWd: c.GetString("password")}
		userPasswd, err := user.SelectUser(user.UserName)
		if err != nil {
			fmt.Println(err)
		}
		if user.PassWd == userPasswd {
			jsonResponse.Code = 200
		} else {
			jsonResponse.Code = 400
		}
		jsonResponse.Msg = user
		c.Data["json"] = &jsonResponse
		c.ServeJSON()
	default:
		result := JsonResult{Code: 500}
		c.Data["json"] = &result
		c.ServeJSON()
	}

}
func (c *UserController) Register() {
	method := c.Ctx.Request.Method
	jsonResponse := JsonResult{Code: 200}
	switch {
	case method == "POST":
		user := models.User{}
		user.UserName = c.GetString("username")
		user.PassWd = c.GetString("password")
		err := user.AddUser(user.UserName, user.PassWd)
		if err != nil {
			fmt.Println(err)
			jsonResponse.Code = 500
		}
		jsonResponse.Msg = user
		c.Data["json"] = &jsonResponse
		c.ServeJSON()
	case method == "GET":
		c.TplName = "register.tpl"
	default:
		jsonResponse.Code = 405
		c.ServeJSON()
	}
}
