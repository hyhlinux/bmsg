package controllers

import (
	"bmsg/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"strconv"
	"bmsg/logger"
)

// Operations about object
type FromUserIdController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Messge	true		"The Messge content"
// @Success 200 {string} models.Messge.Id
// @Failure 403 body is empty
// @router / [post]
func (o *FromUserIdController) CreateMessge() {
	var ob models.Messge
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob); err != nil {
		logger.Errorf("err:%v ob:%v", err, ob)
		o.Abort("400")
	}
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
func (o *FromUserIdController) Get() {
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
func (o *FromUserIdController) ShowFromUserMessges() {
	// 查询 toUserId 用户的，所有消息／未读消息
	ids, err := models.AllMessge()
	if err != nil {
		o.Data["json"] = err.Error()
	}else{
		o.Data["json"] = ids
	}
	o.ServeJSON()
}

