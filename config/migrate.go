package main

import (
	"github.com/joho/godotenv"
	"webapp/internal/infrastructure/database"
	"webapp/internal/models/db_models"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	database.ConnectDb()
}

func main() {
	err := database.DB.AutoMigrate(&models.Account{}, &models.Address{}, &models.Theater{}, &models.Movie{}, &models.Room{}, &models.Slot{}, &models.Seat{}, &models.BookedSeat{}, &models.Booking{})

	if err != nil {
		return
	}
}
