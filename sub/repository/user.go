package repository

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model/entity"
)


type UserRepository interface {
	Insert(user *entity.User) error
	UpdateUser(id int, user *entity.User) error
	DeleteUser(id int) error

	UpdatePassword(id int, password string) error
	UpdateUsername(id int, username string) error
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
			USERNAME, 
			PASSWORD
		 ) VALUES($1,$2)`,
		user.Username, 
		user.Password,
	)
	return err
}


func (rep *userRepository) UpdateUser(id int, user *entity.User) error {
	_, err := rep.db.Exec(
		`UPDATE USERS 
		 SET USERNAME = $1 
		 WHERE USER_ID = $2`,
		user.Username, 
		id,
	)
	return err
}


func (rep *userRepository) DeleteUser(id int) error {
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


func (rep *userRepository) UpdateUsername(id int, username string) error {
	_, err := rep.db.Exec(
		`UPDATE USERS 
		 SET USERNAME = $1 
		 WHERE USER_ID = $2`, 
		username, 
		id,
	)
	return err
}
