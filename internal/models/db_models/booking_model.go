package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	AccountID   uint             `json:"account_id"`
	SlotID      uint             `json:"slot_id"`
	IsBooked    string           `json:"is_booked"`
	BookingTime pgtype.Timestamp `json:"booking_time"`
	DueTime     pgtype.Timestamp `json:"due_time"`
	TotalPrice  float64          `json:"total_price"`
	BookedSeats []BookedSeat     `json:"booked_seats"`
}
