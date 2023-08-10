package booking

import (
	"context"
	"placio-app/ent"
)

type BookingService interface {
	GetBooking(ctx context.Context, bookingID string) (*ent.Booking, error)
	CreateBooking(ctx context.Context, bookingData map[string]interface{}) (*ent.Booking, error)
	UpdateBooking(ctx context.Context, bookingID string, bookingData map[string]interface{}) (*ent.Booking, error)
	DeleteBooking(ctx context.Context, bookingID string) error
}

type BookingServiceImpl struct {
	client *ent.Client
}

func NewBookingService(client *ent.Client) *BookingServiceImpl {
	return &BookingServiceImpl{client: client}
}

func (s *BookingServiceImpl) GetBooking(ctx context.Context, bookingID string) (*ent.Booking, error) {
	return s.client.Booking.Get(ctx, bookingID)
}

func (s *BookingServiceImpl) CreateBooking(ctx context.Context, bookingData map[string]interface{}) (*ent.Booking, error) {
	return s.client.Booking.
		Create().
		//SetNillableDateTime(bookingData["date_time"].(time.Time)).
		Save(ctx)
}

func (s *BookingServiceImpl) UpdateBooking(ctx context.Context, bookingID string, bookingData map[string]interface{}) (*ent.Booking, error) {
	return s.client.Booking.
		UpdateOneID(bookingID).
		//SetNillableDateTime(bookingData["date_time"].(time.Time)).
		Save(ctx)
}

func (s *BookingServiceImpl) DeleteBooking(ctx context.Context, bookingID string) error {
	return s.client.Booking.
		DeleteOneID(bookingID).
		Exec(ctx)
}
