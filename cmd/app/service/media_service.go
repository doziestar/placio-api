package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"placio-app/ent"
)

type MediaService interface {
	CreateMedia(ctx context.Context, media *ent.Media) (*ent.Media, error)
	GetMedia(ctx context.Context, mediaID string) (*ent.Media, error)
	UploadFiles(ctx context.Context, files []string, uploadParams uploader.UploadParams) ([]MediaInfo, error)
	//UpdateMedia(media *models.Media) (*models.Media, error)
	//DeleteMedia(mediaID string) error
	//ListMedia(postID string) ([]*models.Media, error)
	//UploadMedia(file *multipart.FileHeader, id string) (interface{}, error)
}

type MediaInfo struct {
	URL       string
	MediaType string
}

type MediaServiceImpl struct {
	client *ent.Client
	cloud  *cloudinary.Cloudinary
}

func NewMediaService(client *ent.Client) MediaService {
	return &MediaServiceImpl{client: client}
}

func (s *MediaServiceImpl) UploadFiles(ctx context.Context, files []string, uploadParams uploader.UploadParams) ([]MediaInfo, error) {
	ch := make(chan MediaInfo)
	errCh := make(chan error)

	for _, file := range files {
		go func(file string) {
			uploadResp, err := s.cloud.Upload.Upload(ctx, file, uploadParams)
			if err != nil {
				errCh <- err
				return
			}

			ch <- MediaInfo{
				URL:       uploadResp.SecureURL,
				MediaType: uploadResp.ResourceType, // TODO: You need to ensure this is the correct field for media type
			}
		}(file)
	}

	mediaInfos := make([]MediaInfo, 0, len(files))
	for range files {
		select {
		case info := <-ch:
			mediaInfos = append(mediaInfos, info)
		case err := <-errCh:
			return nil, err
		}
	}

	return mediaInfos, nil
}

func (s *MediaServiceImpl) CreateMedia(ctx context.Context, media *ent.Media) (*ent.Media, error) {
	if media == nil {
		return nil, errors.New("media cannot be nil")
	}

	// Create builder
	mediaBuilder := s.client.Media.
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

func (s *MediaServiceImpl) GetMedia(ctx context.Context, mediaID string) (*ent.Media, error) {
	media, err := s.client.Media.Get(ctx, mediaID)
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
