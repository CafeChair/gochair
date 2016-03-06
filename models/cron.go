package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	// "time"
)

type Task struct {
	Id       int
	TaskName string
	CronSpec string
	Command  string
	Timeout  int
}

func AddTask(self *Task) (int64, error) {
	if self.TaskName == "" {
		return 0, fmt.Errorf("TaskName can not be null")
	}
	if self.CronSpec == "" {
		return 0, fmt.Errorf("CronSpec can not be null")
	}
	if self.Command == "" {
		return 0, fmt.Errorf("Command can not be null")
	}
	return orm.NewOrm().Insert(self)
}
