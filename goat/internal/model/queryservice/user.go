package queryservice

import (
	"database/sql"

	"goat/internal/core/db"
	"goat/internal/model/entity"
)


type UserQueryService interface {
	QueryUsers() ([]entity.User, error)
	QueryUser(id int) (entity.User, error)
	QueryUserByName(name string) (entity.User, error)
}

type userQueryService struct {
	db *sql.DB
}

func NewUserQueryService() UserQueryService {
	db := db.GetDB()
	return &userQueryService{db}
}


func (qs *userQueryService) QueryUsers() ([]entity.User, error){
	var ret []entity.User

	rows, err := qs.db.Query(
		`SELECT 
			USER_ID,
			USER_NAME, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := entity.User{}
		err = rows.Scan(
			&u.UserId, 
			&u.UserName, 
			&u.CreateAt, 
			&u.UpdateAt,
		)
		if err != nil {
			break
		}
		ret = append(ret, u)
	}

	return ret, err
}


func (qs *userQueryService) QueryUser(id int) (entity.User, error) {
	var ret entity.User

	err := qs.db.QueryRow(
		`SELECT 
			USER_ID, 
			USER_NAME, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS 
		 WHERE USER_ID = ?`, 
		 id,
	).Scan(
		&ret.UserId, 
		&ret.UserName, 
		&ret.CreateAt, 
		&ret.UpdateAt,
	)

	return ret, err
}


func (qs *userQueryService) QueryUserByName(name string) (entity.User, error) {
	var ret entity.User

	err := qs.db.QueryRow(
		`SELECT 
			USER_ID, 
			USER_NAME, 
			PASSWORD, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS 
		 WHERE USER_NAME = ?`, 
		 name,
	).Scan(
		&ret.UserId, 
		&ret.UserName, 
		&ret.Password, 
		&ret.CreateAt, 
		&ret.UpdateAt,
	)

	return ret, err
}
