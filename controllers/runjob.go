package controllers

import (
	"github.com/astaxie/beego"
)

type RunJobController struct {
	beego.Controller
}

func (self *RunJobController) Get() {
	self.Data["IsRunJob"] = true
	self.TplNames = "runjob.html"
}
