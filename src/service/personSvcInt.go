package service

import (
	"github.com/gin-gonic/gin"
	"madronetek.com/gradeit/model"
)

type PersonSvcInt interface {
	Create(ctx *gin.Context, person model.Person)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context, id int64)
	Update(ctx *gin.Context, person model.Person)
	Delete(ctx *gin.Context, id int64)
}