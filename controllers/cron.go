package controllers

import (
	"github.com/astaxie/beego"
)

type CronController struct {
	beego.Controller
}

func (self *CronController) Get() {
	self.Data["IsCron"] = true
	self.TplNames = "cron.html"
}
