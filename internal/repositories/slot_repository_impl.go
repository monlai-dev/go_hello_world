package repositories

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
	"webapp/internal/infrastructure/database"
	models "webapp/internal/models/db_models"
)

type SlotRepository struct {
	db *gorm.DB
}

const (
	ERROR_MESSAGE = "error fetching slots: %w"
)

func NewSlotRepository(db *gorm.DB) SlotRepositoryInterface {
	return &SlotRepository{db: db}
}

func (s SlotRepository) GetAllSlots(page int, pageSize int) ([]models.Slot, error) {

	var slots []models.Slot

	if err := s.db.Scopes(database.Paginate(page, pageSize)).Find(&slots).Error; err != nil {
		return nil, fmt.Errorf("error fetching slots: %w", err)
	}

	if len(slots) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return slots, nil
}

func (s SlotRepository) GetSlotById(id int) (models.Slot, error) {

	var slot models.Slot

	if err := s.db.First(&slot, id).Error; err != nil {
		return models.Slot{}, fmt.Errorf(ERROR_MESSAGE, err)
	}

	return slot, nil

}

func (s SlotRepository) CreateSlot(slot models.Slot) (models.Slot, error) {

	tx := s.db.Begin()
	if err := tx.Create(&slot).Error; err != nil {
		tx.Rollback()
		return models.Slot{}, fmt.Errorf(ERROR_MESSAGE, err)
	}

	return slot, tx.Commit().Error
}

func (s SlotRepository) UpdateSlot(slot models.Slot) error {

	tx := s.db.Begin()
	if err := tx.Updates(&slot).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating slot: %w", err)
	}

	return tx.Commit().Error
}

func (s SlotRepository) DeleteSlot(slot models.Slot) error {

	tx := s.db.Begin()
	if err := tx.Delete(&slot).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting slot: %w", err)
	}

	return tx.Commit().Error
}

func (s SlotRepository) GetSlotsByMovieId(movieId int) ([]models.Slot, error) {

	var slots []models.Slot

	if err := s.db.Where("movie_id = ?", movieId).Find(&slots).Error; err != nil {
		return nil, fmt.Errorf("error fetching slots: %w", err)
	}

	if len(slots) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return slots, nil
}

func (s SlotRepository) GetSlotsInDateRange(startDate pgtype.Timestamp, endDate pgtype.Timestamp) ([]models.Slot, error) {

	var slots []models.Slot

	if err := s.db.Where("start_time >= ? AND end_time <= ?", startDate, endDate).Find(&slots).Error; err != nil {
		return nil, fmt.Errorf("error fetching slots: %w", err)
	}

	if len(slots) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return slots, nil
}

func (s SlotRepository) GetSlotsByRoomId(roomId int) ([]models.Slot, error) {

	var slots []models.Slot

	if err := s.db.Where("room_id = ?", roomId).Find(&slots).Error; err != nil {
		return nil, fmt.Errorf(ERROR_MESSAGE, err)
	}

	if len(slots) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return slots, nil
}
