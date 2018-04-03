package models

import (
	"errors"
	"strings"
	"fmt"
	"bmsg/logger"

	"github.com/astaxie/beego/orm"
	"time"
)

const TABLENAME = "user_message"
// Params stores the Params
type Params map[string]interface{}

// ParamsList stores paramslist
type ParamsList []interface{}

//http json
type MessgeJson struct {
	PageNumber int 		`form:"pageNumber,omitempty"`
	PageSize int		`form:"pageSize,omitempty"`
	MsgType string      `form:"msgType,omitempty"`
	Id int64   			`form:"id,omitempty"`
	FromUserId int64    `form:"fromUserId,omitempty"`
	ToUserId int64    	`form:"toUserId,omitempty"`
	Title       string         `form:"title,omitempty"`
	Message     string         `form:"message,omitempty"`
	Status      string         `form:"status,omitempty"`
	IsDelete    bool         `form:"isDelete,omitempty"`
}

type Messge struct {
	Id          int64          `orm:"column(id)"`
	FromUserId  int64 	   		`orm:"column(from_user_id);"`
	ToUserId    int64 	   		`orm:"column(to_user_id);"`
	CreatedAt   time.Time 	   `orm:"null;auto_now_add;type(datetime);column(created_at)"`
	UpdateAt    time.Time 	   `orm:"null;auto_now;type(datetime);column(update_at)"`
	Title       string         `orm:"null;column(title)"`
	Message     string         `orm:"null;column(message)"`
	IsDelete    bool           `orm:"null;type(bool);column(is_delete);default(false)`
	Status      string         `orm:"null;column(status);"`
}

func (u *Messge) TableName() string {
	return TABLENAME
}

// TODO 后续根据统计需求修改
// 多字段索引
func (u *Messge) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "FromUserId"},
		[]string{"Id", "ToUserId"},
		[]string{"Id", "CreatedAt"},
	}
}

//type CommonTask LogCommonTask
var (
	ErrFetchTaskTimeout = errors.New("Fetch task timeout.")
	StatusMap = map[string]string {
		"0": "UNSEEN",
		"1": "SEEN",
		"UNSEEN": "UNSEEN",
		"SEEN": "SEEN",
		"ALL": "ALL",
		"DETAIL": "DETAIL",
		"DEF": "UNSEEN",
	}
	DefalutIsDelete = false
)

func init() {
	orm.RegisterModel(new(Messge))
}

func NewMessge() *Messge {
	return &Messge{}
}

func AddMessge(t *Messge) (tid int64, e error) {
	if t == nil {
		return 0, errors.New("AddMessge task is nil")
	}

	msg := NewMessge()
	msg.ToUserId = t.ToUserId
	msg.FromUserId = t.FromUserId
	msg.Title = t.Title
	msg.Message = t.Message
	msg.IsDelete = false
	// 0: 未读  1: 已读
	msg.Status = "UNSEEN"
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Insert(msg)
	if err != nil {
		return tid, err
	}
	return msg.Id, nil
}

func AddMulMessge(t []*Messge) (tid int64, e error) {
	if t == nil {
		return 0, errors.New("AddCommonTask task is nil")
	}

	msgList := make([]*Messge, len(t))
	msgList = t

	o := orm.NewOrm()
	o.Using("default")
	tid, err := o.InsertMulti(100, msgList)
	return tid, err
}

// 发信人查询
func GetMessgeByFromUserId(fromUserId int64, isDelete bool) (num int64, msgList []*Messge, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(TABLENAME)
	num, err = qs.Filter("from_user_id", fromUserId).Filter("is_delete", isDelete).All(&msgList)
	if err != nil {
		return -1, msgList, err
	}
	return num, msgList, nil
}

func checkStatus(status string)  string{
	status = strings.ToUpper(status)
	if _, ok := StatusMap[status]; !ok {
		logger.Warnf("src status:%v not exit", status)
		status = StatusMap["DEF"]
	}
	return StatusMap[status]
}

/*
pageSize: [1:1000]
pageNumer: [1:]
*/
func CheckPage(pageNumber int, pageSize int)  (offset int, limit int){
	if pageSize > 1000 {
		pageSize = 1000
	}
	if pageSize <= 0 {
		pageSize = 5
	}

	if pageNumber <= 0 {
		pageNumber = 0
	}

	limit = pageSize
	offset = (pageNumber-1)*pageSize
	// todo delete
	logger.Debugf("offset:%v limit:%v", offset, limit)
	return offset, limit
}

func ReadMessge(statusType string, id, toUserId int64, pageNumber, pageSize int) (num int64, msgList []*Messge, err error) {
	statusType = checkStatus(statusType)
	switch statusType {
	case StatusMap["DETAIL"]:
		msg := &Messge{}
		msg, err = GetMessgeById(id)
		msgList = []*Messge{
			msg,
		}
		break
	default:
		//StatusMap["ALL"], StatusMap["SEEN"], StatusMap["UNSEEN"]:
		offset, limit := CheckPage(pageNumber, pageSize)
		num, msgList, err = GetMessgeByToUser(toUserId, statusType, false, offset, limit)
		break
	}

	return num, msgList, err
}

