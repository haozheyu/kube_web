package models_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(getParentDirectory(appPath))
	initDb()
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func initDb() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbuser, err := beego.AppConfig.String("DBUser")
	checkErr(err)
	dbpassword, err := beego.AppConfig.String("DBPasswd")
	checkErr(err)
	dbtns, err := beego.AppConfig.String("DBTns")
	checkErr(err)
	dbname, err := beego.AppConfig.String("DBName")
	checkErr(err)
	dbURL := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&",
		dbuser,
		dbpassword,
		dbtns,
		dbname)
	// set timezone  , same as db timezone
	dbloc, err := beego.AppConfig.String("DBLoc")
	checkErr(err)
	dbURL += dbloc
	orm.RegisterDataBase("default", "mysql", dbURL)

	orm.Debug = beego.AppConfig.DefaultBool("ShowSql", false)
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
