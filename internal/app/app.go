package app

import (
	firebaseModule "CodeWithAzri/internal/app/module/firebase"
	"CodeWithAzri/internal/app/module/user"
	"CodeWithAzri/internal/pkg"
	"CodeWithAzri/internal/pkg/constant"
	"CodeWithAzri/internal/pkg/middleware"
	"CodeWithAzri/internal/pkg/router"
	"CodeWithAzri/pkg/sqlPkg"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type App struct {
	SqlDB          *sql.DB
	Gin            *pkg.Gin
	Middlewares    []any
	UserModule     *user.Module
	FirebaseModule *firebaseModule.Module

	Validate *validator.Validate
}

func NewApp() *App {
	a := new(App)
	a.initComponents()
	return a
}

func (a *App) initDB() {
	var err error
	a.SqlDB, err = sqlPkg.Initialize()
	if err != nil {
		panic(err)
	}
}

func (a *App) initModules() {
	a.UserModule = user.NewModule(a.SqlDB, a.Validate)
	a.FirebaseModule = firebaseModule.NewModule()
}

func (a *App) initMigrations() {
	a.UserModule.Migration.CreateUsersTable(a.SqlDB)
}

func (a *App) initMiddlewares() {
	firebaseMiddleware := middleware.NewFirebaseMiddleware(a.FirebaseModule.FirebaseApp)
	a.Middlewares = append(a.Middlewares, firebaseMiddleware)
}

func (a *App) initModuleRouters() {
	m := a.Middlewares[0].(*middleware.FirebaseMiddleware)
	router.RegisterUserRoutes(a.Gin.Engine, constant.V1, a.UserModule, m)
}

func (a *App) initComponents() {
	a.initDB()
	a.Gin = pkg.NewGin()
	a.initModules()
	a.initMigrations()
	a.initMiddlewares()
	a.initModuleRouters()
}

func (a *App) Run() {
	a.Gin.Engine.Run(":8080")
}
