package controllers

import (
	"github.com/astaxie/beego"
)

type CommandController struct {
	beego.Controller
}

func (c *CommandController) Get() {
	c.Data["IsCommand"] = true
	c.TplNames = "command.html"
}