package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"webapp/internal/infrastructure/database"
	models "webapp/internal/models/db_models"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepositoryInterface {
	return &BookingRepository{db: db}
}

func (b BookingRepository) GetAllBookings(page int, pageSize int) ([]models.Booking, error) {

	var bookings []models.Booking

	if err := b.db.Scopes(database.Paginate(page, pageSize)).Find(&bookings).Error; err != nil {
		return nil, err
	}

	if len(bookings) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return bookings, nil
}

func (b BookingRepository) GetBookingById(id int) (models.Booking, error) {

	var booking models.Booking

	if err := b.db.First(&booking, id).Error; err != nil {
		return models.Booking{}, fmt.Errorf("error fetching booking: %v", err)
	}

	return booking, nil
}

func (b BookingRepository) CreateBooking(booking models.Booking) (models.Booking, error) {
	tx := b.db.Begin()
	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		return models.Booking{}, fmt.Errorf("error creating booking: %v", err)
	}

	return booking, tx.Commit().Error
}

func (b BookingRepository) UpdateBooking(booking models.Booking) error {

	tx := b.db.Begin()
	if err := tx.Save(&booking).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating booking: %v", err)
	}

	return tx.Commit().Error
}

func (b BookingRepository) DeleteBooking(booking models.Booking) error {

	tx := b.db.Begin()
	if err := tx.Delete(&booking).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting booking: %v", err)
	}

	return tx.Commit().Error
}
