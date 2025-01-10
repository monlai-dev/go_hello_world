package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	UserName  string     `json:"username,omitempty"`
	Password  string     `json:"password,omitempty"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Role      string     `json:"role"`
	AddressId *uint      `json:"address_id,omitempty"`
	Address   Address    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
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
