package models

import (
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
	Address   Address    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
}
