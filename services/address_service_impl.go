package services

import (
	"gorm.io/gorm"
	models "webapp/models/db_models"
	"webapp/models/request_models"
)

type AddressService struct {
	db             *gorm.DB
	accountService AccountServiceInterface
}

func NewAddressService(db *gorm.DB) AddressServiceInterface {
	return &AddressService{db: db}
}

func (service *AddressService) CreateAddress(request request_models.AddressRequest) (models.Address, error) {

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

	return address, nil
}
