package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/krishna102001/extract_image_from_pdf/logic"
)

func main() {
	var app *gin.Engine = gin.Default()

	// logic.Convert_pdf_to_image("robots-war.pdf")
	// logic.Extract_image_from_pdf("Educational_Visit_to_Water_Treatment_Plant_Nashik.pdf")
	logic.Extract_image_from_pdf_unidoc("robots-war.pdf")

	app.Use(cors.Default())

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello Krishna !! Basic setup done"})
	})

	if err := app.Run(":8080"); err != nil {
		log.Printf("Failed to run server on port no. : %s", "8080")
	}
	log.Printf("Server is running on port no. : %s", "8080")
}
