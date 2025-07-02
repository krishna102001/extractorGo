package routes

import (
	"log"

	"github.com/gin-gonic/gin"
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

}
