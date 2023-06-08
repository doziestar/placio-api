package utility

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)


func UploadImage(image string) (string, error) {
	cld, _ := cloudinary.NewFromParams("n07t21i7", "123456789012345", "abcdeghijklmnopqrstuvwxyz12")
	cld, err := cloudinary.NewFromURL("cloudinary://<api_key>:<api_secret>@<cloud_name>")

	if err != nil {
		log.Fatal(err)
	}

	// Upload an image
	uploadResult, err := cld.Upload.Upload(context.Background(), image, uploader.UploadParams{
		Folder: "my_folder/my_sub_folder/",
	})

	if err != nil {
		log.Fatal(err)
	}

	return uploadResult.SecureURL, nil
}