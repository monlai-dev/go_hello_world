package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webapp/internal/services"
)

func GetBookedSeatsHandler(bookedSeatService services.BookedSeatServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {

		slotId, _ := strconv.Atoi(c.Param("slotId"))

		bookedSeats, err := bookedSeatService.FindAllBookedSeatsBySlotID(slotId, 0, 0)

		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Booked seats retrieved successfully", []interface{}{bookedSeats}))
	}
}
