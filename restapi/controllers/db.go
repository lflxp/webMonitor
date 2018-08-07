package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/lflxp/webMonitor/restapi/models"
	"github.com/lflxp/webMonitor/restapi/pkg/db"
)

// Operations about Bucket
type DbController struct {
	BaseController
}

// @Title CreateBucket
// @Description create users
// @Param	body		body 	models.Bucket	true		"body for Bucket content"
// @Success 200 {int} models.Bucket.Key
// @Failure 403 body is empty
// @router /kv [post]
func (u *DbController) Post() {
	var user models.Bucket
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	err := db.AddKeyValueByBucketName(user.Tablename, user.Key, user.Value)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("add %s %s success", user.Tablename, user.Key)
	}
	u.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.Bucket	true		"body for user content"
// @Success 200 {int} models.Tablename
// @Failure 403 body is empty
// @router /tables [post]
func (u *DbController) AddTables() {
	var user models.Bucket
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	err := db.CreateBucket(user.Tablename)
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = fmt.Sprintf("add table %s success", user.Tablename)
	}
	u.ServeJSON()
}

// @Title GetAllTables
// @Description get all Tables
// @Success 200 {string} success
// @router /tables [get]
func (u *DbController) GetAllTables() {
	data, err := db.GetAllTables()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = data
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *DbController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	tablename path 	string	true		"The key for staticblock"
// @Param	key		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Bucket
// @Failure 403 :tablename is empty
// @router /:tablename/:key [get]
func (u *DbController) Get() {
	tablename := u.GetString(":tablename")
	key := u.GetString(":key")
	if tablename != "" && key != "" {
		data, err := db.GetValueByBucketName(tablename, key)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = string(data)
		}
	} else {
		u.Data["json"] = "tablename or key is none"
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *DbController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the Table
// @Param	tablename		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 tablename is empty
// @router /:tablename [delete]
func (u *DbController) Delete() {
	tablename := u.GetString(":tablename")
	err := db.DeleteBucket(tablename)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = "delete success!"
	}
	u.ServeJSON()
}
