package theaterfx

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
	"webapp/internal/repositories"
	"webapp/internal/services"
)

var Module = fx.Provide(
	provieTheaterRepository,
	provideTheaterService)

func provieTheaterRepository(db *gorm.DB) repositories.TheaterRepositoryInterface {
	return repositories.NewTheaterRepository(db)
}

func provideTheaterService(theaterRepository repositories.TheaterRepositoryInterface, db *gorm.DB) services.TheaterServiceInterface {
	return services.NewTheaterService(theaterRepository, db)
}
