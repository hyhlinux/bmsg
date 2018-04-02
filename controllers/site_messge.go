package controllers

import (
	"bmsg/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"strconv"
	"bmsg/logger"
)

// Operations about object
type MessgeController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Messge	true		"The Messge content"
// @Success 200 {string} models.Messge.Id
// @Failure 403 body is empty
// @router / [post]
func (o *MessgeController) Post() {
	var ob models.Messge
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid, err := models.AddMessge(&ob)
	if err != nil {
		o.Data["json"] = err.Error()
	}else{
		o.Data["json"] = map[string]interface{}{ "ObjectId": objectid}
	}
	o.ServeJSON()
}

// @Title Get
// @Description find msg by MessgeId
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Messge
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *MessgeController) Get() {
	//touserid
	id := o.Ctx.Input.Param(":objectId")
	logger.Debugf("id:%v", id)
	mid, err := strconv.Atoi(id)
	if err != nil {
		o.Data["json"] = err.Error()
	}else{
		msg, err := models.GetMessgeById(int64(mid))
		if err != nil {
			o.Data["json"] = err.Error()
		}else{
			o.Data["json"] = msg
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all msg cnt
// @Success 200 {object} models.Messge
// @Failure 403 :objectId is empty
// @router / [get]
func (o *MessgeController) GetAll() {
	// 查询 toUserId 用户的，所有消息／未读消息
	ids, err := models.AllMessge()
	if err != nil {
		o.Data["json"] = err.Error()
	}else{
		o.Data["json"] = ids
	}
	o.ServeJSON()
}

// @Title Update
// @Description update the msg
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Messge	true		"The body"
// @Success 200 {object} models.Messge
// @Failure 403 :objectId is empty
// @router /:Id [put]
func (o *MessgeController) Put() {
	id := o.Ctx.Input.Param(":id")
	var ob models.Messge
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	if mid, err := strconv.Atoi(id); err != nil {
		o.Data["json"] = err.Error()
	}else{
		_, err = models.UpdateMessge(int64(mid), &ob)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = "update success!"
		}
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the msg
// @Param	Id		path 	string	true		"The message Id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (o *MessgeController) Delete() {
	id := o.Ctx.Input.Param(":id")
	mid, err := strconv.Atoi(id)
	if err != nil {
		o.Data["json"] = err.Error()
	}else{
		models.DeleteMessge(int64(mid))
		o.Data["json"] = "delete success!"
	}
	o.ServeJSON()
}

