package controllers

import (
	"bmsg/models"

	"github.com/astaxie/beego"
	"strconv"
	"bmsg/logger"
)

// Operations about object
type ToUserIdController struct {
	beego.Controller
}

// @Title Get
// @Description find msg by MessgeId
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Messge
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *ToUserIdController) Get() {
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

// @Title ShowToUserMessges
// @Description Show msg
// @Param	body		body 	models.MessgeJson	true	"The Messge content"
// @Param	from_user_id	query	int			true	"from_user_id"
// @Param	to_user_id		query	int			true	"to_user_id"
// @Param	status			query	string		true	"status"
// @Success 200 {int} models.Messge.Id
// @Failure 400 arg is err
// @router / [post]
func (o *ToUserIdController) ShowToUserMessges() {
	// 查询 toUserId 用户的，所有消息／未读消息
	var argJson models.MessgeJson
	err := o.ParseForm(&argJson)
	if err != nil {
		logger.Errorf("url: %v", o.Ctx.Input.URI())
		o.Abort("400")
		return
	}
	logger.Debugf("argJson:%v", argJson)

	nums := int64(0)
	msgList := []*models.Messge{}
	if argJson.FromUserId > 0 {
		nums, msgList, err = models.GetMessgeByToUserAll(argJson.ToUserId, argJson.FromUserId, argJson.Status, false)
	}else{
		nums, msgList, err = models.GetMessgeByToUser(argJson.ToUserId, argJson.Status, false)
	}
	if err != nil {
		logger.Errorf("ShowToUserMessges err:%v argJson:%v", err, argJson)
		o.Data["json"] = err.Error()
	}else{
		o.Data["json"] = map[string]interface{}{
			"err": "",
			"nums": nums,
			"data": msgList,
		}
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
// @Param	status			query	int			true	"status"
// @Param	is_delete		query	bool		true	"is_delete"
// @Success 200 {int} models.Messge.Id
// @Failure 400 arg is err
// @router / [post]
func (o *ToUserIdController) UpdateMessge() {
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