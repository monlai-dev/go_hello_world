package unit

import (
	"github.com/stretchr/testify/mock"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
)

type MockSeatService struct{ mock.Mock }

func (m *MockSeatService) GetSeatByIdList(id []int) ([]models.Seat, error) {
	args := m.Called(id)
	return args.Get(0).([]models.Seat), args.Error(1)
}
func (m *MockSeatService) GetSeatByID(id int) (models.Seat, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSeatService) CreateListOfSeats(seats []models.Seat) ([]models.Seat, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSeatService) GetAllSeatsByRoomID(page int, pageSize int, roomID int) ([]models.Seat, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSeatService) UpdateSeatByID(seatID int, request request_models.UpdateSeatRequest) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockSeatService) DisableAndEnableSeat(seatID int) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockSeatService) AutoImportSeatWithRow(roomID int, row int) ([]models.Seat, error) {
	//TODO implement me
	panic("implement me")
}
