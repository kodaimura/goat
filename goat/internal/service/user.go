package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"goat/internal/core/jwt"
	"goat/internal/core/logger"
	"goat/internal/model"
	"goat/internal/repository"
)


type UserService interface {
	Signup(username, password string) error
	Login(username, password string) (model.User, error)
	GenerateJWT(id int) (string, error)
	GetProfile(id int) (model.User, error)
	ChangeUsername(id int, username string) error
	ChangePassword(id int, password string) error
	DeleteUser(id int) error
}


type userService struct {
	userRepository repository.UserRepository
}

func NewUserService() UserService {
	return &userService{
		userRepository: repository.NewUserRepository(),
	}
}


type SignupConflictError struct {}

func (e *SignupConflictError) Error() string {
	return fmt.Sprintf("SignupConflictError")
}

func (us *userService) Signup(username, password string) error {
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


func (us *userService) Login(username, password string) (model.User, error) {
	user, err := us.userRepository.GetByName(username)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return model.User{}, err
	}

	return user, nil
}


func (us *userService) GenerateJWT(id int) (string, error) {
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


func (us *userService) GetProfile(id int) (model.User, error) {
	user, err := us.userRepository.GetById(id)

	if err != nil {
		logger.Error(err.Error())
	}

	return user, err
}


func (us *userService) ChangeUsername(id int, username string) error {
	var user model.User
	user.UserId = id
	user.Username = username
	err := us.userRepository.UpdateName(&user)

	if err != nil {
		logger.Error(err.Error())
	}

	return err
}


func (us *userService) ChangePassword(id int, password string) error {
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


func (us *userService) DeleteUser(id int) error {
	var user model.User
	user.UserId = id
	err := us.userRepository.Delete(&user)

	if err != nil {
		logger.Error(err.Error())
	}

	return err
}
