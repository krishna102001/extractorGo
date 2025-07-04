package logic

import (
	"errors"
	"fmt"
	"image/png"
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

func Extract_image_from_pdf_unidoc(pathFile string) error {
	log.Println("Started Extracting Embedded Images Using Unidoc pkg....")

	var out_dir string = "extract_images_unidoc"    // output directory name
	if err := os.Mkdir(out_dir, 0755); err != nil { // 755 owner can read ,write and execute will other can read
		log.Printf("Failed to create directory %s and error is %v", out_dir, err)
		return errors.New("failed to create a directory")
	}

	doc, err := os.Open(pathFile) // open the pdf file
	if err != nil {
		log.Printf("Failed to open the file %s, error is %v", filepath.Base(pathFile), err)
		return errors.New("failed to open the file")
	}

	defer doc.Close() // close the file when function end

	pdfReader, err := model.NewPdfReader(doc) // Initialize the pdf reader and read the open file
	if err != nil {
		log.Println("Error Initalization pdf reader ", err)
		return errors.New("error in initializing the pdf reader ")
	}

	numPage, err := pdfReader.GetNumPages() // total no. of page in pdf
	if err != nil {
		log.Println("Error in getting total length of pdf page ", err)
		return errors.New("error in getting total no. of pdf pages")
	}
	log.Println("Total no. of pages it contain", numPage)

	total := 0 // to count no of photo and used in creating a name

	for i := 1; i <= numPage; i++ { // loop all the page to get the embedded image in pdf
		page, err := pdfReader.GetPage(i) // get the page content of particular page in pdf file
		if err != nil {                   // if error occur then continue to next page
			log.Printf("Error in getting page %d no.", i)
			continue
		}

		pextract, err := extractor.New(page) // Initialize the extractor for that particular page
		if err != nil {
			log.Println("Error in Initializing extractor ", err)
			return errors.New("error in initializing the extractor")
		}

		pimage, err := pextract.ExtractPageImages(nil) // extracting the embedded image from that page and return as a list of images content
		if err != nil {
			log.Println("Error in Extracting the Images ", err)
			return errors.New("error in extracting the images")
		}

		for _, img := range pimage.Images { // loop all the images
			log.Printf("Started Extracting Image %d No.", total+1)
			gimg, err := img.Image.ToGoImage() // it will give golang image structure type Image interface {ColorModel() color.Model Bounds() Rectangle At(x, y int) color.Color}
			if err != nil {
				log.Println("Failed to load the image ", err)
				return errors.New("failed to load the images")
			}
			// log.Println(gimg)
			filePath := filepath.Join(out_dir, fmt.Sprintf("page_%d_.png", total+1)) // create a file path
			total++
			f, err := os.Create(filePath) // creating a file
			if err != nil {
				log.Println("Failed to create file:", err)
				return errors.New("failed to create a file")
			}

			err = png.Encode(f, gimg) // now encoding the image data on file
			if err != nil {
				log.Println("Failed to encode in png format:", err)
				return errors.New("failed to encode in png format ")
			}

			f.Close() // closing the file immeditately a to prevent from resource leaks

		}
	}
	log.Println("PDF Extracting successfull")
	return nil
}
