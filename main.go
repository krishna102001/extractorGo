package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/krishna102001/extract_image_from_pdf/routes"
)

func main() {
	var app *gin.Engine = gin.Default()

	app.Use(cors.Default())

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello Krishna !! Basic setup done"})
	})

	router := app.Group("/api/v1")

	router.POST("/extract-pdf-image", routes.ExtractPDFImageRoutes)

	router.POST("/convert-pdf-image", routes.ConvertPDFImageRoutes)

	if err := app.Run(":8080"); err != nil {
		log.Printf("Failed to run server on port no. : %s", "8080")
	}
	log.Printf("Server is running on port no. : %s", "8080")
}
