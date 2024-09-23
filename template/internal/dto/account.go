package dto

import (
	"goat/internal/model"
)


func NewAccount(account model.Account) Account {
	return Account{
		Id:        account.Id,
		Name:      account.Name,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

type Account struct {
	Id int `json:"account_id"`
	Name string `json:"account_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}