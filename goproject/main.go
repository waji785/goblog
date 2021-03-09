package main

import (
	_ "goproject/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"goproject/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

