package controllers

import (
	"github.com/astaxie/beego"
)

type AgentController struct {
	beego.Controller
}

func (self *AgentController) Get() {
	self.Data["IsAgent"] = true
	self.TplNames = "agent.html"
}
