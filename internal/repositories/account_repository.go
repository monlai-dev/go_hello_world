package repositories

import (
	accountModels "webapp/internal/models/db_models"
)

type AccountRepositoryInterface interface {
	FindAccountByEmail(email string) (accountModels.Account, error)
	FindAccountByUserName(userName string) (accountModels.Account, error)
	FindAccountByPhone(phone string) (accountModels.Account, error)
	FindAccountById(id int) (accountModels.Account, error)
	CreateAccount(account accountModels.Account) (accountModels.Account, error)
	UpdateAccount(account accountModels.Account) error
	DeleteAccount(account accountModels.Account) error
	GetAllAccounts(page int, pageSize int) ([]accountModels.Account, error)
	GetRandomAccount() (accountModels.Account, error)
	GetAllHomelessAccounts() ([]accountModels.Account, error)
}
