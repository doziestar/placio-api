package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"placio-app/ent"
)

type MediaService interface {
	CreateMedia(ctx context.Context, url, mediaType string) (*ent.Media, error)
	GetMedia(ctx context.Context, mediaID string) (*ent.Media, error)
	UploadFiles(ctx context.Context, files []*multipart.FileHeader) ([]MediaInfo, error)
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

func NewMediaService(client *ent.Client, cloud *cloudinary.Cloudinary) MediaService {
	return &MediaServiceImpl{client: client, cloud: cloud}
}

func (s *MediaServiceImpl) UploadFiles(ctx context.Context, files []*multipart.FileHeader) ([]MediaInfo, error) {
	ch := make(chan MediaInfo)
	errCh := make(chan error)

	for _, file := range files {
		go func(file *multipart.FileHeader) {
			openedFile, err := file.Open() // Open the file to get a stream
			if err != nil {
				log.Println("Error opening file", err)
				errCh <- err
				return
			}
			defer openedFile.Close()

			uploadResp, err := s.cloud.Upload.Upload(ctx, openedFile, uploader.UploadParams{})
			if err != nil {
				log.Println("Error uploading file", err)
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

func (s *MediaServiceImpl) CreateMedia(ctx context.Context, url, mediaType string) (*ent.Media, error) {
	if url == "" || mediaType == "" {
		return nil, errors.New("url and media type cannot be empty")
	}

	// Create builder
	mediaBuilder := s.client.Media.
		Create().
		SetID(uuid.New().String()).
		SetURL(url).
		SetMediaType(mediaType)

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
