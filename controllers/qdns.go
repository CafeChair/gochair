package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gochair/models"
)

type QdnsController struct {
	beego.Controller
}

func (self *QdnsController) Get() {
	self.Data["IsQdns"] = true
	self.TplNames = "qdns.html"
}

func (self *QdnsController) Post() {
	domain := self.Input().Get("domain")
	dnss, err := models.ResolveFromRedis(domain)
	if err != nil {
		self.Ctx.WriteString(fmt.Sprint(err))
	}
	self.Ctx.WriteString(fmt.Sprint(dnss))
}
