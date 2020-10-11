package main

import (
	"DataCerProject/db_mysql"
	_ "DataCerProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.ConnectDB()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()

}

