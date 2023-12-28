package repository

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model"
)


type UserRepository interface {
	Insert(u *model.User) error
	Get() ([]model.User, error)
	GetById(id int) (model.User, error)
	GetByName(name string) (model.User, error)
	Update(u *model.User) error
	UpdateName(u *model.User) error
	UpdatePassword(u *model.User) error
	Delete(u *model.User) error
}


type userRepository struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	db := db.GetDB()
	return &userRepository{db}
}


func (ur *userRepository) Insert(u *model.User) error {
	_, err := ur.db.Exec(
		`INSERT INTO users (
			user_name, 
			password
		 ) VALUES($1,$2)`,
		u.Username, 
		u.Password,
	)
	return err
}


func (ur *userRepository) Get() ([]model.User, error) {
	var ret []model.User

	rows, err := ur.db.Query(
		`SELECT 
			user_id, 
			user_name, 
			created_at, 
			updated_at 
		 FROM users`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := model.User{}
		err = rows.Scan(
			&u.UserId, 
			&u.Username,
			&u.CreatedAt, 
			&u.UpdatedAt,
		)
		if err != nil {
			break
		}
		ret = append(ret, u)
	}

	return ret, err
}


func (ur *userRepository) GetById(id int) (model.User, error) {
	var ret model.User

	err := ur.db.QueryRow(
		`SELECT 
			user_id, 
			user_name, 
			created_at, 
			updated_at 
		 FROM users 
		 WHERE user_id = $1`, 
		 id,
	).Scan(
		&ret.UserId, 
		&ret.Username, 
		&ret.CreatedAt, 
		&ret.UpdatedAt,
	)

	return ret, err
}


func (ur *userRepository) GetByName(name string) (model.User, error) {
	var ret model.User

	err := ur.db.QueryRow(
		`SELECT 
			user_id, 
			user_name, 
			user_password, 
			created_at, 
			updated_at 
		 FROM users 
		 WHERE user_name = $1`, 
		 name,
	).Scan(
		&ret.UserId, 
		&ret.Username, 
		&ret.Password, 
		&ret.CreatedAt, 
		&ret.UpdatedAt,
	)

	return ret, err
}


func (ur *userRepository) Update(u *model.User) error {
	_, err := ur.db.Exec(
		`UPDATE users 
		 SET user_name = $1,
		 	user_password = $2
		 WHERE user_id = $3`,
		u.Username,
		u.Password, 
		u.UserId,
	)
	return err
}


func (ur *userRepository) UpdateName(u *model.User) error {
	_, err := ur.db.Exec(
		`UPDATE users
		 SET user_name = $1 
		 WHERE user_id = $2`, 
		u.Username, 
		u.UserId,
	)
	return err
}


func (ur *userRepository) UpdatePassword(u *model.User) error {
	_, err := ur.db.Exec(
		`UPDATE users 
		 SET user_password = $1 
		 WHERE user_id = $2`, 
		 u.Password, 
		 u.UserId,
	)
	return err
}


func (ur *userRepository) Delete(u *model.User) error {
	_, err := ur.db.Exec(
		`DELETE FROM users WHERE user_id = $1`, 
		u.UserId,
	)

	return err
}