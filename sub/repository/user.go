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


func (rep *userRepository) Select() ([]entity.User, error){
	var users []entity.User

	rows, err := rep.db.Query(
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


func (rep *userRepository) SelectByUserId(userId int) (entity.User, error){
	var user entity.User
	err := rep.db.QueryRow(
		`SELECT 
			USER_ID, 
			USERNAME, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS 
		 WHERE USER_ID = ?`, userId,
	).Scan(
		&user.UserId, 
		&user.Username, 
		&user.CreateAt, 
		&user.UpdateAt,
	)

	return user, err
}


func (rep *userRepository) SelectByUsername(username string) (entity.User, error){
	var user entity.User
	err := rep.db.QueryRow(
		`SELECT 
			USER_ID, 
			USERNAME, 
			PASSWORD, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS 
		 WHERE USERNAME = ?`, 
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


func (rep *userRepository) UpdateByUserId(userId int, user *entity.User) error {
	_, err := rep.db.Exec(
		`UPDATE USERS 
		 SET USERNAME = ? 
		 WHERE USER_ID = ?`,
		user.Username, 
		userId,
	)
	return err
}


func (rep *userRepository) UpdatePasswordByUserId(userId int, password string) error {
	_, err := rep.db.Exec(
		`UPDATE USERS 
		 SET PASSWORD = ? 
		 WHERE USER_ID = ?`, 
		 password, 
		 userId,
	)
	return err
}


func (rep *userRepository) UpdateUsernameByUserId(userId int, username string) error {
	_, err := rep.db.Exec(
		`UPDATE USERS 
		 SET USERNAME = ? 
		 WHERE USER_ID = ?`, 
		 username, 
		 userId,
	)
	return err
}


func (rep *userRepository) DeleteByUserId(userId int) error {
	_, err := rep.db.Exec(`DELETE FROM USERS WHERE USER_ID = ?`, userId)

	return err
}


func (rep *userRepository) Signup(username, password string) error {
	_, err := rep.db.Exec(
		`INSERT INTO USERS (
			USERNAME, 
			PASSWORD
		 ) VALUES(?,?)`,
		username, 
		password,
	)
	return err
}