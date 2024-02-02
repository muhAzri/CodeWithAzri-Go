package router

import (
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
