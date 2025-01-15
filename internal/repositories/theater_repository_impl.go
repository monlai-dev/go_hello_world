package repositories

import (
	"gorm.io/gorm"
	models "webapp/internal/models/db_models"
)

type TheaterRepository struct {
	DB *gorm.DB
}

func NewTheaterRepository(db *gorm.DB) TheaterRepositoryInterface {
	return &TheaterRepository{DB: db}
}

func (r *TheaterRepository) GetAllTheaters() ([]models.Theater, error) {

	var theaters []models.Theater
	if err := r.DB.Preload("Rooms").Find(&theaters).Error; err != nil {
		return nil, err
	}

	return theaters, nil
}

func (r *TheaterRepository) GetTheaterById(id int) (models.Theater, error) {

	var theater models.Theater
	if err := r.DB.First(&theater, id).Error; err != nil {
		return models.Theater{}, err
	}

	return theater, nil
}

func (r *TheaterRepository) CreateTheater(theater models.Theater) (models.Theater, error) {

	if err := r.DB.Create(&theater).Error; err != nil {
		return models.Theater{}, err
	}

	return theater, nil
}

func (r *TheaterRepository) UpdateTheater(theater models.Theater) error {
	//TODO implement me
	panic("implement me")
}

func (r *TheaterRepository) DeleteTheater(theater models.Theater) error {
	//TODO implement me
	panic("implement me")
}
