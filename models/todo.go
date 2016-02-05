package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
)

const (
	_DB_NAME        = "data/gochair.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Todo struct {
	Id     int64
	Title  string
	Finish bool
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Todo))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddTodo(title string) error {
	o := orm.NewOrm()
	todo := &Todo{Title: title}
	qs := o.QueryTable("todo")
	err := qs.Filter("title", title).One(todo)
	if err == nil {
		return err
	}
	_, err = o.Insert(todo)
	if err != nil {
		return err
	}
	return nil
}

func GetALlTodo() ([]*Todo, error) {
	o := orm.NewOrm()
	todos := make([]*Todo, 0)
	qs := o.QueryTable("todo")
	_, err := qs.All(&todos)
	return todos, err
}

func DelTodo(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	todo := &Todo{Id: tid}
	_, err = o.Delete(todo)
	return err
}
