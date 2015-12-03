package routers

import (
	"github.com/astaxie/beego"
	"gochair/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/qdns", &controllers.QdnsController{})
	beego.Router("/runcmd", &controllers.RunCmdController{})
	beego.Router("/agent", &controllers.AgentController{})
}
