package services

import (
	"gorm.io/gorm"
	"webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
)

type AddressService struct {
	db             *gorm.DB
	accountService AccountServiceInterface
}

func NewAddressService(db *gorm.DB) AddressServiceInterface {
	return &AddressService{db: db}
}

func (service *AddressService) CreateAddress(request request_models.AddressRequest) (models.Address, error) {
	tx := service.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	address := models.Address{
		Street: request.Street,
		City:   request.City,
		State:  request.State,
		Zip:    request.Zip,
	}
	result := service.db.Create(&address)

	if result.Error != nil {
		return models.Address{}, result.Error
	}

	return address, tx.Commit().Error
}
