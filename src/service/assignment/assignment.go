package service

import (
	"github.com/gin-gonic/gin"
	"madronetek.com/gradeit/model"
)

type service struct {
	repository AssignmentSvcInt
}

// NewAssignment - is a factory function to inject store in service.
func NewAssignment(a AssignmentSvcInt) AssignmentSvcInt {
	return service{repository: a}
}

func (s service) Create(ctx *gin.Context, Assignment model.Assignment) {
	s.repository.Create(ctx, Assignment)
}

func (s service) GetAll(ctx *gin.Context) {
	s.repository.GetAll(ctx)
}

func (s service) GetByID(ctx *gin.Context, id int64) {
	s.repository.GetByID(ctx, id)
}

func (s service) Update(ctx *gin.Context, Assignment model.Assignment) {
	s.repository.Update(ctx, Assignment)
}

func (s service) Delete(ctx *gin.Context, id int64) {
	s.repository.Delete(ctx, id)
}