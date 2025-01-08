package unit_test

import (
	"testing"
	"webapp/payment_gateway"
)

func TestPaymentGateway(t *testing.T) {
	t.Setenv("STRIPE_SECRET", "sk_test_51Q8jjrRw3INapdnuf67oafDmAQZt1XDLScwcbcaa4T1DH3oakiBVrV0PMaMGBVeJPZX3BKBt7K0eY1FC0JR64r6F0023NdblnC")
	payment_gateway.StripeAddNewProduct(1, "mon", 500)
}

func TestInvoiceIntentPayment(t *testing.T) {
	t.Setenv("STRIPE_SECRET", "sk_test_51Q8jjrRw3INapdnuf67oafDmAQZt1XDLScwcbcaa4T1DH3oakiBVrV0PMaMGBVeJPZX3BKBt7K0eY1FC0JR64r6F0023NdblnC")
	payment_gateway.CreateStripePayment()
}
