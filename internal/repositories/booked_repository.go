package repositories

import models "webapp/internal/models/db_models"

type BookedSeatRepositoryInterface interface {
	GetAllBooked(page int, pageSize int) ([]models.BookedSeat, error)
	GetBookedById(id int) (models.BookedSeat, error)
	CreateBooked(booked []models.BookedSeat) ([]models.BookedSeat, error)
	UpdateBooked(booked []models.BookedSeat) error
	DeleteBooked(booked models.BookedSeat) error
	FindAllBookedSeatWithIds(ids []int) ([]models.BookedSeat, error)
	GetBookedSeatBySlotId(slotId int) ([]models.BookedSeat, error)
	GetBookedSeatBySlotIdAndStatus(slotId int, status []string) ([]models.BookedSeat, error)
	GetAllBookedSeatByBookingId(bookingId int) ([]models.BookedSeat, error)
}