func ReadMessgeWithFromUserId(statusType string, id, toUserId, fromUserId int64, pageNumber, pageSize int) (num int64, msgList []*Messge, err error) {
	statusType = checkStatus(statusType)
	switch statusType {
	case StatusMap["DETAIL"]:
		msg := &Messge{}
		msg, err = GetMessgeByIdWithFromUserId(id, fromUserId)
		msgList = []*Messge{
			msg,
		}
		break
	default:
		//StatusMap["ALL"], StatusMap["SEEN"], StatusMap["UNSEEN"]:
		offset, limit := CheckPage(pageNumber, pageSize)
		num, msgList, err = GetMessgeByToUserAll(toUserId, fromUserId, statusType, false, offset, limit)
		break
	}

	return num, msgList, err
}

// 收信人查询
func GetMessgeByToUser(toUserId int64, status string, isDelete bool, offset, limit int) (num int64, msgList []*Messge, err error) {
	status = checkStatus(status)
	o := orm.NewOrm()
	qs := o.QueryTable(TABLENAME)
	qs = qs.Filter("to_user_id", toUserId).Filter("is_delete", isDelete)
	if status != "ALL" {
		qs = qs.Filter("status", status)
	}
	num, err = qs.OrderBy("id").Limit(limit, offset).All(&msgList)
	if err != nil {
		return -1, msgList, err
	}
	return num, msgList, nil
}

func GetMessgeByToUserAll(toUserId, fromUserId int64, status string, isDelete bool, offset, limit int) (num int64, msgList []*Messge, err error) {
	status = checkStatus(status)
	o := orm.NewOrm()
	qs := o.QueryTable(TABLENAME)
	//status 必须为SEEN/UNSEEN
	if status != "ALL" {
		qs = qs.Filter("to_user_id", toUserId).Filter("from_user_id", fromUserId).Filter("status", status).Filter("is_delete", isDelete)
	}

	num, err = qs.OrderBy("id").Limit(limit, offset).All(&msgList)
	if err != nil {
		return -1, msgList, err
	}

	return num, msgList, nil
}

/*
不返回已经删除的msg
*/
func GetMessgeById(id int64) (msg *Messge, err error) {
	m := Messge{Id: id}
	o := orm.NewOrm()
	o.Using("default")
	err = o.Read(&m)
	if err == orm.ErrNoRows {
		s := fmt.Sprintf("query not found id: %v", id)
		return nil, errors.New(s)
	} else if err == orm.ErrMissPK {
		s := fmt.Sprintf("query pk err:%v", err)
		return nil, errors.New(s)
	} else {
		logger.Debugf("id:%v", m.Id)
	}
	// todo delete
	if m.IsDelete {
		return nil, errors.New(fmt.Sprintf("msg id:%v is not exist", m.Id))
	}
	return &m, nil
}

func GetMessgeByIdWithFromUserId(id int64, fromUserId int64) (msg *Messge, err error) {
	m := Messge{Id: id}
	o := orm.NewOrm()
	o.Using("default")
	err = o.Read(&m)
	if err == orm.ErrNoRows {
		s := fmt.Sprintf("query not found id: %v", id)
		return nil, errors.New(s)
	} else if err == orm.ErrMissPK {
		s := fmt.Sprintf("query pk err:%v", err)
		return nil, errors.New(s)
	} else {
		logger.Debugf("id:%v", m.Id)
	}
	// todo delete
	if m.IsDelete {
		return nil, errors.New(fmt.Sprintf("msg id:%v is not exist", m.Id))
	}
	if m.FromUserId != fromUserId {
		return nil, errors.New(fmt.Sprintf("msg id:%v FromUserId:%v , but you want fromUserId:%v", m.Id, m.FromUserId, fromUserId))
	}
	return &m, nil
}

func UpdateMessge(mid int64, mm *Messge) (msg *Messge, err error) {
	m, err := GetMessgeById(mid);
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Msg Not Exist mid:%v err:%v", mid, err))
	}
	// todo 是否空判断
	if mm.FromUserId >= 0 {
		m.FromUserId = mm.FromUserId
	}

	if mm.ToUserId >= 0 {
		m.ToUserId = mm.ToUserId
	}

	if mm.Title != "" {
		m.Title = mm.Title
	}

	if mm.Message != "" {
		m.Message = mm.Message
	}

	if mm.Status != "" {
		m.Status = checkStatus(mm.Status)
	}

	o := orm.NewOrm()
	o.Using("default")
	_, err = o.Update(m)
	if err != nil{
		return nil, err
	}
	return m, nil
}

func DeleteMessge(id int64) error {
	o := orm.NewOrm()
	o.Using("default")
	_, e := o.Update(&Messge{Id: id, IsDelete: true}, "is_delete")
	return e
}

func AllMessge() (id int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(TABLENAME)
	return qs.Count()
}