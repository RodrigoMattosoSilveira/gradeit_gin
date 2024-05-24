package service

import (
	"github.com/gin-gonic/gin"
	"madronetek.com/gradeit/model"
)

type AssignmentSvcInt interface {
	Create(ctx *gin.Context, entity model.Assignment)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context, id int64)
	Update(ctx *gin.Context,entity model.Assignment)
	Delete(ctx *gin.Context, id int64)
}