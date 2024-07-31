package service

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"

	"goat/internal/core/jwt"
	"goat/internal/core/logger"
	"goat/internal/core/errs"
	"goat/internal/model"
	"goat/internal/dto"
	"goat/internal/repository"
)


type UserService interface {
	Signup(name, password string) error
	Login(name, password string) (dto.User, error)
	GetProfile(id int) (dto.User, error)
	GenerateJwtPayload(id int) (jwt.Payload, error)
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


func (srv *userService) toUserDTO(user model.User) dto.User {
	return dto.User{
		Id:        user.Id,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}


func (srv *userService) Signup(name, password string) error {
	_, err := srv.userRepository.GetOne(&model.User{Name: name})
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


func (srv *userService) Login(name, password string) (dto.User, error) {
	user, err := srv.userRepository.GetOne(&model.User{Name: name})
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug(err.Error())
		} else {
			logger.Error(err.Error())
		}
		return dto.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logger.Error(err.Error())
	}

	return srv.toUserDTO(user), err
}


func (srv *userService) GetProfile(id int) (dto.User, error) {
	user, err := srv.userRepository.GetOne(&model.User{Id: id})
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug(err.Error())
		} else {
			logger.Error(err.Error())
		}
	}

	return srv.toUserDTO(user), err
}


func (srv *userService) GenerateJwtPayload(id int) (jwt.Payload, error) {
	user, err := srv.userRepository.GetOne(&model.User{Id: id})
	if err != nil {
		logger.Error(err.Error())
		return jwt.Payload{}, err
	}

	var cc jwt.CustomClaims
	cc.UserId = user.Id
	cc.UserName = user.Name
	return jwt.NewPayload(cc), nil
}


func (srv *userService) UpdateName(id int, name string) error {
	u, err := srv.userRepository.GetOne(&model.User{Name: name})
	if err == nil && u.Id != id{
		return errs.NewUniqueConstraintError("user_name")
	}

	user, err := srv.userRepository.GetOne(&model.User{Id: id})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	user.Name = name
	if err = srv.userRepository.Update(&user, nil); err != nil {
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

	user, err := srv.userRepository.GetOne(&model.User{Id: id})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	user.Password = string(hashed)
	if err = srv.userRepository.Update(&user, nil); err != nil {
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