package main

import (
	_ "FirstHttpServer/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
