package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"goat/internal/core/jwt"
	"goat/internal/core/logger"
	"goat/internal/model"
	"goat/internal/repository"
)

type SignupConflictError struct {}

func (e *SignupConflictError) Error() string {
	return fmt.Sprintf("SignupConflictError")
}


type UserService struct {
	userRepository *dao.UserRepository
}


func NewUserService() *UserService {
	return &UserService{
		userRepository: dao.NewUserRepository(),
	}
}


func (us *UserService) Signup(username, password string) error {
	_, err := us.userRepository.GetByName(username)

	if err == nil {
		return &SignupConflictError{}
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var user model.User
	user.Username = username
	user.Password = string(hashed)

	err = us.userRepository.Insert(&user)

	if err != nil {
		logger.Error(err.Error())
	}

	return err
}


func (us *UserService) Login(username, password string) (model.User, error) {
	user, err := us.userRepository.GetByName(username)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return model.User{}, err
	}

	return user, nil
}


func (us *UserService) GenerateJWT(id int) (string, error) {
	user, err := us.userRepository.GetById(id)
	
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	var cc jwt.CustomClaims
	cc.UserId = user.UserId
	cc.Username = user.Username
	jwtStr, err := jwt.GenerateJWT(cc)

	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return jwtStr, nil
}


func (us *UserService) GetProfile(id int) (model.User, error) {
	user, err := us.userRepository.GetById(id)

	if err != nil {
		logger.Error(err.Error())
	}

	return user, err
}


func (us *UserService) ChangeUsername(id int, username string) error {
	var user model.User
	user.UserId = id
	user.Username = username
	err := us.userRepository.UpdateName(&user)

	if err != nil {
		logger.Error(err.Error())
	}

	return err
}


func (us *UserService) ChangePassword(id int, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var user model.User
	user.UserId = id
	user.Password = string(hashed)
	err = us.userRepository.UpdatePassword(&user)
	
	if err != nil {
		logger.Error(err.Error())
	}

	return err
}


func (us *UserService) DeleteUser(id int) error {
	var user model.User
	user.UserId = id
	err := us.userRepository.Delete(&user)

	if err != nil {
		logger.Error(err.Error())
	}

	return err
}
