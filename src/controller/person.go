package controller

import (
	"strconv"

	"madronetek.com/gradeit/model"
	"madronetek.com/gradeit/service"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.PersonSvcInt
}

// NewPerson - is a factory function to inject service in handler.
//
//nolint:revive // handler has to be unexported
func NewPerson(s service.PersonSvcInt) controller {
	return controller{service: s}
}

func (c controller) Create(ctx *gin.Context) {
	var body model.Person
	ctx.BindJSON(&body)

	// Might need this later when authenticating and authorizing
	person := model.Person{Name: body.Name, Email: body.Email, Password: body.Password}

	c.service.Create(ctx, person)
}

func (c controller) GetAll(ctx *gin.Context) {
	c.service.GetAll(ctx)
}

func (c controller) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}
	c.service.GetByID(ctx, idInt)
}

func (c controller) Update(ctx *gin.Context) {
	var body model.Person
	ctx.BindJSON(&body)
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}

	// Might need this later when authenticating and authorizing
	person := model.Person{ID: idInt, Name: body.Name, Email: body.Email, Password: body.Password}
	c.service.Update(ctx, person)
}

func (c controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}
	c.service.Delete(ctx, idInt)
}