package admin

import (
	admin_datamodel "CodeWithAzri/internal/pkg/admin/datamodel"
	"fmt"
	"os"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/chi"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/go-chi/chi"
)

func InitializeAdmin(r *chi.Mux) {
	eng := engine.Default()

	fmt.Println(admin_datamodel.Generators)

	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:            os.Getenv("DB_HOST"),
				Port:            os.Getenv("DB_PORT"),
				User:            os.Getenv("DB_USER"),
				Pwd:             os.Getenv("DB_PASS"),
				Name:            os.Getenv("DB_NAME"),
				MaxIdleConns:    1,
				MaxOpenConns:    1,
				ConnMaxLifetime: time.Minute * 15,
				Driver:          config.DriverPostgresql,
			},
		},
		UrlPrefix: "admin",
		Language:  language.EN,
		IndexUrl:  "/",
		Debug:     true,
	}

	if err := eng.AddConfig(&cfg).
		AddGenerators(admin_datamodel.Generators).
		AddDisplayFilterXssJsFilter().
		Use(r); err != nil {
		panic(err)
	}
}
