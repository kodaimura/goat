package repository

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model/entity"
)


type UserRepository interface {
	Insert(user *entity.User) error
	Update(id int, user *entity.User) error
	Delete(id int) error

	UpdatePassword(id int, password string) error
	UpdateName(id int, name string) error
}


type userRepository struct {
	db *sql.DB
}


func NewUserRepository() UserRepository {
	db := db.GetDB()
	return &userRepository{db}
}


func (rep *userRepository) Insert(user *entity.User) error {
	_, err := rep.db.Exec(
		`INSERT INTO USERS (
			USER_NAME, 
			PASSWORD
		 ) VALUES($1,$2)`,
		user.UserName, 
		user.Password,
	)
	return err
}


func (rep *userRepository) Update(id int, user *entity.User) error {
	_, err := rep.db.Exec(
		`UPDATE USERS 
		 SET USER_NAME = $1 
		 WHERE USER_ID = $2`,
		user.UserName, 
		id,
	)
	return err
}


func (rep *userRepository) Delete(id int) error {
	_, err := rep.db.Exec(
		`DELETE FROM USERS WHERE USER_ID = $1`, 
		id,
	)

	return err
}


func (rep *userRepository) UpdatePassword(id int, password string) error {
	_, err := rep.db.Exec(
		`UPDATE USERS 
		 SET PASSWORD = $1 
		 WHERE USER_ID = $2`, 
		 password, 
		 id,
	)
	return err
}


func (rep *userRepository) UpdateName(id int, name string) error {
	_, err := rep.db.Exec(
		`UPDATE USERS 
		 SET USER_NAME = $1 
		 WHERE USER_ID = $2`, 
		name, 
		id,
	)
	return err
}
