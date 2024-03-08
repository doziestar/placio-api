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
	"placio-app/ent/ticket"
	"placio-app/ent/ticketoption"
	"placio-app/ent/user"
	"placio-pkg/errors"
	"time"
)

type ITicketService interface {
	CreateTicketOption(ctx context.Context, eventId string, option *ent.TicketOption) (*ent.TicketOption, error)
	UpdateTicketOption(ctx context.Context, optionId string, update *ent.TicketOption) (*ent.TicketOption, error)
	AddMediaToTicketOption(ctx context.Context, ticketOptionID string, files []*multipart.FileHeader) (*ent.TicketOption, error)
	RemoveMediaFromTicketOption(ctx context.Context, ticketOptionID string, mediaID string) error
	DeleteTicketOption(ctx context.Context, optionId string) error
	GetTicketOptionsForEvent(ctx context.Context, eventId string) ([]*ent.TicketOption, error)
	PurchaseTicket(ctx context.Context, userId string, purchaseDetails *TicketPurchaseDTO) ([]*ent.Ticket, error)
	ValidateTicket(ctx context.Context, ticketId string) (*ent.Ticket, error)
	CancelTicket(ctx context.Context, ticketId string) error
	TransferTicket(ctx context.Context, ticketId string, toUserId string) error
	GetTicketsByUser(ctx context.Context, userId string) ([]*ent.Ticket, error)
	GetTicketDetails(ctx context.Context, ticketId string) (*ent.Ticket, error)
	ListAttendeesForEvent(ctx context.Context, eventId string) ([]*ent.User, error)
	AssignTicketToAttendee(ctx context.Context, ticketId string, attendeeId string) error

	//ApplyPromotionToTicketOption(ctx context.Context, optionId string, promotion *PromotionDTO) error
	//RemovePromotionFromTicketOption(ctx context.Context, optionId string) error
	//ListCurrentPromotionsForEvent(ctx context.Context, eventId string) ([]*PromotionDTO, error)
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
	return &TicketService{client: client}
}

func (t *TicketService) CreateTicketOption(ctx context.Context, eventId string, option *ent.TicketOption) (*ent.TicketOption, error) {
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

func (t *TicketService) UpdateTicketOption(ctx context.Context, optionId string, update *ent.TicketOption) (*ent.TicketOption, error) {
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

func (t *TicketService) DeleteTicketOption(ctx context.Context, optionId string) error {
	err := t.client.TicketOption.DeleteOneID(optionId).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete ticket option with ID %s: %v", optionId, err)
	}
	return nil
}

func (t *TicketService) GetTicketOptionsForEvent(ctx context.Context, eventId string) ([]*ent.TicketOption, error) {
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

func (t *TicketService) AddMediaToTicketOption(ctx context.Context, ticketOptionID string, files []*multipart.FileHeader) (*ent.TicketOption, error) {
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

func (t *TicketService) RemoveMediaFromTicketOption(ctx context.Context, ticketOptionID string, mediaID string) error {
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

func (t *TicketService) PurchaseTicket(ctx context.Context, userId string, purchaseDetails *TicketPurchaseDTO) ([]*ent.Ticket, error) {
	// Start a transaction
	tx, err := t.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting transaction failed: %v", err)
	}

	// Verify the event exists
	exists, err := tx.Event.Query().Where(event.ID(purchaseDetails.EventID)).Exist(ctx)
	if err != nil || !exists {
		tx.Rollback()
		return nil, fmt.Errorf("event does not exist: %v", err)
	}

	var createdTickets []*ent.Ticket

	// Process each ticket option purchase
	for _, optionPurchase := range purchaseDetails.TicketOptions {
		// Fetch the ticket option to check availability and details
		option, err := tx.TicketOption.Get(ctx, optionPurchase.OptionID)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("ticket option not found: %v", err)
		}

		// Check if there are enough tickets available
		if option.QuantityAvailable < optionPurchase.Quantity {
			tx.Rollback()
			return nil, fmt.Errorf("not enough tickets available for option %s", optionPurchase.OptionID)
		}

		// Update the ticket option's sold and available quantities
		_, err = option.Update().
			AddQuantitySold(optionPurchase.Quantity).
			SetQuantityAvailable(option.QuantityAvailable - optionPurchase.Quantity).
			Save(ctx)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("updating ticket option failed: %v", err)
		}

		// Create the tickets
		for i := 0; i < optionPurchase.Quantity; i++ {
			createdTicket, err := tx.Ticket.Create().
				SetEventID(purchaseDetails.EventID).
				SetTicketOptionID(optionPurchase.OptionID).
				SetPurchaserID(userId).
				SetStatus(ticket.StatusSold). // Assuming you have a status field to track ticket state
				Save(ctx)
			if err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("creating ticket failed: %v", err)
			}
			createdTickets = append(createdTickets, createdTicket)
		}
	}

	// Commit the transaction if everything is successful
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction failed: %v", err)
	}

	// Return the first created ticket as a confirmation, or modify to return all tickets or a summary
	if len(createdTickets) > 0 {
		return createdTickets, nil
	}

	return nil, fmt.Errorf("no tickets were created")
}

