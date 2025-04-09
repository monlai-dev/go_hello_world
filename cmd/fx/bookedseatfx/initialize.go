package bookedseatfx

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
	"webapp/internal/repositories"
	"webapp/internal/services"
)

var Module = fx.Provide(
	provideBookedSeatRepository,
	provideBookedSeatService)

func provideBookedSeatRepository(db *gorm.DB) repositories.BookedSeatRepositoryInterface {
	return repositories.NewBookedRepository(db)
}

func provideBookedSeatService(
	bookedSeatRepository repositories.BookedSeatRepositoryInterface,
) services.BookedSeatServiceInterface {
	return services.NewBookedService(bookedSeatRepository)
}
