package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"webapp/internal/infrastructure/database"
	models "webapp/internal/models/db_models"
)

type SeatRepository struct {
	db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) SeatRepositoryInterface {

	if db == nil {
		panic("NewSeatRepository: db is nil")
	}

	return &SeatRepository{db}
}

func (s SeatRepository) UpdateSeat(seat models.Seat) error {
	tx := s.db.Begin()

	result := s.db.Updates(&seat)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return tx.Commit().Error
}

func (s SeatRepository) FindSeatById(id int) (models.Seat, error) {
	var seat models.Seat
	result := s.db.First(&seat, id)
	if result.Error != nil {
		return models.Seat{}, result.Error
	}
	return seat, nil
}

func (s SeatRepository) CreateListOfSeats(seats []models.Seat) ([]models.Seat, error) {
	tx := s.db.Begin()

	result := tx.Create(&seats)
	if result.Error != nil {
		tx.Rollback()
		return []models.Seat{}, errors.New("Error when create list of seats " + result.Error.Error())
	}
	return seats, tx.Commit().Error
}

func (s SeatRepository) GetAllSeatsByRoomID(page int, pageSize int, roomID int) ([]models.Seat, error) {

	var seats []models.Seat
	result := s.db.Scopes(database.Paginate(page, pageSize)).Where("room_id = ?", roomID).Find(&seats)

	if result.Error != nil {
		return []models.Seat{}, result.Error
	}
	return seats, nil

}

func (s SeatRepository) GetAllSeatsBySlotID(page int, pageSize int, slotID int) ([]models.Seat, error) {
	var slot models.Slot
	if err := s.db.Preload("Rooms").First(&slot, slotID).Error; err != nil {
		return nil, err
	}

	var seats []models.Seat
	if err := s.db.Scopes(database.Paginate(page, pageSize)).Where("room_id = ?", slot.RoomID).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}

func (s SeatRepository) FindAllSeatsWithIds(ids []int) ([]models.Seat, error) {

	var seats []models.Seat
	if err := s.db.Where("id IN (?)", ids).Find(&seats).Error; err != nil {
		log.Printf("error fetching seats: %v", err)
		return nil, fmt.Errorf("error when fetching seats")
	}
	return seats, nil
}
