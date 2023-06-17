package controller

import "placio-app/service"

type ReservationController struct {
	reservationService service.ReservationService
}

func NewReservationController(reservationService service.ReservationService) *ReservationController {
	return &ReservationController{reservationService: reservationService}
}
