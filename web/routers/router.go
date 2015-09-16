package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/Project", &controllers.ProjectController{})
    beego.Router("/Command", &controllers.CommandController{})
    beego.Router("/Loginfo", &controllers.LoginfoController{})
    beego.Router("/Task", &controllers.TaskController{})
}