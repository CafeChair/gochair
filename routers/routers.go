package routers

import (
	"gochair/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
    beego.Router("/qdns", &controllers.QdnsController{})
}