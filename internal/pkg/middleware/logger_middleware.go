package middleware

import (
	"net/http"

	"github.com/MadAppGang/httplog"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return httplog.Logger(next)
}
