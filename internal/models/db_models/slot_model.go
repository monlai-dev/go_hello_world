package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type Slot struct {
	gorm.Model
	StartTime pgtype.Timestamp `json:"start_time"`
	EndTime   pgtype.Timestamp `json:"end_time"`
	Price     float64          `json:"price"`
	RoomID    uint             `json:"room_id"`
	MovieID   uint             `json:"movie_id"`
	Bookings  []Booking        `json:"booking"`
}
