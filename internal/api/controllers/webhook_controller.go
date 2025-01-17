package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/payOSHQ/payos-lib-golang"
	"io"
	"log"
	"net/http"
	"strconv"
	"webapp/internal/services"
)

func WebhookHandler(bookingService services.BookingServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := payos.Key("e47ad6bc-40f4-4fc9-a723-cd8ddae9e3a8",
			"0ca2d2e9-d6fa-4c13-9885-848080846bd7",
			"1589dd37d76b03b0968bcd445dd4ed0e38fa63b0d5811bba8dd72b7ea88f95c0"); err != nil {
			log.Fatalf("Error setting payos key: %v", err)
		}

		// Read the raw request body
		rawBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read request body",
			})
			return
		}

		// Parse the webhook payload
		var body payos.WebhookType
		if err := json.Unmarshal(rawBody, &body); err != nil {
			log.Printf("Error parsing webhook data: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid webhook payload",
			})
			return
		}

		// Verify the webhook data
		data, payosErr := payos.VerifyPaymentWebhookData(body)

		if payosErr != nil {
			log.Printf("Error verifying webhook data: %v", payosErr)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "Failed to verify webhook data",
			})
			return
		}

		if err := bookingService.ConfirmBookingByID(int(data.OrderCode)); err != nil {
			log.Printf("Error confirming booking: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to confirm booking",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	}
}

func CreatePaymentLink(paymentService services.PaymentServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {

		bookingId, _ := strconv.Atoi(c.Param("bookingId"))

		paymentUrl, err := paymentService.CreatePaymentLinkWithPayOsUsingBookingId(bookingId)
		if err != nil {
			log.Printf("Error creating payment link: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create payment link",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": paymentUrl,
		})
	}
}
