package request_models

type UpdateRoomRequest struct {
	RoomID    uint   `json:"room_id"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	TheaterID uint   `json:"theater_id"`
}
