package services

import (
	"fmt"
	"log"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
	"webapp/internal/repositories"
)

type SeatService struct {
	seatRepository repositories.SeatRepositoryInterface
	roomService    RoomServiceInterface
}

func NewSeatService(seatRepository repositories.SeatRepository, roomService RoomServiceInterface) SeatServiceInterface {
	return &SeatService{
		seatRepository: seatRepository,
		roomService:    roomService,
	}
}

func (s SeatService) GetSeatByID(id int) (models.Seat, error) {

	seat, err := s.seatRepository.FindSeatById(id)

	if err != nil {
		return models.Seat{}, fmt.Errorf("error fetching seat: %v", err)
	}

	return seat, nil
}

func (s SeatService) CreateListOfSeats(seats []models.Seat) ([]models.Seat, error) {

	createdSeats, err := s.seatRepository.CreateListOfSeats(seats)

	if err != nil {
		return []models.Seat{}, fmt.Errorf("error creating list of seats: %v", err)
	}

	return createdSeats, nil
}

func (s SeatService) GetAllSeatsByRoomID(page int, pageSize int, roomID int) ([]models.Seat, error) {

	seats, err := s.seatRepository.GetAllSeatsByRoomID(page, pageSize, roomID)

	if err != nil {
		return []models.Seat{}, fmt.Errorf("error fetching seats: %v", err)
	}

	return seats, nil
}

func (s SeatService) UpdateSeatByID(seatID int, request request_models.UpdateSeatRequest) error {

	seat, err := s.GetSeatByID(seatID)
	if err != nil {
		return fmt.Errorf("error fetching seat: %v", err)
	}

	seat.RoomID = request.RoomID
	seat.Name = request.Name

	if err := s.seatRepository.UpdateSeat(seat); err != nil {
		log.Printf("error updating seat: %v", err)
		return fmt.Errorf("error updating seat: %v", err)
	}

	return nil
}

func (s SeatService) DisableAndEnableSeat(seatID int) error {

	seat, err := s.GetSeatByID(seatID)
	if err != nil {
		return fmt.Errorf("error fetching seat: %v", err)
	}

	seat.IsWorking = !seat.IsWorking

	if err := s.seatRepository.UpdateSeat(seat); err != nil {
		log.Printf("error updating seat: %v", err)
		return fmt.Errorf("error updating seat: %v", err)
	}

	return nil
}

func (s SeatService) AutoImportSeatWithRow(roomID int, row int) ([]models.Seat, error) {
	room, err := s.roomService.GetRoomByID(roomID)
	if err != nil {
		return nil, fmt.Errorf("error fetching room: %v", err)
	}

	roomCapacity := room.Capacity
	cols := (roomCapacity + row - 1) / row // Calculate columns, rounding up

	seats := generateSeats(roomCapacity, row, cols, room.ID)

	createdSeats, err := s.seatRepository.CreateListOfSeats(seats)
	if err != nil {
		log.Printf("error creating list of seats: %v", err)
		return nil, fmt.Errorf("error creating list of seats: %v", err)
	}

	return createdSeats, nil
}

func generateSeats(roomCapacity, row, cols int, roomID uint) []models.Seat {
	var seats []models.Seat
	for i := 0; i < row; i++ {
		for j := 1; j <= cols && len(seats) < roomCapacity; j++ {
			seatName := fmt.Sprintf("%c%d", 'A'+i, j) // Generate seat name
			seats = append(seats, models.Seat{
				Name:      seatName,
				IsWorking: true,
				RoomID:    roomID})
		}
	}
	return seats
}
