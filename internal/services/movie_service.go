package services

import (
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
)

type MovieServiceInterface interface {
	GetMovieByID(id int) (models.Movie, error)
	CreateMovie(request request_models.CreateMovieRequest) (models.Movie, error)
	UpdateMovie(request request_models.UpdateMovieRequest) error
	DeleteMovie(id int) error
	GetAllMovies(page int, pageSize int) ([]models.Movie, error)
}
