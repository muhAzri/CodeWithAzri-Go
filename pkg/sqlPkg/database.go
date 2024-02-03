package sqlPkg

import (
	"database/sql"
	"fmt"
	"CodeWithAzri/pkg/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Initialize() (*sql.DB, error) {
	dbUsername := config.GetEnvValue("DB_USER")
	dbPassword := config.GetEnvValue("DB_PASS")
	dbName := config.GetEnvValue("DB_NAME")
	dbHost := config.GetEnvValue("DB_HOST")
	dbPort := config.GetEnvValue("DB_PORT")
	dbSSLMode := config.GetEnvValue("DB_SSL_MODE")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbSSLMode)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
