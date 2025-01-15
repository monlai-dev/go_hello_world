package request_models

type CreateMovieRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
}

type UpdateMovieRequest struct {
	MovieID     int    `json:"movie_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
}

type DeleteMovieRequest struct {
	MovieID int `json:"movie_id"`
}
