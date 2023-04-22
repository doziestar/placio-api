package service

import (
	"gorm.io/gorm"
	"placio-app/models"
)

type RatingService interface {
	CreateRating(rating *models.Rating) error
	GetRatingsByEvent(eventID string) ([]models.Rating, error)
	UpdateRating(rating *models.Rating) error
	DeleteRating(ratingID string) error
}

type RatingServiceImpl struct {
	db          *gorm.DB
	ratingStore *models.Rating
}

func NewRatingService(db *gorm.DB) RatingService {
	return &RatingServiceImpl{db: db, ratingStore: &models.Rating{}}
}

func (rs *RatingServiceImpl) CreateRating(rating *models.Rating) error {
	return rs.db.Create(rating).Error
}

func (rs *RatingServiceImpl) GetRatingsByEvent(eventID string) ([]models.Rating, error) {
	var ratings []models.Rating
	err := rs.db.Where("event_id = ?", eventID).Find(&ratings).Error
	return ratings, err
}

func (rs *RatingServiceImpl) UpdateRating(rating *models.Rating) error {
	return rs.db.Save(rating).Error
}

func (rs *RatingServiceImpl) DeleteRating(ratingID string) error {
	return rs.db.Delete(&models.Rating{}, "id = ?", ratingID).Error
}
