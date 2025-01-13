package services

import (
	models "webapp/internal/models/db_models"
	"webapp/internal/models/request_models"
)

type RoomServiceInterface interface {
	FindAllRoomsByTheaterID(theaterId int, page int, pageSize int) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
	CreateRoom(room request_models.CreateRoomRequest) (models.Room, error)
	UpdateRoom(room request_models.UpdateRoomRequest) error
	DeleteRoom(room models.Room) error
}
