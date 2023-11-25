package dao

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model"
)


type UserRepository struct {
	db *sql.DB
}


func NewUserRepository() *UserRepository {
	db := db.GetDB()
	return &UserRepository{db}
}


func (ur *UserRepository) SelectAll() ([]model.User, error) {
	var ret []model.User

	rows, err := ur.db.Query(
		`SELECT 
			id, 
			username, 
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


func (ur *UserRepository) Select(u *model.User) (model.User, error){
	var ret model.User

	err := ur.db.QueryRow(
		`SELECT 
			id, 
			username, 
			created_at, 
			updated_at 
		 FROM users 
		 WHERE id = ?`, 
		 u.UserId,
	).Scan(
		&ret.UserId, 
		&ret.Username, 
		&ret.CreatedAt, 
		&ret.UpdatedAt,
	)

	return ret, err
}


func (ur *UserRepository) Insert(u *model.User) error {
	_, err := ur.db.Exec(
		`INSERT INTO users (
			username, 
			password
		 ) VALUES(?,?)`,
		u.Username, 
		u.Password,
	)
	return err
}


func (ur *UserRepository) Update(u *model.User) error {
	_, err := ur.db.Exec(
		`UPDATE users 
		 SET username = ? 
			  password = ?
		 WHERE id = ?`,
		u.Username,
		u.Password, 
		u.UserId,
	)
	return err
}


func (ur *UserRepository) Delete(u *model.User) error {
	_, err := ur.db.Exec(
		`DELETE FROM users WHERE id = ?`, 
		u.UserId,
	)

	return err
}


func (ur *UserRepository) UpdatePassword(u *model.User) error {
	_, err := ur.db.Exec(
		`UPDATE users 
		 SET password = ? 
		 WHERE id = ?`, 
		 u.Password, 
		 u.UserId,
	)
	return err
}


func (ur *UserRepository) UpdateName(u *model.User) error {
	_, err := ur.db.Exec(
		`UPDATE users
		 SET username = ? 
		 WHERE id = ?`, 
		u.Username, 
		u.UserId,
	)
	return err
}


func (ur *UserRepository) SelectByName(name string) (model.User, error) {
	var ret model.User

	err := ur.db.QueryRow(
		`SELECT 
			id, 
			username, 
			password, 
			created_at, 
			updated_at 
		 FROM users 
		 WHERE username = ?`, 
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
