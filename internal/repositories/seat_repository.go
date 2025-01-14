package repositories

import models "webapp/internal/models/db_models"

type SeatRepositoryInterface interface {
	FindSeatById(id int) (models.Seat, error)
	CreateListOfSeats(seats []models.Seat) ([]models.Seat, error)
	GetAllSeatsByRoomID(page int, pageSize int, roomID int) ([]models.Seat, error)
	GetAllSeatsBySlotID(page int, pageSize int, slotID int) ([]models.Seat, error)
	UpdateSeat(seat models.Seat) error
}
