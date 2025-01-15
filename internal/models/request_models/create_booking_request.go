package request_models

type CreateBookingRequest struct {
	SeatID []int `json:"seat_id"`
	SlotID int   `json:"slot_id"`
}
