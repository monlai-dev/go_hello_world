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
