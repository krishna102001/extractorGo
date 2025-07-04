package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/krishna102001/extract_image_from_pdf/logic"
)

func ExtractPDFImageRoutes(c *gin.Context) {
	var req struct {
		Input_Pdf_File string `json:"input_pdf_file"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Printf("Failed to bind the request body %v", err)
		c.JSON(400, gin.H{"msg": "Failed to Bind the request body"})
		return
	}

	//call the logic of decoding the string into base64
	if err := logic.DecodeBase64(req.Input_Pdf_File); err != nil {
		c.JSON(400, gin.H{"msg": "Check the base64 file"})
		return
	}

	//call the logic to extract the pdf image from pdf file
	if err := logic.Extract_image_from_pdf_unidoc("out_pdf_file/sample.pdf"); err != nil {
		c.JSON(400, gin.H{"msg": "Failed to extract the pdf file"})
	}

	c.JSON(201, gin.H{"msg": "Image extracted successfully!!"})
}
