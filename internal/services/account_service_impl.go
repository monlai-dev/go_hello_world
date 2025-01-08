package services

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
	"webapp/internal/models/db_models"
	request_models2 "webapp/internal/models/request_models"
	"webapp/internal/repositories"
	utils2 "webapp/pkg/utils"
)

type accountService struct {
	db                *gorm.DB
	addressService    AddressServiceInterface
	redisClient       *redis.Client
	accountRepository repositories.AccountRepositoryInterface
}

func NewAccountService(db *gorm.DB, addressService AddressServiceInterface, redisClient *redis.Client, accountRepository repositories.AccountRepositoryInterface) AccountServiceInterface {
	return &accountService{
		db:                db,
		addressService:    addressService,
		redisClient:       redisClient,
		accountRepository: accountRepository,
	}

}

// CreateAccount implements AccountService.
// @Summary Create a new account
// @Description Create a new account
// @Tags account
// @Accept json
// @Produce json
// @Param userName body string true "UserName"
// @Param password body string true "Password"
// @Param email body string true "Email"
// @Param phone body string true "Phone"
// @Success 200 {object} models.Account
// @Router /account [post]
func (service *accountService) CreateAccount(request request_models2.RegisterRequest) (account models.Account, err error) {

	tx := service.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	hashedPassword, err := utils2.HashPassword(request.Password)

	if err != nil {
		return account, err
	}

	account = models.Account{
		UserName: request.UserName,
		Password: hashedPassword,
		Email:    request.Email,
		Phone:    request.Phone,
	}

	if err := tx.Omit("Address").Create(&account).Error; err != nil {
		return account, err
	}

	return account, tx.Commit().Error
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
	account, err := service.accountRepository.FindAccountByEmail(email)

	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

// GetAccountById implements AccountService.
func (service *accountService) GetAccountById(id int) (models.Account, error) {
	var account models.Account
	result := service.db.Where("id = ?", id).First(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

// GetAccountByPhone implements AccountService.
func (service *accountService) GetAccountByPhone(phone string) (models.Account, error) {
	var account models.Account

	account, err := service.accountRepository.FindAccountByPhone(phone)

	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

// GetAccountByUserName implements AccountService.
// @Summary Get an account by username
// @Description Get an account by username
// @Tags account
// @Accept json
// @Produce json
// @Param userName path string true "UserName"
// @Success 200 {object} models.Account
// @Router /account/{userName} [get]
func (service *accountService) GetAccountByUserName(userName string) (models.Account, error) {
	var account models.Account

	account, err := service.accountRepository.FindAccountByUserName(userName)

	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

// GetAllAccounts implements AccountService.
func (service *accountService) GetAllAccounts(page int, pageSize int) ([]models.Account, error) {

	result, err := service.accountRepository.GetAllAccounts(page, pageSize)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateAccount implements AccountService.
func (service *accountService) UpdateAccount(id uint, userName string, password string, email string, phone string) error {
	hashedPassword, err := utils2.HashPassword(password)

	if err != nil {
		return err
	}

	account, err := service.accountRepository.FindAccountByEmail(email)

	if err != nil {
		return err
	}

	account.UserName = userName
	account.Password = hashedPassword
	account.Email = email
	account.Phone = phone

	result := service.accountRepository.UpdateAccount(account)

	if result != nil {
		return result
	}

	return nil
}

func (service *accountService) Login(request request_models2.LoginRequest) (result string, err error) {

	if request.Email == "" || request.Password == "" {
		return "", nil
	}

	account, err := service.GetAccountByEmail(request.Email)

	if err != nil {
		return "", err
	}

	if err := utils2.ComparePasswords(account.Password, request.Password); err != nil {
		return "", err
	}

	token, _ := utils2.CreateToken(account.Email)
	return token, nil

}

func (service *accountService) GetRandomAccount() (models.Account, error) {
	var account models.Account
	result := service.db.Last(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

func (service *accountService) GetAllHomelessAccounts() ([]models.Account, error) {
	var accounts []models.Account

	result := service.db.Preload("Address").Where("address_id IS NULL").Find(&accounts)

	if result.Error != nil {
		return nil, result.Error
	}

	return accounts, nil
}

func (service *accountService) UpdateAddress(email string, addressRequest request_models2.AddressRequest) error {
	// Start a transaction
	tx := service.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch account
	account, err := service.GetAccountByEmail(email)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch account with email %s: %w", email, err)
	}

	// Create new address
	newAddress, err := service.addressService.CreateAddress(addressRequest)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create address: %w", err)
	}

	// Update account with new address
	account.Address = newAddress
	account.AddressId = &newAddress.ID

	if err := tx.Model(&account).Updates(&account).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update account with email %s: %w", email, err)
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (service *accountService) Logout(token string) error {

	ctx := context.Background()
	val := service.redisClient.Set(ctx, "logged_out"+token, "", time.Minute*15)

	if val.Err() != nil {
		return val.Err()
	}

	return nil
}
