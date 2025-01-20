package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webapp/internal/models/request_models"
	"webapp/internal/services"
)

type SeatResponse struct {
	ID int `json:"id"`
}

func CreateSeatHandler(seatService services.SeatServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request_models.CreateSeatRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		_, err := seatService.AutoImportSeatWithRow(req.RoomID, req.Row)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusCreated, responseSuccess("Seat created successfully", nil))
	}
}
