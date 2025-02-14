package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webapp/internal/infrastructure/rabbitMq"
	"webapp/internal/models/request_models"
	"webapp/internal/services"
)

type BookingResponse struct {
	ID         int     `json:"id"`
	SlotID     int     `json:"slot_id"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"total_price"`
}

type BookingTestingRequest struct {
	Email   string `json:"email"`
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

func CreateBookingHandler(bookingService services.BookingServiceInterface, rabbitClient *rabbitMq.RabbitMq) gin.HandlerFunc {
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

		booking, err := bookingService.CreateBooking(request, c.GetString("email"))
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusCreated, responseSuccess("Booking created successfully", []interface{}{BookingResponse{
			ID:         int(booking.ID),
			SlotID:     int(booking.SlotID),
			Status:     booking.IsBooked,
			TotalPrice: booking.TotalPrice,
		}}))
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

func TestingRabbitMq(bookingService services.BookingServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request BookingTestingRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		err := bookingService.SendNotiEmail(request.Subject, request.Email, request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("RabbitMQ testing successfully", nil))
	}
}
