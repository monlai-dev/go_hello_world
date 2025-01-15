package request_models

type CreateBookingRequest struct {
	SeatID []int `json:"seat_id" validate:"required,gt=0,dive,gt=0"`
	SlotID int   `json:"slot_id" validate:"required,gt=0"`
}

func (r *CreateBookingRequest) Validate() error {
	return validate.Struct(r)
}
