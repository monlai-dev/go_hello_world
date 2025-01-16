package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webapp/internal/models/request_models"
	"webapp/internal/services"
)

func CreateMovieHandler(movieService services.MovieServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request_models.CreateMovieRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		createdMovie, err := movieService.CreateMovie(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Movie created successfully", []interface{}{createdMovie}))
	}
}

func UpdateMovieHandler(movieService services.MovieServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request_models.UpdateMovieRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		err := movieService.UpdateMovie(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Movie updated successfully", nil))
	}
}

func DeleteMovieHandler(movieService services.MovieServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieID, _ := strconv.Atoi(c.Param("id"))

		err := movieService.DeleteMovie(movieID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Movie deleted successfully", nil))
	}
}

func GetMovieByIDHandler(movieService services.MovieServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieID, _ := strconv.Atoi(c.Param("id"))

		movie, err := movieService.GetMovieByID(movieID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Movie retrieved successfully", []interface{}{movie}))
	}
}

func GetAllMoviesHandler(movieService services.MovieServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, _ := strconv.Atoi(c.Query("page"))
		pageSize, _ := strconv.Atoi(c.Query("page_size"))

		movies, err := movieService.GetAllMovies(page, pageSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Movies retrieved successfully", []interface{}{movies}))
	}
}
