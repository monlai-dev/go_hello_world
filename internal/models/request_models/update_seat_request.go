package request_models

type UpdateSeatRequest struct {
	RoomID uint   `json:"room_id"`
	Name   string `json:"name"`
}
