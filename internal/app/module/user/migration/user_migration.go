package migration

import (
	"CodeWithAzri/internal/app/module/user/entity"
	"CodeWithAzri/pkg/migrator"
	"database/sql"
)

type UserMigration struct{}

func (m UserMigration) CreateUsersTable(db *sql.DB) error {
	migrationDB := migrator.CreateMigrateDB(db)

	migrationDB.AutoMigrate(
		&entity.User{},
	)

	return nil
}
