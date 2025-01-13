package models

import "gorm.io/gorm"

type Theater struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Rooms   []Room `json:"rooms"`
}
