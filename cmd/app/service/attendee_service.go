package service

import (
	"gorm.io/gorm"
	"placio-app/models"
)

type AttendeeService interface {
	AddAttendee(attendee *models.Attendee) error
	GetAttendeesByEvent(eventID string) ([]models.Attendee, error)
	UpdateAttendee(attendee *models.Attendee) error
	RemoveAttendee(attendeeID string) error
}

type AttendeeServiceImpl struct {
	db            *gorm.DB
	attendeeStore *models.Attendee
}

func NewAttendeeService(db *gorm.DB) AttendeeService {
	return &AttendeeServiceImpl{db: db, attendeeStore: &models.Attendee{}}
}

func (as *AttendeeServiceImpl) AddAttendee(attendee *models.Attendee) error {
	return as.db.Create(attendee).Error
}

func (as *AttendeeServiceImpl) GetAttendeesByEvent(eventID string) ([]models.Attendee, error) {
	var attendees []models.Attendee
	err := as.db.Where("event_id = ?", eventID).Find(&attendees).Error
	return attendees, err
}

func (as *AttendeeServiceImpl) UpdateAttendee(attendee *models.Attendee) error {
	return as.db.Save(attendee).Error
}

func (as *AttendeeServiceImpl) RemoveAttendee(attendeeID string) error {
	return as.db.Delete(&models.Attendee{}, "id = ?", attendeeID).Error
}
