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

type SeatController struct {
	SeatService services.SeatServiceInterface
}

func NewSeatController(seatService services.SeatServiceInterface) *SeatController {
	return &SeatController{
		SeatService: seatService,
	}
}

func (sc *SeatController) CreateSeatHandler(c *gin.Context) {

	var req request_models.CreateSeatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	_, err := sc.SeatService.AutoImportSeatWithRow(int(req.RoomID), req.Row)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, responseSuccess("Seat created successfully", nil))
}
