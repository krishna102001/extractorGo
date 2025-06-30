package logic

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
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

func Extract_image_from_pdf_unidoc(pathFile string) {
	log.Println("Started Extracting Embedded Images Using Unidoc pkg....")
	var out_dir string = "extract_images_unidoc"
	if err := os.Mkdir(out_dir, 0755); err != nil {
		log.Printf("Failed to create directory %s and error is %v", out_dir, err)
	}

	doc, err := os.Open(pathFile)
	if err != nil {
		log.Printf("Failed to open the file %s, error is %v", filepath.Base(pathFile), err)
	}
	defer doc.Close()
	pdfReader, err := model.NewPdfReader(doc)
	if err != nil {
		log.Println("Error Initalization pdf reader ", err)
	}

	numPage, err := pdfReader.GetNumPages()
	if err != nil {
		log.Println("Error in getting ", err)
	}
	log.Println("Total no. of pages it contain", numPage)

	for i := 1; i <= numPage; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			log.Printf("Error in getting page %d no.", i)
			continue
		}

		pextract, err := extractor.New(page)
		if err != nil {
			log.Println("Error in Initializing extractor ", err)
		}

		pimage, err := pextract.ExtractPageImages(nil)
		if err != nil {
			log.Println("Error in Extracting the Images ", err)
		}

		for idx, img := range pimage.Images {
			gimg, err := img.Image.ToGoImage()
			if err != nil {
				log.Println("Failed to load the image ", err)
			}
			// log.Println(gimg)
			filePath := filepath.Join(out_dir, fmt.Sprintf("page_%d.jpg", idx+1))
			f, err := os.Create(filePath)
			if err != nil {
				log.Println("Failed to create file:", err)
				return
			}
			defer f.Close()

			err = jpeg.Encode(f, gimg, nil)
			if err != nil {
				log.Println("Failed to encode jpeg format:", err)
			}
		}
	}
	log.Println("Exxtracting successfull")
}
