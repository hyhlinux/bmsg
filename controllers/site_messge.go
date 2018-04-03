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
	// TODO 分页
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
func (o *MessgeController) DeleteMessge() {
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

// @Title UpdateMessge
// @Description update msg
// @Param	body		body 	models.Messge	true	"The Messge content"
// @Param	id				query	int			true	"id"
// @Param	from_user_id	query	int			true	"from_user_id"
// @Param	to_user_id		query	int			true	"to_user_id"
// @Param	title			query	string		true	"title"
// @Param	message			query	string		true	"message"
// @Param	status			query	string		true	"status"
// @Param	is_delete		query	bool		true	"is_delete"
// @Success 200 {int} models.Messge.Id
// @Failure 400 arg is err
// @router / [post]
func (o *MessgeController) UpdateMessge() {
	var argJson models.MessgeJson
	err := o.ParseForm(&argJson)
	if err != nil {
		logger.Errorf("url: %v", o.Ctx.Input.URI())
		o.Abort("400")
		return
	}

	msg, err := models.UpdateMessge(int64(argJson.Id), &models.Messge{
		Id: argJson.Id,
		FromUserId: argJson.FromUserId,
		ToUserId: argJson.ToUserId,
		Title: argJson.Title,
		Message: argJson.Message,
		IsDelete: argJson.IsDelete,
	})
	if err != nil {
		logger.Errorf("UpdateMessge err:%v argJson:%v", err, argJson)
		o.Data["json"] = err.Error()
	}else{
		o.Data["json"] = msg
	}
	o.ServeJSON()
}