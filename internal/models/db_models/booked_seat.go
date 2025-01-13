package models

import "gorm.io/gorm"

type BookedSeat struct {
	gorm.Model
	BookingID uint `json:"booking_id"`
	SeatID    uint `json:"seat_id"`
	SlotID    uint `json:"slot_id"`
	IsBooked  bool `json:"is_booked"`
}
