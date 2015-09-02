package controllers

import (
	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

func (self *TaskController) Get() {
	self.Data["IsTask"] = true
	self.TplNames = "task.html"
}
