package controllers

import (
	"github.com/astaxie/beego"
)

type TodoController struct {
	beego.Controller
}

func (self *TodoController) Get() {
	self.Data["IsTodo"] = true
	self.TplNames = "todo.html"
}
