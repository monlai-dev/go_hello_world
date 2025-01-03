package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	AccountID uint   `json:"account_id"`
}
