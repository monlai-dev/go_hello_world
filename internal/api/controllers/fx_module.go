package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewBookedSeatController),
	fx.Provide(NewAccountController),
	fx.Provide(NewBookingController),
	fx.Provide(NewMovieController),
	fx.Provide(NewRoomController),
	fx.Provide(NewSlotController),
	fx.Provide(NewTheaterController),
	fx.Provide(NewWebHookController))
