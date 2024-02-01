package router

import (
	"CodeWithAzri/internal/pkg/middleware"

	"github.com/go-chi/chi"
)

type Router struct {
	Mux *chi.Mux
}

func NewRouter() *Router {
	r := &Router{}
	r.Mux = chi.NewRouter()
	return r
}

func (r *Router) RegisterGlobalMiddlewares(firebaseMiddleware *middleware.FirebaseMiddleware) {
	r.Mux.Use(
		firebaseMiddleware.AuthMiddleware,
		middleware.LoggerMiddleware,
	)
}
