package logic

import (
	"errors"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func Convert_pdf_to_image(pathFile string) error {
	out_dir := "converted_img"
	if err := os.MkdirAll(out_dir, 0755); err != nil { //creating a directory
		log.Println("Failed to create a directory")
		return errors.New("failed to create dir")
	}

	doc, err := fitz.New(pathFile) //opening a pdf file
	if err != nil {
		log.Println("Error occur during opening file", err)
		return errors.New("failed to open file")
	}

	defer doc.Close() //closing the file when done with processing automatic

	for i := 0; i < doc.NumPage(); i++ { //running loop till pdf page
		img, err := doc.Image(i) // converting each pdf page into a image
		if err != nil {
			log.Println("Page error :", err)
			continue
		}
		filePath := filepath.Join(out_dir, fmt.Sprintf("page_%d.png", i+1)) // joining a file path to save the image to desired location
		f, err := os.Create(filePath)                                       //creating a in converted_image directory file
		if err != nil {
			log.Println("failed to create the file ", err)
			return errors.New("failed to create the file")
		}
		if err := png.Encode(f, img); err != nil { //creating a image from RGBA Code and encode it to png
			log.Println("Failed to encode it into the png file ", err)
			return errors.New("failed to encode the png")
		}
		f.Close()
	}
	return nil
}
