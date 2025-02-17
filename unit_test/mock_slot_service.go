package unit

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
)

type MockSlotService struct{ mock.Mock }

func (m *MockSlotService) GetSlotByID(slotId int) (models.Slot, error) {
	args := m.Called(slotId)
	return args.Get(0).(models.Slot), args.Error(1)
}

func (m *MockSlotService) FindAllSlotsByRoomID(roomId int, page int, pageSize int) ([]models.Slot, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSlotService) FindAllSlotByMovieID(movieId int, page int, pageSize int) ([]models.Slot, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSlotService) FindAllSlotByMovieIDAndBetweenDates(movieId int, startDate pgtype.Timestamp, endDate pgtype.Timestamp, page int, pageSize int) ([]models.Slot, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSlotService) CreateSlot(createSlotRequest request_models.CreateSlotRequest) (models.Slot, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSlotService) UpdateSlot(updateSlotRequest request_models.UpdateSlotRequest) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockSlotService) DeleteSlot(slotId int) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockSlotService) GetSlotByRoomIDAndTime(roomId int, startTime pgtype.Timestamp, endTime pgtype.Timestamp) ([]models.Slot, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSlotService) FindAllSlotBetweenDates(startDate pgtype.Timestamp, endDate pgtype.Timestamp, page int, pageSize int) ([]models.Slot, error) {
	//TODO implement me
	panic("implement me")
}
