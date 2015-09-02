package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
    beego.Router("/project", &controllers.ProjectController{})
    beego.Router("/task", &controllers.TaskController{})
    beego.AutoRouter(&controllers.ProjectController{})
}