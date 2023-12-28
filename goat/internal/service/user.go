package service

import (
	"golang.org/x/crypto/bcrypt"

	"goat/internal/core/jwt"
	"goat/internal/core/logger"
	"goat/internal/core/errs"
	"goat/internal/model"
	"goat/internal/repository"
)


type UserService interface {
	Signup(username, password string) error
	Login(username, password string) (model.User, error)
	GenerateJWT(id int) (string, error)
	GetProfile(id int) (model.User, error)
	UpdateUsername(id int, username string) error
	UpdatePassword(id int, password string) error
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


func (us *userService) Signup(username, password string) error {
	_, err := us.userRepository.GetByName(username)

	if err == nil {
		return errs.NewUniqueConstraintError("username")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var user model.User
	user.Username = username
	user.Password = string(hashed)

	if err = us.userRepository.Insert(&user); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}


func (us *userService) Login(username, password string) (model.User, error) {
	user, err := us.userRepository.GetByName(username)
	if err != nil {
		return model.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
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
	cc.UserId = user.Id
	cc.UserName = user.Username
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


func (us *userService) UpdateUsername(id int, username string) error {
	u, err := us.userRepository.GetByName(username)

	if err == nil && u.Id != id{
		return errs.NewUniqueConstraintError("username")
	}

	var user model.User
	user.Id = id
	user.Username = username

	if err = us.userRepository.UpdateName(&user); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}


func (us *userService) UpdatePassword(id int, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var user model.User
	user.Id = id
	user.Password = string(hashed)
	
	if err = us.userRepository.UpdatePassword(&user); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}


func (us *userService) DeleteUser(id int) error {
	var user model.User
	user.Id = id

	if err := us.userRepository.Delete(&user); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
