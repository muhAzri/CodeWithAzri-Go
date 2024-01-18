package router

import (
	"CodeWithAzri/internal/pkg/middleware"
	"net/http"
)

func RegisterGlobalMiddleware(mux *http.ServeMux, firebaseMiddleware *middleware.FirebaseMiddleware) http.Handler {
	combinedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your logic for combining the middlewares goes here
		firebaseMiddleware.AuthMiddleware(middleware.LoggerMiddleware(mux)).ServeHTTP(w, r)
	})

	return combinedHandler
}
