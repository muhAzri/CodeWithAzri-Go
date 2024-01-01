package pkg

import (
	"CodeWithAzri/internal/pkg/router"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	Engine *gin.Engine
}

func NewGin() *Gin {
	g := &Gin{}
	g.Engine = gin.Default()
	router.RegisterGlobalMiddlewares(g.Engine)
	return g
}
