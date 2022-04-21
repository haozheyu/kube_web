package initial

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"kube_web/database/initial"
	"kube_web/util/logs"
)

const DbDriverName = "mysql"

func InitDb() {
	orm.RegisterDriver(DbDriverName, orm.DRMySQL)
	// ensure database exist
	err := ensureDatabase()
	if err != nil {
		panic(err)
	}
	db, err := orm.GetDB()
	if err != nil {
		panic(err)
	}

	ttl := beego.AppConfig.DefaultInt("DBConnTTL", 30)

	db.SetConnMaxLifetime(time.Duration(ttl) * time.Second)

	orm.Debug = beego.AppConfig.DefaultBool("ShowSql", false)
}

func ensureDatabase() error {
	needInit := false
	dbName, err := beego.AppConfig.String("DBName")
	if err != nil {
		logs.Error("配置文件读取数据库名称失败", err.Error())
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

	dbURL := fmt.Sprintf("%s:%s@%s/",
		dbuser,
		dbpasswd,
		dbtns,
	)

	db, err := sql.Open(DbDriverName, fmt.Sprintf("%s%s", dbURL, dbName))
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		switch e := err.(type) {
		case *mysql.MySQLError:
			// MySQL error unkonw database;
			// refer https://dev.mysql.com/doc/refman/5.6/en/error-messages-server.html
			if e.Number == 1049 {
				needInit = true
				dbForCreateDatabase, err := sql.Open(DbDriverName, addLocation(dbURL))
				if err != nil {
					return err
				}
				defer dbForCreateDatabase.Close()
				_, err = dbForCreateDatabase.Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8 COLLATE utf8_general_ci;", dbName))
				if err != nil {
					return err
				}

			} else {
				return err
			}
		default:
			return err
		}
	}
	err = orm.RegisterDataBase("default", "mysql", addLocation(fmt.Sprintf("%s%s", dbURL, dbName)))
	if err != nil {
		return err
	}
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		return err
	}
	if needInit {
		o := orm.NewOrm()
		for _, insertSql := range initial.InitialData {
			_, err = o.Raw(insertSql).Exec()
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func addLocation(dbURL string) string {
	// https://stackoverflow.com/questions/30074492/what-is-the-difference-between-utf8mb4-and-utf8-charsets-in-mysql
	return fmt.Sprintf("%s?charset=utf8mb4&loc=%s", dbURL, beego.AppConfig.DefaultString("DBLoc", "Asia%2FShanghai"))
}
