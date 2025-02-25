package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v5/pgtype"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"time"
	"webapp/internal/infrastructure/rabbitMq"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
	"webapp/internal/repositories"
)

type BookingCache struct {
	BookingID int       `json:"booking_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type EmailRequest struct {
	Subject string `json:"subject"`
	Email   string `json:"email"`
	Body    string `json:"body"`
}

type BookingService struct {
	bookingRepository repositories.BookingRepositoryInterface
	movieService      MovieServiceInterface
	bookedSeatService BookedSeatServiceInterface
	redisClient       *redis.Client
	seatService       SeatServiceInterface
	slotService       SlotServiceInterface
	cronJobService    *CronJobService
	accountService    AccountServiceInterface
	rabbitClient      *rabbitMq.RabbitMq
}

func NewBookingService(
	bookingRepository repositories.BookingRepositoryInterface,
	movieService MovieServiceInterface,
	bookedSeatService BookedSeatServiceInterface,
	redisClient *redis.Client,
	seatService SeatServiceInterface,
	slotService SlotServiceInterface,
	cronjobService *CronJobService,
	accountService AccountServiceInterface,
	rabbitClient *rabbitMq.RabbitMq,
) BookingServiceInterface {
	return &BookingService{
		bookingRepository: bookingRepository,
		movieService:      movieService,
		bookedSeatService: bookedSeatService,
		redisClient:       redisClient,
		seatService:       seatService,
		slotService:       slotService,
		cronJobService:    cronjobService,
		accountService:    accountService,
		rabbitClient:      rabbitClient,
	}
}

func (b BookingService) CreateBooking(request request_models.CreateBookingRequest, email string) (models.Booking, error) {
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

	account, err := b.accountService.GetAccountByEmail(email)
	if err != nil {
		log.Printf("Error fetching account with email %s: %v", email, err)
		return models.Booking{}, fmt.Errorf("error fetching account: %w", err)
	}

	booking := models.Booking{
		AccountID:   account.ID,
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

	cacheBookingData := BookingCache{
		BookingID: int(bookingResult.ID),
		StartTime: bookingResult.BookingTime.Time,
		EndTime:   bookingResult.DueTime.Time,
	}

	if redisErr := cacheBooking(cacheBookingData, b.redisClient); redisErr != nil {
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

	b.updateBookedSeatsStatus(bookedSeats, "CANCELED")

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

	if err := removeBookingFromCache(bookingID, b.redisClient); err != nil {
		log.Printf("Error removing booking with ID %d from cache: %v", bookingID, err)
		return fmt.Errorf("error removing booking from cache: %w", err)
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

func cacheBooking(booking BookingCache, redisClient *redis.Client) error {
	// Cache the booking
	redisKey := fmt.Sprintf("booking:%d", booking.BookingID)

	bookingBytes, err := json.Marshal(booking)

	if err != nil {
		log.Printf("Error caching booking with ID %d: %v", booking.BookingID, err)
		return fmt.Errorf("error caching booking: %w", err)
	}

	redisErr := redisClient.Set(context.Background(), redisKey, bookingBytes, time.Minute*15).Err()

	return redisErr
}

func getBookingFromCache(redisClient *redis.Client) ([]BookingCache, error) {
	// Get all bookings from cache
	redisKeys, err := redisClient.Keys(context.Background(), "booking:*").Result()

	if err != nil {
		log.Printf("Error fetching booking keys: %v", err)
		return nil, fmt.Errorf("error fetching booking keys: %w", err)
	}

	var bookings []BookingCache
	for _, key := range redisKeys {
		bookingBytes, err := redisClient.Get(context.Background(), key).Bytes()
		if err != nil {
			log.Printf("Error fetching booking with key %s: %v", key, err)
			return nil, fmt.Errorf("error fetching booking with key: %w", err)
		}

		var booking BookingCache
		if err := json.Unmarshal(bookingBytes, &booking); err != nil {
			log.Printf("Error unmarshalling booking with key %s: %v", key, err)
			return nil, fmt.Errorf("error unmarshalling booking: %w", err)
		}

		bookings = append(bookings, booking)
	}

	return bookings, nil

}

func (b BookingService) ScanExpiredBooking() error {
	expiredBookings, err := getBookingFromCache(b.redisClient)
	if err != nil {
		log.Printf("Error fetching expired bookings: %v", err)
		return fmt.Errorf("error fetching expired bookings: %w", err)
	}

	log.Printf("Expired bookings: %v", expiredBookings)

	for _, booking := range expiredBookings {
		if booking.EndTime.Before(time.Now()) {
			err := b.CancelBookingByID(booking.BookingID)
			if err != nil {
				log.Printf("Error cancelling booking with ID %d: %v", booking.BookingID, err)
				return fmt.Errorf("error cancelling booking: %w", err)
			}
		}
	}

	return nil
}

func (b BookingService) Scheduler() error {
	err := b.ScanExpiredBooking()
	if err != nil {
		log.Printf("Error scanning expired booking: %v", err)
		return fmt.Errorf("error scanning expired booking: %w", err)
	}
	return nil
}

func removeBookingFromCache(bookingID int, redisClient *redis.Client) error {
	redisKey := fmt.Sprintf("booking:%d", bookingID)
	redisErr := redisClient.Del(context.Background(), redisKey).Err()
	if redisErr != nil {
		log.Printf("Error deleting booking with ID %d: %v", bookingID, redisErr)
		return fmt.Errorf("error deleting booking: %w", redisErr)
	}
	return nil
}

func (b BookingService) SendNotiEmail(data request_models.TestingEmailFormat) error {

	emailRequest := EmailRequest{
		Subject: data.Subject,
		Email:   data.Email,
		Body:    data.Body,
	}

	jsonBody, err := json.Marshal(emailRequest)
	if err != nil {
		log.Println("Error marshalling email request: ", err)
		return errors.New("error marshalling email request")
	}

	rabbitErr := b.rabbitClient.Publish(b.rabbitClient, "email_exchange", "email", amqp.Publishing{
		ContentType:  "text/html",
		Body:         jsonBody,
		DeliveryMode: amqp.Persistent,
	})

	if rabbitErr != nil {
		log.Println("Error publishing email request: ", rabbitErr)
		return errors.New("error publishing email request")
	}

	return nil
}
