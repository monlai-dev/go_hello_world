package repositories

import (
	"gorm.io/gorm"
	"webapp/initializer"
	"webapp/models/db_models"
)

type AccountRepository struct {
	db *gorm.DB
}

func (a AccountRepository) FindAccountByEmail(email string) (models.Account, error) {

	var account models.Account
	result := a.db.Where("email = ?", email).First(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil

}

func (a AccountRepository) FindAccountByUserName(userName string) (models.Account, error) {

	var account models.Account
	result := a.db.Where("userName = ?", userName).First(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

func (a AccountRepository) FindAccountByPhone(phone string) (models.Account, error) {
	var account models.Account
	result := a.db.Where("phone = ?", phone).First(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

func (a AccountRepository) FindAccountById(id int) (models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) CreateAccount(account models.Account) (models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) UpdateAccount(account models.Account) error {
	tx := a.db.Begin()

	result := a.db.Save(&account)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return tx.Commit().Error
}

func (a AccountRepository) DeleteAccount(account models.Account) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) GetAllAccounts(page int, pageSize int) ([]models.Account, error) {

	var accounts []models.Account
	result := a.db.Scopes(initializer.Paginate(page, pageSize)).Find(&accounts).Find(&accounts)

	if result.Error != nil {
		return nil, result.Error
	}

	return accounts, nil
}

func (a AccountRepository) GetRandomAccount() (models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) GetAllHomelessAccounts() ([]models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountRepository(db *gorm.DB) AccountRepositoryInterface {
	return &AccountRepository{db: db}
}
