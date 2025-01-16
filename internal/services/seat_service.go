package services

import (
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
)

type SeatServiceInterface interface {
	GetSeatByID(id int) (models.Seat, error)
	CreateListOfSeats(seats []models.Seat) ([]models.Seat, error)
	GetAllSeatsByRoomID(page int, pageSize int, roomID int) ([]models.Seat, error)
	UpdateSeatByID(seatID int, request request_models.UpdateSeatRequest) error
	DisableAndEnableSeat(seatID int) error
	AutoImportSeatWithRow(roomID int, row int) ([]models.Seat, error)
	GetSeatByIdList(id []int) ([]models.Seat, error)
}
