package controllers

import (
	"github.com/astaxie/beego"
)


type IndexController struct {
	beego.Controller
}

func (self *IndexController) Get() {
	self.Data["IsIndex"] = true
	self.TplNames = "index.html"
	
}