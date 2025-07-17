package logic

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryConfig struct {
	Client *cloudinary.Cloudinary
}

func Cloudinarycredentials() *CloudinaryConfig {
	cld_name := os.Getenv("CLOUDNIARY_NAME")
	if cld_name == "" {
		log.Fatalln("Failed to get the env of cloudinary-name")
	}
	cld_api_key := os.Getenv("CLOUDINARY_API_KEY")
	if cld_api_key == "" {
		log.Fatalln("Failed to get the env of cloudinary-api-key")
	}
	cld_secret := os.Getenv("CLOUDINARY_SECRET")
	if cld_secret == "" {
		log.Fatalln("Failed to get the env of cloudinary-secret")
	}

	cld, err := cloudinary.NewFromParams(cld_name, cld_api_key, cld_secret)
	if err != nil {
		log.Fatal("Failed to Initialize Cloudinary Client : ", err)
	}
	return &CloudinaryConfig{
		Client: cld,
	}
}

func (cc *CloudinaryConfig) UploadFile(fileContent interface{}, folder string) (string, error) {
	uploadParams := uploader.UploadParams{
		Folder:       folder,
		ResourceType: "raw",
	}
	uploadResult, err := cc.Client.Upload.Upload(context.Background(), fileContent, uploadParams)
	if err != nil {
		log.Printf("upload file for %s and error %v", fileContent, err)
		return "", errors.New("failed to upload the file")
	}
	return uploadResult.SecureURL, nil
}
