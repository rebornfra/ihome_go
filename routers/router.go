package routers

import (
	"github.com/astaxie/beego"
	"ihome_go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
