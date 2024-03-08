package tickets

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"mime/multipart"
	"placio-app/domains/media"
	"placio-app/ent"
	"placio-app/ent/event"
	"placio-app/ent/ticketoption"
	"placio-pkg/errors"
)

type ITicketService interface {
	CreateTicketOption(ctx context.Context, eventId string, option *ent.TicketOption) (*ent.TicketOption, error)
	UpdateTicketOption(ctx context.Context, optionId string, update *ent.TicketOption) (*ent.TicketOption, error)
	AddMediaToTicketOption(ctx context.Context, ticketOptionID string, files []*multipart.FileHeader) (*ent.TicketOption, error)
	RemoveMediaFromTicketOption(ctx context.Context, ticketOptionID string, mediaID string) error
	DeleteTicketOption(ctx context.Context, optionId string) error
	GetTicketOptionsForEvent(ctx context.Context, eventId string) ([]*ent.TicketOption, error)
	PurchaseTicket(ctx context.Context, userId string, purchaseDetails *TicketPurchaseDTO) (*ent.Ticket, error)
	ValidateTicket(ctx context.Context, ticketId string) (*ent.Ticket, error)
	CancelTicket(ctx context.Context, ticketId string) error
	TransferTicket(ctx context.Context, ticketId string, toUserId string) error
	GetTicketsByUser(ctx context.Context, userId string) ([]*ent.Ticket, error)
	GetTicketDetails(ctx context.Context, ticketId string) (*ent.Ticket, error)

	//ApplyPromotionToTicketOption(ctx context.Context, optionId string, promotion *PromotionDTO) error
	//RemovePromotionFromTicketOption(ctx context.Context, optionId string) error
	//ListCurrentPromotionsForEvent(ctx context.Context, eventId string) ([]*PromotionDTO, error)
	//
	//ListAttendeesForEvent(ctx context.Context, eventId string) ([]*ent.User, error)
	//AssignTicketToAttendee(ctx context.Context, ticketId string, attendeeId string) error
	//GenerateAttendeeReportForEvent(ctx context.Context, eventId string) (*AttendeeReportDTO, error)
	//
	//GenerateSalesReportForEvent(ctx context.Context, eventId string) (*SalesReportDTO, error)
	//GenerateTicketOptionSalesReport(ctx context.Context, optionId string) (*TicketOptionSalesReportDTO, error)
	//
	//EnableSeatSelectionForTicketOption(ctx context.Context, optionId string, seatMapDetails *SeatMapDTO) error
	//DisableSeatSelectionForTicketOption(ctx context.Context, optionId string) error
	//ReserveSeats(ctx context.Context, ticketId string, seats []SeatSelectionDTO) error
	//ReleaseSeatReservation(ctx context.Context, ticketId string) error
	//
	//InitiateRealTimeTicketAvailability(ctx context.Context, eventId string) (*RealTimeAvailabilityDTO, error)
	//UpdateRealTimeTicketAvailability(ctx context.Context, eventId string, update *RealTimeUpdateDTO) error
	//
	//CustomizeTicketAppearance(ctx context.Context, ticketId string, appearanceDetails *TicketAppearanceDTO) error
	//PersonalizeTicketOffering(ctx context.Context, userId string, personalizationDetails *PersonalizationDetailsDTO) ([]*ent.TicketOption, error)
	//
	//SynchronizeTicketsWithEventSchedule(ctx context.Context, eventId string, scheduleDetails *EventScheduleDTO) error
	//
	//IntegrateWithPaymentGateway(ctx context.Context, gatewayDetails *PaymentGatewayDTO) error
	//ExportTicketSalesData(ctx context.Context, eventId string, exportOptions *DataExportOptionsDTO) error
	//
	//ProvideAccessibilityOptionsForTicketing(ctx context.Context, eventId string, options *AccessibilityOptionsDTO) error
}

type TicketService struct {
	client       *ent.Client
	mediaService media.MediaService
}

func NewTicketService(client *ent.Client, mediaService media.MediaService) ITicketService {
	return TicketService{client: client}
}

