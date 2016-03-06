package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Todo struct {
	Id     int64
	Title  string
	Finish bool
}

func AddTodo(title string) error {
	o := orm.NewOrm()
	todo := &Todo{Title: title}
	_, err := o.Insert(todo)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTodo() ([]*Todo, error) {
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
