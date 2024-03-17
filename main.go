package main

import (
	"log"
	"os"

	"github.com/beezy-dev/traveler-crud-service/inits"
	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DataBackEnd()
}

func main() {
	ginPort := os.Getenv("GIN_PORT")
	ginMode := os.Getenv("GIN_MODE")

	if ginMode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	} else if ginMode != "release" {
		log.Println("Unknown mode:", ginMode)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + ginPort)
}
