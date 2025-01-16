package services

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
	"time"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
	"webapp/internal/repositories"
)

const (
	movieKeyPrefix = "movie:%d"
	cacheTTL       = 24 * time.Hour
)

type MovieService struct {
	movieRepository repositories.MovieRepositoryInterface
	redisClient     *redis.Client
}

func NewMovieService(movieRepository repositories.MovieRepositoryInterface,
	client *redis.Client) MovieServiceInterface {
	return MovieService{
		movieRepository: movieRepository,
		redisClient:     client,
	}
}

func (m MovieService) GetMovieByID(id int) (models.Movie, error) {
	ctx := context.Background()

	// Redis key
	movie, err := m.getFromCache(ctx, id)
	if err == nil {
		return movie, nil
	}

	if !errors.Is(err, redis.Nil) {
		log.Printf("failed to get movie from cache: %v", err)
	}

	movie, err = m.movieRepository.GetMovieById(id)
	if err != nil {
		return models.Movie{}, fmt.Errorf("failed to get movie from database: %w", err)
	}

	// Cache the result
	if err := m.cacheMovie(ctx, movie); err != nil {
		log.Printf("Failed to cache movie: %v", err)
	}

	return movie, nil
}

func (m MovieService) CreateMovie(request request_models.CreateMovieRequest) (models.Movie, error) {

	ctx := context.Background()

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

	redisKey := fmt.Sprintf("movie:%d", createdMovie.ID)
	movieBytes, err := json.Marshal(createdMovie)
	if err == nil {
		// Store with a TTL of 24 hours
		setErr := m.redisClient.Set(ctx, redisKey, movieBytes, 24*time.Hour).Err()
		if setErr != nil {
			fmt.Printf("Failed to cache movie in Redis: %v\n", setErr)
		}
	} else {
		fmt.Printf("Failed to marshal movie for caching: %v\n", err)
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

	movies, err := m.getAllFromCache(context.Background())

	if err == nil {
		return movies, nil
	}
	if !errors.Is(err, redis.Nil) {
		log.Printf("Failed to get movies from cache: %v", err)
	}

	movies, dbErr := m.movieRepository.GetAllMovies()
	if dbErr != nil {
		return []models.Movie{}, err
	}

	return movies, nil
}

func (m MovieService) getFromCache(ctx context.Context, id int) (models.Movie, error) {
	redisKey := fmt.Sprintf(movieKeyPrefix, id)
	movieJSON, err := m.redisClient.Get(ctx, redisKey).Result()
	if err != nil {
		log.Printf("Failed to get movie from cache: %v", err)
		return models.Movie{}, err
	}

	var movie models.Movie
	if err := json.Unmarshal([]byte(movieJSON), &movie); err != nil {
		log.Printf("Failed to unmarshal movie from cache: %v", err)
		return models.Movie{}, err
	}

	return movie, nil
}

func (m MovieService) cacheMovie(ctx context.Context, movie models.Movie) error {
	movieBytes, err := json.Marshal(movie)
	if err != nil {
		return fmt.Errorf("failed to marshal movie for caching: %w", err)
	}

	redisKey := fmt.Sprintf(movieKeyPrefix, movie.ID)
	return m.redisClient.Set(ctx, redisKey, movieBytes, cacheTTL).Err()
}

func (m MovieService) getAllFromCache(ctx context.Context) ([]models.Movie, error) {
	keys, err := m.redisClient.Keys(ctx, "movie:*").Result()
	if err != nil {
		log.Printf("Failed to get keys from cache: %v", err)
		return nil, err
	}

	movies := make([]models.Movie, 0, len(keys))
	for _, key := range keys {
		movieJSON, err := m.redisClient.Get(ctx, key).Result()
		if err != nil {
			log.Printf("Failed to get movie from cache: %v", err)
			return nil, err
		}

		var movie models.Movie
		if err := json.Unmarshal([]byte(movieJSON), &movie); err != nil {
			log.Printf("Failed to unmarshal movie from cache: %v", err)
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}
