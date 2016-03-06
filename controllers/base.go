package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (self *BaseController) Prepare() {
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["curRoute"] = self.controllerName + "." + self.actionName
	self.Data["curController"] = self.controllerName
	self.Data["curAction"] = self.actionName
}
