package services

type PaymentServiceInterface interface {
	CreatePaymentLinkWithPayOsUsingBookingId(bookingId int) (string, error)
}
