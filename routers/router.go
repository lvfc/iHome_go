package routers

import (
	"github.com/astaxie/beego"
	"iHome_go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreasController{}, "get:GetAreas")
}
