package repository

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model"
)


type AccountRepository interface {
	Get(u *model.Account) ([]model.Account, error)
	GetOne(u *model.Account) (model.Account, error)
	Insert(u *model.Account, tx *sql.Tx) error
	Update(u *model.Account, tx *sql.Tx) error
	Delete(u *model.Account, tx *sql.Tx) error
}


type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository() AccountRepository {
	db := db.GetDB()
	return &accountRepository{db}
}

func (rep *accountRepository) Get(u *model.Account) ([]model.Account, error) {
	where, binds := db.BuildWhereClause(u)
	query := 
	`SELECT
		account_id,
		account_name,
		account_password,
		created_at,
		updated_at
	 FROM account ` + where
	rows, err := rep.db.Query(query, binds...)
	defer rows.Close()

	if err != nil {
		return []model.Account{}, err
	}

	ret := []model.Account{}
	for rows.Next() {
		u := model.Account{}
		err = rows.Scan(
			&u.Id, 
			&u.Name,
			&u.Password,
			&u.CreatedAt, 
			&u.UpdatedAt,
		)
		if err != nil {
			return []model.Account{}, err
		}
		ret = append(ret, u)
	}

	return ret, nil
}


func (rep *accountRepository) GetOne(u *model.Account) (model.Account, error) {
	var ret model.Account
	where, binds := db.BuildWhereClause(u)
	query := 
	`SELECT
		account_id,
		account_name,
		account_password,
		created_at,
		updated_at
	 FROM account ` + where

	err := rep.db.QueryRow(query, binds...).Scan(
		&ret.Id, 
		&ret.Name, 
		&ret.Password,
		&ret.CreatedAt, 
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *accountRepository) Insert(u *model.Account, tx *sql.Tx) error {
	cmd := 
	`INSERT INTO account (
		account_name, 
		account_password
	 ) VALUES(?,?)`
	binds := []interface{}{u.Name, u.Password}

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = rep.db.Exec(cmd, binds...)
    }
	
	return err
}


func (rep *accountRepository) Update(u *model.Account, tx *sql.Tx) error {
	cmd := 
	`UPDATE account 
	 SET account_name = ?,
	 	 account_password = ?
	 WHERE account_id = ?`
	binds := []interface{}{u.Name, u.Password, u.Id}
	
	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = rep.db.Exec(cmd, binds...)
    }
	
	return err
}


func (rep *accountRepository) Delete(u *model.Account, tx *sql.Tx) error {
	where, binds := db.BuildWhereClause(u)
	cmd := "DELETE FROM account " + where

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = rep.db.Exec(cmd, binds...)
    }
	
	return err
}