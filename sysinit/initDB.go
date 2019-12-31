package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
)

func initDB() {
	dbAlias := beego.AppConfig.String("db_alias")
	dbName := beego.AppConfig.String("db_name")
	dbUser := beego.AppConfig.String("db_user")
	dbPwd := beego.AppConfig.String("db_pwd")
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbCharset := beego.AppConfig.String("db_charset")

	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")"+"/"+dbName+"?charset="+dbCharset)

	isDev := (beego.AppConfig.String("runmode") == "dev")

	//自动创建数据表 参数:数据库别名,数据库是否强制刷新(false的话只会在有重大变化的时候刷新(比如新增字段),是否显示创建过程
	orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}
}
