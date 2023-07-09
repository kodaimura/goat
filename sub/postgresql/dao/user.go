package dao

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model/entity"
)


type userDao struct {
	db *sql.DB
}


func NewUserDao() *userDao {
	db := db.GetDB()
	return &userDao{db}
}


func (rep *userDao) SelectAll() ([]entity.User, error) {
	var ret []entity.User

	rows, err := rep.db.Query(
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
		u := entity.User{}
		err = rows.Scan(
			&u.UserId, 
			&u.UserName,
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


func (rep *userDao) Select(u *entity.User) (entity.User, error){
	var ret entity.User

	err := rep.db.QueryRow(
		`SELECT 
			user_id, 
			user_name, 
			created_at, 
			updated_at 
		 FROM users 
		 WHERE user_id = $1`, 
		 u.UserId,
	).Scan(
		&ret.UserId, 
		&ret.UserName, 
		&ret.CreatedAt, 
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *userDao) Insert(u *entity.User) error {
	_, err := rep.db.Exec(
		`INSERT INTO users (
			user_name, 
			password
		 ) VALUES($1,$2)`,
		u.UserName, 
		u.Password,
	)
	return err
}


func (rep *userDao) Update(u *entity.User) error {
	_, err := rep.db.Exec(
		`UPDATE users 
		 SET user_name = $1 
			  password = $2
		 WHERE user_id = $3`,
		u.UserName,
		u.Password, 
		u.UserId,
	)
	return err
}


func (rep *userDao) Delete(u *entity.User) error {
	_, err := rep.db.Exec(
		`DELETE FROM users WHERE user_id = $1`, 
		u.UserId,
	)

	return err
}


func (rep *userDao) UpdatePassword(u *entity.User) error {
	_, err := rep.db.Exec(
		`UPDATE users 
		 SET password = $1 
		 WHERE user_id = $2`, 
		 u.Password, 
		 u.UserId,
	)
	return err
}


func (rep *userDao) UpdateName(u *entity.User) error {
	_, err := rep.db.Exec(
		`UPDATE users
		 SET user_name = $1 
		 WHERE user_id = $2`, 
		u.UserName, 
		u.UserId,
	)
	return err
}


func (rep *userDao) SelectByName(name string) (entity.User, error) {
	var ret entity.User

	err := rep.db.QueryRow(
		`SELECT 
			user_id, 
			user_name, 
			password, 
			created_at, 
			updated_at 
		 FROM users 
		 WHERE user_name = $1`, 
		 name,
	).Scan(
		&ret.UserId, 
		&ret.UserName, 
		&ret.Password, 
		&ret.CreatedAt, 
		&ret.UpdatedAt,
	)

	return ret, err
}
