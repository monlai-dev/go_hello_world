package payment_gateway

import (
	"fmt"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentlink"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/product"
	"log"
	"strconv"
)

func init() {
	stripe.Key = "sk_test_51Q8jjrRw3INapdnuf67oafDmAQZt1XDLScwcbcaa4T1DH3oakiBVrV0PMaMGBVeJPZX3BKBt7K0eY1FC0JR64r6F0023NdblnC"
}

func StripeAddNewProduct(id int, name string, productPrice int64) {
	params := &stripe.ProductParams{
		ID:   stripe.String(strconv.Itoa(id)),
		Name: stripe.String(name),
	}

	productValue, err := product.New(params)
	if err != nil {
		panic(err)
	}

	priceParams := &stripe.PriceParams{
		Product:    stripe.String(productValue.ID),
		UnitAmount: stripe.Int64(productPrice),
		Currency:   stripe.String("usd"),
	}

	_, err = price.New(priceParams)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product created with ID:", productValue.ID+"with price: ", productPrice)
}

func CreateStripePayment() {

	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			{
				Price:    stripe.String("1"),
				Quantity: stripe.Int64(1),
			},
		},
	}

	result, err := paymentlink.New(params)
	if err != nil {
		log.Printf("Error creating payment link: %v", err)
		return
	}

	log.Printf("Created payment link: %v", result.URL)
}
