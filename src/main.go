package main

import (
	"github.com/gin-gonic/gin"

	"madronetek.com/gradeit/config"
	"madronetek.com/gradeit/model"
	"madronetek.com/gradeit/repository"
	"madronetek.com/gradeit/service"
	"madronetek.com/gradeit/controller"
)

func init() {
	config.LoadEnv()
	config.DBInit()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "Hello World!",
		})
	})

	//  Set up tables
	config.DB.AutoMigrate(&model.Person{})

	// Set up routes
	repo := repository.NewPerson()
	svc := service.NewPerson(repo)
	ctrlr := controller.NewPerson(svc)
	r.POST("/person", ctrlr.Create)
	r.GET("/person", ctrlr.GetAll)
	r.GET("/person/:id", ctrlr.GetByID)
	r.PUT("/person/:id", ctrlr.Update)
	r.DELETE("/person/:id", ctrlr.Delete)

	// Launch service
	r.Run()
}