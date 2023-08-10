package booking

import (
	"context"
	"placio-app/ent"
)

type ReservationService interface {
	GetReservation(ctx context.Context, reservationID string) (*ent.Reservation, error)
	CreateReservation(ctx context.Context, reservationData map[string]interface{}) (*ent.Reservation, error)
	UpdateReservation(ctx context.Context, reservationID string, reservationData map[string]interface{}) (*ent.Reservation, error)
	DeleteReservation(ctx context.Context, reservationID string) error
}

type ReservationServiceImpl struct {
	client *ent.Client
}

func NewReservationService(client *ent.Client) *ReservationServiceImpl {
	return &ReservationServiceImpl{client: client}
}

func (s *ReservationServiceImpl) GetReservation(ctx context.Context, reservationID string) (*ent.Reservation, error) {
	return s.client.Reservation.Get(ctx, reservationID)
}

func (s *ReservationServiceImpl) CreateReservation(ctx context.Context, reservationData map[string]interface{}) (*ent.Reservation, error) {
	return s.client.Reservation.
		Create().
		//SetNillableDateTime(reservationData["date_time"].(time.Time)).
		Save(ctx)
}

func (s *ReservationServiceImpl) UpdateReservation(ctx context.Context, reservationID string, reservationData map[string]interface{}) (*ent.Reservation, error) {
	return s.client.Reservation.
		UpdateOneID(reservationID).
		//SetNillableDateTime(reservationData["date_time"].(time.Time)).
		Save(ctx)
}

func (s *ReservationServiceImpl) DeleteReservation(ctx context.Context, reservationID string) error {
	return s.client.Reservation.
		DeleteOneID(reservationID).
		Exec(ctx)
}
