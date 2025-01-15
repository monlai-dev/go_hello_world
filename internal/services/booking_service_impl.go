package services

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
	"webapp/internal/repositories"
)

type BookingService struct {
	bookingRepository repositories.BookingRepositoryInterface
	movieService      MovieServiceInterface
	bookedSeatService BookedSeatServiceInterface
	redisClient       *redis.Client
}

func NewBookingService(
	bookingRepository repositories.BookingRepositoryInterface,
	movieService MovieServiceInterface,
	bookedSeatService BookedSeatServiceInterface,
	redisClient *redis.Client) BookingServiceInterface {
	return &BookingService{
		bookingRepository: bookingRepository,
		movieService:      movieService,
		bookedSeatService: bookedSeatService,
		redisClient:       redisClient,
	}
}

func (b BookingService) CreateBooking(request request_models.CreateBookingRequest, accountID int) (models.Booking, error) {
	// Check if booking seat is available
	isSeatsAvailable, err := b.bookedSeatService.IsSeatsAvailable(request.SeatID, request.SlotID)

	if err != nil {
		log.Printf("Error checking seat availability: %v", err)
		return models.Booking{}, fmt.Errorf("error checking seat availability: %w", err)
	}

	// if seat is not available return error
	if !isSeatsAvailable {
		return models.Booking{}, fmt.Errorf("seats are not available")
	}

	// Get all booked seats to parse into booking
	bookedSeats, err := b.bookedSeatService.FindAllBookedSeatWithSeatIDs(request.SeatID)

	booking := models.Booking{
		AccountID:   uint(accountID),
		SlotID:      uint(request.SlotID),
		IsBooked:    "ON_HOLD",
		BookingTime: pgtype.Timestamp{Time: time.Now(), Valid: true},
		DueTime:     pgtype.Timestamp{Time: time.Now().Add(time.Minute * 10), Valid: true},
		BookedSeats: bookedSeats,
	}

	bookingResult, err := b.bookingRepository.CreateBooking(booking)

	if err != nil {
		log.Printf("Error creating booking: %v", err)
		return models.Booking{}, fmt.Errorf("error creating booking: %w", err)
	}

	return bookingResult, nil
}

func (b BookingService) GetBookingByID(bookingID int) (models.Booking, error) {

	booking, err := b.bookingRepository.GetBookingById(bookingID)

	if err != nil {
		log.Printf("error fetching booking with ID %d: %v", bookingID, err)
		return models.Booking{}, fmt.Errorf("error fetching booking: %w", err)
	}

	return booking, nil
}

func (b BookingService) GetAllBookingsByAccountID(accountID int, page int, pageSize int) ([]models.Booking, error) {

	panic("implement me")
}

func (b BookingService) GetAllBookingsBySlotID(slotID int, page int, pageSize int) ([]models.Booking, error) {
	bookings, err := b.bookingRepository.GetAllBookingsBySlotID(slotID, page, pageSize)

	if err != nil {
		log.Printf("Error fetching bookings by slot ID %d: %v", slotID, err)
		return nil, fmt.Errorf("error fetching bookings by slot ID: %w", err)
	}

	return bookings, nil
}

func (b BookingService) UpdateBookingByID(bookingID int, status string) (models.Booking, error) {

	booking, err := b.bookingRepository.GetBookingById(bookingID)

	if err != nil {
		log.Printf("Error fetching booking with ID %d: %v", bookingID, err)
		return models.Booking{}, fmt.Errorf("error fetching booking: %w", err)
	}

	booking.IsBooked = status

	updateErr := b.bookingRepository.UpdateBooking(booking)

	if updateErr != nil {
		log.Printf("Error updating booking with ID %d: %v", bookingID, err)
		return models.Booking{}, fmt.Errorf("error updating booking: %w", err)
	}

	return booking, nil
}

func (b BookingService) CancelBookingByID(bookingID int) error {
	bookedSeats, err := b.bookedSeatService.FindAllBookedSeatWithBookingId(bookingID)
	if err != nil {
		log.Printf("Error fetching booking with ID %d: %v", bookingID, err)
		return fmt.Errorf("error fetching booking: %w", err)
	}

	for _, seat := range bookedSeats {
		seat.Status = "CANCELED"

		if err := b.bookedSeatService.UpdateBookedSeat(seat); err != nil {
			log.Printf("Error updating booked seat with ID %d: %v", seat.ID, err)
			return fmt.Errorf("error updating booked seat: %w", err)
		}
	}

	return nil
}
