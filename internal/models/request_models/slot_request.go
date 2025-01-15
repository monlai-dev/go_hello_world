package request_models

import "github.com/jackc/pgx/v5/pgtype"

type CreateSlotRequest struct {
	MovieID   uint             `json:"movie_id" validate:"required,gt=0"`
	RoomID    uint             `json:"room_id" validate:"required,gt=0"`
	Price     float64          `json:"price" validate:"required,gt=0"`
	StartTime pgtype.Timestamp `json:"start_time" validate:"required"`
	EndTime   pgtype.Timestamp `json:"end_time" validate:"required"`
}

type UpdateSlotRequest struct {
	SlotID    uint             `json:"slot_id" validate:"required,gt=0"`
	StartTime pgtype.Timestamp `json:"start_time" validate:"required"`
	EndTime   pgtype.Timestamp `json:"end_time" validate:"required"`
	Price     float64          `json:"price" validate:"required,gt=0,float"`
	RoomID    uint             `json:"room_id" validate:"required,gt=0"`
	MovieID   uint             `json:"movie_id" validate:"required,gt=0"`
}

func (r *CreateSlotRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateSlotRequest) Validate() error {
	return validate.Struct(r)
}