func (t *TicketService) ValidateTicket(ctx context.Context, ticketId string) (*ent.Ticket, error) {
	// Begin a transaction
	tx, err := t.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting transaction failed: %v", err)
	}

	// Retrieve the ticket
	ticket, err := tx.Ticket.Get(ctx, ticketId)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("ticket not found: %v", err)
	}

	// Check if the ticket is already used
	if ticket.Status == "validated" {
		tx.Rollback()
		return nil, fmt.Errorf("ticket already validated")
	}

	// Update the ticket status to validated
	validatedTicket, err := ticket.Update().
		SetStatus("validated").
		SetValidationTime(time.Now()).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("validating ticket failed: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction failed: %v", err)
	}

	return validatedTicket, nil
}

func (t *TicketService) CancelTicket(ctx context.Context, ticketId string) error {
	ticket, err := t.client.Ticket.Get(ctx, ticketId)
	if err != nil {
		return fmt.Errorf("ticket not found: %v", err)
	}

	if ticket.Status != "available" && ticket.Status != "reserved" {
		return fmt.Errorf("ticket cannot be cancelled in its current state")
	}

	_, err = ticket.Update().
		SetStatus("cancelled").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("cancelling ticket failed: %v", err)
	}

	return nil
}

func (t *TicketService) TransferTicket(ctx context.Context, ticketId string, toUserId string) error {
	tx, err := t.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction failed: %v", err)
	}

	_, err = tx.Ticket.UpdateOneID(ticketId).
		SetPurchaserID(toUserId).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("transferring ticket failed: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction failed: %v", err)
	}

	return nil
}

func (t *TicketService) GetTicketsByUser(ctx context.Context, userId string) ([]*ent.Ticket, error) {
	tickets, err := t.client.Ticket.Query().
		Where(ticket.HasPurchaserWith(user.ID(userId))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching tickets for user failed: %v", err)
	}

	return tickets, nil
}

func (t *TicketService) GetTicketDetails(ctx context.Context, ticketId string) (*ent.Ticket, error) {
	ticket, err := t.client.Ticket.Get(ctx, ticketId)
	if err != nil {
		return nil, fmt.Errorf("ticket not found: %v", err)
	}

	return ticket, nil
}

func (t *TicketService) ListAttendeesForEvent(ctx context.Context, eventId string) ([]*ent.User, error) {
	// Fetch all tickets for the given event
	tickets, err := t.client.Ticket.
		Query().
		Where(ticket.HasEventWith(event.ID(eventId))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tickets for event %s: %v", eventId, err)
	}

	// Using ticket IDs, fetch all unique users (attendees)
	ticketIDs := make([]string, len(tickets))
	for i, ticket := range tickets {
		ticketIDs[i] = ticket.ID
	}

	users, err := t.client.User.
		Query().
		Where(user.HasPurchasedTicketsWith(ticket.IDIn(ticketIDs...))).
		Unique(true).
		//Distinct().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch attendees for event %s: %v", eventId, err)
	}

	return users, nil
}

func (t *TicketService) AssignTicketToAttendee(ctx context.Context, ticketId string, attendeeId string) error {
	// Begin a transaction to ensure data consistency
	tx, err := t.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction failed: %v", err)
	}

	// Verify the ticket exists
	ticket, err := tx.Ticket.Get(ctx, ticketId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ticket not found: %v", err)
	}

	// Verify the attendee (user) exists
	exists, err := tx.User.Query().Where(user.ID(attendeeId)).Exist(ctx)
	if err != nil || !exists {
		tx.Rollback()
		return fmt.Errorf("attendee not found: %v", err)
	}

	// Update the ticket to assign it to the new attendee
	_, err = ticket.Update().
		SetPurchaserID(attendeeId).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("assigning ticket to attendee failed: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction failed: %v", err)
	}

	return nil
}
