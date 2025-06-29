package logic

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func Convert_pdf_to_image(pathFile string) {
	out_dir := "converted_img"
	if err := os.MkdirAll(out_dir, 0755); err != nil { //creating a directory
		log.Println("Failed to create a directory")
	}

	doc, err := fitz.New(pathFile) //opening a pdf file
	if err != nil {
		log.Println("Error occur during opening file", err)
	}

	defer doc.Close() //closing the file when done with processing automatic

	for i := 0; i < doc.NumPage(); i++ { //running loop till pdf page
		img, err := doc.Image(i) // converting each pdf page into a image
		if err != nil {
			log.Println("Page error :", err)
			continue
		}
		filePath := filepath.Join(out_dir, fmt.Sprintf("page_%d.png", i+1)) // joining a file path to save the image to desired location
		f, _ := os.Create(filePath)                                         //creating a in converted_image directory file
		png.Encode(f, img)                                                  //creating a image from RGBA Code and encode it to png
		f.Close()
	}
}
