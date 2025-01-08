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
	UserName  string     `gorm:"unique" json:"username"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Role      string     `json:"role"`
	AddressId *uint      `json:"address_id"`
	Address   Address    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
}

func (account *Account) BeforeCreate(tx *gorm.DB) (err error) {

	if account.UserName == "" {
		return errors.New("Username is required")
	}

	if account.Password == "" {
		return errors.New("Password is required")
	}

	if account.Email == "" {
		return errors.New("Email is required")
	}

	if account.Phone == "" {
		return errors.New("Phone is required")
	}

	if account.Role == "" {
		return errors.New("Role is required")
	}

	return nil
}
