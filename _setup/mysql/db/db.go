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
			conditions = append(conditions, columnName+" = ?")
			args = append(args, fieldValue)
		}
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	return whereClause, args
}