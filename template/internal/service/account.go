package service

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"

	"goat/internal/core/jwt"
	"goat/internal/core/logger"
	"goat/internal/core/errs"
	"goat/internal/core/utils"
	"goat/internal/dto"
	"goat/internal/model"
	"goat/internal/repository"
)


type AccountService interface {
	GetOne(id int) (dto.Account, error)
	Delete(id int) error
	UpdateName(id int, name string) error
	UpdatePassword(id int, password string) error

	Login(input dto.Login) (dto.Account, error)
	Signup(input dto.Signup) (int, error)
	GenerateJwtPayload(id int) (jwt.Payload, error)
}


type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService() AccountService {
	return &accountService{
		accountRepository: repository.NewAccountRepository(),
	}
}


func (srv *accountService) GetOne(id int) (dto.Account, error) {
	account, err := srv.accountRepository.GetOne(&model.Account{Id: id})
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug(err.Error())
		} else {
			logger.Error(err.Error())
		}
	}

	var ret dto.Account
	utils.MapFields(&ret, account)
	return ret, err
}


func (srv *accountService) UpdateName(id int, name string) error {
	u, err := srv.accountRepository.GetOne(&model.Account{Name: name})
	if err == nil && u.Id != id{
		return errs.NewUniqueConstraintError("account_name")
	}

	account, err := srv.accountRepository.GetOne(&model.Account{Id: id})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	account.Name = name
	if err = srv.accountRepository.Update(&account, nil); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}


func (srv *accountService) UpdatePassword(id int, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	account, err := srv.accountRepository.GetOne(&model.Account{Id: id})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	account.Password = string(hashed)
	if err = srv.accountRepository.Update(&account, nil); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}


func (srv *accountService) Delete(id int) error {
	if err := srv.accountRepository.Delete(&model.Account{Id: id}, nil); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}


func (srv *accountService) Login(input dto.Login) (dto.Account, error) {
	account, err := srv.accountRepository.GetOne(&model.Account{Name: input.Name})
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug(err.Error())
		} else {
			logger.Error(err.Error())
		}
		return dto.Account{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(input.Password))
	if err != nil {
		logger.Error(err.Error())
	}

	var ret dto.Account
	utils.MapFields(&ret, account)
	return ret, err
}


func (srv *accountService) Signup(input dto.Signup) (int, error) {
	_, err := srv.accountRepository.GetOne(&model.Account{Name: input.Name})
	if err == nil {
		return 0, errs.NewUniqueConstraintError("account_name")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	var account model.Account
	account.Name = input.Name
	account.Password = string(hashed)

	accountId, err := srv.accountRepository.Insert(&account, nil);
	if err != nil {
		logger.Error(err.Error())
	}

	return accountId, err
}


func (srv *accountService) GenerateJwtPayload(id int) (jwt.Payload, error) {
	account, err := srv.accountRepository.GetOne(&model.Account{Id: id})
	if err != nil {
		logger.Error(err.Error())
		return jwt.Payload{}, err
	}

	var cc jwt.CustomClaims
	cc.AccountId = account.Id
	cc.AccountName = account.Name
	return jwt.NewPayload(cc), nil
}