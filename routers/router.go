package routers

import (
	"github.com/astaxie/beego"
	"studyBeego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
