package services

import (
	models "webapp/models/db_models"
	"webapp/models/request_models"
)

type AddressServiceInterface interface {
	CreateAddress(request request_models.AddressRequest) (models.Address, error)
}
