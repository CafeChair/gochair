package controllers

import (
	"github.com/astaxie/beego"
)

type JobsController struct {
	beego.Controller
}

func (self *JobsController) Get() {
	self.Data["IsJobs"] = true
	self.TplNames = "jobs.html"
}
