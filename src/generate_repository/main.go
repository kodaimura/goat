package main
 
import (
	"io"
	"os"
	"fmt"
	"strings"
	
	"github.com/kodaimura/ddlparse"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Error: The path to the DDL must be provided as the first argument.")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	
	tables, _ := ddlparse.ParseForce(string(data))
	
	for _, table := range tables {
		tn := strings.ToLower(table.Name)
		s := generateRepositoryCode(table)

		file, err := os.Create(tn + ".go")
		if err != nil {
			fmt.Println("ファイルの作成に失敗しました:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(s)
		if err != nil {
			fmt.Println("ファイルへの書き込みに失敗しました:", err)
			return
		}
	}
}

func snakeToPascal(snake string) string {
	ls := strings.Split(strings.ToLower(snake), "_")
	for i, s := range ls {
		ls[i] = strings.ToUpper(s[0:1]) + s[1:]
	}
	return strings.Join(ls, "")
}

func snakeToCamel(snake string) string {
	ls := strings.Split(strings.ToLower(snake), "_")
	for i, s := range ls {
		if i != 0 {
			ls[i] = strings.ToUpper(s[0:1]) + s[1:]
		}
	}
	return strings.Join(ls, "")
}


func getSnakeInitial(snake string) string {
	ls := strings.Split(strings.ToLower(snake), "_")
	ret := ""
	for _, s := range ls {
		ret += s[0:1]
	}
	return ret
}

func dataTypeToGoType(dataType string) string {
	dataType = strings.ToUpper(dataType)

	if (strings.Contains(dataType, "INT") || strings.Contains(dataType, "BIT")) {
		return "int"
	} else if strings.Contains(dataType, "NUMERIC") ||
		strings.Contains(dataType, "DECIMAL") ||
		strings.Contains(dataType, "FLOAT") ||
		strings.Contains(dataType, "REAL") ||
		strings.Contains(dataType, "DOUBLE") {
		return "float64"
	} else {
		return "string"
	}
}

const TEMPLATE = 
`package repository

import (
	"database/sql"

	"xxxxx/internal/core/db"
	"xxxxx/internal/model"
)


type %sRepository interface {
	Get(%s *model.%s) ([]model.%s, error)
	GetOne(%s *model.%s) (model.%s, error)
	Insert(%s *model.%s, tx *sql.Tx) error
	Update(%s *model.%s, tx *sql.Tx) error
	Delete(%s *model.%s, tx *sql.Tx) error
}


type %sRepository struct {
	db *sql.DB
}

func New%sRepository() %sRepository {
	db := db.GetDB()
	return &%sRepository{db}
}


%s


%s


%s


%s


%s`

const TEMPLATE_GET =
`func (rep *%sRepository) Get(%s *model.%s) ([]model.%s, error) {
	where, binds := db.BuildWhereClause(%s)
	query := %s + where
	rows, err := rep.db.Query(query, binds...)
	defer rows.Close()

	if err != nil {
		return []model.%s{}, err
	}

	ret := []model.%s{}
	for rows.Next() {
		%s := model.%s{}
		err = rows.Scan(%s)
		if err != nil {
			return []model.%s{}, err
		}
		ret = append(ret, %s)
	}

	return ret, nil
}`

const TEMPLATE_GETONE =
`func (rep *%sRepository) GetOne(%s *model.%s) (model.%s, error) {
	var ret model.%s
	where, binds := db.BuildWhereClause(%s)
	query := %s + where

	err := rep.db.QueryRow(query, binds...).Scan(%s)

	return ret, err
}`

const TEMPLATE_INSERT =
`func (rep *%sRepository) Insert(%s *model.%s, tx *sql.Tx) error {
	cmd := %s
	binds := []interface{}{%s}

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = rep.db.Exec(cmd, binds...)
    }
	
	return err
}`

const TEMPLATE_UPDATE =
`func (rep *%sRepository) Update(%s *model.%s, tx *sql.Tx) error {
	cmd := %s
	binds := []interface{}{%s}
	
	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = rep.db.Exec(cmd, binds...)
    }
	
	return err
}`

const TEMPLATE_DELETE =
`func (rep *%sRepository) Delete(%s *model.%s, tx *sql.Tx) error {
	where, binds := db.BuildWhereClause(%s)
	cmd := "DELETE FROM %s " + where

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = rep.db.Exec(cmd, binds...)
    }
	
	return err
}`

func generateRepositoryCode(table ddlparse.Table) string {
	tn := strings.ToLower(table.Name)
	tnc := snakeToCamel(tn)
	tnp := snakeToPascal(tn)
	tni := getSnakeInitial(tn)

	return fmt.Sprintf(
		TEMPLATE,
		tnp, tni, tnp, tnp, tni, tnp, tnp, tni, tnp, tni, tnp, tni, tnp,
		tnc, tnp, tnp, tnc,
		generateRepositoryGetCode(table),
		generateRepositoryGetOneCode(table),
		//generateRepositoryInserCode(table),
		//generateRepositoryUpdateCode(table),
		//generateRepositoryDeleteCode(table),
	)
}

func generateRepositoryGetCode(table ddlparse.Table) string {
	tn := strings.ToLower(table.Name)
	tnc := snakeToCamel(tn)
	tnp := snakeToPascal(tn)
	tni := getSnakeInitial(tn)

	query := "\n\t`SELECT\n"
	for i, c := range table.Columns {
		if i == 0 {
			query += fmt.Sprintf("\t\t%s", c.Name)
		} else {
			query += fmt.Sprintf("\n\t\t,%s", c.Name)
		}
	}
	query += fmt.Sprintf("\n\t FROM %s `", tn)

	scan := "\n"
	for _, c := range table.Columns {
		scan += fmt.Sprintf("\t\t\t&%s.%s,\n", tni, snakeToPascal(c.Name))
	}
	scan += "\t\t"

	return fmt.Sprintf(
		TEMPLATE_GET,
		tnc, tni, tnp, tnp, tni, 
		query,
		tnp, tnp, tni, tnp,
		scan,
		tnp, tni,
	) 
}


func generateRepositoryGetOneCode(table ddlparse.Table) string {
	tn := strings.ToLower(table.Name)
	tnc := snakeToCamel(tn)
	tnp := snakeToPascal(tn)
	tni := getSnakeInitial(tn)

	query := "\n\t`SELECT\n"
	for i, c := range table.Columns {
		if i == 0 {
			query += fmt.Sprintf("\t\t%s", c.Name)
		} else {
			query += fmt.Sprintf("\n\t\t,%s", c.Name)
		}
	}
	query += fmt.Sprintf("\n\t FROM %s `", tn)

	scan := "\n"
	for _, c := range table.Columns {
		scan += fmt.Sprintf("\t\t&ret.%s,\n", snakeToPascal(c.Name))
	}
	scan += "\t"

	return fmt.Sprintf(
		TEMPLATE_GETONE,
		tnc, tni, tnp, tnp, tnp, tni, 
		query,
		scan,
	) 
}