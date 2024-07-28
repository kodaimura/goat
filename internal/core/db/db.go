package db

import (
	"log"
	"fmt"
	"reflect"
	"strings"	
	"database/sql"
	
	_ "github.com/mattn/go-sqlite3"

	"goat/config"
	"goat/internal/core/utils"
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

func BuildWhereClause(filter interface{}) (string, []interface{}) {
	var conditions []string
	var binds []interface{}

	val := reflect.ValueOf(filter)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		columnName := typ.Field(i).Tag.Get("db")
		fieldValue := val.Field(i).Interface()

		if columnName == "" { 
			continue
		}
		if utils.IsZero(fieldValue) {
			continue
		}

		conditions = append(conditions, fmt.Sprintf("%s = ?", columnName))
		binds = append(binds, fieldValue)
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	return whereClause, binds
}