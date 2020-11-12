package routers

import (
	"github.com/astaxie/beego"
	"myproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
