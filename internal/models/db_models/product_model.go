package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Price       float64    `json:"price" binding:"required"`
	Quantity    int        `json:"quantity" binding:"required"`
	Images      []string   `json:"images"`
	CategoryID  uint       `json:"category_id" binding:"required"`
	Category    []Category `json:"category" binding:"required"`
}
