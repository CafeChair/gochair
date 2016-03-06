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
	self.TplNames = "todo.html"
	todos, err := models.GetAllTodo()
	if err != nil {
		beego.Error(err)
	} else {
		self.Data["Todos"] = todos
	}
}
