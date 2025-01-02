package repository

import (
	"gorm.io/gorm"
	models "webapp/models/db_models"
)

type AccountRepo interface {
	CreateAccount(userName string, password string, email string, phone string) error
	DeleteAccount(id uint) error
	GetAccountByEmail(email string) (string, error)
	GetAccountById(id uint) (string, error)
	GetAccountByPhone(phone string) (string, error)
	GetAccountByUserName(userName string) (string, error)
	GetAllAccounts() (string, error)
	UpdateAccount(id uint, userName string, password string, email string, phone string) error
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) AccountRepo {
	return &accountRepo{db: db}
}

// DeleteAccount implements AccountRepo.
func (repo *accountRepo) DeleteAccount(id uint) error {
	return repo.db.Delete(&models.Account{}, id).Error
}

// GetAccountByEmail implements AccountRepo.
func (repo *accountRepo) GetAccountByEmail(email string) (string, error) {

	var account models.Account
	result := repo.db.Where("email = ?", email).Find(&account)
	if result.Error != nil {
		return "", result.Error
	}

	return account.Email, nil
}

// GetAccountById implements AccountRepo.
func (repo *accountRepo) GetAccountById(id uint) (string, error) {
	var account models.Account
	result := repo.db.Where("id = ?", id).Find(&account)
	if result.Error != nil {
		return "", result.Error
	}

	return account.Email, nil
}

// GetAccountByPhone implements AccountRepo.
func (repo *accountRepo) GetAccountByPhone(phone string) (string, error) {
	var account models.Account
	result := repo.db.Where("phone = ?", phone).Find(&account)
	if result.Error != nil {
		return "", result.Error
	}

	return account.Email, nil
}

// GetAccountByUserName implements AccountRepo.
func (repo *accountRepo) GetAccountByUserName(userName string) (string, error) {
	var account models.Account
	result := repo.db.Where("userName = ?", userName).Find(&account)
	if result.Error != nil {
		return "", result.Error
	}

	return account.Email, nil
}

// GetAllAccounts implements AccountRepo.
func (repo *accountRepo) GetAllAccounts() (string, error) {

	var accounts []models.Account
	result := repo.db.Find(&accounts)
	if result.Error != nil {
		return "", result.Error
	}

	return "", nil
}

// UpdateAccount implements AccountRepo.
func (repo *accountRepo) UpdateAccount(id uint, userName string, password string, email string, phone string) error {
	panic("unimplemented")
}

// CreateAccount implements AccountRepo.
func (repo *accountRepo) CreateAccount(userName string, password string, email string, phone string) error {
	panic("unimplemented")
}
