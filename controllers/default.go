package controllers

import (
	_ "github.com/anyingiit/xcms_beego/models"
	modeles "github.com/anyingiit/xcms_beego/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
