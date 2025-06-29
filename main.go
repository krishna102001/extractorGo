package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var app *gin.Engine = gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello Krishna !! Basic setup done"})
	})

	if err := app.Run(":8080"); err != nil {
		log.Printf("Failed to run server on port no. : %s", "8080")
	}
	log.Printf("Server is running on port no. : %s", "8080")
}
