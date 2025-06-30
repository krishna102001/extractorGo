package logic

import (
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func Extract_image_from_pdf(pathFile string) {
	log.Println("Started Extracting Embedded Images......")
	var out_dir string = "extract_images"
	if err := os.MkdirAll(out_dir, 0755); err != nil { // creating a directory
		log.Printf("Failed to create out_dir %s and error is %v", out_dir, err)
	}

	if err := api.ExtractImagesFile(pathFile, out_dir, nil, nil); err != nil { //extracting all the embedded images in one go from pdfcpu package
		log.Printf("Failed to extract the embedded image %v: ", err)
	}

	log.Println("Finished Extracting Embedded Images.......")
}
