package services

import (
	models "webapp/models/db_models"
	"webapp/models/request_models"
)

type AccountServiceInterface interface {
	Login(request_models.LoginRequest) (string, error)
	CreateAccount(request request_models.RegisterRequest) (models.Account, error)
	GetAccountByEmail(email string) (models.Account, error)
	GetAccountByUserName(userName string) (models.Account, error)
	GetAccountByPhone(phone string) (models.Account, error)
	GetAccountById(id int) (models.Account, error)
	UpdateAccount(id uint, userName string, password string, email string, phone string) error
	DeleteAccount(id uint) error
	GetAllAccounts() ([]models.Account, error)
	GetRandomAccount() (models.Account, error)
	GetAllHomelessAccounts() ([]models.Account, error)
	UpdateAddress(email string, address request_models.AddressRequest) error
}
