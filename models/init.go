package models

import (
	"net/url"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //必须导入mysql 驱动
)
/**
连接数据库
 */
func Init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	//实例化 与数据看连接的 实例
	//orm.RegisterModel(new(User), new(Task), new(TaskGroup), new(TaskLog))
	orm.RegisterModel(new(Demo))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}