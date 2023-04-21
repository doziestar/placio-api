package service

import (
	"errors"
	"gorm.io/gorm"
	"mime/multipart"
	"placio-app/models"
)

type MediaService interface {
	CreateMedia(media *models.Media) (*models.Media, error)
	GetMedia(mediaID string) (*models.Media, error)
	UpdateMedia(media *models.Media) (*models.Media, error)
	DeleteMedia(mediaID string) error
	ListMedia(postID string) ([]*models.Media, error)
	UploadMedia(file *multipart.FileHeader, id string) (interface{}, error)
}

type mediaServiceImpl struct {
	db *gorm.DB
}

func NewMediaService(db *gorm.DB) MediaService {
	return &mediaServiceImpl{db: db}
}

func (ms *mediaServiceImpl) CreateMedia(media *models.Media) (*models.Media, error) {
	if err := ms.db.Create(&media).Error; err != nil {
		return nil, err
	}
	return media, nil
}

func (ms *mediaServiceImpl) GetMedia(mediaID string) (*models.Media, error) {
	var media models.Media
	if err := ms.db.First(&media, "id = ?", mediaID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &media, nil
}

func (ms *mediaServiceImpl) UpdateMedia(media *models.Media) (*models.Media, error) {
	if err := ms.db.Model(&models.Media{}).Where("id = ?", media.ID).Updates(media).Error; err != nil {
		return nil, err
	}
	return media, nil
}

func (ms *mediaServiceImpl) DeleteMedia(mediaID string) error {
	if err := ms.db.Delete(&models.Media{}, "id = ?", mediaID).Error; err != nil {
		return err
	}
	return nil
}

func (ms *mediaServiceImpl) ListMedia(postID string) ([]*models.Media, error) {
	var mediaList []*models.Media
	if err := ms.db.Where("post_id = ?", postID).Find(&mediaList).Error; err != nil {
		return nil, err
	}
	return mediaList, nil
}

func (ms *mediaServiceImpl) UploadMedia(file *multipart.FileHeader, id string) (interface{}, error) {
	return nil, nil
}
