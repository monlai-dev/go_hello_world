package services

import (
	"github.com/jackc/pgx/v5/pgtype"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
)

type SlotServiceInterface interface {
	FindAllSlotsByRoomID(roomId int, page int, pageSize int) ([]models.Slot, error)
	FindAllSlotByMovieID(movieId int, page int, pageSize int) ([]models.Slot, error)
	FindAllSlotByMovieIDAndBetweenDates(movieId int, startDate pgtype.Timestamp, endDate pgtype.Timestamp, page int, pageSize int) ([]models.Slot, error)
	CreateSlot(createSlotRequest request_models.CreateSlotRequest) (models.Slot, error)
	UpdateSlot(updateSlotRequest request_models.UpdateSlotRequest) error
	DeleteSlot(slotId int) error
	GetSlotByID(slotId int) (models.Slot, error)
	GetSlotByRoomIDAndTime(roomId int, startTime pgtype.Timestamp, endTime pgtype.Timestamp) ([]models.Slot, error)
	FindAllSlotBetweenDates(startDate pgtype.Timestamp, endDate pgtype.Timestamp, page int, pageSize int) ([]models.Slot, error)
}
