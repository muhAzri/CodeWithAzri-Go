package gormPkg

import (
	"CodeWithAzri/pkg/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initiliaze() *gorm.DB {
	dbUsername := config.GetEnvValue("DB_USER")
	dbPassword := config.GetEnvValue("DB_PASS")
	dbName := config.GetEnvValue("DB_NAME")
	dbHost := config.GetEnvValue("DB_HOST")
	dbPort := config.GetEnvValue("DB_PORT")
	dbSSLMode := config.GetEnvValue("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", dbHost, dbUsername, dbPassword, dbName, dbPort, dbSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}

	return db
}
