package slotfx

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"webapp/internal/repositories"
	"webapp/internal/services"
)

var Module = fx.Provide(
	provideSlotRepository,
	provideSlotService)

func provideSlotRepository(db *gorm.DB) repositories.SlotRepositoryInterface {
	return repositories.NewSlotRepository(db)
}

func provideSlotService(slotRepo repositories.SlotRepositoryInterface, roomService services.RoomServiceInterface, movieSerivce services.MovieServiceInterface, redis *redis.Client) services.SlotServiceInterface {
	return services.NewSlotService(slotRepo, roomService, movieSerivce, redis)
}
