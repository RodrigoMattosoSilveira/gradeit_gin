package main

import (
	"github.com/gin-gonic/gin"

	"madronetek.com/gradeit/config"
	"madronetek.com/gradeit/model"
	personRep     "madronetek.com/gradeit/repository/person"
	personSvc     "madronetek.com/gradeit/service/person"
	personCtl     "madronetek.com/gradeit/controller/person"
	assignmentRep "madronetek.com/gradeit/repository/assignment"
	assignmentSvc "madronetek.com/gradeit/service/assignment"
	assignmentCtl "madronetek.com/gradeit/controller/assignment"

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

	// Set up Person routess
	repo := personRep.NewPerson()
	svc := personSvc.NewPerson(repo)
	ctrlr := personCtl.NewPerson(svc)
	r.POST("/person", ctrlr.Create)
	r.GET("/person", ctrlr.GetAll)
	r.GET("/person/:id", ctrlr.GetByID)
	r.PUT("/person/:id", ctrlr.Update)
	r.DELETE("/person/:id", ctrlr.Delete)

	// Set up Assignment routes
	repoAssignment := assignmentRep.NewAssignment()
	svcAssignment := assignmentSvc.NewAssignment(repoAssignment)
	ctrlrAssignment := assignmentCtl.NewAssignment(svcAssignment)
	r.POST("/assignment", ctrlrAssignment.Create)
	r.GET("/assignment", ctrlrAssignment.GetAll)
	r.GET("/assignment/:id", ctrlrAssignment.GetByID)
	r.PUT("/assignment/:id", ctrlrAssignment.Update)
	r.DELETE("/assignment/:id", ctrlrAssignment.Delete)

	// Launch service
	r.Run()
}