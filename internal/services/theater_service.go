package services

import models "webapp/internal/models/db_models"

type TheaterServiceInterface interface {
	GetAllTheaters() ([]models.Theater, error)
	GetTheaterById(id int) (models.Theater, error)
	CreateTheater(theater models.Theater) (models.Theater, error)
	UpdateTheater(theater models.Theater) error
	DeleteTheater(theater models.Theater) error
}
