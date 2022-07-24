package repository

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model/entity"
)


type UserRepository interface {
	Select() ([]entity.User, error)
	SelectByUserId(userId int) (entity.User, error)
	UpdateByUserId(userId int, user *entity.User) error
	DeleteByUserId(userId int) error

	Signup(username, password string) error
	SelectByUsername(username string) (entity.User, error)
	UpdatePasswordByUserId(userId int, password string) error
	UpdateUsernameByUserId(userId int, username string) error
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
			USER_ID,
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
			&user.UserId, 
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


func (ur *userRepository) SelectByUserId(userId int) (entity.User, error){
	var user entity.User
	err := ur.db.QueryRow(
		`SELECT 
			USER_ID, 
			USERNAME, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS 
		 WHERE UID = $1`, userId,
	).Scan(
		&user.UserId, 
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
			USER_ID, 
			USERNAME, 
			PASSWORD, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS 
		 WHERE USERNAME = $1`, 
		 username,
	).Scan(
		&user.UserId, 
		&user.Username, 
		&user.Password, 
		&user.CreateAt, 
		&user.UpdateAt,
	)

	return user, err
}


func (ur *userRepository) UpdateByUserId(userId int, user *entity.User) error {
	_, err := ur.db.Exec(
		`UPDATE USERS 
		 SET USERNAME = $1 
		 WHERE USER_ID = $2`,
		user.Username, 
		userId,
	)
	return err
}


func (ur *userRepository) UpdatePasswordByUserId(userId int, password string) error {
	_, err := ur.db.Exec(
		`UPDATE USERS 
		 SET PASSWORD = $1 
		 WHERE USER_ID = $2`, 
		 password, 
		 userId,
	)
	return err
}


func (ur *userRepository) UpdateUsernameByUserId(userId int, username string) error {
	_, err := ur.db.Exec(
		`UPDATE USERS 
		 SET USERNAME = $1 
		 WHERE USER_ID = $2`, 
		 username, 
		 userId,
	)
	return err
}


func (ur *userRepository) DeleteByUserId(userId int) error {
	_, err := ur.db.Exec(
		`DELETE FROM USERS WHERE USER_ID = $1`, userId,
	)

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
	return err
}