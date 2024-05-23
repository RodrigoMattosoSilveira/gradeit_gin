package main

import (
	"github.com/gin-gonic/gin"

	"madronetek.com/gradeit/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "Hello World!",
		})
	})
	r.Run()
}