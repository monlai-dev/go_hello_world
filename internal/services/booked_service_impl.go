package services

import (
	"fmt"
	"log"
	models "webapp/internal/models/db_models"
	"webapp/internal/repositories"
)

type BookedService struct {
	bookedSeatRepository repositories.BookedSeatRepositoryInterface
}

func NewBookedService(bookedSeatRepository repositories.BookedSeatRepositoryInterface) BookedSeatServiceInterface {
	return &BookedService{
		bookedSeatRepository: bookedSeatRepository,
	}
}

func (b BookedService) FindAllBookedSeatsBySlotID(slotId int, page int, pageSize int) ([]models.BookedSeat, error) {

	bookedSeats, err := b.bookedSeatRepository.GetBookedSeatBySlotId(slotId)

	if err != nil {
		log.Printf("Error while fetching booked seats by slot id: %v", err)
		return nil, fmt.Errorf("error while fetching booked seats by slot id: %w", err)
	}

	return bookedSeats, nil
}

func (b BookedService) CreateBookedSeat(seat []models.BookedSeat) ([]models.BookedSeat, error) {

	bookedSeats, err := b.bookedSeatRepository.CreateBooked(seat)

	if err != nil {
		log.Printf("Error while creating booked seat: %v", err)
		return nil, fmt.Errorf("error while creating booked seat: %w", err)
	}

	return bookedSeats, nil
}

func (b BookedService) UpdateBookedSeat(seat models.BookedSeat) error {
	//TODO implement me
	panic("implement me")
}

func (b BookedService) IsSeatsAvailable(seatIds []int, slotId int) (bool, error) {

	bookedSeats, err := b.bookedSeatRepository.GetBookedSeatBySlotIdAndStatus(
		slotId,
		[]string{"COMPLETED_PAYMENT", "ON_HOLD"})

	if err != nil {
		log.Printf("Error while fetching booked seats by slot id: %v", err)
		return false, fmt.Errorf("error while fetching booked seats by slot id: %w", err)
	}

	//integrate over the seatIds and check if the seat is already booked
	for _, seatId := range seatIds {
		for _, bookedSeat := range bookedSeats {
			if uint(seatId) == bookedSeat.SeatID {
				return false, nil
			}
		}
	}

	return true, nil
}

func (b BookedService) FindAllBookedSeatWithSeatIDs(seatIds []int) ([]models.BookedSeat, error) {

	bookedSeats, err := b.bookedSeatRepository.FindAllBookedSeatWithIds(seatIds)

	if err != nil {
		log.Printf("Error while fetching booked seats by seat ids: %v", err)
		return nil, fmt.Errorf("error while fetching booked seats by seat ids: %w", err)
	}

	return bookedSeats, nil
}

func (b BookedService) FindAllBookedSeatWithBookingId(bookingId int) ([]models.BookedSeat, error) {

	bookedSeats, err := b.bookedSeatRepository.GetAllBookedSeatByBookingId(bookingId)

	if err != nil {
		log.Printf("Error while fetching booked seats by booking id: %v", err)
		return nil, fmt.Errorf("error while fetching booked seats by booking id: %w", err)
	}

	return bookedSeats, nil
}
