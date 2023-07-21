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
		u := entity.User{}
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


func (rep *userDao) Select(u *entity.User) (entity.User, error){
	var ret entity.User

	err := rep.db.QueryRow(
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


func (rep *userDao) Insert(u *entity.User) error {
	_, err := rep.db.Exec(
		`INSERT INTO users (
			username, 
			password
		 ) VALUES(?,?)`,
		u.Username, 
		u.Password,
	)
	return err
}


func (rep *userDao) Update(u *entity.User) error {
	_, err := rep.db.Exec(
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


func (rep *userDao) Delete(u *entity.User) error {
	_, err := rep.db.Exec(
		`DELETE FROM users WHERE id = ?`, 
		u.UserId,
	)

	return err
}


func (rep *userDao) UpdatePassword(u *entity.User) error {
	_, err := rep.db.Exec(
		`UPDATE users 
		 SET password = ? 
		 WHERE id = ?`, 
		 u.Password, 
		 u.UserId,
	)
	return err
}


func (rep *userDao) UpdateName(u *entity.User) error {
	_, err := rep.db.Exec(
		`UPDATE users
		 SET username = ? 
		 WHERE id = ?`, 
		u.Username, 
		u.UserId,
	)
	return err
}


func (rep *userDao) SelectByName(name string) (entity.User, error) {
	var ret entity.User

	err := rep.db.QueryRow(
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
