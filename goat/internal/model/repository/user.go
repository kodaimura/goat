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
		`INSERT INTO user (
			user_name, 
			password
		 ) VALUES(?,?)`,
		user.UserName, 
		user.Password,
	)
	return err
}


func (rep *userRepository) Update(id int, user *entity.User) error {
	_, err := rep.db.Exec(
		`UPDATE user 
		 SET user_name = ? 
		 WHERE user_id = ?`,
		user.UserName, 
		id,
	)
	return err
}


func (rep *userRepository) Delete(id int) error {
	_, err := rep.db.Exec(
		`DELETE FROM user WHERE user_id = ?`, 
		id,
	)

	return err
}


func (rep *userRepository) UpdatePassword(id int, password string) error {
	_, err := rep.db.Exec(
		`UPDATE user 
		 SET password = ? 
		 WHERE user_id = ?`, 
		 password, 
		 id,
	)
	return err
}


func (rep *userRepository) UpdateName(id int, name string) error {
	_, err := rep.db.Exec(
		`UPDATE user
		 SET user_name = ? 
		 WHERE user_id = ?`, 
		name, 
		id,
	)
	return err
}
