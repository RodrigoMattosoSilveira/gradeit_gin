package controller

import (
    "fmt"

	"github.com/gin-gonic/gin"

	"madronetek.com/gradeit/model"
)

func validPerson(ctx *gin.Context) bool {
	_, exists := ctx.Get("person")
     return exists
}

func ValidateAssignment(ctx *gin.Context, entity model.Assignment) (bool, []string) {
	validAssignment := true
	errors := make([]string, 0)

	if !validPerson(ctx) {
		validAssignment = false
		errors = append(errors, fmt.Sprintf("Invalid assignment:  person %d does not exist", entity.PersonId))
	}

	return validAssignment, errors
}