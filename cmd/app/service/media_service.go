package service

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent"
)

type MediaService interface {
	CreateMedia(ctx context.Context, media *ent.Media) (*ent.Media, error)
	GetMedia(ctx context.Context, mediaID string) (*ent.Media, error)
	//UpdateMedia(media *models.Media) (*models.Media, error)
	//DeleteMedia(mediaID string) error
	//ListMedia(postID string) ([]*models.Media, error)
	//UploadMedia(file *multipart.FileHeader, id string) (interface{}, error)
}

type MediaServiceImpl struct {
	client *ent.Client
}

func NewMediaService(client *ent.Client) MediaService {
	return &MediaServiceImpl{client: client}
}

func (ms *MediaServiceImpl) CreateMedia(ctx context.Context, media *ent.Media) (*ent.Media, error) {
	if media == nil {
		return nil, errors.New("media cannot be nil")
	}

	// Create builder
	mediaBuilder := ms.client.Media.
		Create().
		SetID(media.ID).
		SetURL(media.URL).
		SetMediaType(media.MediaType)

	// Save media
	createdMedia, err := mediaBuilder.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating media: %w", err)
	}

	return createdMedia, nil
}

func (ms *MediaServiceImpl) GetMedia(ctx context.Context, mediaID string) (*ent.Media, error) {
	media, err := ms.client.Media.Get(ctx, mediaID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed retrieving media: %w", err)
	}
	return media, nil
}

//
//func (ms *MediaServiceImpl) GetMedia(mediaID string) (*models.Media, error) {
//	var media models.Media
//	if err := ms.db.First(&media, "id = ?", mediaID).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, nil
//		}
//		return nil, err
//	}
//	return &media, nil
//}

//func (ms *MediaServiceImpl) UpdateMedia(media *models.Media) (*models.Media, error) {
//	if err := ms.db.Model(&models.Media{}).Where("id = ?", media.ID).Updates(media).Error; err != nil {
//		return nil, err
//	}
//	return media, nil
//}
//
//func (ms *MediaServiceImpl) DeleteMedia(mediaID string) error {
//	if err := ms.db.Delete(&models.Media{}, "id = ?", mediaID).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (ms *MediaServiceImpl) ListMedia(postID string) ([]*models.Media, error) {
//	var mediaList []*models.Media
//	if err := ms.db.Where("post_id = ?", postID).Find(&mediaList).Error; err != nil {
//		return nil, err
//	}
//	return mediaList, nil
//}
//
//func (ms *MediaServiceImpl) UploadMedia(file *multipart.FileHeader, id string) (interface{}, error) {
//	return nil, nil
//}
