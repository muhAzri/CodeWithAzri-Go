package router

import (
	"CodeWithAzri/internal/app/module/user"
	"CodeWithAzri/internal/pkg/constant"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux, version string, module *user.Module) {
	pathPrefix := constant.ApiPattern + version
	userHandler := http.HandlerFunc(module.Handler.Create)

	mux.Handle(pathPrefix+"/users", userHandler)
}
