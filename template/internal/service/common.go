package service

import (
	"regexp"
    "github.com/lib/pq"
    "github.com/go-sql-driver/mysql"
    "github.com/mattn/go-sqlite3"
)


func GetConflictColumn(err error) (string, bool) {
    if err == nil {
        return "", false
    }

	if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
        re := regexp.MustCompile(`Key \((\w+)\)=`)
		match := re.FindStringSubmatch(pgErr.Detail)
		if len(match) > 1 {
			return match[1], true
		}
    } else if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
        re := regexp.MustCompile(`for key '([\w.]+)'`)
		match := re.FindStringSubmatch(mysqlErr.Message)
		if len(match) > 1 {
			return match[1], true
		}
    } else if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
        re := regexp.MustCompile(`UNIQUE constraint failed: ([\w.]+)`)
		match := re.FindStringSubmatch(sqliteErr.Error())
		if len(match) > 1 {
			return match[1], true
		}
    }

	return "", false
}