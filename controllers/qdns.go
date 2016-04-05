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
	dnss, err := models.ResolveFromRedis2(domain)

	if err != nil {
		self.Ctx.WriteString(fmt.Sprint(err))
	} else {
		self.Data["QDNS"] = dnss
		self.TplNames = "qdns.html"
	}
	// self.Ctx.WriteString(fmt.Sprint(dnss))
}
