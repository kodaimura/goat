package repository

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model/entity"
)


type UserRepository interface {
	Insert(u *entity.User) error
	Update(id int, u *entity.User) error
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


func (rep *userRepository) Insert(u *entity.User) error {
	_, err := rep.db.Exec(
		`INSERT INTO users (
			user_name, 
			password
		 ) VALUES(?,?)`,
		u.UserName, 
		u.Password,
	)
	return err
}


func (rep *userRepository) Update(id int, u *entity.User) error {
	_, err := rep.db.Exec(
		`UPDATE users 
		 SET user_name = ? 
		 WHERE user_id = ?`,
		u.UserName, 
		id,
	)
	return err
}


func (rep *userRepository) Delete(id int) error {
	_, err := rep.db.Exec(
		`DELETE FROM users WHERE user_id = ?`, 
		id,
	)

	return err
}


func (rep *userRepository) UpdatePassword(id int, password string) error {
	_, err := rep.db.Exec(
		`UPDATE users 
		 SET password = ? 
		 WHERE user_id = ?`, 
		 password, 
		 id,
	)
	return err
}


func (rep *userRepository) UpdateName(id int, name string) error {
	_, err := rep.db.Exec(
		`UPDATE users
		 SET user_name = ? 
		 WHERE user_id = ?`, 
		name, 
		id,
	)
	return err
}
