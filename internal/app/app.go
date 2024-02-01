package app

import (
	"CodeWithAzri/internal/app/module/course"
	firebaseModule "CodeWithAzri/internal/app/module/firebase"
	"CodeWithAzri/internal/app/module/user"
	"CodeWithAzri/internal/pkg/constant"
	"CodeWithAzri/internal/pkg/middleware"
	"CodeWithAzri/internal/pkg/router"
	"CodeWithAzri/pkg/sqlPkg"
	"database/sql"
	"log"
	"net/http"

	_ "CodeWithAzri/docs"

	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type App struct {
	SqlDB          *sql.DB
	Router         *router.Router
	Middlewares    []any
	Validate       *validator.Validate
	UserModule     *user.Module
	FirebaseModule *firebaseModule.Module
	CourseModule   *course.Module
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
	a.CourseModule = course.NewModule(a.SqlDB, a.Validate)
}

func (a *App) initMigrations() {
	var err error

	err = a.UserModule.Migration.CreateUsersTable(a.SqlDB)
	if err != nil {
		log.Fatal(err)
	}
	err = a.CourseModule.Migration.CreateCourseTables(a.SqlDB)
	if err != nil {
		log.Fatal(err)
	}
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

	a.Router.RegisterGlobalMiddlewares(m)
	router.RegisterUserRoutes(a.Router, constant.V1, a.UserModule)
	router.RegisterCourseRoutes(a.Router, constant.V1, a.CourseModule)

}

func (a *App) initComponents() {
	a.initDB()
	a.Router = router.NewRouter()
	a.Validate = validator.New()
	a.initModules()
	a.initMigrations()
	a.initMiddlewares()
	a.initModuleRouters()
}

func (a *App) Run() {
	err := http.ListenAndServe(
		":8080",
		a.Router.Mux,
	)
	if err != nil {
		log.Fatal(err)
	}
}
