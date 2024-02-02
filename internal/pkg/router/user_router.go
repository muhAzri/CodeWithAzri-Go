package router

import (
	"CodeWithAzri/internal/app/module/user"
	"CodeWithAzri/internal/pkg/constant"
	"CodeWithAzri/internal/pkg/middleware"

	"github.com/go-chi/chi"
)

func RegisterUserRoutes(router *Router, version string, module *user.Module, firebaseMiddleware *middleware.FirebaseMiddleware) {
	router.Mux.Group(
		func(r chi.Router) {
			r.Use(firebaseMiddleware.AuthMiddleware)
			r.Use(middleware.LoggerMiddleware)
			r.Route(
				constant.ApiPattern+version+constant.UsersPattern,
				func(r chi.Router) {
					r.Post(constant.RootPattern, module.Handler.Create)
					r.Get(constant.RootPattern+"profile", module.Handler.GetProfile)
				},
			)
		},
	)
}
