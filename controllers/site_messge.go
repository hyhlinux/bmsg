package controllers

import (
	"bmsg/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"bmsg/logger"
	"fmt"
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
func (o *MessgeController) CreateMessge() {
	var ob models.MessgeJson
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logger.Errorf("err:%v", err)
		o.Abort("500")
	}
	logger.Debugf("argJson:%v", ob)
	objectid, err := models.AddMessge(&models.Messge{
		Id: ob.Id,
		FromUserId: ob.FromUserId,
		ToUserId: ob.ToUserId,
		Title: ob.Title,
		Message: ob.Message,
	})
	if err != nil {
		o.Data["json"] = err.Error()
	}else{
		o.Data["json"] = map[string]interface{}{ "ObjectId": objectid}
	}
	o.ServeJSON()
}

func (o *MessgeController) ReadMessge() {
	var ob models.MessgeJson
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logger.Errorf("err:%v", err)
		o.Abort("500")
	}
	logger.Debugf("argJson:%v", ob)

	nums := int64(0)
	msgList := []*models.Messge{}
	// todo
	//models.ReadMessage()
	if ob.FromUserId > 0 {
		nums, msgList, err = models.ReadMessgeWithFromUserId(ob.MsgType, ob.Id, ob.ToUserId, ob.FromUserId, ob.PageNumber, ob.PageSize)
	}else{
		nums, msgList, err = models.ReadMessge(ob.MsgType,  ob.Id, ob.ToUserId, ob.PageNumber, ob.PageSize)
	}
	if err != nil {
		logger.Errorf("ShowToUserMessges err:%v argJson:%v", err, ob)
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


// @Title Delete
// @Description delete the msg
// @Param	Id		path 	string	true		"The message Id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (o *MessgeController) DeleteMessge() {
	var ob models.MessgeJson
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logger.Errorf("err:%v", err)
		o.Abort("500")
	}
	logger.Debugf("argJson:%v", ob)
	if err := models.DeleteMessge(ob.Id); err != nil {
		logger.Errorf("err:%v", err)
		o.Data["json"] = fmt.Errorf("DeleteMessge src err:%v ob.id:%v", err, ob.Id)
	}else{
		o.Data["json"] = "delete success!"
	}
	o.ServeJSON()
}


func (o *MessgeController) UpdateMessge() {
	var ob models.MessgeJson
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logger.Errorf("err:%v", err)
		o.Abort("500")
	}
	logger.Debugf("argJson:%v", ob)
	msg, err := models.UpdateMessge(ob.Id, &models.Messge{
		Id: ob.Id,
		//FromUserId: ob.FromUserId,
		//ToUserId: ob.ToUserId,
		//Title: ob.Title,
		Status: ob.Status,
		//Message: ob.Message,
		// 更新接口，用户端不能更新isDelete
		//IsDelete: ob.IsDelete,
	})
	if err != nil {
		logger.Errorf("UpdateMessge err:%v argJson:%v", err, ob)
		o.Data["json"] = err.Error()
	}else{
		o.Data["json"] = msg
	}
	o.ServeJSON()
}

