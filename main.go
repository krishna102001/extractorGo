package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/krishna102001/extract_image_from_pdf/database"
	"github.com/krishna102001/extract_image_from_pdf/middleware"
	"github.com/krishna102001/extract_image_from_pdf/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to intitalized the env file")
	}
}

func main() {
	var app *gin.Engine = gin.Default()

	database.InitializedDB()

	app.Use(cors.Default())

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello Krishna !! Basic setup done"})
	})

	router := app.Group("/api/v1")

	router.POST("/extract-pdf-image", middleware.RateLimiter(), routes.ExtractPDFImageRoutes)

	router.GET("/get/extract/:id", middleware.RateLimiter(), routes.GetExtract)

	router.POST("/convert-pdf-image", middleware.RateLimiter(), routes.ConvertPDFImageRoutes)

	router.GET("/get/convert/:id", middleware.RateLimiter(), routes.GetConvert)

	if err := app.Run(":8080"); err != nil {
		log.Printf("Failed to run server on port no. : %s", "8080")
	}
	log.Printf("Server is running on port no. : %s", "8080")
}
