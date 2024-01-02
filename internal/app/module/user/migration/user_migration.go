package migration

import (
	"CodeWithAzri/internal/app/module/user/entity"

	"gorm.io/gorm"
)

type UserMigration struct{}

func (m UserMigration) CreateUsersTable(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}
