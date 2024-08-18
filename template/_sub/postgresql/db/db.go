package db

import (
	"log"
	"fmt"
	"reflect"
	"strings"	
	"database/sql"
	
	_ "github.com/lib/pq"

	"goat/config"
	"goat/internal/core/utils"
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

func BuildWhereClause(filter interface{}) (string, []interface{}) {
	var conditions []string
	var binds []interface{}

	val := reflect.ValueOf(filter)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	seq := 1
	for i := 0; i < val.NumField(); i++ {
		columnName := typ.Field(i).Tag.Get("db")
		fieldValue := val.Field(i).Interface()

		if columnName == "" { 
			continue
		}
		if utils.IsZero(fieldValue) {
			continue
		}

		conditions = append(conditions, fmt.Sprintf("%s = $%d", columnName, seq))
		binds = append(binds, fieldValue)
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	return whereClause, binds
}