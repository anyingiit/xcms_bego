package routers

import (
	"github.com/anyingiit/xcms_beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
