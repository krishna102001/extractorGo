package logic

import (
	"errors"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
	"github.com/google/uuid"
	yw "github.com/yeka/zip"
)

func Convert_pdf_to_image(pathFile string) (string, error) {
	password := os.Getenv("ENCRYPT_PASS")
	if password == "" {
		log.Fatalln("Failed to get the encrypt pass")
	}
	log.Println("Started Converting pdf file into images .................")
	out_dir := "converted_img"
	_, err := os.Stat(out_dir)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(out_dir, 0755); err != nil { //creating a directory
			log.Println("Failed to create a directory")
			return "", errors.New("failed to create dir")
		}
	}

	doc, err := fitz.New(pathFile) //opening a pdf file
	if err != nil {
		log.Println("Error occur during opening file", err)
		return "", errors.New("failed to open file")
	}

	//generating unique name
	uniqueId, err := uuid.NewRandom()
	if err != nil {
		log.Println("failed to generate the uuid ", err)
		return "", errors.New("failed to generate the uuid")
	}

	// ------------------- zip file logic -------------------
	zipFile, err := os.Create(filepath.Join(out_dir, fmt.Sprintf("converted_%v_.zip", uniqueId)))
	if err != nil {
		log.Println("Failed to create the zip file in converted-image")
		return "", errors.New("failed to create the zip file")
	}
	defer zipFile.Close()

	zipWriter := yw.NewWriter(zipFile)
	defer zipWriter.Close()
	// -------------------------------------------------------

	defer doc.Close() //closing the file when done with processing automatic

	for i := 0; i < doc.NumPage(); i++ { //running loop till pdf page
		img, err := doc.Image(i) // converting each pdf page into a image
		if err != nil {
			log.Println("Page error :", err)
			continue
		}
		// -------------------------- encrypting each file using password ------------------------
		f, err := zipWriter.Encrypt(fmt.Sprintf("page_%d.png", i+1), password, yw.AES256Encryption) //creating a in converted_image directory file
		log.Printf("extracted page_%d.png", i+1)
		if err != nil {
			log.Println("failed to create the file ", err)
			return "", errors.New("failed to create the file")
		}
		if err := png.Encode(f, img); err != nil { //creating a image from RGBA Code and encode it to png
			log.Println("Failed to encode it into the png file ", err)
			return "", errors.New("failed to encode the png")
		}
	}
	log.Println("Successfully converted into images...............")

	log.Println("Uploading file to cloudinary.....................")
	cld := Cloudinarycredentials()
	upld_url, err := cld.UploadFile(zipFile.Name(), "converted_image")
	if err != nil {
		log.Println("Failed to upload the file", err)
		return "", errors.New("failed to upload the file")
	}
	log.Println("Upload successfull...................")
	return upld_url, nil
}
