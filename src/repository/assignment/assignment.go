package assignment

import (
	"github.com/gin-gonic/gin"

	"madronetek.com/gradeit/config"
	"madronetek.com/gradeit/model"
)

type repository struct{}

// NewAssignment is a factory function for store layer that returns a interface type
func NewAssignment() AssignmentRepoInt {
	return repository{}
}

// cURL validation command
// curl -X POST --json '{"PersonId": 1, "Description": "Find 4th law", "Due": "2025-10-31T09:00:00.594Z"}' localhost:3000/assignment
// 
func (repo repository) Create(ctx *gin.Context, entity model.Assignment) {
	result := config.DB.Create(&entity)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": entity})
}

// cURL validation command
// curl -X GET localhost:3000/assignment
//
func (repo repository) GetAll(ctx *gin.Context) {
	var entity []model.Assignment

	result := config.DB.Find(&entity)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
	return
	}

	ctx.JSON(200, gin.H{"data": entity})
}

// cURL validation command
// curl -X GET localhost:3000/assignment/1
//
func (repo repository) GetByID(ctx *gin.Context, id int64) {
	var entity model.Assignment

	result := config.DB.First(& entity, id)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data":  entity})
}


// cURL validation command
// curl -X PUT --json '{"PersonId": 1, "Description": "Find 4th law", "Due": "2025-10-31T09:00:00.594Z"}' localhost:3000/assignment/1
// 
func (repo repository) Update(ctx *gin.Context, entity model.Assignment) {
	config.DB.Model(&entity).Updates(model.Assignment{PersonId: entity.PersonId, Description: entity.Description, Due: entity.Due})

	ctx.JSON(200, gin.H{"data": entity})
}

// cURL validation command
// curl -X DELETE localhost:3000/assignment/1
//
func (repo repository) Delete(ctx *gin.Context, id int64) {
	config.DB.Delete(&model.Assignment{}, id)

 	ctx.JSON(200, gin.H{"data": "assignment has been deleted successfully"})
}