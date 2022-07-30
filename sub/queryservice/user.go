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
			USERNAME, 
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
			&u.Username, 
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
			USERNAME, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS 
		 WHERE USER_ID = $1`, 
		 id,
	).Scan(
		&ret.UserId, 
		&ret.Username, 
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
			USERNAME, 
			PASSWORD, 
			CREATE_AT, 
			UPDATE_AT 
		 FROM USERS 
		 WHERE USERNAME = $1`, 
		 name,
	).Scan(
		&ret.UserId, 
		&ret.Username, 
		&ret.Password, 
		&ret.CreateAt, 
		&ret.UpdateAt,
	)

	return ret, err
}
