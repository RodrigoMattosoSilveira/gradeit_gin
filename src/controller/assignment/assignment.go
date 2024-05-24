package controller

import (
	"strconv"

	"madronetek.com/gradeit/model"
	service "madronetek.com/gradeit/service/assignment"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.AssignmentSvcInt
}

// NewAssignment - is a factory function to inject service in handler.
//
//nolint:revive // handler has to be unexported
func NewAssignment(s service.AssignmentSvcInt) controller {
	return controller{service: s}
}

func (c controller) Create(ctx *gin.Context) {
	var body model.AssignmentBody
	ctx.BindJSON(&body)

	valid, errors := ValidateAssignment(ctx, body)
	if !valid {
		ctx.JSON(500, gin.H{"error": errors})
		return
	}

	assignment := model.Assignment{PersonId: body.PersonId, Description: body.Description, Due: body.Due}
	assignment.Due = assignment.Due.UTC()

	c.service.Create(ctx, assignment)
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
	var body model.AssignmentBody
	ctx.BindJSON(&body)
	
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}

	valid, errors := ValidateAssignment(ctx, body)
	if !valid {
		ctx.JSON(500, gin.H{"error": errors})
		return
	}

	assignment := model.Assignment{ID: idInt,PersonId: body.PersonId, Description: body.Description, Due: body.Due}
	c.service.Update(ctx, assignment)
}

func (c controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}
	c.service.Delete(ctx, idInt)
}