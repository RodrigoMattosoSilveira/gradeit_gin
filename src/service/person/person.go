package service

import (
	"github.com/gin-gonic/gin"
	"madronetek.com/gradeit/model"
)

type service struct {
	repository PersonSvcInt
}

// NewPerson - is a factory function to inject store in service.
func NewPerson(p PersonSvcInt) PersonSvcInt {
	return service{repository: p}
}

func (s service) Create(ctx *gin.Context, person model.Person) {
	s.repository.Create(ctx, person)
}

func (s service) GetAll(ctx *gin.Context) {
	s.repository.GetAll(ctx)
}

func (s service) GetByID(ctx *gin.Context, id int64) {
	s.repository.GetByID(ctx, id)
}

func (s service) Update(ctx *gin.Context, person model.Person) {
	s.repository.Update(ctx, person)
}

func (s service) Delete(ctx *gin.Context, id int64) {
	s.repository.Delete(ctx, id)
}