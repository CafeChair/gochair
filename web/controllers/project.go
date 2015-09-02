package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type ProjectController struct {
	beego.Controller
}

func (self *ProjectController) Get() {
	self.Data["IsProject"] = true
	self.TplNames = "project.html"
	key := "/"
	projectnames, err := models.GetAllProject(key)
	fi err !=nil {
		beego.Error(err)
	} else {
		self.Data["ProjectNames"] = projectnames
	}
}

func (self *ProjectController) Post() {
	projectname := self.Input().Get("projectname")
	var err error
	err = models.AddProject(projectname)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/project", 302)
}

func (self *ProjectController) Add() {
	self.TplNames = "project_add.html"
}