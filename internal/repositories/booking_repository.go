package repositories

import models "webapp/internal/models/db_models"

type BookingRepositoryInterface interface {
	GetAllBookings(page int, pageSize int) ([]models.Booking, error)
	GetBookingById(id int) (models.Booking, error)
	CreateBooking(booking models.Booking) (models.Booking, error)
	UpdateBooking(booking models.Booking) error
	DeleteBooking(booking models.Booking) error
}
