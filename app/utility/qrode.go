package utility

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/skip2/go-qrcode"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GenerateAndUploadQRCode(ctx context.Context, url string, identifier string) (string, error) {
	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		log.Println("Error generating QR code:", err)
		return "", fmt.Errorf("failed to generate QR code: %w", err)
	}

	qr.ForegroundColor = color.RGBA{R: 139, G: 0, B: 0, A: 255}
	qr.BackgroundColor = color.White

	png, err := qr.PNG(256)
	if err != nil {
		log.Println("Error converting QR code to PNG:", err)
		return "", fmt.Errorf("failed to convert QR code to PNG: %w", err)
	}

	tmpFile, err := ioutil.TempFile("", "qr-code-*.png")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}

	_, err = tmpFile.Write(png)
	if err != nil {
		tmpFile.Close()
		return "", fmt.Errorf("failed to write to temporary file: %w", err)
	}
	tmpFile.Close()

	// Adjust the uploadQRCodeToDigitalOceanSpace function to be more generic
	signedURL, err := uploadQRCodeToDigitalOceanSpace(ctx, identifier, tmpFile.Name(), "image/png")
	if err != nil {
		return "", err
	}

	os.Remove(tmpFile.Name())

	return signedURL, nil
}

func uploadQRCodeToDigitalOceanSpace(ctx context.Context, identifier, filePath, contentType string) (string, error) {
	// DigitalOcean Spaces credentials and configuration
	accessKeyID := "DO00YJ68Y7KMTYP3J7HE"
	secretAccessKey := "P55ReutOGyn1d4qThoPCMj+O7qCUggr/Y+DQIUwYtjc"
	spaceName := "placio"
	endpoint := "https://placio.fra1.digitaloceanspaces.com"
	cdnEndpoint := "https://placio.fra1.cdn.digitaloceanspaces.com"

	// Create a new session without specifying a region
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("fra1"),
		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return "", fmt.Errorf("error creating session: %w", err)
	}

	// Open the file to upload
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Construct the unique file name
	uniqueFileName := fmt.Sprintf("qrcode/%s-%s", identifier, filepath.Base(filePath))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file and get the URL in return
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(spaceName),
		Key:                aws.String(uniqueFileName),
		ACL:                aws.String("public-read"),
		Body:               file,
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String("inline"),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload to DigitalOcean Space: %w", err)
	}

	// Direct URL of the uploaded file
	directURL := result.Location

	// Construct the CDN URL by replacing the direct URL's domain with the CDN endpoint
	cdnURL := strings.Replace(directURL, endpoint, cdnEndpoint, 1)

	return cdnURL, nil
}
