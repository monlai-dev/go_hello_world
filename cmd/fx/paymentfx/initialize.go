package paymentfx

import (
	"go.uber.org/fx"
	"webapp/internal/services"
)

var Module = fx.Provide(providePaymentService)

func providePaymentService(slotService services.SlotServiceInterface,
	bookingService services.BookingServiceInterface,
	seatService services.SeatServiceInterface,
	bookedSeatService services.BookedSeatServiceInterface) services.PaymentServiceInterface {
	return services.NewPaymentService(slotService, bookingService, seatService, bookedSeatService)
}
