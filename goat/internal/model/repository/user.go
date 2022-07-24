package repository

import (
	"fmt"
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model/entity"
)


type UserRepository interface {
	Select() ([]entity.User, error)
	SelectByUId(uid int) (entity.User, error)
	UpdateByUId(uid int, user *entity.User) error
	DeleteByUId(uid int) error

	Signup(username, password string) error
	SelectByUsername(username string) (entity.User, error)
	UpdatePasswordByUId(uid int, password string) error
	UpdateUsernameByUId(uid int, username string) error
}


type userRepository struct {
	db *sql.DB
}


func NewUserRepository() UserRepository {
	db := db.GetDB()
	return &userRepository{db}
}


func (ur *userRepository) Select() ([]entity.User, error){
	var users []entity.User

	rows, err := ur.db.Query(
		`SELECT 
			UID,
			USERNAME, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(
			&user.UId, 
			&user.Username, 
			&user.CreateAt, 
			&user.UpdateAt,
		)
		if err != nil {
			break
		}
		users = append(users, user)
	}

	return users, err
}


func (ur *userRepository) SelectByUId(uid int) (entity.User, error){
	var user entity.User
	err := ur.db.QueryRow(
		`SELECT 
			UID, 
			USERNAME, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS 
		 WHERE UID = $1`, uid,
	).Scan(
		&user.UId, 
		&user.Username, 
		&user.CreateAt, 
		&user.UpdateAt,
	)

	return user, err
}


func (ur *userRepository) SelectByUsername(username string) (entity.User, error){
	var user entity.User
	err := ur.db.QueryRow(
		`SELECT 
			UID, 
			USERNAME, 
			PASSWORD, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS 
		 WHERE USERNAME = $1`, 
		 username,
	).Scan(
		&user.UId, 
		&user.Username, 
		&user.Password, 
		&user.CreateAt, 
		&user.UpdateAt,
	)

	return user, err
}


func (ur *userRepository) UpdateByUId(uid int, user *entity.User) error {
	_, err := ur.db.Exec(
		`UPDATE USERS 
		 SET USERNAME = $1 
		 WHERE UID = $2`,
		user.Username, 
		uid,
	)
	return err
}


func (ur *userRepository) UpdatePasswordByUId(uid int, password string) error {
	_, err := ur.db.Exec(
		`UPDATE USERS 
		 SET PASSWORD = $1 
		 WHERE UID = $2`, 
		 password, 
		 uid,
	)
	return err
}


func (ur *userRepository) UpdateUsernameByUId(uid int, username string) error {
	_, err := ur.db.Exec(
		`UPDATE USERS 
		 SET USERNAME = $1 
		 WHERE UID = $2`, 
		 username, 
		 uid,
	)
	return err
}


func (ur *userRepository) DeleteByUId(uid int) error {
	_, err := ur.db.Exec(`DELETE FROM USERS WHERE UID = $1`, uid)

	return err
}


func (ur *userRepository) Signup(username, password string) error {
	_, err := ur.db.Exec(
		`INSERT INTO USERS (
			USERNAME, 
			PASSWORD
		 ) VALUES($1,$2)`,
		username, 
		password,
	)
	fmt.Println(err)
	return err
}