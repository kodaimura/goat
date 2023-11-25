package db

import (
	"fmt"
	"log"
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"

	"goat/config"
)


var db *sql.DB

func init() {
	var err error

	cf := config.GetConfig()

	db, err = sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			cf.DBUser, cf.DBPass, cf.DBHost, cf.DBPort, cf.DBName,
		),
	)

	if err != nil {
		log.Panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}