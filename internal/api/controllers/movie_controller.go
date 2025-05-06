package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webapp/internal/models/request_models"
	"webapp/internal/services"
)

type MovieResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
}

type MovieController struct {
	movieService services.MovieServiceInterface
}

func NewMovieController(movieService services.MovieServiceInterface) *MovieController {
	return &MovieController{
		movieService: movieService,
	}
}

func (mc *MovieController) CreateMovieHandler(c *gin.Context) {

	var req request_models.CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	createdMovie, err := mc.movieService.CreateMovie(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("Movie created successfully", []interface{}{MovieResponse{
		ID:          int(createdMovie.ID),
		Title:       createdMovie.Title,
		Description: createdMovie.Description,
		Duration:    int(createdMovie.Duration),
	}}))

}

func (mc *MovieController) UpdateMovieHandler(c *gin.Context) {

	var req request_models.UpdateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	err := mc.movieService.UpdateMovie(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("Movie updated successfully", nil))

}

func (mc *MovieController) DeleteMovieHandler(c *gin.Context) {

	movieID, _ := strconv.Atoi(c.Param("id"))

	err := mc.movieService.DeleteMovie(movieID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("Movie deleted successfully", nil))

}

func (mc *MovieController) GetMovieByIDHandler(c *gin.Context) {

	movieID, _ := strconv.Atoi(c.Param("id"))

	movie, err := mc.movieService.GetMovieByID(movieID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("Movie retrieved successfully", []interface{}{MovieResponse{
		ID:          int(movie.ID),
		Title:       movie.Title,
		Description: movie.Description,
		Duration:    int(movie.Duration),
	}}))
}

func (mc *MovieController) GetAllMoviesHandler(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	movies, err := mc.movieService.GetAllMovies(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	var movieList []MovieResponse

	for _, movie := range movies {
		movieList = append(movieList, MovieResponse{
			ID:          int(movie.ID),
			Title:       movie.Title,
			Description: movie.Description,
			Duration:    int(movie.Duration),
		})

		c.JSON(http.StatusOK, responseSuccess("Movies retrieved successfully", []interface{}{movieList}))
	}

}
