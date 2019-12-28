package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/logs"
	_ "studyBeego/routers"
	_ "studyBeego/sysinit"
)

func main() {
	logs.SetLevel(logs.LevelInformational)
	logs.SetLogger("file", `{"filename":"./logs/test.log"}`)
	beego.Run()
}
