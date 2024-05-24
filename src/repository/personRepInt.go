package repository

import (
	"github.com/gin-gonic/gin"
	"madronetek.com/gradeit/model"
)

type PersonRepoInt interface {
	Create(ctx *gin.Context, person model.Person)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context, id int64)
	Update(ctx *gin.Context, person model.Person)
	Delete(ictx *gin.Context, id int64)
}