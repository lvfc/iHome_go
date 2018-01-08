package routers

import (
	"github.com/astaxie/beego"
	"iHome_go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
