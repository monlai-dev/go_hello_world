package services

import (
	"webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
)

type AddressServiceInterface interface {
	CreateAddress(request request_models.AddressRequest) (models.Address, error)
}
