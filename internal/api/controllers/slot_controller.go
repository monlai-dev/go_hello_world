package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webapp/internal/models/request_models"
	"webapp/internal/services"
)

func CreateSlotHandler(slotService services.SlotServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request_models.CreateSlotRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		if err := validate.Struct(&req); err != nil {
			c.JSON(http.StatusOK, responseError(err.Error()))
			return
		}

		createdSlot, err := slotService.CreateSlot(req)
		if err != nil {
			c.JSON(http.StatusOK, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Slot created successfully", []interface{}{createdSlot}))
	}
}

func GetAllSlotsByMovieIdHandler(slotService services.SlotServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId, _ := strconv.Atoi(c.Param("movieId"))
		slots, err := slotService.FindAllSlotByMovieID(movieId, 1, 10)
		if err != nil {
			c.JSON(http.StatusOK, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Slots fetched successfully", []interface{}{slots}))
	}
}
