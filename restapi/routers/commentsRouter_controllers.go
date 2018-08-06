package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetInfo",
			Router: `/info`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:AdminController"],
		beego.ControllerComments{
			Method: "LogOut",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:tablename`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:tablename/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/kv`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"],
		beego.ControllerComments{
			Method: "AddTables",
			Router: `/tables`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:DbController"],
		beego.ControllerComments{
			Method: "GetAllTables",
			Router: `/tables`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lflxp/webMonitor/restapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
