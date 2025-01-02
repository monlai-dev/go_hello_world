package main

import (
	"webapp/initializer"
	model "webapp/models"

	"github.com/joho/godotenv"
)

func init(){
	godotenv.Load()
	initializer.ConnectDb()
}

func main() {
	initializer.DB.AutoMigrate(&model.Student{}, &model.Address{})
}