package models

import (
	"errors"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserName  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	AddressId *uint     `json:"address_id,omitempty"`
	Address   Address   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
	Bookings  []Booking `json:"bookings"`
}

func (account *Account) BeforeCreate(tx *gorm.DB) error {
	requiredFields := map[string]string{
		"Username": account.UserName,
		"Password": account.Password,
		"Email":    account.Email,
		"Phone":    account.Phone,
		"Role":     account.Role,
	}

	for field, value := range requiredFields {
		if value == "" {
			return errors.New(field + " is required")
		}
	}

	return nil
}
