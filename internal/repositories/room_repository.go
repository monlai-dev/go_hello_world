package repositories

import models "webapp/internal/models/db_models"

type RoomRepositoryInterface interface {
	FindAllRoomsByTheaterID(theaterId int) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
	CreateRoom(room models.Room) (models.Room, error)
	UpdateRoom(room models.Room) error
	DeleteRoom(room models.Room) error
}
