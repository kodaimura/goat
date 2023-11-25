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


type UserService struct {
	userDao *dao.UserDao
}


func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}


func (us *UserService) Signup(username, password string) error {
	_, err := us.userDao.SelectByName(username)

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

	err = us.userDao.Insert(&user)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (us *UserService) Login(username, password string) (entity.User, error) {
	user, err := us.userDao.SelectByName(username)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return entity.User{}, err
	}

	return user, nil
}


func (us *UserService) GenerateJWT(id int) (string, error) {
	var user entity.User
	user.UserId = id
	user, err := us.userDao.Select(&user)
	
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


func (us *UserService) GetProfile(id int) (entity.User, error) {
	var user entity.User
	user.UserId = id
	user, err := us.userDao.Select(&user)

	if err != nil {
		logger.LogError(err.Error())
	}

	return user, err
}


func (us *UserService) ChangeUsername(id int, username string) error {
	var user entity.User
	user.UserId = id
	user.Username = username
	err := us.userDao.UpdateName(&user)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (us *UserService) ChangePassword(id int, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.LogError(err.Error())
		return err
	}

	var user entity.User
	user.UserId = id
	user.Password = string(hashed)
	err = us.userDao.UpdatePassword(&user)
	
	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (us *UserService) DeleteUser(id int) error {
	var user entity.User
	user.UserId = id
	err := us.userDao.Delete(&user)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}
