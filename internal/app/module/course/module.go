package course

import (
	"CodeWithAzri/internal/app/module/course/handler"
	"CodeWithAzri/internal/app/module/course/migration"
	"CodeWithAzri/internal/app/module/course/repository"
	"CodeWithAzri/internal/app/module/course/service"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type Module struct {
	Handler    *handler.Handler
	Service    service.CourseService
	Repository repository.CourseRepository
	Migration  *migration.CourseMigration
}

func NewModule(db *sql.DB, validate *validator.Validate) *Module {
	m := new(Module)
	m.Repository = repository.NewRepository(db)
	m.Service = service.NewCourseService(m.Repository)
	m.Handler = handler.NewHandler(m.Service, validate)
	m.Migration = &migration.CourseMigration{}

	return m
}
