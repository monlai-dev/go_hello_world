package repositories

import (
	"github.com/jackc/pgx/v5/pgtype"
	models "webapp/internal/models/db_models"
)

type SlotRepositoryInterface interface {
	GetAllSlots(page int, pageSize int) ([]models.Slot, error)
	GetSlotById(id int) (models.Slot, error)
	CreateSlot(slot models.Slot) (models.Slot, error)
	UpdateSlot(slot models.Slot) error
	DeleteSlot(slot models.Slot) error
	GetSlotsByMovieId(movieId int) ([]models.Slot, error)
	GetSlotsInDateRange(startDate pgtype.Timestamp, endDate pgtype.Timestamp) ([]models.Slot, error)
	GetSlotsByRoomId(roomId int) ([]models.Slot, error)
	GetSlotByMovieIdAndBetweenDates(movieId int, startDate pgtype.Timestamp, endDate pgtype.Timestamp) ([]models.Slot, error)
	GetSlotByRoomIdAndBetweenDates(roomId int, startDate pgtype.Timestamp, endDate pgtype.Timestamp) ([]models.Slot, error)
}
