package logic

import (
	"context"
	"errors"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryConfig struct {
	Client *cloudinary.Cloudinary
}

func Cloudinarycredentials() *CloudinaryConfig {

	cld, err := cloudinary.NewFromParams("cloud_name", "cloud_api_key", "cloud_secret")
	if err != nil {
		log.Fatal("Failed to Initialize Cloudinary Client : ", err)
	}
	return &CloudinaryConfig{
		Client: cld,
	}
}

func (cc *CloudinaryConfig) UploadFile(fileContent interface{}, folder string) (string, error) {
	uploadParams := uploader.UploadParams{
		Folder: folder,
	}
	uploadResult, err := cc.Client.Upload.Upload(context.Background(), fileContent, uploadParams)
	if err != nil {
		log.Printf("upload file for %s and error %v", fileContent, err)
		return "", errors.New("failed to upload the file")
	}
	return uploadResult.SecureURL, nil
}
