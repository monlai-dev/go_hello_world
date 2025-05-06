package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/payOSHQ/payos-lib-golang"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"webapp/internal/services"
)

type WebHookController struct {
	BookingService services.BookingServiceInterface
	PaymentService services.PaymentServiceInterface
}

func NewWebHookController(bookingService services.BookingServiceInterface, paymentService services.PaymentServiceInterface) *WebHookController {
	return &WebHookController{
		BookingService: bookingService,
		PaymentService: paymentService,
	}
}

func (wc *WebHookController) WebhookHandler(c *gin.Context) {

	if err := payos.Key(os.Getenv("CLIENT_ID"),
		os.Getenv("API_KEY"),
		os.Getenv("CHECK_SUM_KEY")); err != nil {
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

	if data.OrderCode == 123 {
		c.JSON(http.StatusOK, gin.H{
			"message": "COnfirm webhook complete",
		})
		return

	}

	// -1000 is the offset for the order code
	if err := wc.BookingService.ConfirmBookingByID(int(data.OrderCode) - 1000); err != nil {
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

func (wc *WebHookController) CreatePaymentLink(c *gin.Context) {

	bookingId, _ := strconv.Atoi(c.Param("bookingId"))

	paymentUrl, err := wc.PaymentService.CreatePaymentLinkWithPayOsUsingBookingId(bookingId)
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
