package controllers

import (
	"github.com/astaxie/beego"
)

type DocController struct {
	beego.Controller
}

func (self *DocController) Get() {
	self.Data["IsDoc"] = true
	self.TplNames = "doc.html"
}
