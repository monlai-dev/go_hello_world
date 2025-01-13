package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"webapp/internal/infrastructure/database"
	models "webapp/internal/models/db_models"
)

type BookedRepository struct {
	db *gorm.DB
}

func (b BookedRepository) GetAllBooked(page int, pageSize int) ([]models.BookedSeat, error) {

	var booked []models.BookedSeat

	if err := b.db.Scopes(database.Paginate(page, pageSize)).Find(&booked).Error; err != nil {
		return nil, fmt.Errorf("error fetching booked seats: %w", err)
	}

	if len(booked) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return booked, nil
}

func (b BookedRepository) GetBookedById(id int) (models.BookedSeat, error) {

	var booked models.BookedSeat
	if err := b.db.First(&booked, id).Error; err != nil {
		return models.BookedSeat{}, fmt.Errorf("error fetching booked seat: %w", err)
	}

	return booked, nil
}

func (b BookedRepository) CreateBooked(booked models.BookedSeat) ([]models.BookedSeat, error) {

	tx := b.db.Begin()

	if err := tx.Create(&booked).Error; err != nil {
		return nil, fmt.Errorf("error creating booked seat: %w", err)
	}

	return []models.BookedSeat{booked}, tx.Commit().Error
}

func (b BookedRepository) UpdateBooked(booked []models.BookedSeat) error {

	tx := b.db.Begin()

	for _, b := range booked {
		if err := tx.Updates(&b).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("error updating booked seat: %w", err)
		}
	}

	return tx.Commit().Error
}

func (b BookedRepository) DeleteBooked(booked models.BookedSeat) error {

	tx := b.db.Begin()

	if err := tx.Delete(&booked).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting booked seat: %w", err)
	}

	return tx.Commit().Error
}

func (b BookedRepository) GetBookedSeatBySlotId(slotId int) ([]models.BookedSeat, error) {

	var booked []models.BookedSeat

	if err := b.db.Where("slot_id = ?", slotId).Find(&booked).Error; err != nil {
		return nil, fmt.Errorf("error fetching booked seats: %w", err)
	}

	if len(booked) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return booked, nil
}

func (b BookedRepository) GetBookedSeatBySlotIdAndStatus(slotId int, status []string) ([]models.BookedSeat, error) {

	var booked []models.BookedSeat

	if err := b.db.Where("slot_id = ? AND status IN ?", slotId, status).Find(&booked).Error; err != nil {
		return nil, fmt.Errorf("error fetching booked seats: %w", err)
	}

	if len(booked) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return booked, nil
}

func NewBookedRepository(db *gorm.DB) BookedSeatRepositoryInterface {
	return &BookedRepository{db: db}
}
