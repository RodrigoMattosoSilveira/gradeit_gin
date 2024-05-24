package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"madronetek.com/gradeit/config"
	"madronetek.com/gradeit/model"
)

func validPerson(personId int64) bool {
	var person model.Person

	result := config.DB.First(&person, personId)
    return result.Error == nil
}

func dueNotOld(due time.Time) bool {
	//  Reduce everything to the midnight UTC prior the time
	timeNowMidNight := time.Now()
	timeNowMidNight = time.Date(timeNowMidNight.Year(), timeNowMidNight.Month(), timeNowMidNight.Day(), 0, 0, 0, 0, time.UTC)
	dueMidNight := time.Date(due.Year(), due.Month(), due.Day(), 0, 0, 0, 0, time.UTC)
	return !dueMidNight.Before(timeNowMidNight)
}

func ValidateAssignment(ctx *gin.Context, body model.AssignmentBody) (bool, []string) {
	validAssignment := true
	errors := make([]string, 0)

	// The Person linked to the Assignment must be in the People table
	if !validPerson(body.PersonId) {
		validAssignment = false
		errors = append(errors, fmt.Sprintf("Invalid Assignment:  person %d does not exist", body.PersonId))
	}

	// The Assignment date must equal or greater than today's date
	if !dueNotOld(body.Due) {
		validAssignment = false
		errors = append(errors, fmt.Sprintf("Invalid Assignment:  time %s less than today", body.Due))
	}

	return validAssignment, errors
}