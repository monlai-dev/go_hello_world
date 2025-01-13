package request_models

type CreateRoomRequest struct {
	Name      string `json:"name,required"`
	Capacity  int    `json:"capacity,required"`
	TheaterID uint   `json:"theater_id,required"`
}
