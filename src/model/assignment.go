package model

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	ID        	int64   `gorm:"primaryKey"`
	PersonId    int64	`gorm:"person_id"` 		
	Description	string
	Due 		time.Time
}