package tickets

import (
	"gorm.io/gorm"
	"placio-app/models"
)

type TicketService interface {
	CreateTicket(ticket *models.Ticket) error
	GetTicketByEvent(eventID string) ([]models.Ticket, error)
	UpdateTicket(ticket *models.Ticket) error
	DeleteTicket(ticketID string) error
}

type TicketServiceImpl struct {
	db          *gorm.DB
	ticketStore *models.Ticket
}

func NewTicketService(db *gorm.DB) TicketService {
	return &TicketServiceImpl{db: db, ticketStore: &models.Ticket{}}
}

func (ts *TicketServiceImpl) CreateTicket(ticket *models.Ticket) error {
	return ts.db.Create(ticket).Error
}

func (ts *TicketServiceImpl) GetTicketByEvent(eventID string) ([]models.Ticket, error) {
	var ticketOptions []models.Ticket
	err := ts.db.Where("event_id = ?", eventID).Find(&ticketOptions).Error
	return ticketOptions, err
}

func (ts *TicketServiceImpl) UpdateTicket(ticket *models.Ticket) error {
	return ts.db.Save(ticket).Error
}

func (ts *TicketServiceImpl) DeleteTicket(ticketID string) error {
	return ts.db.Delete(&models.Ticket{}, "id = ?", ticketID).Error
}
