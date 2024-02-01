package router

import (
	"CodeWithAzri/internal/app/module/user"
	"CodeWithAzri/internal/pkg/constant"

	"github.com/go-chi/chi"
)

func RegisterUserRoutes(router *Router, version string, module *user.Module) {
	router.Mux.Group(
		func(r chi.Router) {
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
