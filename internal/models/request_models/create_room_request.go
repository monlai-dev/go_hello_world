package request_models

type CreateRoomRequest struct {
	Name      string `json:"name,required"`
	Capacity  int    `json:"capacity,required"`
	TheaterID uint   `json:"theater_id,required"`
}

type UpdateRoomRequest struct {
	RoomID    uint   `json:"room_id"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	TheaterID uint   `json:"theater_id"`
}
