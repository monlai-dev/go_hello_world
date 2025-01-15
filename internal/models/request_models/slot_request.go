package request_models

import "github.com/jackc/pgx/v5/pgtype"

type CreateSlotRequest struct {
	MovieID   uint             `json:"movie_id"`
	RoomID    uint             `json:"room_id"`
	Price     float64          `json:"price"`
	StartTime pgtype.Timestamp `json:"start_time"`
	EndTime   pgtype.Timestamp `json:"end_time"`
}

type UpdateSlotRequest struct {
	SlotID    uint             `json:"slot_id"`
	StartTime pgtype.Timestamp `json:"start_time"`
	EndTime   pgtype.Timestamp `json:"end_time"`
	Price     float64          `json:"price"`
	RoomID    uint             `json:"room_id"`
	MovieID   uint             `json:"movie_id"`
}
