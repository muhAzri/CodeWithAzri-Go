package router

import (
	"CodeWithAzri/internal/app/module/user"
	"CodeWithAzri/internal/pkg/constant"
	"CodeWithAzri/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(e *gin.Engine, version string, module *user.Module, firebaseMiddleware *middleware.FirebaseMiddleware) {
	routes := e.Group(constant.ApiPattern + version)
	routes.Use(firebaseMiddleware.AuthMiddleware())

	routes.POST(constant.UsersPattern, module.Handler.Create)

}
