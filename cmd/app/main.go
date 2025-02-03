// filepath: /c:/Go_Tutorial/main.go
package main

import (
	"github.com/joho/godotenv"
	"log"
	"webapp/internal/infrastructure/cache"
	"webapp/internal/infrastructure/database"
	models "webapp/internal/models/db_models"
)

func init() {
	err := godotenv.Load()
	database.ConnectDb()
	cache.ConnectRedis()
	database.DB.AutoMigrate(&models.Account{}, &models.Address{}, &models.Theater{}, &models.Movie{}, &models.Room{}, &models.Slot{}, &models.Seat{}, &models.BookedSeat{}, &models.Booking{})
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {

	r, err := InitializeApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	if err := r.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
