package bookingfx

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"webapp/internal/infrastructure/rabbitMq"
	"webapp/internal/repositories"
	"webapp/internal/services"
)

var Module = fx.Provide(
	provideBookingRepository,
	provideBookingService)

func provideBookingRepository(db *gorm.DB) repositories.BookingRepositoryInterface {
	return repositories.NewBookingRepository(db)
}

func provideBookingService(
	repo repositories.BookingRepositoryInterface,
	movieService services.MovieServiceInterface,
	bookedSeatService services.BookedSeatServiceInterface,
	redisClient *redis.Client,
	seatService services.SeatServiceInterface,
	slotService services.SlotServiceInterface,
	cronjobService *services.CronJobService,
	accountService services.AccountServiceInterface,
	rabbitClient *rabbitMq.RabbitMq,
) services.BookingServiceInterface {
	return services.NewBookingService(repo, movieService, bookedSeatService, redisClient, seatService, slotService, cronjobService, accountService, rabbitClient)
}
