package user

import (
	"CodeWithAzri/internal/app/module/user/handler"
	"CodeWithAzri/internal/app/module/user/migration"
	"CodeWithAzri/internal/app/module/user/repository"
	"CodeWithAzri/internal/app/module/user/service"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type Module struct {
	Handler    *handler.Handler
	Service    service.UserService
	Repository repository.UserRepository
	Migration  *migration.UserMigration
}

func NewModule(db *sql.DB, validate *validator.Validate) *Module {
	m := new(Module)
	m.Repository = repository.NewRepository(db)
	m.Service = service.NewService(m.Repository)
	m.Handler = handler.NewHandler(m.Service, validate)
	m.Migration = &migration.UserMigration{}

	return m
}
