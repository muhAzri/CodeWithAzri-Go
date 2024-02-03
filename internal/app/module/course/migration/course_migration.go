package migration

import (
	"CodeWithAzri/internal/app/module/course/entity"
	"CodeWithAzri/pkg/migrator"
	"database/sql"
)

type CourseMigration struct{}

func (m CourseMigration) CreateCourseTables(db *sql.DB) error {
	migrationDB := migrator.CreateMigrateDB(db)

	migrationDB.AutoMigrate(
		entity.Course{},
		entity.CourseGallery{},
		entity.CourseSection{},
		entity.CourseLesson{},
		entity.CourseReviews{},
		entity.CourseTags{},
	)

	return nil
}
