package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/krishna102001/extract_image_from_pdf/database"
	"github.com/krishna102001/extract_image_from_pdf/logic"
)

func ExtractPDFImageRoutes(c *gin.Context) {
	var pdf_addr string

	switch c.ContentType() {
	case "application/json":
		log.Println("Base64 PDF detected...........decoding")
		var req struct {
			Input_Pdf_File string `json:"input_pdf_file"`
		}

		if err := c.ShouldBindJSON(&req); err != nil || req.Input_Pdf_File == "" {
			c.JSON(400, gin.H{"msg": "input_pdf_file is required in JSON"})
			return
		}

		//call the logic of decoding the string into base64
		decode_path, err := logic.DecodeBase64(req.Input_Pdf_File)
		if err != nil {
			c.JSON(400, gin.H{"msg": "Invalid the base64 file"})
			return
		}
		pdf_addr = decode_path
	case "multipart/form-data":
		file, err := c.FormFile("file")
		if err != nil {
			log.Println("No file or base64 input provided", err)
			c.JSON(400, gin.H{"msg": "Provide either base64 or file input"})
			return
		}
		//generating unique name
		uniqueId, err := uuid.NewRandom()
		if err != nil {
			log.Println("failed to generate the uuid ", err)
			c.JSON(500, gin.H{"msg": "failed to generate unique name"})
			return
		}

		savePath := fmt.Sprintf("./out_pdf_file/ %v_%s", uniqueId, file.Filename)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			log.Println("Failed to save uploaded file:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to save file"})
			return
		}
		pdf_addr = savePath
	default:
		c.JSON(415, gin.H{"msg": "Unsupported Content‑Type"})
		return
	}

	var ExtractId string
	var err error
	//call the logic to extract the pdf image from pdf file
	if ExtractId, err = logic.Extract_image_from_pdf_unidoc(pdf_addr); err != nil {
		c.JSON(400, gin.H{"msg": "Failed to extract the pdf file"})
		return
	}

	// --------------------------- fetching the data ------------------------
	var savedData struct {
		ExtractId   uuid.UUID `json:"extract_id"`
		ResponseUrl string    `json:"respsone_url"`
	}

	if err = database.DB.Model(&database.ExtractsTable{}).Select("extract_id", "response_url").
		Where("extract_id = ?", uuid.MustParse(ExtractId)).
		First(&savedData).Error; err != nil {
		log.Println("Failed to fetch the data from database", err)
		c.JSON(500, gin.H{"msg": "Failed to fetch the data"})
		return
	}

	err = os.Remove(pdf_addr)
	if err != nil {
		log.Printf("Failed to remove the file %s from disk after upload error is %v ", pdf_addr, err)
	}
	c.JSON(201, gin.H{"data": savedData})
}

func GetExtract(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("Param is empty")
		c.JSON(400, gin.H{"msg": "invalid param id"})
		return
	}
	var savedData = struct {
		ExtractId   uuid.UUID `json:"extract_id"`
		ResponseUrl string    `json:"response_url"`
	}{}

	if err := database.DB.Model(&database.ExtractsTable{}).Select("extract_id", "response_url").Where("extract_id =?", uuid.MustParse(id)).First(&savedData).Error; err != nil {
		log.Println("Failed to find the data from the database ", err)
		c.JSON(500, gin.H{"msg": "failed to get the data"})
		return
	}
	c.JSON(200, gin.H{"data": savedData})
}
