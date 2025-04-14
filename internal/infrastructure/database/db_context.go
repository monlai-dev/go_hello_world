package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

var dataBaseUrl *string

func InitPostgres() {
	var url string
	if os.Getenv("ENV") == "staging" {
		url = os.Getenv("RENDER_DATABASE_URL")
		dataBaseUrl = &url
		return
	}

	url = os.Getenv("DATABASE_URL")
	dataBaseUrl = &url
}

func ConnectDb() *gorm.DB {

	var err error
	dsn := *dataBaseUrl

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	return DB
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 5
	}

	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
