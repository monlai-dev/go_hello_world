package services

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
	"webapp/internal/repositories"
)

type SlotService struct {
	slotRepository repositories.SlotRepositoryInterface
}

func NewSlotService(slotRepository repositories.SlotRepositoryInterface) SlotServiceInterface {
	return &SlotService{
		slotRepository: slotRepository,
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
	slots, err := s.GetSlotByRoomIDAndTime(int(createSlotRequest.RoomID), createSlotRequest.StartTime, createSlotRequest.EndTime)
	if err != nil {
		return models.Slot{}, fmt.Errorf("error fetching slots: %w", err)
	}

	available, err := isRequestTimeAvailable(createSlotRequest.StartTime, createSlotRequest.EndTime, slots)
	if err != nil {
		return models.Slot{}, err
	}

	if !available {
		return models.Slot{}, fmt.Errorf("slot is not available")
	}

	slotModel := models.Slot{
		RoomID:    createSlotRequest.RoomID,
		MovieID:   createSlotRequest.MovieID,
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
		if (startTime.Time.After(slot.StartTime.Time) && startTime.Time.Before(slot.EndTime.Time)) || (endTime.Time.After(slot.StartTime.Time) && endTime.Time.Before(slot.EndTime.Time)) {
			return false, fmt.Errorf("slot is not available")
		}
	}

	return true, nil
}
