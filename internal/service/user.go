package service

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"

	"goat/internal/core/jwt"
	"goat/internal/core/logger"
	"goat/internal/core/errs"
	"goat/internal/model"
	"goat/internal/repository"
)


type UserService interface {
	Signup(name, password string) error
	Login(name, password string) (model.User, error)
	GenerateJwtPayload(id int) (jwt.Payload, error)
	GetProfile(id int) (model.User, error)
	UpdateName(id int, name string) error
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


func (us *userService) Signup(name, password string) error {
	_, err := us.userRepository.GetByName(name)

	if err == nil {
		return errs.NewUniqueConstraintError("user_name")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var user model.User
	user.Name = name
	user.Password = string(hashed)

	_, err = us.userRepository.Insert(&user);
	if err != nil {
		logger.Error(err.Error())
	}

	return err
}


func (us *userService) Login(name, password string) (model.User, error) {
	user, err := us.userRepository.GetByName(name)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug(err.Error())
		} else {
			logger.Error(err.Error())
		}
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logger.Error(err.Error())
	}

	return user, err
}


func (us *userService) GenerateJwtPayload(id int) (jwt.Payload, error) {
	user, err := us.userRepository.GetById(id)
	
	if err != nil {
		logger.Error(err.Error())
		return jwt.Payload{}, err
	}

	var cc jwt.CustomClaims
	cc.UserId = user.Id
	cc.UserName = user.Name
	return jwt.NewPayload(cc), nil
}


func (us *userService) GetProfile(id int) (model.User, error) {
	user, err := us.userRepository.GetById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug(err.Error())
		} else {
			logger.Error(err.Error())
		}
	}

	return user, err
}


func (us *userService) UpdateName(id int, name string) error {
	u, err := us.userRepository.GetByName(name)

	if err == nil && u.Id != id{
		return errs.NewUniqueConstraintError("user_name")
	}

	var user model.User
	user.Id = id
	user.Name = name

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
