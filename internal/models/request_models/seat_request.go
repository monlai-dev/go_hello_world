package request_models

type CreateSeatRequest struct {
	RoomID uint `json:"room_id" validate:"required,gt=0"`
	Row    int  `json:"row" validate:"required,gt=0"`
}

type UpdateSeatRequest struct {
	RoomID uint   `json:"room_id"`
	Name   string `json:"name"`
}
