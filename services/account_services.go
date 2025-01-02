package services

import ()

type AccountService interface {
	Login(email string, password string) (string, error)
	CreateAccount(userName string, password string, email string, phone string) error
	GetAccountByEmail(email string) (string, error)
	GetAccountByUserName(userName string) (string, error)
	GetAccountByPhone(phone string) (string, error)
	GetAccountById(id uint) (string, error)
	UpdateAccount(id uint, userName string, password string, email string, phone string) error
	DeleteAccount(id uint) error
	GetAllAccounts() (string, error)
}



type accountService struct {
}

// CreateAccount implements AccountService.
func (service *accountService) CreateAccount(userName string, password string, email string, phone string) error {
	
}

// DeleteAccount implements AccountService.
func (service *accountService) DeleteAccount(id uint) error {
	panic("unimplemented")
}

// GetAccountByEmail implements AccountService.
func (service *accountService) GetAccountByEmail(email string) (string, error) {
	panic("unimplemented")
}

// GetAccountById implements AccountService.
func (service *accountService) GetAccountById(id uint) (string, error) {
	panic("unimplemented")
}

// GetAccountByPhone implements AccountService.
func (service *accountService) GetAccountByPhone(phone string) (string, error) {
	panic("unimplemented")
}

// GetAccountByUserName implements AccountService.
func (service *accountService) GetAccountByUserName(userName string) (string, error) {
	panic("unimplemented")
}

// GetAllAccounts implements AccountService.
func (service *accountService) GetAllAccounts() (string, error) {
	panic("unimplemented")
}

// UpdateAccount implements AccountService.
func (service *accountService) UpdateAccount(id uint, userName string, password string, email string, phone string) error {
	panic("unimplemented")
}

func NewAccountService() AccountService {
	return &accountService{}
}

func (service *accountService) Login(email string, password string) (string, error) {
	return "", nil
}
