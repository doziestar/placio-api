package controller

import "placio-app/service"

type BookingController struct {
	bookingService service.BookingService
}

func NewBookingController(bookingService service.BookingService) *BookingController {
	return &BookingController{bookingService: bookingService}
}
