package modeles

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	Id      int
	Website string
	Email   string
}

func init() {
	//orm.RegisterModel(new(Page))
}
func GetPage() Page {
	o := orm.NewOrm()
	data := Page{Id: 1}
	err := o.Read(&data)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(data)
	return data
}

func UpdatePage() {
	o := orm.NewOrm()
	p := Page{Id: 1, Email: "Email_Updated2"}
	o.Update(&p, "Email")
}
