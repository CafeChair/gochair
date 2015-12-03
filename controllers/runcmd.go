package controllers

import (
	"github.com/astaxie/beego"
)

type RunCmdController struct {
	beego.Controller
}

func (self *RunCmdController) Get() {
	self.Data["IsRunCmd"] = true
	self.TplNames = "runcmd.html"
}
