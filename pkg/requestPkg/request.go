package requestPkg

import (
	"CodeWithAzri/internal/pkg/middleware"
	"net/http"

	"github.com/go-chi/chi"
)

func GetUserID(r *http.Request) string {
	userID := r.Context().Value(middleware.UserIDContextKey).(string)
	return userID
}

func GetURLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
