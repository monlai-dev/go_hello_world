package services

import (
	"fmt"
	"github.com/payOSHQ/payos-lib-golang"
	"log"
	"os"
)

type PaymentService struct {
	slotService       SlotServiceInterface
	bookingService    BookingServiceInterface
	seatService       SeatServiceInterface
	bookedSeatService BookedSeatServiceInterface
}

func NewPaymentService(slotService SlotServiceInterface, bookingService BookingServiceInterface, seatService SeatServiceInterface, bookedSeatService BookedSeatServiceInterface) PaymentServiceInterface {
	return &PaymentService{
		slotService:       slotService,
		bookingService:    bookingService,
		seatService:       seatService,
		bookedSeatService: bookedSeatService,
	}
}

func (p PaymentService) CreatePaymentLinkWithPayOsUsingBookingId(bookingId int) (string, error) {
	var items []payos.Item

	// Get booking by booking id
	booking, err := p.bookingService.GetBookingByID(bookingId)
	if err != nil {
		log.Printf("error fetching booking: %v", err)
		return "", fmt.Errorf("error fetching booking")
	}

	// Get all booked seats by booking id
	bookedSeats, err := p.bookedSeatService.FindAllBookedSeatWithBookingId(bookingId)
	if err != nil {
		log.Printf("error fetching booked seats with id: %d, error: %v", bookingId, err)
		return "", fmt.Errorf("error fetching booked seats")
	}

	// Get slot by slot id
	slot, err := p.slotService.GetSlotByID(int(booking.SlotID))
	if err != nil {
		log.Printf("error fetching slot with id: %d, error: %v", booking.SlotID, err)
		return "", fmt.Errorf("error fetching slot")
	}

	// Extract seat ids from booked seats
	var seatIds []int
	for _, seat := range bookedSeats {
		seatIds = append(seatIds, int(seat.SeatID))
	}

	// Get seats by seat ids
	seatList, err := p.seatService.GetSeatByIdList(seatIds)
	if err != nil {
		log.Printf("error fetching seats: %v", err)
		return "", fmt.Errorf("error fetching seats")
	}

	// Initialize payos
	if payOsErr := payos.Key(os.Getenv("CLIENT_ID"),
		os.Getenv("API_KEY"),
		os.Getenv("CHECK_SUM_KEY")); payOsErr != nil {
		log.Fatal(payOsErr)
	}

	// Create items for payment
	for _, seat := range seatList {
		item := payos.Item{
			Name:     seat.Name,
			Price:    int(slot.Price),
			Quantity: 1,
		}
		items = append(items, item)
	}

	// Create payment link
	body := payos.CheckoutRequestType{
		OrderCode:   int64(booking.ID) + 1000,
		Amount:      int(booking.TotalPrice),
		Items:       items,
		Description: "Thanh toán đơn hàng",
		CancelUrl:   "http://localhost:8080/cancel/",
		ReturnUrl:   "http://localhost:8080/success/",
	}

	// Create payment link
	data, err := payos.CreatePaymentLink(body)
	if err != nil {
		log.Printf("error creating payment link: %v", err)
	}

	return data.CheckoutUrl, nil
}
