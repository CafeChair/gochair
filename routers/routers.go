package routers

import (
	"github.com/astaxie/beego"
	"gochair/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/qdns", &controllers.QdnsController{})
	beego.Router("/runcmd", &controllers.RunCmdController{})
	beego.Router("/cron", &controllers.CronController{})
	beego.Router("/agent", &controllers.AgentController{})
	beego.Router("/runjob", &controllers.RunJobController{})
	beego.Router("/todo", &controllers.TodoController{})
	beego.Router("/doc", &controllers.DocController{})
	beego.AutoRouter(&controllers.TodoController{})
}
