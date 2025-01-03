package main

import (
	"webapp/initializer"
	model "webapp/models/db_models"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	initializer.ConnectDb()
}

func main() {
	err := initializer.DB.AutoMigrate(&model.Account{}, &model.Address{})

	if err != nil {
		return
	}
}
