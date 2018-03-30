package models

import (
	"errors"
	"bmsg/logger"
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

const TABLENAME = "messge_test1"

type Messge struct {
	Id          int64          `orm:"column(id)"`
	FromUserId  string 	   		`orm:"column(from_user_id);"`
	ToUserId    string 	   		`orm:"column(to_user_id);""`
	CreatedAt   time.Time 	   `orm:"null;auto_now_add;type(datetime);column(created_at)"`
	UpdateAt    time.Time 	   `orm:"null;auto_now;type(datetime);column(update_at)"`
	Title       string         `orm:"null;column(title)"`
	Text        string         `orm:"null;column(text)"`
	IsDelete    bool           `orm:"null;type(bool);column(is_delete);default(false)`
	Status      int            `orm:"null;column(status);default(0)"`
}

func (u *Messge) TableName() string {
	return TABLENAME
}

// 多字段索引
func (u *Messge) TableIndex() [][]string {
	return [][]string{
		[]string{"FromUserId", "ToUserId", "Status", "IsDelete"},
	}
}

//type CommonTask LogCommonTask
var (
	ErrFetchTaskTimeout = errors.New("Fetch task timeout.")
	dORM     orm.Ormer
)

func init() {
	ormInit()
}

func ormInit()  {
	orm.RegisterModel(new(Messge))
	dORM = NewOrm()
}

func NewOrm() orm.Ormer {
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		panic(err)
	}
	return o
}


func NewMessge() *Messge {
	return &Messge{}
}

func AddMessge(t *Messge) (tid int64, e error) {
	if t == nil {
		return 0, errors.New("AddCommonTask task is nil")
	}

	msg := new(Messge)
	msg.ToUserId = t.ToUserId
	msg.FromUserId = t.FromUserId
	msg.Title = t.Title
	msg.Text = t.Text
	msg.IsDelete = false
	// 0: 未读  1: 已读
	msg.Status = 0
	_, err := dORM.Insert(msg)
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
	successNums, err := dORM.InsertMulti(100, msgList)
	return successNums, err
}

func GetMessge(id int64, toUserId string, status int, isDelete bool) (msg *Messge, err error) {
	m := Messge{Id: id, ToUserId: toUserId, Status: status, IsDelete: isDelete }
	err = dORM.Read(&m)
	if err == orm.ErrNoRows {
		//logger.Debug("查询不到")
		s := fmt.Sprintf("query not found id: %v", id)
		return nil, errors.New(s)
	} else if err == orm.ErrMissPK {
		s := fmt.Sprintf("query pk err , id:%v", id)
		return nil, errors.New(s)
	} else {
		logger.Debugf("id:%v", m.Id)
	}
	return &m, nil
}

func GetMessgeById(id int64) (msg *Messge, err error) {
	m := Messge{Id: id}

	err = dORM.Read(&m)
	if err == orm.ErrNoRows {
		//logger.Debug("查询不到")
		s := fmt.Sprintf("query not found id: %v", id)
		return nil, errors.New(s)
	} else if err == orm.ErrMissPK {
		s := fmt.Sprintf("query pk err , id:%v", id)
		return nil, errors.New(s)
	} else {
		logger.Debugf("id:%v", m.Id)
	}
	return &m, nil
}

func UpdateMessge(mid int64, mm *Messge) (a *Messge, err error) {
	if m, err  := GetMessgeById(mid); err != nil {
		if mm.FromUserId != "" {
			m.FromUserId = mm.FromUserId
		}

		if mm.ToUserId != "" {
			m.ToUserId = mm.ToUserId
		}

		if mm.Title != "" {
			m.Title = mm.Title
		}

		if mm.Text != "" {
			m.Text = mm.Text
		}

		_, err = dORM.Update(m)
		if err != nil{
			return nil, err
		}
		return m, nil
	}
	return nil, errors.New(fmt.Sprintf("Msg Not Exist mid:%v", mid))
}

func DeleteMessge(id int64) error {
	_, e := dORM.Update(Messge{Id: id, IsDelete: true}, "is_delete")
	return e
}

func AllMessge() (id int64, err error) {
	qs := dORM.QueryTable(TABLENAME)
	return qs.Count()
}
