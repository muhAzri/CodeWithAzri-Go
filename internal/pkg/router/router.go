package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterGlobalMiddlewares(e *gin.Engine) {

	e.Use(gin.Logger())

	e.Use(gin.Recovery())

}
