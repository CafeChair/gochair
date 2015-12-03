package controllers

import (
	"fmt"
	// "gochair/models"
	"github.com/miekg/dns"
	"github.com/astaxie/beego"
)


type QdnsController struct {
	beego.Controller
}

func (self *QdnsController) Get() {
	self.Data["IsQdns"] = true
	self.TplNames = "qdns.html"
}

func (self *QdnsController) Post() {
	domain:=self.Input().Get("domain")
	dnss,err:= resolveDNS(domain)
	if err !=nil {
		self.Ctx.WriteString(fmt.Sprint(err))
	}
	self.Ctx.WriteString(fmt.Sprint(dnss))
}

func resolveDNS(domain string) ([]string,error) {
	answer := make([]string,0)
	m:= new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain),dns.TypeA)
	c := new(dns.Client)
	in,_,err := c.Exchange(m,"119.29.29.29"+":53")
	if err !=nil {
		return answer,err
	}
	for _,ain := range in.Answer {
		if a,ok := ain.(*dns.A);ok {
			answer = append(answer,a.A.String())
		}
	}
	return answer,nil
}