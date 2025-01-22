package services

import (
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
)

type BookingServiceInterface interface {
	CreateBooking(request request_models.CreateBookingRequest, email string) (models.Booking, error)
	GetBookingByID(bookingID int) (models.Booking, error)
	GetAllBookingsByAccountID(accountID int, page int, pageSize int) ([]models.Booking, error)
	GetAllBookingsBySlotID(slotID int, page int, pageSize int) ([]models.Booking, error)
	UpdateBookingByID(bookingID int, status string) (models.Booking, error)
	CancelBookingByID(bookingID int) error
	ConfirmBookingByID(bookingID int) error
	Scheduler() error
}
