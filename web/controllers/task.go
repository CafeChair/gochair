package controllers

import (
	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

func (c *TaskController) Get() {
	c.Data["IsTask"] = true
	c.TplNames = "task.html"
}