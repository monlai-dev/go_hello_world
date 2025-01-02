package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

