package tickets

import (
	"context"
	"placio-app/ent"
)

type ITicketService interface {
	CreateTicketOption(ctx context.Context, eventId string, option *ent.TicketOption) (*ent.TicketOption, error)
	UpdateTicketOption(ctx context.Context, optionId string, update *ent.TicketOption) (*ent.TicketOption, error)
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
	client *ent.Client
}

func NewTicketService(client *ent.Client) ITicketService {
	return TicketService{client: client}
}

func (t TicketService) CreateTicketOption(ctx context.Context, eventId string, option *ent.TicketOption) (*ent.TicketOption, error) {
	//TODO implement me
	panic("implement me")
}

func (t TicketService) UpdateTicketOption(ctx context.Context, optionId string, update *ent.TicketOption) (*ent.TicketOption, error) {
	//TODO implement me
	panic("implement me")
}

func (t TicketService) DeleteTicketOption(ctx context.Context, optionId string) error {
	//TODO implement me
	panic("implement me")
}

func (t TicketService) GetTicketOptionsForEvent(ctx context.Context, eventId string) ([]*ent.TicketOption, error) {
	//TODO implement me
	panic("implement me")
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
