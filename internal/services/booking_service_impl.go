package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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
	seatService       SeatServiceInterface
	slotService       SlotServiceInterface
}

func NewBookingService(
	bookingRepository repositories.BookingRepositoryInterface,
	movieService MovieServiceInterface,
	bookedSeatService BookedSeatServiceInterface,
	redisClient *redis.Client,
	seatService SeatServiceInterface,
	slotService SlotServiceInterface) BookingServiceInterface {
	return &BookingService{
		bookingRepository: bookingRepository,
		movieService:      movieService,
		bookedSeatService: bookedSeatService,
		redisClient:       redisClient,
		seatService:       seatService,
		slotService:       slotService,
	}
}

func (b BookingService) CreateBooking(request request_models.CreateBookingRequest, accountID int) (models.Booking, error) {
	// Check if booking seat is available
	isSeatsAvailable, err := b.bookedSeatService.IsSeatsAvailable(request.SeatID, request.SlotID)

	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		log.Printf("Error checking seat availability: %v", err)
		return models.Booking{}, fmt.Errorf("error checking seat availability: %w", err)
	}

	// if seat is not available return error
	if !isSeatsAvailable {
		return models.Booking{}, fmt.Errorf("seats are not available")
	}

	// Get all booked seats to parse into booking
	bookedSeats, err := b.seatService.GetSeatByIdList(request.SeatID)

	slot, err := b.slotService.GetSlotByID(request.SlotID)
	if err != nil {
		log.Printf("Error fetching slot with ID %d: %v", request.SlotID, err)
		return models.Booking{}, fmt.Errorf("slot with id not found %d", request.SlotID)
	}

	// integrate over the booked seats and set the status to ON_HOLD
	var bookedSeat []models.BookedSeat
	for _, val := range bookedSeats {
		bookedSeat = append(bookedSeat, models.BookedSeat{
			SeatID: val.ID,
			Status: "ON_HOLD",
			SlotID: slot.ID,
		})
	}

	booking := models.Booking{
		AccountID:   uint(accountID),
		SlotID:      slot.ID,
		IsBooked:    "ON_HOLD",
		BookingTime: pgtype.Timestamp{Time: time.Now(), Valid: true},
		DueTime:     pgtype.Timestamp{Time: time.Now().Add(time.Minute * 10), Valid: true},
		BookedSeats: bookedSeat,
		TotalPrice:  float64(len(bookedSeat)) * slot.Price,
	}

	bookingResult, err := b.bookingRepository.CreateBooking(booking)

	if err != nil {
		log.Printf("Error creating booking: %v", err)
		return models.Booking{}, fmt.Errorf("error creating booking: %w", err)
	}

	if redisErr := cacheBooking(bookingResult, b.redisClient); redisErr != nil {
		log.Printf("Error caching booking with ID %d: %v", bookingResult.ID, redisErr)
		return models.Booking{}, fmt.Errorf("error caching booking: %w", redisErr)
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
		log.Printf("error fetching booking with ID %d: %v", bookingID, err)
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
		log.Printf("error fetching booking with ID %d: %v", bookingID, err)
		return fmt.Errorf("error fetching booking: %w", err)
	}

	for _, seat := range bookedSeats {
		seat.Status = "CANCELED"
	}

	if err := b.bookedSeatService.UpdateBookedSeat(bookedSeats); err != nil {
		log.Printf("Error updating booked seat with ID %d: %v", bookingID, err)
		return fmt.Errorf("error updating booked seat: %w", err)
	}

	return nil
}

func (b BookingService) ConfirmBookingByID(bookingID int) error {
	booking, err := b.fetchBookingByID(bookingID)
	if err != nil {
		return err
	}

	bookedSeats, err := b.fetchBookedSeatsByBookingID(bookingID)
	if err != nil {
		return err
	}

	b.updateBookingStatus(&booking, "BOOKED")
	b.updateBookedSeatsStatus(bookedSeats, "BOOKED")

	if err := b.bookedSeatService.UpdateBookedSeat(bookedSeats); err != nil {
		log.Printf("Error updating booked seat with ID %d: %v", bookingID, err)
		return fmt.Errorf("error updating booked seat: %w", err)
	}

	if err := b.bookingRepository.UpdateBooking(booking); err != nil {
		log.Printf("Error updating booking with ID %d: %v", bookingID, err)
		return fmt.Errorf("error updating booking: %w", err)
	}

	return nil
}

func (b BookingService) fetchBookingByID(bookingID int) (models.Booking, error) {
	booking, err := b.bookingRepository.GetBookingById(bookingID)
	if err != nil {
		log.Printf("Error fetching booking with ID %d: %v", bookingID, err)
		return models.Booking{}, fmt.Errorf("error fetching booking: %w", err)
	}
	return booking, nil
}

func (b BookingService) fetchBookedSeatsByBookingID(bookingID int) ([]models.BookedSeat, error) {
	bookedSeats, err := b.bookedSeatService.FindAllBookedSeatWithBookingId(bookingID)
	if err != nil {
		log.Printf("Error fetching booking with ID %d: %v", bookingID, err)
		return nil, fmt.Errorf("error fetching booking: %w", err)
	}
	return bookedSeats, nil
}

func (b BookingService) updateBookingStatus(booking *models.Booking, status string) {
	booking.IsBooked = status
}

func (b BookingService) updateBookedSeatsStatus(bookedSeats []models.BookedSeat, status string) {
	for i := range bookedSeats {
		bookedSeats[i].Status = status
	}
}

func cacheBooking(booking models.Booking, redisClient *redis.Client) error {
	// Cache the booking
	redisKey := fmt.Sprintf("booking:%d", booking.ID)

	bookingBytes, err := json.Marshal(booking)

	if err != nil {
		log.Printf("Error caching booking with ID %d: %v", booking.ID, err)
		return fmt.Errorf("error caching booking: %w", err)
	}

	redisErr := redisClient.Set(context.Background(), redisKey, bookingBytes, time.Minute*10).Err()

	return redisErr
}
