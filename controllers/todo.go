package controllers

import (
	"github.com/astaxie/beego"
	"gochair/models"
)

type TodoController struct {
	beego.Controller
}

func (self *TodoController) Add() {
	self.TplNames = "todo_add.html"
}

func (self *TodoController) Post() {
	// tid := self.Input().Get("tid")
	title := self.Input().Get("title")
	err := models.AddTodo(title)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/todo", 302)
}

func (self *TodoController) Get() {
	self.Data["IsTodo"] = true
	var err error
	self.Data["Todos"], err = models.GetALlTodo()
	if err != nil {
		beego.Error(err)
	}
	self.TplNames = "todo.html"
}
