package payment_gateway

import (
	"fmt"
	"github.com/payOSHQ/payos-lib-golang"
	"log"
	models "webapp/internal/models/db_models"
)

func TestCreatePaymentLink() {
	payos.Key("e47ad6bc-40f4-4fc9-a723-cd8ddae9e3a8", "0ca2d2e9-d6fa-4c13-9885-848080846bd7", "1589dd37d76b03b0968bcd445dd4ed0e38fa63b0d5811bba8dd72b7ea88f95c0")

	body := payos.CheckoutRequestType{
		OrderCode: 10002,
		Amount:    10000,
		Items: []payos.Item{
			{
				Name:     "Mỳ tôm Hảo Hảo ly",
				Price:    10000,
				Quantity: 1,
			},
		},
		Description: "Thanh toán đơn hàng",
		CancelUrl:   "http://localhost:8080/cancel/",
		ReturnUrl:   "http://localhost:8080/success/",
	}

	data, err := payos.CreatePaymentLink(body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.CheckoutUrl)
}

func CreatePaymentLink(booking models.Booking, bookedSeats []models.Seat, slot models.Slot) (string, error) {
	payos.Key("e47ad6bc-40f4-4fc9-a723-cd8ddae9e3a8",
		"0ca2d2e9-d6fa-4c13-9885-848080846bd7",
		"1589dd37d76b03b0968bcd445dd4ed0e38fa63b0d5811bba8dd72b7ea88f95c0")

	var items []payos.Item

	for _, seat := range bookedSeats {
		item := payos.Item{
			Name:     seat.Name,
			Price:    int(slot.Price),
			Quantity: 1,
		}
		items = append(items, item)
	}

	body := payos.CheckoutRequestType{
		OrderCode:   int64(booking.ID),
		Amount:      int(booking.TotalPrice),
		Items:       items,
		Description: "Thanh toán đơn hàng",
		CancelUrl:   "http://localhost:8080/cancel/",
		ReturnUrl:   "http://localhost:8080/success/",
	}

	data, err := payos.CreatePaymentLink(body)
	if err != nil {
		log.Fatal(err)
	}

	return data.CheckoutUrl, nil
}
