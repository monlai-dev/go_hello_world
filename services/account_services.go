package services

import (
	"fmt"
	models "webapp/models/db_models"
	request_models "webapp/models/request_models"
	utils "webapp/utils"

	"gorm.io/gorm"
)

type AccountService interface {
	Login(request_models.LoginRequest) (string, error)
	CreateAccount(request request_models.RegisterRequest) (models.Account, error)
	GetAccountByEmail(email string) (models.Account, error)
	GetAccountByUserName(userName string) (models.Account, error)
	GetAccountByPhone(phone string) (models.Account, error)
	GetAccountById(id uint) (models.Account, error)
	UpdateAccount(id uint, userName string, password string, email string, phone string) error
	DeleteAccount(id uint) error
	GetAllAccounts() ([]models.Account, error)
}

type accountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) AccountService {
	return &accountService{db : db}
}

// CreateAccount implements AccountService.
// @Summary Create a new account
// @Description Create a new account
// @Tags account
// @Accept json
// @Produce json
// @Param userName body string true "User Name"
// @Param password body string true "Password"
// @Param email body string true "Email"
// @Param phone body string true "Phone"
// @Success 200 {object} models.Account
// @Router /account [post]
func (service *accountService) CreateAccount( request request_models.RegisterRequest) (account models.Account, err error) {

	hashedPassword, err := utils.HashPassword(request.Password)
	
	if err != nil {
		return account, err
	}

	account = models.Account{
		UserName: request.UserName,
		Password: hashedPassword,
		Email:    request.Email,
		Phone:    request.Phone,
	}

	fmt.Println(account)

	if err := service.db.Create(&account).Error; err != nil {
		return account, err
	}

	return account, nil
}

// DeleteAccount implements AccountService.
// @Summary Delete an account
// @Description Delete an account
// @Tags account
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {string} string "Account deleted"
// @Router /account/{id} [delete]
func (service *accountService) DeleteAccount(id uint) error {
	result := service.db.Where("id = ?", id).Delete(&models.Account{})

	if err := result.Error; err != nil {
		return err
	}

	return nil
}


// GetAccountByEmail implements AccountService.
func (service *accountService) GetAccountByEmail(email string) (models.Account, error) {

	var account models.Account

	if err := service.db.Where("email = ?", email).First(&account).Error; err != nil {
		return models.Account{}, err
	}

	return account, nil
}

// GetAccountById implements AccountService.
func (service *accountService) GetAccountById(id uint) (models.Account, error) {
	var account models.Account
	result := service.db.Where("id = ?", id).Find(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

// GetAccountByPhone implements AccountService.
func (service *accountService) GetAccountByPhone(phone string) (models.Account, error) {
	var account models.Account
	result := service.db.Where("phone = ?", phone).Find(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

// GetAccountByUserName implements AccountService.
// @Summary Get an account by user name
// @Description Get an account by user name
// @Tags account
// @Accept json
// @Produce json
// @Param userName path string true "User Name"
// @Success 200 {object} models.Account
// @Router /account/{userName} [get]
func (service *accountService) GetAccountByUserName(userName string) (models.Account, error) {
	var account models.Account
	result := service.db.Where("userName = ?", userName).Find(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

// GetAllAccounts implements AccountService.
func (service *accountService) GetAllAccounts() ([]models.Account, error) {
	var accounts []models.Account
	result := service.db.Find(&accounts)

	if result.Error != nil {
		return nil, result.Error
	}

	return accounts, nil
}

// UpdateAccount implements AccountService.
func (service *accountService) UpdateAccount(id uint, userName string, password string, email string, phone string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	result := service.db.Model(&models.Account{}).Where("id = ?", id).Updates(models.Account{
		UserName: userName,
		Password: hashedPassword,
		Email:    email,
		Phone:    phone,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *accountService) Login(request request_models.LoginRequest) (result string, err error) {
	
	if request.Email == "" || request.Password == "" {
		return "", nil
	}

	account, err := service.GetAccountByEmail(request.Email)

	if err != nil {
		return "", err
	}

	if err := utils.ComparePasswords(account.Password, request.Password); err != nil {
		return "", err
	}

	token, _ := utils.CreateToken(account.Email)
	return token, nil

}
