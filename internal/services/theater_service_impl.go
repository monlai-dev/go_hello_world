package services

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
	models "webapp/internal/models/db_models"
	"webapp/internal/repositories"
)

type TheaterService struct {
	theaterRepo repositories.TheaterRepositoryInterface
	db          *gorm.DB
}

func NewTheaterService(theaterRepo repositories.TheaterRepositoryInterface, db *gorm.DB) TheaterServiceInterface {
	return &TheaterService{
		theaterRepo: theaterRepo,
		db:          db,
	}
}

func (t TheaterService) GetAllTheaters() ([]models.Theater, error) {

	var theaters []models.Theater
	theaters, err := t.theaterRepo.GetAllTheaters()

	if err != nil {
		log.Printf("error fetching theaters: %v", err)
		return nil, errors.New("error fetching theaters")
	}

	return theaters, nil
}

func (t TheaterService) GetTheaterById(id int) (models.Theater, error) {

	theater, err := t.theaterRepo.GetTheaterById(id)

	if err != nil {
		log.Printf("error fetching theater: %v", err)
		return models.Theater{}, errors.New("error fetching theater")
	}

	return theater, nil
}

func (t TheaterService) CreateTheater(theater models.Theater) (models.Theater, error) {

	theater, err := t.theaterRepo.CreateTheater(theater)

	if err != nil {
		log.Printf("error creating theater: %v", err)
		return models.Theater{}, errors.New("error creating theater")
	}

	return theater, nil
}

func (t TheaterService) UpdateTheater(theater models.Theater) error {

	err := t.theaterRepo.UpdateTheater(theater)

	if err != nil {
		log.Printf("error updating theater: %v", err)
		return errors.New("error updating theater")
	}

	return nil
}

func (t TheaterService) DeleteTheater(theater models.Theater) error {

	err := t.theaterRepo.DeleteTheater(theater)

	if err != nil {
		log.Printf("error deleting theater: %v, at: %v", err, time.Now())
		return errors.New("error deleting theater")
	}

	return nil
}
