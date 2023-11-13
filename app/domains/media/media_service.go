package media

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"placio-app/ent"
	"sync"

	firebase "firebase.google.com/go"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"google.golang.org/api/option"
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
	s3Client    *s3.S3
    firebase    *firebase.App
}

func NewMediaService(client *ent.Client, cloud *cloudinary.Cloudinary) MediaService {
	return &MediaServiceImpl{client: client, cloud: cloud}
}

func (s *MediaServiceImpl) uploadToS3(ctx context.Context, file *multipart.FileHeader) ([]MediaInfo, error) {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("your-region"), // e.g., "us-west-2"
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create AWS session: %w", err)
    }

    uploader := s3manager.NewUploader(sess)

    var mediaInfos []MediaInfo
	openedFile, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	key := "your-path-prefix/" + filepath.Base(file.Filename) // Customize your S3 key path
	// Define the bucket and key
	uploadOutput, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String("your-bucket-name"),
		Key:    aws.String(key),
		Body:   openedFile,
	})
	openedFile.Close() // Make sure to close the file

	if err != nil {
		log.Printf("failed to upload file to S3: %s, error: %v", file.Filename, err)
	}

	mediaInfos = append(mediaInfos, MediaInfo{
		URL:       uploadOutput.Location,
		MediaType: file.Header.Get("Content-Type"), // Or determine the type in another way
	})

    return mediaInfos, nil
}


func (s *MediaServiceImpl) uploadToFirebase(ctx context.Context, file *multipart.FileHeader) ([]MediaInfo, error) {
    conf := &firebase.Config{
        StorageBucket: "your-firebase-storage-bucket.appspot.com",
    }
    opt := option.WithCredentialsFile("path/to/your/serviceAccountKey.json")
    app, err := firebase.NewApp(ctx, conf, opt)
    if err != nil {
        return nil, fmt.Errorf("error initializing firebase app: %w", err)
    }

    client, err := app.Storage(ctx)
    if err != nil {
        return nil, fmt.Errorf("error getting Storage client: %w", err)
    }

    var mediaInfos []MediaInfo
	openedFile, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer openedFile.Close()

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, fmt.Errorf("error getting default bucket: %w", err)
	}

	object := bucket.Object("your-path-prefix/" + filepath.Base(file.Filename)) // Customize your Firebase Storage path
	wc := object.NewWriter(ctx)

	if _, err = io.Copy(wc, openedFile); err != nil {
		return nil, fmt.Errorf("error writing to Firebase Storage: %w", err)
	}
	if err = wc.Close(); err != nil {
		return nil, fmt.Errorf("error closing writer: %w", err)
	}

	attrs, err := object.Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting object attributes: %w", err)
	}

	mediaInfos = append(mediaInfos, MediaInfo{
		URL:       attrs.MediaLink,
		MediaType: file.Header.Get("Content-Type"), // Or determine the type in another way
	})
    

    return mediaInfos, nil
}


func (s *MediaServiceImpl) uploadToCloudinary(ctx context.Context, files []*multipart.FileHeader) ([]MediaInfo, error) {
    var mediaInfos []MediaInfo

    for _, file := range files {
        openedFile, err := file.Open()
        if err != nil {
            return nil, fmt.Errorf("failed to open file: %w", err)
        }
        defer openedFile.Close()

        uploadParams := uploader.UploadParams{Folder: "your-folder-path"} // Customize your path in Cloudinary
        uploadResult, err := s.cloud.Upload.Upload(ctx, openedFile, uploadParams)
        if err != nil {
            return nil, fmt.Errorf("error uploading file to Cloudinary: %w", err)
        }

        mediaInfos = append(mediaInfos, MediaInfo{
            URL:       uploadResult.SecureURL,
            MediaType: uploadResult.ResourceType, // This is typically 'image' or 'video'
        })
    }

    return mediaInfos, nil
}


func (s *MediaServiceImpl) UploadFiles(ctx context.Context, files []*multipart.FileHeader) ([]MediaInfo, error) {
    const maxWorkers = 5 // Adjust this to a suitable number

    ch := make(chan MediaInfo)
    errCh := make(chan error, len(files)) // Buffer this channel to prevent goroutine leaks
    wg := sync.WaitGroup{}
    sem := make(chan struct{}, maxWorkers)

    for _, file := range files {
        wg.Add(1)
        go func(file *multipart.FileHeader) {
            defer wg.Done()

            sem <- struct{}{}
            defer func() { <-sem }()

            var mediaInfo []MediaInfo
            var err error

            // Randomly select a storage service
            switch rand.Intn(2) { 
            case 0:
                mediaInfo, err = s.uploadToS3(ctx, file)
            case 1:
                mediaInfo, err = s.uploadToFirebase(ctx, file)
            // case 2:
            //     mediaInfo, err = s.uploadToCloudinary(ctx, []*multipart.FileHeader{file})
            }

            if err != nil {
                log.Println("Error uploading file: ", err)
                errCh <- err
                return
            }

            // Assuming mediaInfo is a slice, even though it contains only one element
            for _, info := range mediaInfo {
                ch <- info
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
        case info := <-ch:
            mediaInfos = append(mediaInfos, info)
        case err := <-errCh:
            if firstError == nil {
                firstError = err
            }
        }
    }

    if firstError != nil {
        return nil, firstError
    }

    log.Println("Uploaded media info: ", mediaInfos)
    return mediaInfos, nil
}


// func (s *MediaServiceImpl) UploadFiles(ctx context.Context, files []*multipart.FileHeader) ([]MediaInfo, error) {
// 	const maxWorkers = 5 // adjust this to a suitable number

// 	ch := make(chan MediaInfo)
// 	errCh := make(chan error, len(files)) // buffer this channel to prevent goroutine leaks
// 	wg := sync.WaitGroup{}
// 	sem := make(chan struct{}, maxWorkers)

// 	for _, file := range files {
// 		wg.Add(1)
// 		go func(file *multipart.FileHeader) {
// 			defer wg.Done()

// 			sem <- struct{}{}
// 			defer func() { <-sem }()

// 			openedFile, err := file.Open() // Open the file to get a stream
// 			if err != nil {
// 				log.Println("Error opening file", err)
// 				errCh <- err
// 				return
// 			}
// 			defer openedFile.Close()

// 			log.Println("uploading media: ", openedFile)
// 			uploadResp, err := s.cloud.Upload.Upload(ctx, openedFile, uploader.UploadParams{})
// 			if err != nil {
// 				log.Println("Error uploading file", err)
// 				errCh <- err
// 				return
// 			}

// 			ch <- MediaInfo{
// 				URL:       uploadResp.SecureURL,
// 				MediaType: uploadResp.ResourceType,
// 			}
// 		}(file)
// 	}

// 	// Close the channels after all goroutines complete
// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 		close(errCh)
// 	}()

// 	mediaInfos := make([]MediaInfo, 0, len(files))
// 	var firstError error
// 	for i := 0; i < len(files); i++ {
// 		select {
// 		case info, ok := <-ch:
// 			if ok {
// 				// log.Println("media uploaded: ", mediaInfos[0].URL)
// 				mediaInfos = append(mediaInfos, info)
// 			}
// 		case err, ok := <-errCh:
// 			if ok && firstError == nil {
// 				firstError = err
// 			}
// 		}
// 	}

// 	if firstError != nil {
// 		return nil, firstError
// 	}

// 	log.Println("uploaded media info: ", mediaInfos)
// 	return mediaInfos, nil
// }

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

	log.Println("image uploaded successfully: ", uploadedFiles)

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

	log.Println("upload complete", mediaList)

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
