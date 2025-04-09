package seatfx

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
	"webapp/internal/repositories"
	"webapp/internal/services"
)

var Module = fx.Provide(
	provideSeatRepository,
	provideSeatService)

func provideSeatRepository(db *gorm.DB) repositories.SeatRepositoryInterface {
	return repositories.NewSeatRepository(db)
}

func provideSeatService(seatRepo repositories.SeatRepositoryInterface, roomSerive services.RoomServiceInterface) services.SeatServiceInterface {
	return services.NewSeatService(seatRepo, roomSerive)
}
