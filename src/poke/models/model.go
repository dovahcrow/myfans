package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
	. "tools"
)

type User struct {
	Id       int
	Username string
	Password string
}

var O orm.Ormer

func init() {
	fmt.Print(``)
	orm.RegisterDataBase("default", "sqlite3", "database.db")
	orm.RegisterModel(new(User), new(Threads), new(AvFids), new(AvTids), new(Record), new(ThreadRecord))
	//orm.Debug = true
	orm.RunCommand()
	O = orm.NewOrm()
}

func GetAllUsers() (ret []*User, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	_, err = O.QueryTable(new(User)).All(&ret)
	E(`获取所有用户错误`, err)
	return
}
func GetUserById(id int) (ret *User, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	ret = &User{}
	ret.Id = id
	err = O.Read(ret)
	E(`根据id读取用户错误`, err)
	return
}

func CreateUser(username, password string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	ret := &User{}
	ret.Password = password
	ret.Username = username
	_, err = O.Insert(ret)
	E(`新建用户错误`, err)
	return
}

func UpdateUser(id int, username, password string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	ret := &User{}
	ret.Password = password
	ret.Username = username
	ret.Id = id
	_, err = O.Update(ret)
	E(`更新用户错误`, err)
	return
}
func DeleteUser(id int) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	ret := &User{}
	ret.Id = id
	_, err = O.Delete(ret)
	E(`删除用户错误`, err)
	return
}

type Threads struct {
	Id    int
	Title string
	Text  string
	Sends int
}

func GetAllThreads() (ret []*Threads, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	_, err = O.QueryTable(new(Threads)).All(&ret)
	E(`获取所有帖子错误`, err)
	return
}
func GetThreadById(id int) (ret *Threads, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	ret = &Threads{}
	ret.Id = id
	err = O.Read(ret)
	E(`根据id读取帖子错误`, err)
	return
}

func CreateThread(title, text string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	ret := &Threads{}
	ret.Text = text
	ret.Title = title
	ret.Sends = 0
	_, err = O.Insert(ret)
	E(`新建帖子错误`, err)
	return
}
func AddThreadSends(id int) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	t := new(Threads)
	t.Id = id
	err = O.Read(t)
	E(`修改帖子发帖数失败`, err)
	t.Sends++
	_, err = O.Update(t)
	E(`修改帖子发帖数失败`, err)
	return

}
func UpdateThread(id int, title, text string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	ret := &Threads{}
	ret.Text = text
	ret.Title = title
	ret.Id = id
	_, err = O.Update(ret)
	E(`更新帖子错误`, err)
	return
}
func DeleteThread(id int) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	ret := &Threads{}
	ret.Id = id
	_, err = O.Delete(ret)
	E(`删除帖子错误`, err)
	return
}

type AvFids struct {
	Id   int
	Fids string
}

func DeleteAllFids() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	_, err = O.Raw("delete from av_fids").Exec()
	return
}

func InsertFids(fids []string) (err error) {
	insertion := []AvFids{}

	for _, v := range fids {
		insertion = append(insertion, AvFids{Fids: v})
	}
	if len(insertion) == 0 {
		return
	}

	_, err = O.InsertMulti(1, insertion)
	return
}
func GetAllFids() (fids []*AvFids, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	_, err = O.QueryTable(new(AvFids)).All(&fids)
	return
}

type AvTids struct {
	Id   int
	Tids string
}

func DeleteAllTids() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	_, err = O.Raw("delete from av_tids").Exec()
	return
}

func InsertTids(tids []string) (err error) {

	insertion := []AvTids{}

	for _, v := range tids {
		insertion = append(insertion, AvTids{Tids: v})
	}
	if len(insertion) == 0 {
		return
	}
	num := len(insertion)
	if num >= 10 {
		num = 10
	}
	_, err = O.InsertMulti(1, insertion)
	return
}
func GetAllTids() (tids []*AvTids, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	_, err = O.QueryTable(new(AvTids)).All(&tids)
	return
}

type Record struct {
	Id       int
	Title    string
	Username string
	Ret      string `orm:"null"`
	Succ     bool
	Time     time.Time `orm:"auto_now"`
}

func AddRecord(username string, ret string, Title string, succ bool) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	r := &Record{}
	r.Username = username
	r.Ret = ret
	r.Title = Title
	r.Succ = succ
	_, err = O.Insert(r)
	beego.Info("insert a record")
	E("insert record fail", err)
	return
}

func GetAllRecords() (r []*Record, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	_, err = O.QueryTable(new(Record)).All(&r)
	beego.Info("get all record")
	E("get all record fail", err)
	return
}
func GetRecordsCountSuccOrNot(succ bool) (count int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	i, err := O.QueryTable(new(Record)).Filter("Succ", succ).Count()
	E("get record count succ fail", err)
	count = i
	return
}
func GetRecordsRange(offset, num int) (r []*Record, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	_, err = O.QueryTable(new(Record)).OrderBy(`-id`).Limit(num).Offset(offset).All(&r)
	beego.Info("get record range", offset, num)
	E("get record range fail", err)
	return
}
func GetRecordsCounts() (i int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	i, err = O.QueryTable(new(Record)).Count()
	E("get record count fail", err)
	return
}
func SearchRecord(tp []string, val []interface{}) (r []*Record, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	qst := O.QueryTable(new(Record))

	for i, _ := range tp {
		qst = qst.Filter(tp[i], val[i])
	}
	_, err = qst.All(&r)
	return
}

type ThreadRecord struct {
	Id       int
	Title    string
	Username string
	Ret      string `orm:"null"`
	Succ     bool
	Time     time.Time `orm:"auto_now"`
}

func AddThreadRecord(username string, ret string, Title string, succ bool) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	r := &ThreadRecord{}
	r.Username = username
	r.Ret = ret
	r.Title = Title
	r.Succ = succ
	_, err = O.Insert(r)
	beego.Info("insert a record")
	E("insert record fail", err)
	return
}

func GetAllThreadRecords() (r []*ThreadRecord, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	_, err = O.QueryTable(new(ThreadRecord)).All(&r)
	beego.Info("get all record")
	E("get all record fail", err)
	return
}
func GetThreadRecordsCountSuccOrNot(succ bool) (count int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	i, err := O.QueryTable(new(ThreadRecord)).Filter("Succ", succ).Count()
	E("get record count succ fail", err)
	count = i
	return
}
func GetThreadRecordsRange(offset, num int) (r []*ThreadRecord, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	_, err = O.QueryTable(new(ThreadRecord)).OrderBy(`-id`).Limit(num).Offset(offset).All(&r)
	beego.Info("get record range", offset, num)
	E("get record range fail", err)
	return
}
func GetThreadRecordsCounts() (i int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	i, err = O.QueryTable(new(ThreadRecord)).Count()
	E("get record count fail", err)
	return
}
func SearchThreadRecord(tp []string, val []interface{}) (r []*ThreadRecord, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	qst := O.QueryTable(new(ThreadRecord))

	for i, _ := range tp {
		qst = qst.Filter(tp[i], val[i])
	}
	_, err = qst.All(&r)
	return
}
