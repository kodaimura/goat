package db

import (
	"log"
	"fmt"
	"reflect"
	"strings"	
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

func BuildWhereClause(filter interface{}) (string, []interface{}) {
	var conditions []string
	var args []interface{}

	val := reflect.ValueOf(filter)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i).Interface()
		columnName := field.Tag.Get("db")

		if columnName == "" {
			continue
		}

		isZero := false
		switch fieldValue.(type) {
		case string:
			isZero = fieldValue == ""
		case int:
			isZero = fieldValue == 0
		}

		if !isZero {
			conditions = append(conditions, dmt.Sprintf("%s = $%d", columnName, i++))
			args = append(args, fieldValue)
		}
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	return whereClause, args
}