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


func (srv *userService) Signup(name, password string) error {
	var u model.User
	u.Name = name

	_, err := srv.userRepository.GetOne(&u)
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

	err = srv.userRepository.Insert(&user, nil);
	if err != nil {
		logger.Error(err.Error())
	}

	return err
}


func (srv *userService) Login(name, password string) (model.User, error) {
	var u model.User
	u.Name = name

	user, err := srv.userRepository.GetOne(&u)
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


func (srv *userService) GenerateJwtPayload(id int) (jwt.Payload, error) {
	var u model.User
	u.Id = id

	user, err := srv.userRepository.GetOne(&u)
	if err != nil {
		logger.Error(err.Error())
		return jwt.Payload{}, err
	}

	var cc jwt.CustomClaims
	cc.UserId = user.Id
	cc.UserName = user.Name
	return jwt.NewPayload(cc), nil
}


func (srv *userService) GetProfile(id int) (model.User, error) {
	var u model.User
	u.Id = id

	user, err := srv.userRepository.GetOne(&u)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug(err.Error())
		} else {
			logger.Error(err.Error())
		}
	}

	return user, err
}


func (srv *userService) UpdateName(id int, name string) error {
	var u model.User
	u.Name = name

	u, err := srv.userRepository.GetOne(&u)
	if err == nil && u.Id != id{
		return errs.NewUniqueConstraintError("user_name")
	}

	var user model.User
	user.Id = id
	user.Name = name

	if err = srv.userRepository.UpdateName(&user, nil); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}


func (srv *userService) UpdatePassword(id int, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var user model.User
	user.Id = id
	user.Password = string(hashed)
	
	if err = srv.userRepository.UpdatePassword(&user, nil); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}


func (srv *userService) DeleteUser(id int) error {
	var user model.User
	user.Id = id

	if err := srv.userRepository.Delete(&user, nil); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
