package service

import (
	"gorm.io/gorm"
	"placio-app/Dto"
	"placio-app/models"
)

type IEventService interface {
	CreateEvent(data *Dto.EventDto) (*models.Event, error)
	GetEventByID(eventId string) (*models.Event, error)
	GetEventByLocation(locationId string) (*[]models.Event, error)
	GetEventByCategory(categoryId string) (*[]models.Event, error)
	GetEventByDate(date string) (*[]models.Event, error)
	DeleteEvent(eventId string) error
	UpdateEvent(eventId string, data *Dto.EventDto) (*models.Event, error)
	GetEventParticipants(eventId string) (*[]models.User, error)
	GetEventsByAccount(accountID string) ([]models.Event, error)
}

type EventService struct {
	db      *gorm.DB
	event   *models.Event
	account *models.Account
}

func NewEventService(db *gorm.DB, event *models.Event, account *models.Account) *EventService {
	return &EventService{db: db, event: event, account: account}
}

func (s *EventService) CreateEvent(data *Dto.EventDto) (*models.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventByID(eventId string) (*models.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventByLocation(locationId string) (*[]models.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventByCategory(categoryId string) (*[]models.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventByDate(date string) (*[]models.Event, error) {
	return nil, nil
}

func (s *EventService) DeleteEvent(eventId string) error {
	return nil
}

func (s *EventService) UpdateEvent(eventId string, data *Dto.EventDto) (*models.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventParticipants(eventId string) (*[]models.User, error) {
	return nil, nil
}

func (s *EventService) GetEventsByAccount(accountID string) ([]models.Event, error) {
	var events []models.Event
	err := s.db.Where("account_id = ?", accountID).Find(&events).Error
	return events, err
}
