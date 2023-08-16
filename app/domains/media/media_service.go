package media

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
	"sync"
)

type MediaService interface {
	CreateMedia(ctx context.Context, url, mediaType string) (*ent.Media, error)
	GetMedia(ctx context.Context, mediaID string) (*ent.Media, error)
	UploadFiles(ctx context.Context, files []*multipart.FileHeader) ([]MediaInfo, error)
	UploadAndCreateMedia(ctx context.Context, files []*multipart.FileHeader) ([]*ent.Media, error)
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
	const maxWorkers = 5 // adjust this to a suitable number

	ch := make(chan MediaInfo)
	errCh := make(chan error, len(files)) // buffer this channel to prevent goroutine leaks
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, maxWorkers)

	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

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
				MediaType: uploadResp.ResourceType,
			}
		}(file)
	}

	// Close the channels after all goroutines complete
	go func() {
		wg.Wait()
		close(ch)
		close(errCh)
	}()

	mediaInfos := make([]MediaInfo, 0, len(files))
	var firstError error
	for i := 0; i < len(files); i++ {
		select {
		case info, ok := <-ch:
			if ok {
				mediaInfos = append(mediaInfos, info)
			}
		case err, ok := <-errCh:
			if ok && firstError == nil {
				firstError = err
			}
		}
	}

	if firstError != nil {
		return nil, firstError
	}

	log.Println("uploaded media info: ", mediaInfos)
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

func (s *MediaServiceImpl) UploadAndCreateMedia(ctx context.Context, files []*multipart.FileHeader) ([]*ent.Media, error) {
	// Upload files
	uploadedFiles, err := s.UploadFiles(ctx, files)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed uploading files: %w", err)
	}

	// Prepare media creations
	mediaCreations := make([]*ent.MediaCreate, len(uploadedFiles))

	for i, file := range uploadedFiles {
		mediaID := uuid.New().String()
		mediaCreations[i] = s.client.Media.
			Create().
			SetID(mediaID).
			SetMediaType(file.MediaType).
			SetURL(file.URL)
	}

	// Bulk create media
	mediaList, err := s.client.Media.CreateBulk(mediaCreations...).Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("upload complete")

	return mediaList, nil
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
