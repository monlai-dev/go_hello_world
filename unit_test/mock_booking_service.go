package unit

import (
	"github.com/stretchr/testify/mock"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
)

type MockBookingService struct{ mock.Mock }

func (m *MockBookingService) GetBookingByID(bookingID int) (models.Booking, error) {
	args := m.Called(bookingID)
	return args.Get(0).(models.Booking), args.Error(1)
}

func (m *MockBookingService) CreateBooking(request request_models.CreateBookingRequest, email string) (models.Booking, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookingService) GetAllBookingsByAccountID(accountID int, page int, pageSize int) ([]models.Booking, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookingService) GetAllBookingsBySlotID(slotID int, page int, pageSize int) ([]models.Booking, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookingService) UpdateBookingByID(bookingID int, status string) (models.Booking, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookingService) CancelBookingByID(bookingID int) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookingService) ConfirmBookingByID(bookingID int) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookingService) Scheduler() error {
	//TODO implement me
	panic("implement me")
}

func (m *MockBookingService) SendNotiEmail([]request_models.TestingEmailFormat) error {
	//TODO implement me
	panic("implement me")
}
