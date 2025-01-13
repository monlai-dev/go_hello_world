package repositories

import models "webapp/internal/models/db_models"

type MovieRepositoryInterface interface {
	GetAllMovies() ([]models.Movie, error)
	GetMovieById(id int) (models.Movie, error)
	CreateMovie(movie models.Movie) (models.Movie, error)
	UpdateMovie(movie models.Movie) error
	DeleteMovie(movie models.Movie) error
}
