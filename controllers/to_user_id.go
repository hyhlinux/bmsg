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

// @Title GetAll
// @Description get all msg cnt
// @Success 200 {object} models.Messge
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ToUserIdController) ShowToUserMessges() {
	// 查询 toUserId 用户的，所有消息／未读消息
	ids, err := models.AllMessge()
	if err != nil {
		o.Data["json"] = err.Error()
	}else{
		o.Data["json"] = ids
	}
	o.ServeJSON()
}

