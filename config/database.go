package config

import (
	"database/sql"
	"fmt"
	"sync"

	// import init mysql
	_ "github.com/go-sql-driver/mysql"
)

// MustInitDB returns DB pointer.
func MustInitDB(cfg *Config) *sql.DB {
	var doOnce sync.Once
	var db *sql.DB
	var err error
	doOnce.Do(func() {
		db, err = sql.Open("mysql", conStr(cfg))
		if err != nil {
			// retry 3 times with 3 seconds here
			panic(err)
		}
	})

	return db
}

func conStr(cfg *Config) string {
	return fmt.Sprintf("%s:%s@/%s", cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Name)
}
