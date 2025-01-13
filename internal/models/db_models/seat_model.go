package models

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	Name      string `json:"name"`
	RoomID    uint   `json:"room_id"`
	IsWorking bool   `json:"is_working"`
}
