package router

import (
	"CodeWithAzri/internal/app/module/user"
	"CodeWithAzri/internal/pkg/constant"
	"CodeWithAzri/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(e *gin.Engine, version string, module *user.Module, firebaseMiddleware *middleware.FirebaseMiddleware) {
	routes := e.Group(constant.ApiPattern + version + constant.UsersPattern)
	routes.Use(firebaseMiddleware.AuthMiddleware())

	//Dummy Endpoint Remove It Later
	routes.GET(constant.RootPattern, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is a dummy endpoint",
		})
	})

}
