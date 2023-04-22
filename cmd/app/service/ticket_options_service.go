package service

import (
	"gorm.io/gorm"
	"placio-app/models"
)

type TicketOptionService interface {
	CreateTicketOption(ticketOption *models.TicketOption) error
	GetTicketOptionsByEvent(eventID string) ([]models.TicketOption, error)
	UpdateTicketOption(ticketOption *models.TicketOption) error
	DeleteTicketOption(ticketOptionID string) error
}

type TicketOptionServiceImpl struct {
	db                *gorm.DB
	ticketOptionStore *models.TicketOption
}

func NewTicketOptionService(db *gorm.DB) TicketOptionService {
	return &TicketOptionServiceImpl{db: db, ticketOptionStore: &models.TicketOption{}}
}

func (ts *TicketOptionServiceImpl) CreateTicketOption(ticketOption *models.TicketOption) error {
	return ts.db.Create(ticketOption).Error
}

func (ts *TicketOptionServiceImpl) GetTicketOptionsByEvent(eventID string) ([]models.TicketOption, error) {
	var ticketOptions []models.TicketOption
	err := ts.db.Where("event_id = ?", eventID).Find(&ticketOptions).Error
	return ticketOptions, err
}

func (ts *TicketOptionServiceImpl) UpdateTicketOption(ticketOption *models.TicketOption) error {
	return ts.db.Save(ticketOption).Error
}

func (ts *TicketOptionServiceImpl) DeleteTicketOption(ticketOptionID string) error {
	return ts.db.Delete(&models.TicketOption{}, "id = ?", ticketOptionID).Error
}
