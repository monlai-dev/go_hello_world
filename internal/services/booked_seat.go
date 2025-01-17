package services

import (
	models "webapp/internal/models/db_models"
)

type BookedSeatServiceInterface interface {
	FindAllBookedSeatsBySlotID(slotId int, page int, pageSize int) ([]models.BookedSeat, error)
	CreateBookedSeat(seat []models.BookedSeat) ([]models.BookedSeat, error)
	UpdateBookedSeat(seat []models.BookedSeat) error
	IsSeatsAvailable(seatIds []int, slotId int) (bool, error)
	FindAllBookedSeatWithSeatIDs(seatIds []int) ([]models.BookedSeat, error)
	FindAllBookedSeatWithBookingId(bookingId int) ([]models.BookedSeat, error)
}
