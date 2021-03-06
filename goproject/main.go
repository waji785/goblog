package main

import (
	_ "goproject/routers"
	"github.com/astaxie/beego"
	"goproject/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

