package router

import (
	"CodeWithAzri/internal/app/module/course"
	"CodeWithAzri/internal/pkg/constant"
	"CodeWithAzri/internal/pkg/middleware"

	"github.com/go-chi/chi"
)

func RegisterCourseRoutes(router *Router, version string, module *course.Module, firebaseMiddleware *middleware.FirebaseMiddleware) {
	router.Mux.Group(
		func(r chi.Router) {
			r.Use(firebaseMiddleware.AuthMiddleware)
			r.Use(middleware.LoggerMiddleware)
			r.Route(
				constant.ApiPattern+version+constant.CoursesPattern,
				func(r chi.Router) {
					r.Get(constant.RootPattern+"{id}", module.Handler.GetCourseDetail)
					r.Get(constant.RootPattern, module.Handler.GetPaginatedCourses)
				},
			)
		},
	)
}
