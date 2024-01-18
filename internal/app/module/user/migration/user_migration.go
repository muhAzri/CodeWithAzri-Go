package migration

import (
	"database/sql"
)

type UserMigration struct{}

func (m UserMigration) CreateUsersTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255) UNIQUE,
			created_at BIGINT,
			updated_at BIGINT
		);
	`
	_, err := db.Exec(query)
	return err
}
