package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webapp/internal/models/request_models"
	"webapp/internal/services"
)

func CreateBookingHandler(bookingService services.BookingServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request_models.CreateBookingRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		if err := request.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		booking, err := bookingService.CreateBooking(request, 14)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusCreated, responseSuccess("Booking created successfully", []interface{}{booking}))
	}
}

func ConfirmBookingHandler(bookingService services.BookingServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingID, err := strconv.Atoi(c.Param("bookingID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Invalid booking ID"))
			return
		}

		err = bookingService.ConfirmBookingByID(bookingID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Booking confirmed successfully", nil))
	}
}
