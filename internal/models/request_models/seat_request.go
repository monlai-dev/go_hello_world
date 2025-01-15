package request_models

type CreateSeatRequest struct {
	RoomID uint `json:"room_id" validate:"required,gt=0"`
	Row    int  `json:"row" validate:"required,gt=0"`
}

type UpdateSeatRequest struct {
	RoomID uint   `json:"room_id"`
	Name   string `json:"name"`
}

func (r *CreateSeatRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateSeatRequest) Validate() error {
	return validate.Struct(r)
}
