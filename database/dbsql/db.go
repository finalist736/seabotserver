package dbsql

import (
	"fmt"
	"time"

	"github.com/finalist736/seabotserver/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

var db *dbr.Connection

func session() *dbr.Session {
	if db == nil {
		conf := config.GetConfiguration()
		connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			conf.DB.User,
			conf.DB.Pass,
			conf.DB.Host,
			conf.DB.Port,
			conf.DB.Name)
		var err error
		db, err = dbr.Open("mysql", connectString, nil)
		if err != nil {
			return &dbr.Session{}
		}
		db.SetConnMaxLifetime(time.Minute * 10)
		db.SetMaxIdleConns(1)
		db.SetMaxOpenConns(5)
	}
	return db.NewSession(nil)
}
