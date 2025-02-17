package unit

import (
	"github.com/stretchr/testify/mock"
	models "webapp/internal/models/db_models"
)

type MockBookedSeatService struct{ mock.Mock }

func (m *MockBookedSeatService) FindAllBookedSeatWithBookingId(bookingId int) ([]models.BookedSeat, error) {
	args := m.Called(bookingId)
	return args.Get(0).([]models.BookedSeat), args.Error(1)
}

func (m *MockBookedSeatService) FindAllBookedSeatsBySlotID(slotId int, page int, pageSize int) ([]models.BookedSeat, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookedSeatService) CreateBookedSeat(seat []models.BookedSeat) ([]models.BookedSeat, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookedSeatService) UpdateBookedSeat(seat []models.BookedSeat) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookedSeatService) IsSeatsAvailable(seatIds []int, slotId int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookedSeatService) FindAllBookedSeatWithSeatIDs(seatIds []int) ([]models.BookedSeat, error) {
	//TODO implement me
	panic("implement me")
}
