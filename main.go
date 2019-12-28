package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/logs"
	_ "github.com/anyingiit/xcms_beego/routers"
	_ "github.com/anyingiit/xcms_beego/sysinit"
)

func main() {
	logs.SetLevel(logs.LevelInformational)
	logs.SetLogger("file", `{"filename":"./logs/test.log"}`)
	beego.Run()
}
