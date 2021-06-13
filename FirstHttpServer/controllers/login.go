package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) post() {

	l.Ctx.Output.Header("Content-Type", "message/http")
	l.Ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
}
