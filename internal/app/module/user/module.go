package user

import (
	"CodeWithAzri/internal/app/module/user/handler"
	"CodeWithAzri/internal/app/module/user/repository"
	"CodeWithAzri/internal/app/module/user/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Module struct {
	Handler    *handler.Handler
	Service    *service.Service
	Repository *repository.Repository
}

func NewModule(db *gorm.DB, validate *validator.Validate) *Module {
	m := new(Module)
	m.Repository = repository.NewRepository(db)
	m.Service = service.NewService(*m.Repository)
	m.Handler = handler.NewHandler(m.Service, validate)
	return m
}
