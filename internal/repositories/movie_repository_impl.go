package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	models "webapp/internal/models/db_models"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepositoryInterface {
	return &MovieRepository{
		db: db,
	}
}

func (m MovieRepository) GetAllMovies() ([]models.Movie, error) {

	var movies []models.Movie
	if err := m.db.Find(&movies); err.Error != nil {
		return nil, fmt.Errorf("error fetching movies: %v", err.Error)
	}

	return movies, nil
}

func (m MovieRepository) GetMovieById(id int) (models.Movie, error) {

	var movie models.Movie
	if err := m.db.First(&movie, id); err.Error != nil {
		log.Printf("error fetching movie: %v", err)
		return models.Movie{}, fmt.Errorf("error fetching movie: %v", err.Error)
	}

	return movie, nil
}

func (m MovieRepository) CreateMovie(movie models.Movie) (models.Movie, error) {

	tx := m.db.Begin()

	if err := tx.Create(&movie); err.Error != nil {
		tx.Rollback()
		log.Printf("error creating movie: %v", err)
		return models.Movie{}, fmt.Errorf("error creating movie: %v", err.Error)
	}

	return movie, tx.Commit().Error
}

func (m MovieRepository) UpdateMovie(movie models.Movie) error {
	tx := m.db.Begin()

	if err := tx.Save(&movie); err.Error != nil {
		log.Printf("error updating movie: %v", err)
		return errors.New("error updating movie")
	}

	return tx.Commit().Error
}

func (m MovieRepository) DeleteMovie(movie models.Movie) error {

	if err := m.db.Delete(&movie); err.Error != nil {
		log.Printf("error deleting movie: %v", err)
		return errors.New("error deleting movie")
	}

	return nil
}
