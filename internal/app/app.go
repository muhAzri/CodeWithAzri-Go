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
	"net/http"

	_ "CodeWithAzri/docs"

	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type App struct {
	SqlDB          *sql.DB
	Server         *pkg.Server
	Middlewares    []any
	UserModule     *user.Module
	FirebaseModule *firebaseModule.Module
	Validate       *validator.Validate
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
	http.Handle("/swagger/", http.StripPrefix("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)))

	m := a.Middlewares[0].(*middleware.FirebaseMiddleware)
	globalMiddlewares := router.RegisterGlobalMiddleware(a.Server.Mux, m)
	router.RegisterUserRoutes(a.Server.Mux, constant.V1, a.UserModule)

	http.Handle("/api/v1/", globalMiddlewares)
}

func (a *App) initComponents() {
	a.initDB()
	a.Server = pkg.NewServer()
	a.Validate = validator.New()
	a.initModules()
	a.initMigrations()
	a.initMiddlewares()
	a.initModuleRouters()
}

func (a *App) Run() {
	http.ListenAndServe(":8080", nil)
}
