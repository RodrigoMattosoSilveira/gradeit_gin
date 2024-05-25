package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"madronetek.com/gradeit/config"
	"madronetek.com/gradeit/model"
)

type repository struct{}

// NewPerson is a factory function for store layer that returns a interface type, UserInt
func NewPerson() PersonRepoInt {
	return repository{}
}

// cURL validation command: 
// curl -X POST --json '{"Name": "Albert Einstein", "Email": "einstein@mail.com", "Password": "einstein124"}' localhost:3000/person
// 
func (repo repository) Create(ctx *gin.Context, person model.Person) {
	result := config.DB.Create(&person)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": person})
}

// cURL validation command: 
// curl -X GET -localhost:3000/person
// 
func (repo repository) GetAll(ctx *gin.Context) {
	var people []model.Person

	// result := config.DB.Find(&people)
	// if result.Error != nil {
	// 	ctx.JSON(500, gin.H{"error": result.Error})
	// 	return
	// }
	err := config.DB.Model(&model.Person{}).Preload("Assignments").Find(&people).Error
	if err != nil {
		fmt.Println(err)
		ctx.JSON(500, gin.H{"error": "error getting people"})
		return
	}
	ctx.JSON(200, gin.H{"data": people})
}

// cURL validation command: 
// curl -X GET -localhost:3000/person/1
// 
func (repo repository) GetByID(ctx *gin.Context, id int64) {
	var person model.Person
	var assignments []model.Assignment

	result := config.DB.First(&person, id)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	result = config.DB.Where("person_id = ?", id).Find(&assignments)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	person.Assignments = append(person.Assignments, assignments...)

	ctx.JSON(200, gin.H{"data": person})
}

// cURL validation command: 
// curl -X PUT --json '{"Name": "Albert Einstein", "Email": "einstein@mail.com", "Password": "einstein124"}' localhost:3000/person/1
// 
func (repo repository) Update(ctx *gin.Context, person model.Person) {
	config.DB.Model(&person).Updates(model.Person{Name: person.Name, Email: person.Email, Password: person.Password})

	ctx.JSON(200, gin.H{"data": person})
}

// cURL validation command: 
// curl -X DELETE -localhost:3000/person/1
// 
func (repo repository) Delete(ctx *gin.Context, id int64) {
	config.DB.Delete(&model.Person{}, id)

 	ctx.JSON(200, gin.H{"data": "post has been deleted successfully"})
}