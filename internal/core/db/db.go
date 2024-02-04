package db

import (
	"log"
	"database/sql"
	
	_ "github.com/mattn/go-sqlite3"

	"goat-base/config"
)


var db *sql.DB

func init() {
	var err error

	cf := config.GetConfig()

	db, err = sql.Open("sqlite3", cf.DBName)

	if err != nil {
		log.Panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}