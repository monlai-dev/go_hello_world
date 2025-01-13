package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	TheaterID uint   `json:"theater_id"`
	Slots     []Slot `json:"slot"`
	Seats     []Seat `json:"seat"`
}