func (t TicketService) CreateTicketOption(ctx context.Context, eventId string, option *ent.TicketOption) (*ent.TicketOption, error) {
	if option.Price < 0 {
		return nil, errors.New("ticket option price must be non-negative")
	}

	exists, err := t.client.Event.Query().Where(event.ID(eventId)).Exist(ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to query event existence: %v", err))
	}
	if !exists {
		return nil, errors.New(fmt.Sprintf("event with ID %s does not exist", eventId))
	}

	createdOption, err := t.client.TicketOption.
		Create().
		SetName(option.Name).
		SetID(uuid.New().String()).
		SetDescription(option.Description).
		SetPrice(option.Price).
		SetQuantityAvailable(option.QuantityAvailable).
		SetEventID(eventId).
		Save(ctx)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create ticket option: %v", err))
	}

	return createdOption, nil
}

func (t TicketService) UpdateTicketOption(ctx context.Context, optionId string, update *ent.TicketOption) (*ent.TicketOption, error) {
	existingOption, err := t.client.TicketOption.Get(ctx, optionId)
	if err != nil {
		return nil, fmt.Errorf("failed to find ticket option with ID %s: %v", optionId, err)
	}

	updatedOption, err := existingOption.Update().
		SetNillableName(&update.Name).
		SetNillableDescription(&update.Description).
		SetNillablePrice(&update.Price).
		SetNillableQuantityAvailable(&update.QuantityAvailable).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to update ticket option: %v", err)
	}

	return updatedOption, nil
}

func (t TicketService) DeleteTicketOption(ctx context.Context, optionId string) error {
	err := t.client.TicketOption.DeleteOneID(optionId).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete ticket option with ID %s: %v", optionId, err)
	}
	return nil
}

func (t TicketService) GetTicketOptionsForEvent(ctx context.Context, eventId string) ([]*ent.TicketOption, error) {
	exists, err := t.client.Event.Query().Where(event.ID(eventId)).Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query event existence: %v", err)
	}
	if !exists {
		return nil, fmt.Errorf("event with ID %s does not exist", eventId)
	}

	options, err := t.client.TicketOption.Query().
		Where(ticketoption.HasEventWith(event.ID(eventId))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ticket options for event ID %s: %v", eventId, err)
	}
	return options, nil
}

func (t TicketService) AddMediaToTicketOption(ctx context.Context, ticketOptionID string, files []*multipart.FileHeader) (*ent.TicketOption, error) {
	// Start a transaction
	tx, err := t.client.Tx(ctx)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	// Fetch ticketOption within transaction
	ticketOption, err := tx.TicketOption.Get(ctx, ticketOptionID)
	if err != nil {
		sentry.CaptureException(err)
		tx.Rollback()
		return nil, err
	}

	uploadedFiles, err := t.mediaService.UploadAndCreateMedia(ctx, files)
	if err != nil {
		sentry.CaptureException(err)
		tx.Rollback()
		return nil, err
	}

	ticketOption, err = tx.TicketOption.UpdateOne(ticketOption).
		AddMedia(uploadedFiles...).
		Save(ctx)
	if err != nil {
		sentry.CaptureException(err)
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	ticketOption, err = t.client.TicketOption.Query().
		Where(ticketoption.IDEQ(ticketOptionID)).
		WithMedia().
		First(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return ticketOption, nil
}

func (t TicketService) RemoveMediaFromTicketOption(ctx context.Context, ticketOptionID string, mediaID string) error {
	// Start a transaction
	tx, err := t.client.Tx(ctx)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	media, err := tx.Media.Get(ctx, mediaID)
	if err != nil {
		tx.Rollback()
		sentry.CaptureException(err)
		return err
	}

	err = tx.TicketOption.UpdateOneID(ticketOptionID).
		RemoveMedia(media).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		sentry.CaptureException(err)
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func (t TicketService) PurchaseTicket(ctx context.Context, userId string, purchaseDetails *TicketPurchaseDTO) (*ent.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (t TicketService) ValidateTicket(ctx context.Context, ticketId string) (*ent.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (t TicketService) CancelTicket(ctx context.Context, ticketId string) error {
	//TODO implement me
	panic("implement me")
}

func (t TicketService) TransferTicket(ctx context.Context, ticketId string, toUserId string) error {
	//TODO implement me
	panic("implement me")
}

func (t TicketService) GetTicketsByUser(ctx context.Context, userId string) ([]*ent.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (t TicketService) GetTicketDetails(ctx context.Context, ticketId string) (*ent.Ticket, error) {
	//TODO implement me
	panic("implement me")
}
