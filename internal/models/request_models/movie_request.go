package request_models

type CreateMovieRequest struct {
	Title       string `json:"title" validate:"required,ne=0"`
	Description string `json:"description" validate:"required,ne=0"`
	Duration    int    `json:"duration" validate:"required,gt=0"`
}

type UpdateMovieRequest struct {
	MovieID     int    `json:"movie_id" validate:"required,gt=0"`
	Title       string `json:"title" validate:"required,ne=0"`
	Description string `json:"description" validate:"required,ne=0"`
	Duration    int    `json:"duration" validate:"required,gt=0"`
}

type DeleteMovieRequest struct {
	MovieID int `json:"movie_id" validate:"required,gt=0"`
}

func (r *CreateMovieRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateMovieRequest) Validate() error {
	return validate.Struct(r)
}

func (r *DeleteMovieRequest) Validate() error {
	return validate.Struct(r)
}
