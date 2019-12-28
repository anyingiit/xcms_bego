package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "studyBeego/models"
	modeles "studyBeego/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	logs.Informational("user login in")
	modeles.UpdatePage()
	data := modeles.GetPage()
	c.Data["Website"] = data.Website
	c.Data["Email"] = data.Email
	c.TplName = "index.tpl"
}
