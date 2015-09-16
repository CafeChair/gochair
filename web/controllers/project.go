package controllers

import (
	"github.com/astaxie/beego"
)

type ProjectController struct {
	beego.Controller
}

func (c *ProjectController) Get() {
	c.Data["IsProject"] = true
	c.TplNames = "project.html"
}