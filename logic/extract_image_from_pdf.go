package logic

import (
	"errors"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/krishna102001/extract_image_from_pdf/database"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
	yw "github.com/yeka/zip"
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

func Extract_image_from_pdf_unidoc(pathFile string) (string, error) {
	password := os.Getenv("ENCRYPT_PASS")
	if password == "" {
		log.Fatal("Failed to get the encrypt pass")
	}
	log.Println("paswword .......................", password)
	log.Println("Started Extracting Embedded Images Using Unidoc pkg....")

	var out_dir string = "extract_images_unidoc" // output directory name
	_, err := os.Stat(out_dir)
	if os.IsNotExist(err) {
		if err := os.Mkdir(out_dir, 0755); err != nil { // 755 owner can read ,write and execute will other can read
			log.Printf("Failed to create directory %s and error is %v", out_dir, err)
			return "", errors.New("failed to create a directory")
		}
	}

	//generating unique name
	uniqueId, err := uuid.NewRandom()
	if err != nil {
		log.Println("failed to generate the uuid ", err)
		return "", errors.New("failed to generate the uuid")
	}

	// zip file creation program
	zipFile, err := os.Create(filepath.Join(out_dir, fmt.Sprintf("extracted_file_%v_.zip", uniqueId))) //created a zip file
	if err != nil {
		log.Printf("Failed to create zip file %v", err)
		return "", errors.New("failed to create a zip file")
	}
	defer zipFile.Close()

	zipWriter := yw.NewWriter(zipFile) //zip file writer initilaizied
	defer zipWriter.Close()

	doc, err := os.Open(pathFile) // open the pdf file
	if err != nil {
		log.Printf("Failed to open the file %s, error is %v", filepath.Base(pathFile), err)
		return "", errors.New("failed to open the file")
	}

	defer doc.Close() // close the file when function end

	pdfReader, err := model.NewPdfReader(doc) // Initialize the pdf reader and read the open file
	if err != nil {
		log.Println("Error Initalization pdf reader ", err)
		return "", errors.New("error in initializing the pdf reader ")
	}

	numPage, err := pdfReader.GetNumPages() // total no. of page in pdf
	if err != nil {
		log.Println("Error in getting total length of pdf page ", err)
		return "", errors.New("error in getting total no. of pdf pages")
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
			return "", errors.New("error in initializing the extractor")
		}

		pimage, err := pextract.ExtractPageImages(nil) // extracting the embedded image from that page and return as a list of images content
		if err != nil {
			log.Println("Error in Extracting the Images ", err)
			return "", errors.New("error in extracting the images")
		}

		for _, img := range pimage.Images { // loop all the images
			log.Printf("Started Extracting Image %d No.", total+1)
			gimg, err := img.Image.ToGoImage() // it will give golang image structure type Image interface {ColorModel() color.Model Bounds() Rectangle At(x, y int) color.Color}
			if err != nil {
				log.Println("Failed to load the image ", err)
				return "", errors.New("failed to load the images")
			}

			w, err := zipWriter.Encrypt(fmt.Sprintf("page_%d.png", total+1), password, yw.AES256Encryption)
			if err != nil {
				return "", fmt.Errorf("failed to create a zip file %d : %v", total, err)
			}

			err = png.Encode(w, gimg) // now encoding the image data on file
			if err != nil {
				log.Println("Failed to encode in png format:", err)
				return "", errors.New("failed to encode in png format")
			}
			total++
		}
	}

	log.Println("PDF Extracting successfull")

	log.Println("uploading file started.............")

	// --------------------cloudinary upload------------------
	cld := Cloudinarycredentials()
	zip_url, err := cld.UploadFile(zipFile.Name(), "extracted-data")
	if err != nil {
		log.Println("Failed to upload the file", err)
		return "", errors.New("failed to upload the file")
	}
	log.Println("uploading file successfull..........")

	var insertData = &database.ExtractsTable{
		DocName:     zipFile.Name(),
		ResponseUrl: zip_url,
	}

	//----------------- saving to database ----------------
	if err = database.DB.Model(&database.ExtractsTable{}).Create(&insertData).Error; err != nil {
		log.Printf("Error in Inserting the data into database %s", err.Error())
		return "", errors.New("failed to save in the database")
	}

	if err := os.Remove(zipFile.Name()); err != nil {
		log.Printf("Failed to delete the file %s and error is %v", zipFile.Name(), err)
	}

	return insertData.ExtractId.String(), nil
}
