package controllers

import (
	"github.com/astaxie/beego"
)

type LoginfoController struct {
	beego.Controller
}

func (c *LoginfoController) Get() {
	c.Data["IsLoginfo"] = true
	c.TplNames = "loginfo.html"
}