package repositories

import (
	"gorm.io/gorm"
	models "webapp/internal/models/db_models"
)

type TheaterRepository struct {
	DB *gorm.DB
}

func NewTheaterRepository(db *gorm.DB) *TheaterRepository {
	return &TheaterRepository{DB: db}
}

func (r *TheaterRepository) FindAll() ([]models.Theater, error) {
	var theaters []models.Theater

	if err := r.DB.Find(&theaters).Error; err != nil {
		return nil, err
	}

	if len(theaters) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return theaters, nil
}

func (r *TheaterRepository) FindById(id int) (models.Theater, error) {
	var theater models.Theater

	if err := r.DB.First(&theater, id).Error; err != nil {
		return models.Theater{}, err
	}

	return theater, nil
}

func (r *TheaterRepository) Create(theater models.Theater) (models.Theater, error) {

	if err := r.DB.Create(&theater).Error; err != nil {
		return models.Theater{}, err
	}

	return theater, nil
}

func (r *TheaterRepository) Update(theater models.Theater) error {

	if err := r.DB.Save(&theater).Error; err != nil {
		return err
	}

	return nil
}

func (r *TheaterRepository) Delete(theater models.Theater) error {

	if err := r.DB.Delete(&theater).Error; err != nil {
		return err
	}

	return nil
}
