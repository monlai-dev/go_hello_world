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

type BookingController struct {
	BookingService services.BookingServiceInterface
	RabbitClient   *rabbitMq.RabbitMq
}

func NewBookingController(bookingService services.BookingServiceInterface, rabbitClient *rabbitMq.RabbitMq) *BookingController {
	return &BookingController{
		BookingService: bookingService,
		RabbitClient:   rabbitClient,
	}
}

func (bc *BookingController) CreateBookingHandler(c *gin.Context) {

	var request request_models.CreateBookingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	booking, err := bc.BookingService.CreateBooking(request, c.GetString("email"))
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

func (bc *BookingController) ConfirmBookingHandler(c *gin.Context) {

	bookingID, err := strconv.Atoi(c.Param("bookingID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError("Invalid booking ID"))
		return
	}

	err = bc.BookingService.ConfirmBookingByID(bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("Booking confirmed successfully", nil))
}

func (bc *BookingController) TestingRabbitMq(c *gin.Context) {

	var request request_models.TestingEmailFormat
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	err := bc.BookingService.SendNotiEmail(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("RabbitMQ testing successfully", nil))

}
