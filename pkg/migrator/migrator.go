package migrator

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"CodeWithAzri/pkg/sqlPkg"
)

func CreateMigrateDB(db *sql.DB) *gorm.DB {

	dsn := sqlPkg.GetConnectionString()
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}

	dbSql := stdlib.OpenDB(*config)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbSql,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return gormDB
}
