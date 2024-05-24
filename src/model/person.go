package model

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID 			int64  	`gorm:"primaryKey"`
	Name		string
	Email		string	`gorm:"unique"`
	Password 	string
	Assignments []Assignment
}

type PersonBody struct {
	ID 			int64
	Name		string
	Email		string
	Password 	string
}