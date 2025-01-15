package request_models

type CreateSeatRequest struct {
}

type UpdateSeatRequest struct {
	RoomID uint   `json:"room_id"`
	Name   string `json:"name"`
}
