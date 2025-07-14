package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

		savePath := "./out_pdf_file/" + file.Filename
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

	//call the logic to extract the pdf image from pdf file
	if err := logic.Extract_image_from_pdf_unidoc(pdf_addr); err != nil {
		c.JSON(400, gin.H{"msg": "Failed to extract the pdf file"})
	}

	c.JSON(201, gin.H{"msg": "Image extracted successfully!!"})
}
