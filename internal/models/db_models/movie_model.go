package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `json:"title" gorm:"uniqueIndex"`
	Description string `json:"description"`
	MediaLink   string `json:"media_link"`
	Duration    int64  `json:"duration"`
	Slots       []Slot `json:"slot"`
}
