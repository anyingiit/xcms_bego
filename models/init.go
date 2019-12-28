package modeles

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
)

func init() { //所有的RegisterModule都在这里进行
	orm.RegisterModel(new(MenuModel))
}
