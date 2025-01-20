package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"webapp/internal/models/request_models"
	"webapp/internal/services"
)

type RoomResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	TheaterId int    `json:"theater_id"`
}

var validate = validator.New()

// CreateRoomHandler creates a new room
// CreateRoomHandler godoc
// @Summary Create a new room
// @Description Create a new room
// @Tags rooms
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param capacity body int true "Capacity"
// @Param theaterId body int true "TheaterId"
// @Success 200 {object} models.Room
// @Router /rooms [post]
func CreateRoomHandler(roomService services.RoomServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request_models.CreateRoomRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		if err := validate.Struct(&req); err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		createdRoom, err := roomService.CreateRoom(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Room created successfully", []interface{}{RoomResponse{
			ID:        int(createdRoom.ID),
			Name:      createdRoom.Name,
			Capacity:  createdRoom.Capacity,
			TheaterId: int(createdRoom.TheaterID),
		}}))
	}
}
