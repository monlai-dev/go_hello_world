package services

import (
	"fmt"
	"gorm.io/gorm"
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
	"webapp/internal/repositories"
)

type RoomService struct {
	db             *gorm.DB
	roomRepository repositories.RoomRepositoryInterface
}

func NewRoomService(db *gorm.DB, roomRepository repositories.RoomRepositoryInterface) RoomServiceInterface {
	return RoomService{
		db:             db,
		roomRepository: roomRepository,
	}
}

func (r RoomService) FindAllRoomsByTheaterID(theaterId int, page int, pageSize int) ([]models.Room, error) {

	rooms, err := r.roomRepository.FindAllRoomsByTheaterID(theaterId)

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r RoomService) GetRoomByID(id int) (models.Room, error) {

	room, err := r.roomRepository.GetRoomByID(id)

	if err != nil {
		return models.Room{}, err
	}

	return room, nil
}

func (r RoomService) CreateRoom(room request_models.CreateRoomRequest) (models.Room, error) {

	roomModel := models.Room{
		Name:      room.Name,
		TheaterID: room.TheaterID,
	}

	roomCreated, err := r.roomRepository.CreateRoom(roomModel)

	if err != nil {
		return models.Room{}, err
	}

	return roomCreated, nil
}

func (r RoomService) UpdateRoom(room request_models.UpdateRoomRequest) error {

	roomModel, err := r.roomRepository.GetRoomByID(int(room.RoomID))
	if err != nil {
		return err
	}

	if err := validateRoomData(room); err != nil {
		return err
	}

	updateRoomModel(&roomModel, room)

	if err := r.roomRepository.UpdateRoom(roomModel); err != nil {
		return err
	}

	return nil
}

func (r RoomService) DeleteRoom(roomId int) error {
	tx := r.db.Begin()

	result, err := r.roomRepository.GetRoomByID(roomId)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error getting room by id: %w", err)
	}

	if err := r.roomRepository.DeleteRoom(result); err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting room: %w", err)
	}

	return tx.Commit().Error
}

func validateRoomData(room request_models.UpdateRoomRequest) error {
	if room.Name == "" || room.Capacity <= 0 || room.TheaterID == 0 {
		return fmt.Errorf("invalid room data: %+v", room)
	}
	return nil
}

func updateRoomModel(roomModel *models.Room, room request_models.UpdateRoomRequest) {
	roomModel.Name = room.Name
	roomModel.Capacity = room.Capacity
	roomModel.TheaterID = room.TheaterID
}
