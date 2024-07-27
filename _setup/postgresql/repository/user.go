package repository

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model"
)


type UserRepository interface {
	Get(u *model.User) ([]model.User, error)
	GetOne(u *model.User) (model.User, error)
	Insert(u *model.User, tx *sql.Tx) error
	Update(u *model.User, tx *sql.Tx) error
	Delete(u *model.User, tx *sql.Tx) error

	UpdateName(u *model.User, tx *sql.Tx) error
	UpdatePassword(u *model.User, tx *sql.Tx) error
}


type userRepository struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	db := db.GetDB()
	return &userRepository{db}
}

func (ur *userRepository) Get(u *model.User) ([]model.User, error) {
	where, binds := db.BuildWhereClause(u)
	query := "SELECT * FROM users " + where
	rows, err := ur.db.Query(query, binds...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	ret := []model.User{}
	for rows.Next() {
		u := model.User{}
		err = rows.Scan(
			&u.Id, 
			&u.Name,
			&u.Password,
			&u.CreatedAt, 
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		ret = append(ret, u)
	}

	return ret, nil
}


func (ur *userRepository) GetOne(u *model.User) (model.User, error) {
	var ret model.User
	where, binds := db.BuildWhereClause(u)
	query := "SELECT * FROM users " + where

	err := ur.db.QueryRow(query, binds...).Scan(
		&ret.Id, 
		&ret.Name, 
		&ret.Password,
		&ret.CreatedAt, 
		&ret.UpdatedAt,
	)

	return ret, err
}


func (ur *userRepository) Insert(u *model.User, tx *sql.Tx) error {
	cmd := 
	`INSERT INTO users (
		user_name, 
		user_password
	 ) VALUES($1,$2)`
	binds := []interface{}{u.Name, u.Password}

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = ur.db.Exec(cmd, binds...)
    }
	
	return err
}


func (ur *userRepository) Update(u *model.User, tx *sql.Tx) error {
	cmd := 
	`UPDATE users 
	 SET user_name = $1,
	 user_password = $2
	 WHERE user_id = $3`
	binds := []interface{}{u.Name, u.Password, u.Id}
	
	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = ur.db.Exec(cmd, binds...)
    }
	
	return err
}


func (ur *userRepository) Delete(u *model.User, tx *sql.Tx) error {
	cmd := "DELETE FROM users WHERE user_id = $1"
	binds := []interface{}{u.Id}

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = ur.db.Exec(cmd, binds...)
    }
	
	return err
}


func (ur *userRepository) UpdateName(u *model.User, tx *sql.Tx) error {
	cmd := 
	`UPDATE users
	 SET user_name = $1 
	 WHERE user_id = $2`
	binds := []interface{}{u.Name, u.Id}

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = ur.db.Exec(cmd, binds...)
    }
	
	return err
}


func (ur *userRepository) UpdatePassword(u *model.User, tx *sql.Tx) error {
	cmd := 
	`UPDATE users 
	 SET user_password = $1 
	 WHERE user_id = $2`
	binds := []interface{}{u.Password, u.Id}

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = ur.db.Exec(cmd, binds...)
    }

	return err
}