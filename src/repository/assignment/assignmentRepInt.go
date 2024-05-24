package assignment

import (
	"github.com/gin-gonic/gin"
	"madronetek.com/gradeit/model"
)

type AssignmentRepoInt interface {
	Create(ctx *gin.Context, person model.Assignment)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context, id int64)
	Update(ctx *gin.Context, person model.Assignment)
	Delete(ictx *gin.Context, id int64)
}