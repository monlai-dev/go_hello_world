package request_models

import "github.com/go-playground/validator/v10"

type CreateRoomRequest struct {
	Name      string `json:"name" validate:"required,ne=0"`
	Capacity  int    `json:"capacity" validate:"required,gt=0"`
	TheaterID uint   `json:"theater_id" validate:"required,gt=0"`
}

var validate = validator.New()

func (r *CreateRoomRequest) Validate() error {
	return validate.Struct(r)
}

type UpdateRoomRequest struct {
	RoomID    uint   `json:"room_id" validate:"required,gt=0"`
	Name      string `json:"name" validate:"required,ne=0"`
	Capacity  int    `json:"capacity" validate:"required,gt=0"`
	TheaterID uint   `json:"theater_id" validate:"required,gt=0"`
}

func (r *UpdateRoomRequest) Validate() error {
	return validate.Struct(r)
}
