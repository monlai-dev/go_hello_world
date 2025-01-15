package services

import (
	"fmt"
	"log"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
	"webapp/internal/repositories"
)

type MovieService struct {
	movieRepository repositories.MovieRepositoryInterface
}

func NewMovieService(movieRepository repositories.MovieRepositoryInterface) MovieServiceInterface {
	return MovieService{
		movieRepository: movieRepository,
	}
}

func (m MovieService) GetMovieByID(id int) (models.Movie, error) {

	movie, err := m.movieRepository.GetMovieById(id)

	if err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}

func (m MovieService) CreateMovie(request request_models.CreateMovieRequest) (models.Movie, error) {

	movie := models.Movie{
		Title:       request.Title,
		Description: request.Description,
		Duration:    int64(request.Duration),
	}

	createdMovie, err := m.movieRepository.CreateMovie(movie)

	if err != nil {
		log.Printf("error creating movie: %v", err)
		return models.Movie{}, err
	}

	return createdMovie, nil
}

func (m MovieService) UpdateMovie(request request_models.UpdateMovieRequest) error {

	movie, err := m.GetMovieByID(request.MovieID)
	if err != nil {
		return err
	}

	movie.Title = request.Title
	movie.Description = request.Description
	movie.Duration = int64(request.Duration)

	if err := m.movieRepository.UpdateMovie(movie); err != nil {
		log.Printf("error updating movie: %v", err)
		return err
	}

	return nil
}

func (m MovieService) DeleteMovie(id int) error {

	movie, err := m.GetMovieByID(id)
	if err != nil {
		return fmt.Errorf("error fetching movie: %v", err)
	}

	if err := m.movieRepository.DeleteMovie(movie); err != nil {
		log.Printf("error deleting movie: %v", err)
		return err
	}

	return nil
}

func (m MovieService) GetAllMovies(page int, pageSize int) ([]models.Movie, error) {

	movies, err := m.movieRepository.GetAllMovies()

	if err != nil {
		return []models.Movie{}, err
	}

	return movies, nil
}
