package service

import (
	"golang.org/x/crypto/bcrypt"

	"goat-cg/internal/core/jwt"
	
	"goat-cg/internal/core/logger"
	"goat-cg/internal/model/entity"
	"goat-cg/internal/model/repository"
	"goat-cg/internal/model/queryservice"
)


type UserService interface {
	Signup(username, password string) int
	Login(username, password string) int
	GenerateJWT(userId int) string
}


type userService struct {
	ur repository.UserRepository
	uq queryservice.UserQueryService
}


func NewUserService() UserService {
	ur := repository.NewUserRepository()
	uq := queryservice.NewUserQueryService()
	return &userService{ur, uq}
}


// Signup() Return value
/*----------------------------------------*/
const SIGNUP_SUCCESS_INT = 0
const SIGNUP_CONFLICT_INT = 1
const SIGNUP_ERROR_INT = 2
/*----------------------------------------*/

func (us *userService) Signup(username, password string) int {
	_, err := us.uq.QueryUserByName(username)

	if err == nil {
		return SIGNUP_CONFLICT_INT
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.LogError(err.Error())
		return SIGNUP_ERROR_INT
	}

	var user entity.User
	user.Username = username
	user.Password = string(hashed)

	err = us.ur.Insert(&user)

	if err != nil {
		logger.LogError(err.Error())
		return SIGNUP_ERROR_INT
	}

	return SIGNUP_SUCCESS_INT
}


// Login() Return value
/*----------------------------------------*/
const LOGIN_FAILURE_INT = -1
// 正常時: ユーザ識別ID
/*----------------------------------------*/

func (us *userService) Login(username, password string) int {
	user, err := us.uq.QueryUserByName(username)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return LOGIN_FAILURE_INT
	}

	return user.UserId
}


// GenerateJWT() Return value
/*----------------------------------------*/
const GENERATE_JWT_FAILURE_STR = ""
// 正常時: jwt文字列
/*----------------------------------------*/

func (us *userService) GenerateJWT(userId int) string {
	user, err := us.uq.QueryUser(userId)
	
	if err != nil {
		logger.LogError(err.Error())
		return GENERATE_JWT_FAILURE_STR
	}

	var cc jwt.CustomClaims
	cc.UserId = user.UserId
	cc.Username = user.Username
	jwtStr, err := jwt.GenerateJWT(cc)

	if err != nil {
		logger.LogError(err.Error())
		return GENERATE_JWT_FAILURE_STR
	}

	return jwtStr
}