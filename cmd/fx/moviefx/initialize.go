package moviefx

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"webapp/internal/repositories"
	"webapp/internal/services"
)

var Module = fx.Provide(provideMovieRepository,
	provideMovieService)

func provideMovieRepository(db *gorm.DB) repositories.MovieRepositoryInterface {
	return repositories.NewMovieRepository(db)
}

func provideMovieService(repo repositories.MovieRepositoryInterface, redisClient *redis.Client) services.MovieServiceInterface {
	return services.NewMovieService(repo, redisClient)
}
