package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	models "webapp/internal/models/db_models"
)

type RoomRepository struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepositoryInterface {
	return &RoomRepository{DB: db}
}

func (r *RoomRepository) FindAllRoomsByTheaterID(theaterId int) ([]models.Room, error) {

	var rooms []models.Room
	r.DB.Where("theater_id = ?", theaterId).Find(&rooms)

	if len(rooms) == 0 {
		log.Printf("Error getting rooms by theater id: %v", gorm.ErrRecordNotFound)
		return nil, gorm.ErrRecordNotFound
	}

	return rooms, nil
}

func (r *RoomRepository) GetRoomByID(id int) (models.Room, error) {

	var room models.Room
	err := r.DB.Where("id = ?", id).First(&room).Error

	if err != nil {
		log.Printf("Error getting room by id: %v", err)
		return models.Room{}, errors.New("room with with given id not found")
	}

	return room, nil
}

func (r *RoomRepository) CreateRoom(room models.Room) (models.Room, error) {

	if err := r.DB.Create(&room).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Error creating room: %v", err)
			return models.Room{}, errors.New("room with given id already exists")
		}
		return models.Room{}, err
	}

	return room, nil
}

func (r *RoomRepository) UpdateRoom(room models.Room) error {

	if err := r.DB.Save(&room).Error; err != nil {
		log.Printf("Error updating room: %v", err)
		return err
	}

	return nil
}

func (r *RoomRepository) DeleteRoom(room models.Room) error {

	if err := r.DB.Delete(&room).Error; err != nil {
		log.Printf("Error deleting room: %v", err)
		return err
	}

	return nil
}
