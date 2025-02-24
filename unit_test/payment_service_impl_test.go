package unit

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"math/rand"
	"os"
	"regexp"
	"testing"
	"time"
	"webapp/internal/services"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"webapp/internal/models/db_models"
)

func TestPaymentService_CreatePaymentLinkWithPayOsUsingBookingId(t *testing.T) {
	// Set up environment variables required for PayOS API authentication
	os.Setenv("CLIENT_ID", "e47ad6bc-40f4-4fc9-a723-cd8ddae9e3a8")
	os.Setenv("API_KEY", "0ca2d2e9-d6fa-4c13-9885-848080846bd7")
	os.Setenv("CHECK_SUM_KEY", "1589dd37d76b03b0968bcd445dd4ed0e38fa63b0d5811bba8dd72b7ea88f95c0")
	randomOrderCode := generateOrderCode()

	testCases := []struct {
		name          string
		bookingID     int
		expectedError error
		mockBooking   models.Booking
		mockSeats     []models.Seat
		mockSlot      models.Slot
		mockBooked    []models.BookedSeat
	}{
		{
			name:          "booking not found",
			bookingID:     9999,
			expectedError: fmt.Errorf("error fetching booking"),
		},
		{
			name:          "valid payment link",
			bookingID:     randomOrderCode,
			expectedError: nil,
			mockBooking: models.Booking{
				SlotID:    1,
				AccountID: 1,
				BookingTime: pgtype.Timestamp{
					Time: time.Now(),
				},
				DueTime: pgtype.Timestamp{
					Time: time.Now().Add(time.Hour),
				},
				TotalPrice: 10000,
				Model: gorm.Model{
					ID:        uint(randomOrderCode),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			mockSeats: []models.Seat{
				{RoomID: 1, Name: "Seat_1", IsWorking: true},
			},
			mockSlot: models.Slot{
				StartTime: pgtype.Timestamp{
					Time: time.Now(),
				},
				EndTime: pgtype.Timestamp{
					Time: time.Now().Add(time.Hour),
				},
				MovieID: 1,
				Price:   10000,
				RoomID:  1,
			},
			mockBooked: []models.BookedSeat{
				{SeatID: 1, SlotID: 1, Status: "ON_HOLD"},
			},
		}, {
			name:          "duplicate order code",
			bookingID:     1,
			expectedError: errors.New("error creating payment link"),
			mockBooking: models.Booking{
				SlotID:    1,
				AccountID: 1,
				BookingTime: pgtype.Timestamp{
					Time: time.Now(),
				},
				DueTime: pgtype.Timestamp{
					Time: time.Now().Add(time.Hour),
				},
				TotalPrice: 10000,
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			mockSeats: []models.Seat{
				{RoomID: 1, Name: "Seat_1", IsWorking: true},
			},
			mockSlot: models.Slot{
				StartTime: pgtype.Timestamp{
					Time: time.Now(),
				},
				EndTime: pgtype.Timestamp{
					Time: time.Now().Add(time.Hour),
				},
				MovieID: 1,
				Price:   10000,
				RoomID:  1,
			},
			mockBooked: []models.BookedSeat{
				{SeatID: 1, SlotID: 1, Status: "ON_HOLD"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockBookingService := new(MockBookingService)
			mockBookedSeatService := new(MockBookedSeatService)
			mockSlotService := new(MockSlotService)
			mockSeatService := new(MockSeatService)

			if testCase.mockBooking.ID != 0 {
				mockBookingService.On("GetBookingByID", testCase.bookingID).Return(testCase.mockBooking, nil)
				mockBookedSeatService.On("FindAllBookedSeatWithBookingId", int(testCase.mockBooking.ID)).Return(testCase.mockBooked, nil)
				mockSlotService.On("GetSlotByID", int(testCase.mockBooking.SlotID)).Return(testCase.mockSlot, nil)
				mockSeatService.On("GetSeatByIdList", []int{1}).Return(testCase.mockSeats, nil)
			} else {
				mockBookingService.On("GetBookingByID", testCase.bookingID).Return(models.Booking{}, fmt.Errorf("booking with id %d not found", testCase.bookingID))
			}

			paymentService := services.PaymentService{
				BookingService:    mockBookingService,
				BookedSeatService: mockBookedSeatService,
				SlotService:       mockSlotService,
				SeatService:       mockSeatService,
			}

			// When
			paymentLink, err := paymentService.CreatePaymentLinkWithPayOsUsingBookingId(testCase.bookingID)

			// Then
			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				payosPattern := `^https://pay\.payos\.vn/web/[a-f0-9]+$`
				matched, _ := regexp.MatchString(payosPattern, paymentLink)
				require.True(t, matched, "Payment link format invalid")
			}
		})
	}
}

func generateOrderCode() int {
	return rand.Intn(100000) + 100000
}
