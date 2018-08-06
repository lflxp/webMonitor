package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lflxp/webMonitor/restapi/models"
)

// 后端系统认证
type AdminController struct {
	BaseController
}

// @Title Login
// @Description 登陆验证
// @Param	body		body 	models.AdminUser	true		"登陆验证"
// @Success 200 {string} models.AdminUser.Username
// @Failure 403 body is empty
// @router /login [post]
func (u *AdminController) Login() {
	// {"code":20000,"data":{"token":"admin"}}
	var user models.AdminUser
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	data, err := user.GetToken()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = data
	}

	u.ServeJSON()
}

// @Title LogOut
// @Description 登陆验证
// @Success 200 {string} models.AdminUser.Username
// @Failure 403 body is empty
// @router /logout [post]
func (u *AdminController) LogOut() {
	// {"code":20000,"data":{"token":"admin"}}
	vt := models.AdminLogin{}
	vt.Data.Token = "admin"
	data, err := vt.GetInfo()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = data
	}

	u.ServeJSON()
}

// @Title Get
// @Description 获取用户信息
// @Param	token       query 	string	true		"在namespace下的limitrange名称"
// @Success 200 {string} k8s.io/api/core/v1/LimitRange
// @Failure 403 nothing input
// @router /info [get]
func (u *AdminController) GetInfo() {
	// {"code":20000,"data":{"roles":["admin"],"name":"admin","avatar":"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"}}
	// info?token=admin
	token := u.GetString("token")
	if token != "" && token == "admin" {
		vt := &models.AdminLogin{}
		vt.Data.Token = token
		limit, err := vt.GetInfo()
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = limit
		}
	} else {
		u.Data["json"] = fmt.Sprintf("token is error")
	}
	log.Println(u.Data["json"])
	u.ServeJSON()
}
