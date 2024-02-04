package db

import (
	"fmt"
	"log"
	"database/sql"
	
	_ "github.com/lib/pq"

	"goat/config"
)


var db *sql.DB

func init() {
	var err error

	cf := config.GetConfig()

	db, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cf.DBHost, cf.DBPort, cf.DBUser, cf.DBPass, cf.DBName,
		),
	)

	if err != nil {
		log.Panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}