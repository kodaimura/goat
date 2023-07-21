package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"goat/internal/core/jwt"
	"goat/internal/core/logger"
	"goat/internal/model/entity"
	"goat/internal/model/dao"
)

type SignupConflictError struct {}

func (e *SignupConflictError) Error() string {
	return fmt.Sprintf("SignupConflictError")
}


type UserDao interface {
	SelectAll() ([]entity.User, error)
	Select(u *entity.User) (entity.User, error)
	Insert(u *entity.User) error
	Update(u *entity.User) error
	Delete(u *entity.User) error
	UpdatePassword(u *entity.User) error
	UpdateName(u *entity.User) error
	SelectByName(name string) (entity.User, error)
}


type userService struct {
	uDao UserDao
}


func NewUserService() *userService {
	uDao := dao.NewUserDao()
	return &userService{uDao}
}


func (serv *userService) Signup(username, password string) error {
	_, err := serv.uDao.SelectByName(username)

	if err == nil {
		return &SignupConflictError{}
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.LogError(err.Error())
		return err
	}

	var user entity.User
	user.Username = username
	user.Password = string(hashed)

	err = serv.uDao.Insert(&user)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *userService) Login(username, password string) (entity.User, error) {
	user, err := serv.uDao.SelectByName(username)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return entity.User{}, err
	}

	return user, nil
}


func (serv *userService) GenerateJWT(id int) (string, error) {
	var user entity.User
	user.UserId = id
	user, err := serv.uDao.Select(&user)
	
	if err != nil {
		logger.LogError(err.Error())
		return "", err
	}

	var cc jwt.CustomClaims
	cc.UserId = user.UserId
	cc.Username = user.Username
	jwtStr, err := jwt.GenerateJWT(cc)

	if err != nil {
		logger.LogError(err.Error())
		return "", err
	}

	return jwtStr, nil
}


func (serv *userService) GetProfile(id int) (entity.User, error) {
	var user entity.User
	user.UserId = id
	user, err := serv.uDao.Select(&user)

	if err != nil {
		logger.LogError(err.Error())
	}

	return user, err
}


func (serv *userService) ChangeUsername(id int, username string) error {
	var user entity.User
	user.UserId = id
	user.Username = username
	err := serv.uDao.UpdateName(&user)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *userService) ChangePassword(id int, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.LogError(err.Error())
		return err
	}

	var user entity.User
	user.UserId = id
	user.Password = string(hashed)
	err = serv.uDao.UpdatePassword(&user)
	
	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *userService) DeleteUser(id int) error {
	var user entity.User
	user.UserId = id
	err := serv.uDao.Delete(&user)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}
