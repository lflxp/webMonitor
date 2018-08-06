package main

import (
	_ "github.com/lflxp/webMonitor/restapi/pkg/db"
	_ "github.com/lflxp/webMonitor/restapi/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
