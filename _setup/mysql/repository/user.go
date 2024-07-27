package repository

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model"
)


type UserRepository interface {
	Get(u *model.User) ([]model.User, error)
	GetOne(u *model.User) (model.User, error)
	Insert(u *model.User) (int, error)
	Update(u *model.User) error
	Delete(u *model.User) error

	UpdateName(u *model.User) error
	UpdatePassword(u *model.User) error
}


type userRepository struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	db := db.GetDB()
	return &userRepository{db}
}


func (ur *userRepository) Get(u *model.User) ([]model.User, error) {
	where, args := db.BuildWhereClause(u)
	query := "SELECT * FROM users " + where
	rows, err := ur.db.Query(query, args...)
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
	where, args := db.BuildWhereClause(u)
	query := "SELECT * FROM users " + where

	err := ur.db.QueryRow(query, args...).Scan(
		&ret.Id, 
		&ret.Name, 
		&ret.Password,
		&ret.CreatedAt, 
		&ret.UpdatedAt,
	)

	return ret, err
}


func (ur *userRepository) Insert(u *model.User) (int, error) {
	var userId int
	_, err := ur.db.Exec(
		`INSERT INTO users (
			user_name, 
			user_password
		 ) VALUES(?,?)`,
		u.Name, 
		u.Password,
	)
	if err != nil {
		return userId, err
	}

	err = ur.db.QueryRow(
		`SELECT LAST_INSERT_ID() AS user_id`,
	).Scan(
		&userId,
	)
	return userId, err
}


func (ur *userRepository) Update(u *model.User) error {
	_, err := ur.db.Exec(
		`UPDATE users 
		 SET user_name = ?,
		 	user_password = ?
		 WHERE user_id = ?`,
		u.Name,
		u.Password, 
		u.Id,
	)
	return err
}


func (ur *userRepository) Delete(u *model.User) error {
	_, err := ur.db.Exec(
		`DELETE FROM users WHERE user_id = ?`, 
		u.Id,
	)

	return err
}


func (ur *userRepository) UpdateName(u *model.User) error {
	_, err := ur.db.Exec(
		`UPDATE users
		 SET user_name = ? 
		 WHERE user_id = ?`, 
		u.Name, 
		u.Id,
	)
	return err
}


func (ur *userRepository) UpdatePassword(u *model.User) error {
	_, err := ur.db.Exec(
		`UPDATE users 
		 SET user_password = ? 
		 WHERE user_id = ?`, 
		 u.Password, 
		 u.Id,
	)
	return err
}