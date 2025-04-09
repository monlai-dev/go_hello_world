package roomfx

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
	"webapp/internal/repositories"
	"webapp/internal/services"
)

var Module = fx.Provide(
	provideRoomRepository,
	provideRoomService)

func provideRoomRepository(db *gorm.DB) repositories.RoomRepositoryInterface {
	return repositories.NewRoomRepository(db)
}

func provideRoomService(db *gorm.DB, roomRepository repositories.RoomRepositoryInterface, theaterService services.TheaterServiceInterface) services.RoomServiceInterface {
	return services.NewRoomService(db, roomRepository, theaterService)
}
