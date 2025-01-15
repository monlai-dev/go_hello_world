package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webapp/internal/models/request_models"
	"webapp/internal/services"
)

func CreateTheaterHandler(theaterService services.TheaterServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request_models.TheaterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		createdTheater, err := theaterService.CreateTheater(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Theater created successfully", []interface{}{createdTheater}))
	}
}

func GetAllTheatersHandler(theaterService services.TheaterServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		theaters, err := theaterService.GetAllTheaters()

		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error getting theaters"))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Theaters retrieved successfully", []interface{}{theaters}))
	}
}
