package controller

import (
    "fmt"
    "net/mail"

	"madronetek.com/gradeit/model"
)

func validMailAddress(address string) bool {
    _, err := mail.ParseAddress(address)
     return err == nil
}

func ValidatePerson(person model.Person) (bool, []string) {
	validPerson := true
	errors := make([]string, 0)

	if !validMailAddress(person.Email) {
		validPerson = false
		errors = append(errors, fmt.Sprintf("Person %d: invalid email = %s", person.ID, person.Email))
	}

	return validPerson, errors
}