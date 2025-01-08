package services

import (
	"webapp/internal/models/db_models"
	request_models2 "webapp/internal/models/request_models"
)

type AccountServiceInterface interface {
	Login(request_models2.LoginRequest) (string, error)
	CreateAccount(request request_models2.RegisterRequest) (models.Account, error)
	GetAccountByEmail(email string) (models.Account, error)
	GetAccountByUserName(userName string) (models.Account, error)
	GetAccountByPhone(phone string) (models.Account, error)
	GetAccountById(id int) (models.Account, error)
	UpdateAccount(id uint, userName string, password string, email string, phone string) error
	DeleteAccount(id uint) error
	GetAllAccounts(page int, page_size int) ([]models.Account, error)
	GetRandomAccount() (models.Account, error)
	GetAllHomelessAccounts() ([]models.Account, error)
	UpdateAddress(email string, address request_models2.AddressRequest) error
	Logout(token string) error
}
