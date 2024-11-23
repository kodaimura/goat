package service

import (
    "strings"
    "github.com/lib/pq"
    "github.com/go-sql-driver/mysql"
    "github.com/mattn/go-sqlite3"
)


func GetConflictColumn(err error) (string, bool) {
    if err == nil {
        return "", false
    }

	if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
        if pgErr.Detail != "" {
			parts := strings.Split(pgErr.Detail, "violates unique constraint")
			if len(parts) > 0 {
				return strings.TrimSpace(parts[0]), true
			}
		}
    } else if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
        if strings.Contains(mysqlErr.Message, "for key") {
			parts := strings.Split(mysqlErr.Message, "for key")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1]), true
			}
		}
    } else if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
        if strings.Contains(sqliteErr.Error(), "UNIQUE constraint failed") {
			parts := strings.Split(sqliteErr.Error(), ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1]), true
			}
		}
    }

	return "", false
}