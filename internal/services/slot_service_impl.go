package services

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
	"webapp/internal/repositories"
)

type SlotService struct {
	slotRepository repositories.SlotRepositoryInterface
	roomService    RoomServiceInterface
	movieService   MovieServiceInterface
}

func NewSlotService(slotRepository repositories.SlotRepositoryInterface, roomService RoomServiceInterface, movieService MovieServiceInterface) SlotServiceInterface {
	return &SlotService{
		slotRepository: slotRepository,
		roomService:    roomService,
		movieService:   movieService,
	}
}

func (s SlotService) FindAllSlotsByRoomID(roomId int, page int, pageSize int) ([]models.Slot, error) {

	slots, err := s.slotRepository.GetSlotsByRoomId(roomId)

	if err != nil {
		return nil, fmt.Errorf("error fetching slots: %w", err)
	}

	return slots, nil
}

func (s SlotService) FindAllSlotByMovieID(movieId int, page int, pageSize int) ([]models.Slot, error) {

	slots, err := s.slotRepository.GetSlotsByMovieId(movieId)

	if err != nil {
		return nil, fmt.Errorf("error fetching slots: %w", err)
	}

	return slots, nil
}

func (s SlotService) FindAllSlotByMovieIDAndBetweenDates(movieId int, startDate pgtype.Timestamp, endDate pgtype.Timestamp, page int, pageSize int) ([]models.Slot, error) {

	slots, err := s.slotRepository.GetSlotByMovieIdAndBetweenDates(movieId, startDate, endDate)

	if err != nil {
		return nil, fmt.Errorf("error fetching slots: %w", err)
	}

	return slots, nil
}

func (s SlotService) CreateSlot(createSlotRequest request_models.CreateSlotRequest) (models.Slot, error) {
	slots, _ := s.GetSlotByRoomIDAndTime(int(createSlotRequest.RoomID), createSlotRequest.StartTime, createSlotRequest.EndTime)

	// Validate room existence
	room, err := s.roomService.GetRoomByID(int(createSlotRequest.RoomID))
	if err != nil {
		return models.Slot{}, fmt.Errorf("error fetching room: %w", err)
	}

	// Validate movie existence
	movie, err := s.movieService.GetMovieByID(int(createSlotRequest.MovieID))
	if err != nil {
		return models.Slot{}, fmt.Errorf("error fetching movie: %w", err)
	}

	// Validate slot availability
	available, err := isRequestTimeAvailable(createSlotRequest.StartTime, createSlotRequest.EndTime, slots)
	if err != nil {
		return models.Slot{}, err
	}

	// If slot is not available, return error
	if !available {
		return models.Slot{}, fmt.Errorf("slot is not available")
	}

	slotModel := models.Slot{
		RoomID:    room.ID,
		MovieID:   movie.ID,
		StartTime: createSlotRequest.StartTime,
		EndTime:   createSlotRequest.EndTime,
	}

	slotCreated, err := s.slotRepository.CreateSlot(slotModel)
	if err != nil {
		return models.Slot{}, err
	}

	return slotCreated, nil
}

func (s SlotService) UpdateSlot(updateSlotRequest request_models.UpdateSlotRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s SlotService) DeleteSlot(slotId int) error {
	//TODO implement me
	panic("implement me")
}

func (s SlotService) GetSlotByID(slotId int) (models.Slot, error) {
	//TODO implement me
	panic("implement me")
}

func (s SlotService) GetSlotByRoomIDAndTime(roomId int, startTime pgtype.Timestamp, endTime pgtype.Timestamp) ([]models.Slot, error) {

	slots, err := s.slotRepository.GetSlotByRoomIdAndBetweenDates(roomId, startTime, endTime)

	if err != nil {
		return nil, fmt.Errorf("error fetching slots: %w", err)
	}

	return slots, nil
}

func (s SlotService) FindAllSlotBetweenDates(startDate pgtype.Timestamp, endDate pgtype.Timestamp, page int, pageSize int) ([]models.Slot, error) {
	//TODO implement me
	panic("implement me")
}

func isRequestTimeAvailable(startTime pgtype.Timestamp, endTime pgtype.Timestamp, slots []models.Slot) (bool, error) {

	for _, slot := range slots {
		if startTime.Time.Before(slot.EndTime.Time) && endTime.Time.After(slot.StartTime.Time) {
			return false, fmt.Errorf("slot is not available")
		}
	}

	if startTime.Time.Before(time.Now()) || endTime.Time.Before(time.Now()) {
		return false, fmt.Errorf("start time is in the past or end time is in the past or both")
	}

	return true, nil
}
