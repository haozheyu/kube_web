package main

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	err := beego.LoadAppConfig("ini", "src/backend/conf/app.conf")
	if err != nil {
		panic(err)
	}
}

func main() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(err)
	}
	dbuser, err := beego.AppConfig.String("DBUser")
	if err != nil {
		panic(err)
	}
	dbpasswd, err := beego.AppConfig.String("DBPasswd")
	if err != nil {
		panic(err)
	}
	dbtns, err := beego.AppConfig.String("DBTns")
	if err != nil {
		panic(err)
	}
	dbname, err := beego.AppConfig.String("DBName")
	if err != nil {
		panic(err)
	}
	dbloc, err := beego.AppConfig.String("DBLoc")
	if err != nil {
		panic(err)
	}
	dbURL := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&", dbuser,
		dbpasswd,
		dbtns,
		dbname,
	)
	// set timezone  , same as db timezone
	dbURL += dbloc

	err = orm.RegisterDataBase("default", "mysql", dbURL)
	if err != nil {
		panic(err)
	}
	orm.RunCommand()
}
